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
		Delegate:    msg.Delegate,
		CreatedAt:   int32(ctx.BlockHeight()),
		UpdatedAt:   int32(ctx.BlockHeight()),
	}

	k.SetProject(
		ctx,
		project,
	)

	// TODO: EVENTS
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeCreateProject,
			sdk.NewAttribute(AttributeKeyProjectSymbol, msg.Symbol),
			sdk.NewAttribute(AttributeKeyProjectName, msg.Name),
			sdk.NewAttribute(AttributeKeyProjectDescription, msg.Description),
			sdk.NewAttribute(AttributeKeyProjectCreator, msg.Creator),
		),
	)

	return &types.MsgCreateProjectResponse{}, nil
}

func (k msgServer) UpdateProject(goCtx context.Context, msg *types.MsgUpdateProject) (*types.MsgUpdateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Skipping role policy validation as we use Ownership or Delegate array.string

	err := k.ValidateArgsProject(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetProject(
		ctx,
		utils.RegExSymbol(msg.Symbol), // regEx applied,  TODO consider that maybe here to update is not required, as we do not allow changing symbol
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Check if msg.Creator included in valFound.Delegate
	isDelegate := false
	for _, del := range valFound.Delegate {
		if del == msg.Creator {
			isDelegate = true
			break
		}
	}
	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator && !isDelegate {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner nor delegate address")
	}

	// TODO consider delegates removal also
	// appending new delegate values to valFound.Delegate
	for _, delegate := range msg.Delegate {
		bech32, err := sdk.AccAddressFromBech32(delegate)
		if err == nil { // TODO fix this, is not preventing invalid addresses to be appended
			valFound.Delegate = append(valFound.Delegate, bech32.String())
		}
	}

	valFound.Name = msg.Name
	valFound.Description = msg.Description
	valFound.UpdatedAt = int32(ctx.BlockHeight())

	k.SetProject(ctx, valFound)

	// TODO: EVENTS
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			EventTypeUpdateProject,
			sdk.NewAttribute(AttributeKeyProjectSymbol, msg.Symbol),
			sdk.NewAttribute(AttributeKeyProjectName, msg.Name),
			sdk.NewAttribute(AttributeKeyProjectDescription, msg.Description),
			sdk.NewAttribute(AttributeKeyProjectCreator, msg.Creator),
		),
	)

	return &types.MsgUpdateProjectResponse{}, nil
}
