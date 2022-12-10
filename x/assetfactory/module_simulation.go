package assetfactory

import (
	"math/rand"

	"github.com/G4AL-Entertainment/g4al-chain/testutil/sample"
	assetfactorysimulation "github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/simulation"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
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
	_ = assetfactorysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateClass = "op_weight_msg_class"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateClass int = 100

	opWeightMsgUpdateClass = "op_weight_msg_class"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateClass int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	assetfactoryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		ClassList: []types.Class{
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
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&assetfactoryGenesis)
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

	var weightMsgCreateClass int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = defaultWeightMsgCreateClass
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClass,
		assetfactorysimulation.SimulateMsgCreateClass(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateClass int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateClass, &weightMsgUpdateClass, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClass = defaultWeightMsgUpdateClass
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateClass,
		assetfactorysimulation.SimulateMsgUpdateClass(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
