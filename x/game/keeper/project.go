package keeper

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	SymbolMinLength      = 8
	SymbolMaxLength      = 30
	NameMinLength        = 8
	NameMaxLength        = 30
	DescriptionMinLength = 10
	DescriptionMaxLength = 250
)

// SetProject set a specific project in the store from its index
func (k Keeper) SetProject(ctx sdk.Context, project types.Project) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKeyPrefix))
	b := k.cdc.MustMarshal(&project)
	store.Set(types.ProjectKey(
		project.Symbol,
	), b)
}

// GetProject returns a project from its index
func (k Keeper) GetProject(
	ctx sdk.Context,
	symbol string,

) (val types.Project, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKeyPrefix))

	b := store.Get(types.ProjectKey(
		symbol,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProject removes a project from the store
func (k Keeper) RemoveProject(
	ctx sdk.Context,
	symbol string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKeyPrefix))
	store.Delete(types.ProjectKey(
		symbol,
	))
}

// GetAllProject returns all project
func (k Keeper) GetAllProject(ctx sdk.Context) (list []types.Project) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Project
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
