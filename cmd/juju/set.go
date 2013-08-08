// Copyright 2012, 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"errors"
	"fmt"
	"strings"

	"launchpad.net/gnuflag"

	"launchpad.net/juju-core/charm"
	"launchpad.net/juju-core/cmd"
	"launchpad.net/juju-core/juju"
)

// SetCommand updates the configuration of a service
type SetCommand struct {
	EnvCommandBase
	ServiceName     string
	SettingsStrings map[string]string
	SettingsYAML    cmd.FileVar
}

func (c *SetCommand) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "set",
		Args:    "<service> name=value ...",
		Purpose: "set service config options",
		Doc:     "Set one or more configuration options for the specified service.",
	}
}

func (c *SetCommand) SetFlags(f *gnuflag.FlagSet) {
	c.EnvCommandBase.SetFlags(f)
	f.Var(&c.SettingsYAML, "config", "path to yaml-formatted service config")
}

func (c *SetCommand) Init(args []string) error {
	if len(args) == 0 || len(strings.Split(args[0], "=")) > 1 {
		return errors.New("no service name specified")
	}
	if c.SettingsYAML.Path != "" && len(args) > 1 {
		return errors.New("cannot specify --config when using key=value arguments")
	}
	c.ServiceName = args[0]
	settings, err := parse(args[1:])
	if err != nil {
		return err
	}
	c.SettingsStrings = settings
	return nil
}

// Run updates the configuration of a service.
func (c *SetCommand) Run(ctx *cmd.Context) error {
	conn, err := juju.NewConnFromName(c.EnvName)
	if err != nil {
		return err
	}
	defer conn.Close()
	service, err := conn.State.Service(c.ServiceName)
	if err != nil {
		return err
	}
	ch, _, err := service.Charm()
	if err != nil {
		return err
	}
	var settings charm.Settings
	if c.SettingsYAML.Path != "" {
		settingsYAML, err := c.SettingsYAML.Read(ctx)
		if err != nil {
			return err
		}
		settings, err = ch.Config().ParseSettingsYAML(settingsYAML, c.ServiceName)
		if err != nil {
			return err
		}
	} else if len(c.SettingsStrings) > 0 {
		settings, err = ch.Config().ParseSettingsStrings(c.SettingsStrings)
		if err != nil {
			return err
		}
	} else {
		return nil
	}
	return service.UpdateConfigSettings(settings)
}

// parse parses the option k=v strings into a map of options to be
// updated in the config. Keys with empty values are returned separately
// and should be removed.
func parse(options []string) (map[string]string, error) {
	kv := make(map[string]string)
	for _, o := range options {
		s := strings.SplitN(o, "=", 2)
		if len(s) != 2 || s[0] == "" {
			return nil, fmt.Errorf("invalid option: %q", o)
		}
		kv[s[0]] = s[1]
	}
	return kv, nil
}
