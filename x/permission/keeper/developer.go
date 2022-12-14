package keeper

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetDeveloper set a specific developer in the store from its index
func (k Keeper) SetDeveloper(ctx sdk.Context, developer types.Developer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeveloperKeyPrefix))
	b := k.cdc.MustMarshal(&developer)
	store.Set(types.DeveloperKey(
		developer.Address,
	), b)
}

// GetDeveloper returns a developer from its index
func (k Keeper) GetDeveloper(
	ctx sdk.Context,
	address string,

) (val types.Developer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeveloperKeyPrefix))

	b := store.Get(types.DeveloperKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDeveloper removes a developer from the store
func (k Keeper) RemoveDeveloper(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeveloperKeyPrefix))
	store.Delete(types.DeveloperKey(
		address,
	))
}

// GetAllDeveloper returns all developer
func (k Keeper) GetAllDeveloper(ctx sdk.Context) (list []types.Developer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeveloperKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Developer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ValidateDeveloper(ctx sdk.Context, creator string) error {
	// Checking developer role
	val, found := k.GetDeveloper(ctx, creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid developer address (%s)", creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator developer address blocked (%s)", creator)
	}
	return nil
}
