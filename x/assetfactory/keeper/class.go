package keeper

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
)

const (
	SymbolMinLength      = 3
	SymbolMaxLength      = 20
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

// SaveClass saves x/nft class
func (k msgServer) SaveClass(ctx sdk.Context, symbol string, name string, description string, uri string, uriHash string, data string) error {
	// Treating msg.Data any value
	//msgData, err := StringToAny(msg.Data)
	//if err != nil {
	//	return nil, err
	//}

	var nftClass = nft.Class{
		Id:          symbol,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         uri,
		UriHash:     uriHash,
		//Data:        msgData,
	}
	err := k.nftKeeper.SaveClass(ctx, nftClass)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "class creation has not occurred")
	}
	return nil
}
