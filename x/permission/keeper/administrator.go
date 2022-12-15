package keeper

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	EventTypeCreateAdministrator     = "Create"
	EventTypeUpdateAdministrator     = "Update"
	AdministratorAttribute           = "administrator"
	AdministratorBlockedAttribute    = "blocked"
	AttributeKeyAdministratorCreator = "creator"
)

// SetAdministrator set a specific administrator in the store from its index
func (k Keeper) SetAdministrator(ctx sdk.Context, administrator types.Administrator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdministratorKeyPrefix))
	b := k.cdc.MustMarshal(&administrator)
	store.Set(types.AdministratorKey(
		administrator.Address,
	), b)
}

// GetAdministrator returns a administrator from its index
func (k Keeper) GetAdministrator(
	ctx sdk.Context,
	address string,

) (val types.Administrator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdministratorKeyPrefix))

	b := store.Get(types.AdministratorKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAdministrator removes a administrator from the store
func (k Keeper) RemoveAdministrator(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdministratorKeyPrefix))
	store.Delete(types.AdministratorKey(
		address,
	))
}

// GetAllAdministrator returns all administrator
func (k Keeper) GetAllAdministrator(ctx sdk.Context) (list []types.Administrator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdministratorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Administrator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
