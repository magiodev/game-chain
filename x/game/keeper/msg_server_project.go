package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetProject(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var project = types.Project{
		Creator:     msg.Creator,
		Symbol:      msg.Symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Delegate:    msg.Delegate,
	}

	k.SetProject(
		ctx,
		project,
	)
	return &types.MsgCreateProjectResponse{}, nil
}

func (k msgServer) UpdateProject(goCtx context.Context, msg *types.MsgUpdateProject) (*types.MsgUpdateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProject(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var project = types.Project{
		Creator:     msg.Creator,
		Symbol:      msg.Symbol,
		Name:        msg.Name,
		Description: msg.Description,
		Delegate:    msg.Delegate,
	}

	k.SetProject(ctx, project)

	return &types.MsgUpdateProjectResponse{}, nil
}

func (k msgServer) DeleteProject(goCtx context.Context, msg *types.MsgDeleteProject) (*types.MsgDeleteProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetProject(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveProject(
		ctx,
		msg.Symbol,
	)

	return &types.MsgDeleteProjectResponse{}, nil
}
