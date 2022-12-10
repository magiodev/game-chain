package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ClassAll(c context.Context, req *types.QueryAllClassRequest) (*types.QueryAllClassResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var classs []types.Class
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	classStore := prefix.NewStore(store, types.KeyPrefix(types.ClassKeyPrefix))

	pageRes, err := query.Paginate(classStore, req.Pagination, func(key []byte, value []byte) error {
		var class types.Class
		if err := k.cdc.Unmarshal(value, &class); err != nil {
			return err
		}

		classs = append(classs, class)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClassResponse{Class: classs, Pagination: pageRes}, nil
}

func (k Keeper) Class(c context.Context, req *types.QueryGetClassRequest) (*types.QueryGetClassResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetClass(
		ctx,
		req.Symbol,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetClassResponse{Class: val}, nil
}
