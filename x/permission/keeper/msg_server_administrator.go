package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAdministrator(goCtx context.Context, msg *types.MsgCreateAdministrator) (*types.MsgCreateAdministratorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateAdministrator(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetAdministrator(
		ctx,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "to create index already set")
	}

	var administrator = types.Administrator{
		Creator:   msg.Creator,
		Address:   msg.Address,
		CreatedAt: ctx.BlockHeight(),
		UpdatedAt: ctx.BlockHeight(),
		Blocked:   false,
	}

	k.SetAdministrator(
		ctx,
		administrator,
	)

	err = ctx.EventManager().EmitTypedEvent(&types.EventCreateAdministrator{
		Address: msg.Address,
		Creator: msg.Creator,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateAdministratorResponse{}, nil
}

func (k msgServer) UpdateAdministrator(goCtx context.Context, msg *types.MsgUpdateAdministrator) (*types.MsgUpdateAdministratorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateAdministrator(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.ValidateUpdateAdministrator(ctx, msg.Creator, msg.Address)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	admin, isFound := k.GetAdministrator(ctx, msg.Address)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "to delete index not set")
	}
	// Checks if the msg creator is the same as the toDelete admin owner

	var administrator = types.Administrator{
		Creator:   admin.Creator,
		Address:   admin.Address,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: ctx.BlockHeight(),
		Blocked:   msg.Blocked,
	}

	k.SetAdministrator(ctx, administrator)

	err = ctx.EventManager().EmitTypedEvent(&types.EventUpdateAdministrator{
		Address: msg.Address,
		Blocked: msg.Blocked,
		Creator: msg.Creator,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateAdministratorResponse{}, nil
}
