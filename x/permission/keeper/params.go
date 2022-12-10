package keeper

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.GenesisAdministrator(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// GenesisAdministrator returns the GenesisAdministrator param
func (k Keeper) GenesisAdministrator(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyGenesisAdministrator, &res)
	return
}
