package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAdministrator(goCtx context.Context, msg *types.MsgCreateAdministrator) (*types.MsgCreateAdministratorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetAdministrator(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var administrator = types.Administrator{
		Creator:   msg.Creator,
		Address:   msg.Address,
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
		Blocked:   msg.Blocked,
	}

	k.SetAdministrator(
		ctx,
		administrator,
	)
	return &types.MsgCreateAdministratorResponse{}, nil
}

func (k msgServer) UpdateAdministrator(goCtx context.Context, msg *types.MsgUpdateAdministrator) (*types.MsgUpdateAdministratorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAdministrator(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var administrator = types.Administrator{
		Creator:   msg.Creator,
		Address:   msg.Address,
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
		Blocked:   msg.Blocked,
	}

	k.SetAdministrator(ctx, administrator)

	return &types.MsgUpdateAdministratorResponse{}, nil
}

func (k msgServer) DeleteAdministrator(goCtx context.Context, msg *types.MsgDeleteAdministrator) (*types.MsgDeleteAdministratorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAdministrator(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAdministrator(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteAdministratorResponse{}, nil
}
