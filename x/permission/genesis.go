package permission

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the administrator
	for _, elem := range genState.AdministratorList {
		k.SetAdministrator(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AdministratorList = k.GetAllAdministrator(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
