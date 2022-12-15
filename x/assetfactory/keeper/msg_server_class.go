package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
)

func (k msgServer) CreateClass(goCtx context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.ValidateProjectOwnershipOrDelegateByProject(ctx, msg.Creator, msg.Project)
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

	// Treating msg.Data any value
	//msgData, err := StringToAny(msg.Data)
	//if err != nil {
	//	return nil, err
	//}

	var nftClass = nft.Class{
		Id:          msg.Symbol,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Description: msg.Description,
		Uri:         msg.Uri,
		UriHash:     msg.UriHash,
		//Data:        msgData,
	}
	err = k.nftKeeper.SaveClass(ctx, nftClass)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "class creation has not occurred")
	}

	k.SetClass(
		ctx,
		class,
	)
	return &types.MsgCreateClassResponse{}, nil
}

func (k msgServer) UpdateClass(goCtx context.Context, msg *types.MsgUpdateClass) (*types.MsgUpdateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetClass(
		ctx,
		msg.Symbol,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	err := k.ValidateArgsClass(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	var class = types.Class{
		Creator:   msg.Creator,
		Symbol:    msg.Symbol,
		MaxSupply: msg.MaxSupply,
	}

	//TODO metadata workflow

	k.SetClass(ctx, class)

	return &types.MsgUpdateClassResponse{}, nil
}
