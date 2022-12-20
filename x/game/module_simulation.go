package game

import (
	"math/rand"

	"github.com/G4AL-Entertainment/g4al-chain/testutil/sample"
	gamesimulation "github.com/G4AL-Entertainment/g4al-chain/x/game/simulation"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = gamesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateProject = "op_weight_msg_project"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProject int = 100

	opWeightMsgUpdateProject = "op_weight_msg_project"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProject int = 100

	opWeightMsgAddDelegate = "op_weight_msg_add_delegate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddDelegate int = 100

	opWeightMsgRemoveDelegate = "op_weight_msg_remove_delegate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveDelegate int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	gameGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		ProjectList: []types.Project{
			{
				Creator: sample.AccAddress(),
				Symbol:  "0",
			},
			{
				Creator: sample.AccAddress(),
				Symbol:  "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&gameGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = defaultWeightMsgCreateProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProject,
		gamesimulation.SimulateMsgCreateProject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateProject, &weightMsgUpdateProject, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProject = defaultWeightMsgUpdateProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProject,
		gamesimulation.SimulateMsgUpdateProject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddDelegate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddDelegate, &weightMsgAddDelegate, nil,
		func(_ *rand.Rand) {
			weightMsgAddDelegate = defaultWeightMsgAddDelegate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddDelegate,
		gamesimulation.SimulateMsgAddDelegate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveDelegate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveDelegate, &weightMsgRemoveDelegate, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveDelegate = defaultWeightMsgRemoveDelegate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveDelegate,
		gamesimulation.SimulateMsgRemoveDelegate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
