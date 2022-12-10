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
	// Set all the developer
	for _, elem := range genState.DeveloperList {
		k.SetDeveloper(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init

	// TODO consider if this is good practice, maybe using genState.AdministratorList is the correct thing to do?
	if len(k.GetAllAdministrator(ctx)) == 0 {
		bech32, err := sdk.AccAddressFromBech32(genState.Params.GetGenesisAdministrator())
		if err != nil {
			return
		}

		k.SetAdministrator(ctx, types.Administrator{
			Address:   bech32.String(),
			CreatedAt: int32(ctx.BlockHeight()), // TODO check if int32 instead64 is fine, less digits!
			UpdatedAt: int32(ctx.BlockHeight()), // TODO check if int32 instead64 is fine, less digits!
			Blocked:   false,
			Creator:   bech32.String(),
		})
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AdministratorList = k.GetAllAdministrator(ctx)
	genesis.DeveloperList = k.GetAllDeveloper(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
