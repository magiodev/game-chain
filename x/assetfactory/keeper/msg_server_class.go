package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateClass(goCtx context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.gameKeeper.ValidateProjectOwnershipOrDelegateByProject(ctx, msg.Creator, msg.Project)
	if err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetClass(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Check if input texts meet requirements
	err = k.ValidateArgsClass(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	var class = types.Class{
		Creator:            msg.Creator,
		Symbol:             msg.Symbol,
		Project:            msg.Project,
		MaxSupply:          msg.MaxSupply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	// Set the x/nft class
	err = k.SaveClass(ctx, msg.Symbol, msg.Name, msg.Description, msg.Uri, msg.UriHash, msg.Data)
	if err != nil {
		return nil, err
	}

	// Set this module's map, not x/nft
	k.SetClass(ctx, class)

	return &types.MsgCreateClassResponse{}, nil
}

func (k msgServer) UpdateClass(goCtx context.Context, msg *types.MsgUpdateClass) (*types.MsgUpdateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate developer owner or delegate
	err := k.ValidateProjectOwnershipOrDelegateByClassId(ctx, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	// Validate input args
	err = k.ValidateArgsClass(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetClass(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// If CanChangeMaxSupply we allow editing MaxSupply
	if valFound.CanChangeMaxSupply {
		currentTotalSupply := k.nftKeeper.GetTotalSupply(ctx, valFound.Symbol)
		if currentTotalSupply > msg.MaxSupply {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "total supply already exceeds the new maxSupply")
		}
		valFound.MaxSupply = msg.MaxSupply
	}

	// Set the x/nft class
	err = k.SaveClass(ctx, valFound.Symbol, msg.Name, msg.Description, msg.Uri, msg.UriHash, msg.Data)
	if err != nil {
		return nil, err
	}

	// Set this module's map, not x/nft
	k.SetClass(ctx, valFound)

	return &types.MsgUpdateClassResponse{}, nil
}
