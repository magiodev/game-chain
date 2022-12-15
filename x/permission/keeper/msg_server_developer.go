package keeper

import (
	"context"
	"strconv"

	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDeveloper(goCtx context.Context, msg *types.MsgCreateDeveloper) (*types.MsgCreateDeveloperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateAdministrator(ctx, msg.Creator)
	if err != nil {
		return nil, err
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

	// TODO: EVENTS
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeCreateDeveloper,
			sdk.NewAttribute(DeveloperAttribute, msg.Address),
			sdk.NewAttribute(AttributeKeyDeveloperCreator, msg.Creator),
		),
	)

	return &types.MsgCreateDeveloperResponse{}, nil
}

func (k msgServer) UpdateDeveloper(goCtx context.Context, msg *types.MsgUpdateDeveloper) (*types.MsgUpdateDeveloperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateAdministrator(ctx, msg.Creator)
	if err != nil {
		return nil, err
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

	// TODO: EVENTS
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeUpdateDeveloper,
			sdk.NewAttribute(DeveloperAttribute, msg.Address),
			sdk.NewAttribute(DeveloperBlockedAttribute, strconv.FormatBool(msg.Blocked)),
			sdk.NewAttribute(AttributeKeyDeveloperCreator, msg.Creator),
		),
	)

	return &types.MsgUpdateDeveloperResponse{}, nil
}
