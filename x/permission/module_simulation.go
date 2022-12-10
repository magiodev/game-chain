package permission

import (
	"math/rand"

	"github.com/G4AL-Entertainment/g4al-chain/testutil/sample"
	permissionsimulation "github.com/G4AL-Entertainment/g4al-chain/x/permission/simulation"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
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
	_ = permissionsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAdministrator = "op_weight_msg_administrator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAdministrator int = 100

	opWeightMsgUpdateAdministrator = "op_weight_msg_administrator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAdministrator int = 100

	opWeightMsgDeleteAdministrator = "op_weight_msg_administrator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAdministrator int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	permissionGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AdministratorList: []types.Administrator{
			{
				Creator: sample.AccAddress(),
				Address: "0",
			},
			{
				Creator: sample.AccAddress(),
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&permissionGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	permissionParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyGenesisAdministrator), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(permissionParams.GenesisAdministrator))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAdministrator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAdministrator, &weightMsgCreateAdministrator, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAdministrator = defaultWeightMsgCreateAdministrator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAdministrator,
		permissionsimulation.SimulateMsgCreateAdministrator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAdministrator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAdministrator, &weightMsgUpdateAdministrator, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAdministrator = defaultWeightMsgUpdateAdministrator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAdministrator,
		permissionsimulation.SimulateMsgUpdateAdministrator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAdministrator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAdministrator, &weightMsgDeleteAdministrator, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAdministrator = defaultWeightMsgDeleteAdministrator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAdministrator,
		permissionsimulation.SimulateMsgDeleteAdministrator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
