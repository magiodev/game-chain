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

	var project = types.Project{
		Creator:     msg.Creator,
		Symbol:      symbol,
		Name:        msg.Name,
		Description: msg.Description,
		CreatedAt:   ctx.BlockHeight(),
		UpdatedAt:   ctx.BlockHeight(),
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

	valFound.Name = msg.Name
	valFound.Description = msg.Description
	valFound.UpdatedAt = ctx.BlockHeight()

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

func (k msgServer) AddDelegate(goCtx context.Context, msg *types.MsgAddDelegate) (*types.MsgAddDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateProjectOwnershipByProject(ctx, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	valFound, isFound := k.GetProject(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "project not found")
	}

	var validDelegate []string
	for _, delegate := range msg.Delegate {
		// Validating that string[] addresses are Bech32 compliant
		bech32, err := sdk.AccAddressFromBech32(delegate)
		if err == nil {
			validDelegate = append(validDelegate, bech32.String())
		}
	}

	valFound.Delegate = Union(valFound.Delegate, validDelegate)
	valFound.UpdatedAt = ctx.BlockHeight()

	k.SetProject(ctx, valFound)

	err = ctx.EventManager().EmitTypedEvent(&types.EventAddDelegateProject{
		Symbol:   msg.Symbol,
		Delegate: validDelegate,
		Creator:  msg.Creator,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgAddDelegateResponse{}, nil
}

func (k msgServer) RemoveDelegate(goCtx context.Context, msg *types.MsgRemoveDelegate) (*types.MsgRemoveDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateProjectOwnershipByProject(ctx, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	valFound, isFound := k.GetProject(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var validDelegate []string
	for _, delegate := range msg.Delegate {
		// Validating that string[] addresses are Bech32 compliant
		bech32, err := sdk.AccAddressFromBech32(delegate)
		if err == nil {
			validDelegate = append(validDelegate, bech32.String())
		}
	}

	valFound.Delegate = Delete(valFound.Delegate, validDelegate)
	valFound.UpdatedAt = ctx.BlockHeight()

	k.SetProject(ctx, valFound)

	err = ctx.EventManager().EmitTypedEvent(&types.EventRemoveDelegateProject{
		Symbol:   msg.Symbol,
		Delegate: validDelegate,
		Creator:  msg.Creator,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveDelegateResponse{}, nil
}

func Union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func Delete(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		m[item] = false
	}
	var c []string

	for key, value := range m {
		if value {
			c = append(c, key)
		}
	}
	return c
}
