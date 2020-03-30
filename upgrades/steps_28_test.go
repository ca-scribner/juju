// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package upgrades_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/mock/gomock"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/version"
	gc "gopkg.in/check.v1"
	"gopkg.in/juju/charm.v6/hooks"
	"gopkg.in/juju/names.v3"
	"gopkg.in/yaml.v2"

	"github.com/juju/juju/agent"
	"github.com/juju/juju/api/base"
	apicallermocks "github.com/juju/juju/api/base/mocks"
	"github.com/juju/juju/apiserver/params"
	k8sprovider "github.com/juju/juju/caas/kubernetes/provider"
	"github.com/juju/juju/testing"
	"github.com/juju/juju/upgrades"
	"github.com/juju/juju/upgrades/mocks"
	"github.com/juju/juju/worker/uniter/hook"
	"github.com/juju/juju/worker/uniter/operation"
	configsettermocks "github.com/juju/juju/worker/upgradedatabase/mocks"
)

var v280 = version.MustParse("2.8.0")

type steps28Suite struct {
	testing.BaseSuite
}

var _ = gc.Suite(&steps28Suite{})

func (s *steps28Suite) TestIncrementTasksSequence(c *gc.C) {
	step := findStateStep(c, v280, "increment tasks sequence by 1")
	c.Assert(step.Targets(), jc.DeepEquals, []upgrades.Target{upgrades.DatabaseMaster})
}

func (s *steps28Suite) TestAddMachineIDToSubordinates(c *gc.C) {
	step := findStateStep(c, v280, "add machine ID to subordinate units")
	c.Assert(step.Targets(), jc.DeepEquals, []upgrades.Target{upgrades.DatabaseMaster})
}

func (s *steps28Suite) TestPopulateRebootHandledFlagsForDeployedUnits(c *gc.C) {
	step := findStep(c, v280, "ensure currently running units do not fire start hooks thinking a reboot has occurred")
	c.Assert(step.Targets(), jc.DeepEquals, []upgrades.Target{upgrades.HostMachine})
}

func (s *steps28Suite) TestMoveUniterStateToControllerStep(c *gc.C) {
	step := findStep(c, v280, "write uniter state to controller for all running units and remove files")
	c.Assert(step.Targets(), jc.DeepEquals, []upgrades.Target{upgrades.HostMachine})
}

type mockSteps28Suite struct {
	testing.BaseSuite

	dataDir    string
	tagOne     names.Tag
	tagTwo     names.Tag
	storTagOne names.Tag

	opStateOne         operation.State
	opStateTwo         operation.State
	opStateOneYaml     string
	opStateTwoYaml     string
	opStateOneFileName string
	opStateTwoFileName string

	opStorOne         bool
	opStorOneYaml     string
	opStorOneFileName string

	mockCtx         *mocks.MockContext
	mockClient      *mocks.MockUpgradeStepsClient
	mockAgentConfig *configsettermocks.MockConfigSetter
	mockAPICaller   *apicallermocks.MockAPICaller
}

var _ = gc.Suite(&mockSteps28Suite{})

func (s *mockSteps28Suite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)

	s.tagOne = names.NewUnitTag("testing/0")
	s.tagTwo = names.NewUnitTag("testing/1")
	s.storTagOne = names.NewStorageTag("data/3")

	s.opStateOne = operation.State{
		Leader: true,
		Kind:   operation.Continue,
		Step:   operation.Pending,
	}

	s.opStateTwo = operation.State{
		Kind: operation.RunHook,
		Step: operation.Pending,
		Hook: &hook.Info{Kind: hooks.ConfigChanged},
	}

	s.opStorOne = true

	s.dataDir = c.MkDir()
	agentDir := agent.BaseDir(s.dataDir)

	unitOneStateDir := filepath.Join(agentDir, s.tagOne.String(), "state")
	err := os.MkdirAll(unitOneStateDir, 0755)
	c.Assert(err, jc.ErrorIsNil)
	s.opStateOneYaml, s.opStateOneFileName = writeUnitStateFile(c, unitOneStateDir, s.opStateOne)

	s.opStorOneYaml, s.opStorOneFileName = writeStorageState(c, unitOneStateDir, s.storTagOne, s.opStorOne)

	unitTwoStateDir := filepath.Join(agentDir, s.tagTwo.String(), "state")
	err = os.MkdirAll(unitTwoStateDir, 0755)
	c.Assert(err, jc.ErrorIsNil)
	s.opStateTwoYaml, s.opStateTwoFileName = writeUnitStateFile(c, unitTwoStateDir, s.opStateTwo)
}

// writeUnitStateFile writes the operation.State in yaml format to the
// path/uniter/state file.  It returns the yaml in string form and the
// full path to the file written.
func writeUnitStateFile(c *gc.C, path string, st operation.State) (string, string) {
	filePath := filepath.Join(path, "uniter")

	err := st.Validate()
	c.Assert(err, jc.ErrorIsNil)
	content, err := yaml.Marshal(st)
	c.Assert(err, jc.ErrorIsNil)

	err = ioutil.WriteFile(filePath, content, 0644)
	c.Assert(err, jc.ErrorIsNil)

	return string(content), filePath
}

func writeStorageState(c *gc.C, path string, storTag names.Tag, attached bool) (string, string) {
	storDir := filepath.Join(path, "storage")
	err := os.MkdirAll(storDir, 0755)
	c.Assert(err, jc.ErrorIsNil)

	storFileName := strings.Replace(storTag.Id(), "/", "-", -1)
	filePath := filepath.Join(storDir, storFileName)

	data := diskInfo{Attached: &attached}
	content, err := yaml.Marshal(data)
	c.Assert(err, jc.ErrorIsNil)

	err = ioutil.WriteFile(filePath, content, 0644)
	c.Assert(err, jc.ErrorIsNil)

	expectedYaml, err := yaml.Marshal(map[string]bool{storTag.Id(): attached})
	c.Assert(err, jc.ErrorIsNil)

	return string(expectedYaml), filePath
}

type diskInfo struct {
	Attached *bool `yaml:"attached,omitempty"`
}

func (s *mockSteps28Suite) TestMoveUniterStateToControllerNotMachine(c *gc.C) {
	defer s.setup(c).Finish()
	s.expectAPIState()
	s.expectAgentConfigUnitTag()
	s.patchClient()
	err := upgrades.MoveUniterStateToController(s.mockCtx)
	c.Assert(err, jc.ErrorIsNil)
}
func (s *mockSteps28Suite) TestMoveUniterStateToControllerIAAS(c *gc.C) {
	defer s.setup(c).Finish()
	s.expectAPIState()
	s.expectAgentConfigMachineTag()
	s.expectAgentConfigValueIAAS()
	s.expectWriteTwoUniterState(c)
	s.patchClient()

	err := upgrades.MoveUniterStateToController(s.mockCtx)
	c.Assert(err, jc.ErrorIsNil)
	_, err = os.Stat(s.opStateOneFileName)
	c.Assert(err, jc.Satisfies, os.IsNotExist)
	_, err = os.Stat(s.opStateTwoFileName)
	c.Assert(err, jc.Satisfies, os.IsNotExist)
	_, err = os.Stat(s.opStorOneFileName)
	c.Assert(err, jc.Satisfies, os.IsNotExist)

	// Check idempotent
	err = upgrades.MoveUniterStateToController(s.mockCtx)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *mockSteps28Suite) TestMoveUniterStateToControllerCAASDoesNothing(c *gc.C) {
	// TODO: (hml) 27-03-2020
	// remove when uniterstate moved for CAAS units and/or relations etc
	// added to move.
	defer s.setup(c).Finish()
	s.expectAgentConfigMachineTag()
	s.expectAgentConfigValueCAAS()
	s.patchClient()

	// Check idempotent
	for i := 0; i < 2; i += 1 {
		c.Logf("round %d", i)
		err := upgrades.MoveUniterStateToController(s.mockCtx)
		c.Assert(err, jc.ErrorIsNil)
	}
}

func (s *mockSteps28Suite) setup(c *gc.C) *gomock.Controller {
	ctlr := gomock.NewController(c)

	s.mockCtx = mocks.NewMockContext(ctlr)
	s.mockAPICaller = apicallermocks.NewMockAPICaller(ctlr)
	s.mockAgentConfig = configsettermocks.NewMockConfigSetter(ctlr)
	s.mockClient = mocks.NewMockUpgradeStepsClient(ctlr)

	s.expectAgentConfig()
	s.expectDataDir()

	return ctlr
}

func (s *mockSteps28Suite) expectAgentConfig() {
	s.mockCtx.EXPECT().AgentConfig().Return(s.mockAgentConfig).AnyTimes()
}

func (s *mockSteps28Suite) expectAPIState() {
	s.mockCtx.EXPECT().APIState().Return(s.mockAPICaller).AnyTimes()
}

func (s *mockSteps28Suite) expectDataDir() {
	s.mockAgentConfig.EXPECT().DataDir().Return(s.dataDir).AnyTimes()
}

func (s *mockSteps28Suite) expectAgentConfigValueCAAS() {
	s.mockAgentConfig.EXPECT().Value(agent.ProviderType).Return(k8sprovider.CAASProviderType).AnyTimes()
}

func (s *mockSteps28Suite) expectAgentConfigValueIAAS() {
	s.mockAgentConfig.EXPECT().Value(agent.ProviderType).Return("IAAS").AnyTimes()
}

func (s *mockSteps28Suite) expectAgentConfigMachineTag() {
	s.mockAgentConfig.EXPECT().Tag().Return(names.NewMachineTag("0")).AnyTimes()
}

func (s *mockSteps28Suite) expectAgentConfigUnitTag() {
	s.mockAgentConfig.EXPECT().Tag().Return(names.NewUnitTag("test/0"))
}

func (s *mockSteps28Suite) patchClient() {
	s.PatchValue(upgrades.GetUpgradeStepsClient, func(_ base.APICaller) upgrades.UpgradeStepsClient {
		return s.mockClient
	})
}

func (s *mockSteps28Suite) expectWriteTwoUniterState(c *gc.C) {
	args := []params.SetUnitStateArg{{
		Tag:          s.tagOne.String(),
		UniterState:  &s.opStateOneYaml,
		StorageState: &s.opStorOneYaml,
	}, {
		Tag:         s.tagTwo.String(),
		UniterState: &s.opStateTwoYaml,
	},
	}
	cExp := s.mockClient.EXPECT()
	cExp.WriteUniterState(unitStateMatcher{c, args}).Return(nil)
}

type unitStateMatcher struct {
	c        *gc.C
	expected []params.SetUnitStateArg
}

func (m unitStateMatcher) Matches(x interface{}) bool {
	obtained, ok := x.([]params.SetUnitStateArg)
	m.c.Assert(ok, jc.IsTrue)
	if !ok {
		return false
	}

	m.c.Assert(obtained, gc.HasLen, len(m.expected))

	for _, obtainedArg := range obtained {
		var matched bool
		for _, expectedArg := range m.expected {
			if obtainedArg.Tag != expectedArg.Tag {
				continue
			}
			assertSetUnitStateArg(m.c, expectedArg, obtainedArg)
			matched = true
		}
		m.c.Assert(matched, jc.IsTrue)
	}
	return true
}

func assertSetUnitStateArg(c *gc.C, expectedArg, obtainedArg params.SetUnitStateArg) {
	if expectedArg.State != nil {
		c.Assert(obtainedArg.State, gc.NotNil)
		c.Assert(*obtainedArg.State, gc.DeepEquals, *expectedArg.State)
	} else {
		c.Assert(obtainedArg.State, gc.IsNil)
	}
	if expectedArg.UniterState != nil {
		c.Assert(*obtainedArg.UniterState, gc.Equals, *expectedArg.UniterState)
	} else {
		c.Assert(obtainedArg.UniterState, gc.IsNil)
	}
	if expectedArg.RelationState != nil {
		c.Assert(*obtainedArg.RelationState, gc.DeepEquals, *expectedArg.RelationState)
	} else {
		c.Assert(obtainedArg.State, gc.IsNil)
	}
	if expectedArg.StorageState != nil {
		c.Assert(*obtainedArg.StorageState, gc.Equals, *expectedArg.StorageState)
	} else {
		c.Assert(obtainedArg.StorageState, gc.IsNil)
	}
}

func (m unitStateMatcher) String() string {
	return "Match the contents of the UniterState pointer in params.SetUnitStateArg"
}
