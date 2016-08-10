// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package machineundertaker

import (
	"github.com/juju/errors"
	"github.com/juju/loggo"
	"gopkg.in/juju/names.v2"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/watcher"
)

var logger = loggo.GetLogger("juju.apiserver.machineundertaker")

func init() {
	common.RegisterStandardFacade("MachineUndertaker", 1, newMachineUndertakerAPIFromState)
}

type MachineUndertakerAPI struct {
	backend        Backend
	resources      facade.Resources
	canManageModel func(modelUUID string) bool
}

// NewMachineUndertakerAPI implements the API used by the machine
// undertaker worker to find out what provider-level resources need to
// be cleaned up when a machine goes away.
func NewMachineUndertakerAPI(backend Backend, resources facade.Resources, authorizer facade.Authorizer) (*MachineUndertakerAPI, error) {
	if !authorizer.AuthModelManager() {
		return nil, errors.Trace(common.ErrPerm)
	}

	api := &MachineUndertakerAPI{
		backend:   backend,
		resources: resources,
		canManageModel: func(modelUUID string) bool {
			return modelUUID == authorizer.ConnectedModel()
		},
	}
	return api, nil
}

func newMachineUndertakerAPIFromState(st *state.State, res facade.Resources, auth facade.Authorizer) (*MachineUndertakerAPI, error) {
	return NewMachineUndertakerAPI(&backendShim{st}, res, auth)
}

// AllMachineRemovals returns tags for all of the machines that have
// been marked for removal in the requested model.
func (m *MachineUndertakerAPI) AllMachineRemovals(models params.Entities) (params.Entities, error) {
	err := m.checkModelAuthorization(models)
	if err != nil {
		return params.Entities{}, errors.Trace(err)
	}
	machineIds, err := m.backend.AllMachineRemovals()
	if err != nil {
		return params.Entities{}, errors.Trace(err)
	}
	var entities []params.Entity
	for _, id := range machineIds {
		entities = append(entities, params.Entity{
			Tag: names.NewMachineTag(id).String(),
		})
	}
	return params.Entities{Entities: entities}, nil
}

// GetMachineProviderInterfaceInfo returns the provider details for
// all network interfaces attached to the machines requested.
func (m *MachineUndertakerAPI) GetMachineProviderInterfaceInfo(args params.Entities) params.ProviderInterfaceInfoResults {
	results := make([]params.ProviderInterfaceInfoResult, len(args.Entities))
	for i, entity := range args.Entities {
		results[i].MachineTag = entity.Tag

		tag, err := names.ParseMachineTag(entity.Tag)
		if err != nil {
			results[i].Error = common.ServerError(err)
			continue
		}
		machine, err := m.backend.Machine(tag.Id())
		if err != nil {
			results[i].Error = common.ServerError(err)
			continue
		}
		interfaces, err := machine.AllProviderInterfaceInfos()
		if err != nil {
			results[i].Error = common.ServerError(err)
			continue
		}

		infos := make([]params.ProviderInterfaceInfo, len(interfaces))
		for i, info := range interfaces {
			infos[i].InterfaceName = info.InterfaceName
			infos[i].MACAddress = info.MACAddress
			infos[i].ProviderId = string(info.ProviderId)
		}

		results[i].Interfaces = infos
	}
	return params.ProviderInterfaceInfoResults{results}
}

// CompleteMachineRemovals removes the specified machines from the
// model database. It should only be called once any provider-level
// cleanup has been done for those machines.
func (m *MachineUndertakerAPI) CompleteMachineRemovals(args params.Entities) error {
	machineIDs, err := collectMachineIDs(args)
	if err != nil {
		return errors.Trace(err)
	}
	return m.backend.CompleteMachineRemovals(machineIDs...)
}

// WatchMachineRemovals returns a watcher that will signal each time a
// machine is marked for removal.
func (m *MachineUndertakerAPI) WatchMachineRemovals(models params.Entities) params.NotifyWatchResult {
	err := m.checkModelAuthorization(models)
	if err != nil {
		return params.NotifyWatchResult{Error: common.ServerError(err)}
	}
	watch := m.backend.WatchMachineRemovals()
	if _, ok := <-watch.Changes(); ok {
		return params.NotifyWatchResult{
			NotifyWatcherId: m.resources.Register(watch),
		}
	}
	return params.NotifyWatchResult{
		Error: common.ServerError(watcher.EnsureErr(watch)),
	}
}

func (m *MachineUndertakerAPI) checkModelAuthorization(entities params.Entities) error {
	if len(entities.Entities) == 0 {
		return errors.New("one model tag is required")
	}
	for _, entity := range entities.Entities {
		modelTag, err := names.ParseModelTag(entity.Tag)
		if err != nil {
			return errors.Trace(err)
		}
		if !m.canManageModel(modelTag.Id()) {
			return errors.Trace(common.ErrPerm)
		}
	}
	return nil
}

func collectMachineIDs(args params.Entities) ([]string, error) {
	machineIDs := make([]string, len(args.Entities))
	for i := range args.Entities {
		tag, err := names.ParseMachineTag(args.Entities[i].Tag)
		if err != nil {
			return nil, errors.Trace(err)
		}
		machineIDs[i] = tag.Id()
	}
	return machineIDs, nil
}
