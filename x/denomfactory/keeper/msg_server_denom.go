package keeper

import (
	"context"
	"cosmossdk.io/math"
	"github.com/G4AL-Entertainment/g4al-chain/utils"
	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate
	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.ValidateProjectOwnershipOrDelegateByProject(ctx, msg.Creator, msg.Project)
	if err != nil {
		return nil, err
	}

	// Check if input texts meet requirements
	err = k.ValidateArgsDenom(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// Regex first of all as we strip characters
	symbol := utils.RegExSymbol(msg.Symbol)

	var denom = types.Denom{
		Creator:            msg.Creator,
		Symbol:             symbol,
		Project:            msg.Project,
		MaxSupply:          msg.MaxSupply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	// Set metadata for Denom
	k.SetCoinMetadata(ctx, symbol, msg.Name, msg.Description)

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate
	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.ValidateProjectOwnershipOrDelegateByProject(ctx, msg.Creator, msg.Project)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	valFound, isFound := k.GetDenom(
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

	// Allowing update maxSupply only if previously declared
	if valFound.CanChangeMaxSupply {
		valFound.MaxSupply = msg.MaxSupply
	}
	var denom = types.Denom{
		Creator:            valFound.Creator,
		Symbol:             valFound.Symbol,
		Project:            valFound.Project,
		MaxSupply:          valFound.MaxSupply,
		CanChangeMaxSupply: valFound.CanChangeMaxSupply,
	}

	// Getting existing metadata
	var existingMetadata, found = k.bankKeeper.GetDenomMetaData(ctx, valFound.Symbol)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "metadata for denom not found")
	}

	// Check if input texts meet requirements
	err = k.ValidateArgsDenom(msg.Symbol, msg.Description, msg.Name)
	if err != nil {
		return nil, err
	}

	// Set metadata for Denom
	k.SetCoinMetadata(ctx, existingMetadata.Symbol, msg.Name, msg.Description)

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}

func (k msgServer) MintDenom(goCtx context.Context, msg *types.MsgMintDenom) (*types.MsgMintDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate
	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	denom, found := k.GetDenom(ctx, msg.Symbol)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "denom not find")
	}

	err = k.ValidateProjectOwnershipOrDelegateByProject(ctx, msg.Creator, denom.Project)
	if err != nil {
		return nil, err
	}

	// Minting coins with x/bank module
	coin := sdk.NewCoin(msg.Symbol, math.NewInt(int64(msg.Amount)))
	err = k.bankKeeper.MintCoins(ctx, "denomfactory", sdk.NewCoins(coin))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "mint has not occurred")
	}
	// Sending minted amount from module to receiver
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "denomfactory", sdk.AccAddress(msg.Receiver), sdk.NewCoins(coin))
	if err != nil {
		return nil, err
	}
	return &types.MsgMintDenomResponse{}, nil
}

func (k msgServer) BurnDenom(goCtx context.Context, msg *types.MsgBurnDenom) (*types.MsgBurnDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Sending minted amount from module to receiver
	coin := sdk.NewCoin(msg.Symbol, math.NewInt(int64(msg.Amount)))
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), "denomfactory", sdk.NewCoins(coin))
	if err != nil {
		return nil, err
	}
	// Burning transferred amount from module
	err = k.bankKeeper.BurnCoins(ctx, "denomfactory", sdk.NewCoins())
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnDenomResponse{}, nil
}

func (k msgServer) TransferDenom(goCtx context.Context, msg *types.MsgTransferDenom) (*types.MsgTransferDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Transferring coins
	coin := sdk.NewCoin(msg.Symbol, math.NewInt(int64(msg.Amount)))
	err := k.bankKeeper.SendCoins(ctx, sdk.AccAddress(msg.Creator), sdk.AccAddress(msg.Receiver), sdk.NewCoins(coin))
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferDenomResponse{}, nil
}
