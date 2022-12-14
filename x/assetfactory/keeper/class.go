package keeper

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
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

// SetClass set a specific class in the store from its index
func (k Keeper) SetClass(ctx sdk.Context, class types.Class) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))
	b := k.cdc.MustMarshal(&class)
	store.Set(types.ClassKey(
		class.Symbol,
	), b)
}

// GetClass returns a class from its index
func (k Keeper) GetClass(
	ctx sdk.Context,
	symbol string,

) (val types.Class, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))

	b := store.Get(types.ClassKey(
		symbol,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveClass removes a class from the store
func (k Keeper) RemoveClass(
	ctx sdk.Context,
	symbol string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))
	store.Delete(types.ClassKey(
		symbol,
	))
}

// GetAllClass returns all class
func (k Keeper) GetAllClass(ctx sdk.Context) (list []types.Class) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClassKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Class
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
