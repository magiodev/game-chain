package keeper

import (
	"context"
	"fmt"

	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAdministrator(goCtx context.Context, msg *types.MsgCreateAdministrator) (*types.MsgCreateAdministratorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := validateCreateAdministrator(ctx, k, msg); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("wrappedError: invalid administrator address (%s)", msg.Creator))
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
		CreatedAt: int32(ctx.BlockHeight()),
		UpdatedAt: int32(ctx.BlockHeight()),
		Blocked:   false,
	}

	k.SetAdministrator(
		ctx,
		administrator,
	)
	return &types.MsgCreateAdministratorResponse{}, nil
}

func (k msgServer) UpdateAdministrator(goCtx context.Context, msg *types.MsgUpdateAdministrator) (*types.MsgUpdateAdministratorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := validateUpdateAdministrator(ctx, k, msg); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("wrappedError: invalid administrator address (%s)", msg.Creator))
	}

	var administrator = types.Administrator{
		//Creator:   msg.Creator,
		//Address:   msg.Address,
		//CreatedAt: msg.CreatedAt,
		UpdatedAt: int32(ctx.BlockHeight()),
		Blocked:   msg.Blocked,
	}

	k.SetAdministrator(ctx, administrator)

	return &types.MsgUpdateAdministratorResponse{}, nil
}

// Private Methods

func validateCreateAdministrator(ctx sdk.Context, k msgServer, msg *types.MsgCreateAdministrator) error {
	// Checking administrator role
	val, found := k.GetAdministrator(ctx, msg.Creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid administrator address (%s)", msg.Creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator administrator address blocked (%s)", msg.Creator)
	}
	return nil
}

func validateUpdateAdministrator(ctx sdk.Context, k msgServer, msg *types.MsgUpdateAdministrator) error {
	// Checking administrator role
	val, found := k.GetAdministrator(ctx, msg.Creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator invalid administrator address (%s)", msg.Creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator administrator address blocked (%s)", msg.Creator)
	}
	// Checking if last admin
	if len(k.GetAllAdministrator(ctx)) < 2 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "deletion of the last administrator is not permitted")
	}
	// Check if the value exists
	toDelete, isFound := k.GetAdministrator(ctx, msg.Address)
	if !isFound {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "to delete index not set")
	}
	// Checks if the msg creator is the same as the toDelete admin owner
	if msg.Creator != toDelete.Address {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "an administrator cannot delete himself")
	}
	return nil
}
