package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/utils"

	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	err = k.ValidateArgsProject(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	// Regex first of all as we strip characters
	symbol := utils.RegExSymbol(msg.Symbol)

	// Check if the value already exists
	_, isFound := k.GetProject(
		ctx,
		symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Validating that string[] addresses are Bech32 compliant
	err = k.ValidateDelegateIsValid(msg.Delegate)
	if err != nil {
		return nil, err
	}

	var project = types.Project{
		Creator:     msg.Creator,
		Symbol:      symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Delegate:    msg.Delegate,
		CreatedAt:   int32(ctx.BlockHeight()),
		UpdatedAt:   int32(ctx.BlockHeight()),
	}

	k.SetProject(
		ctx,
		project,
	)

	err = ctx.EventManager().EmitTypedEvent(&types.EventCreateProject{
		Symbol:      msg.Symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Creator:     msg.Creator,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateProjectResponse{}, nil
}

func (k msgServer) UpdateProject(goCtx context.Context, msg *types.MsgUpdateProject) (*types.MsgUpdateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate only ownership and delegate, not developer role
	err := k.ValidateProjectOwnershipOrDelegateByProject(ctx, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	err = k.ValidateArgsProject(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetProject(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// appending new delegate values to valFound.Delegate
	for _, delegate := range msg.Delegate {
		valFound.Delegate = append(valFound.Delegate, sdk.AccAddress(delegate).String())
	}

	valFound.Name = msg.Name
	valFound.Description = msg.Description
	valFound.UpdatedAt = int32(ctx.BlockHeight())

	k.SetProject(ctx, valFound)

	err = ctx.EventManager().EmitTypedEvent(&types.EventCreateProject{
		Symbol:      msg.Symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Creator:     msg.Creator,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateProjectResponse{}, nil
}
