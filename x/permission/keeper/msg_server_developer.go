package keeper

import (
	"context"
	"fmt"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDeveloper(goCtx context.Context, msg *types.MsgCreateDeveloper) (*types.MsgCreateDeveloperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := validateCreateDeveloper(ctx, k, msg); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("wrappedError: invalid administrator address (%s)", msg.Creator))
	}

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
		CreatedAt: int32(ctx.BlockHeight()),
		UpdatedAt: int32(ctx.BlockHeight()),
		Blocked:   false,
	}

	k.SetDeveloper(
		ctx,
		developer,
	)
	return &types.MsgCreateDeveloperResponse{}, nil
}

func (k msgServer) UpdateDeveloper(goCtx context.Context, msg *types.MsgUpdateDeveloper) (*types.MsgUpdateDeveloperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := validateUpdateDeveloper(ctx, k, msg); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("wrappedError: invalid administrator address (%s)", msg.Creator))
	}

	// Check if the value exists
	dev, isFound := k.GetDeveloper(
		ctx,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var developer = types.Developer{
		Creator:   dev.Creator,
		Address:   dev.Address,
		CreatedAt: dev.CreatedAt,
		UpdatedAt: int32(ctx.BlockHeight()),
		Blocked:   msg.Blocked,
	}

	k.SetDeveloper(ctx, developer)

	return &types.MsgUpdateDeveloperResponse{}, nil
}

// Private Methods

func validateUpdateDeveloper(ctx sdk.Context, k msgServer, msg *types.MsgUpdateDeveloper) error {
	// Checking administrator role
	val, found := k.GetAdministrator(ctx, msg.Creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid administrator address (%s)", msg.Creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "administrator address blocked (%s)", msg.Creator)
	}
	return nil
}

func validateCreateDeveloper(ctx sdk.Context, k msgServer, msg *types.MsgCreateDeveloper) error {
	// Checking administrator role
	val, found := k.GetAdministrator(ctx, msg.Creator)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid administrator address (%s)", msg.Creator)
	}
	if val.Blocked {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "administrator address blocked (%s)", msg.Creator)
	}
	return nil
}
