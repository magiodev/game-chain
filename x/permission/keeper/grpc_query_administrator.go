package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AdministratorAll(c context.Context, req *types.QueryAllAdministratorRequest) (*types.QueryAllAdministratorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var administrators []types.Administrator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	administratorStore := prefix.NewStore(store, types.KeyPrefix(types.AdministratorKeyPrefix))

	pageRes, err := query.Paginate(administratorStore, req.Pagination, func(key []byte, value []byte) error {
		var administrator types.Administrator
		if err := k.cdc.Unmarshal(value, &administrator); err != nil {
			return err
		}

		administrators = append(administrators, administrator)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAdministratorResponse{Administrator: administrators, Pagination: pageRes}, nil
}

func (k Keeper) Administrator(c context.Context, req *types.QueryGetAdministratorRequest) (*types.QueryGetAdministratorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAdministrator(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAdministratorResponse{Administrator: val}, nil
}
