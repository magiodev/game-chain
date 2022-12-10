package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDeveloper(goCtx context.Context, msg *types.MsgCreateDeveloper) (*types.MsgCreateDeveloperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDeveloper(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var developer = types.Developer{
		Creator:   msg.Creator,
		Address:   msg.Address,
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
		Blocked:   msg.Blocked,
	}

	k.SetDeveloper(
		ctx,
		developer,
	)
	return &types.MsgCreateDeveloperResponse{}, nil
}

func (k msgServer) UpdateDeveloper(goCtx context.Context, msg *types.MsgUpdateDeveloper) (*types.MsgUpdateDeveloperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDeveloper(
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

	var developer = types.Developer{
		Creator:   msg.Creator,
		Address:   msg.Address,
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
		Blocked:   msg.Blocked,
	}

	k.SetDeveloper(ctx, developer)

	return &types.MsgUpdateDeveloperResponse{}, nil
}

func (k msgServer) DeleteDeveloper(goCtx context.Context, msg *types.MsgDeleteDeveloper) (*types.MsgDeleteDeveloperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDeveloper(
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

	k.RemoveDeveloper(
		ctx,
		msg.Address,
	)

	return &types.MsgDeleteDeveloperResponse{}, nil
}
