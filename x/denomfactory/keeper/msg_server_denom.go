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

	// Regex first of all as we strip characters
	symbol := utils.RegExSymbol(msg.Symbol)

	// Check if the value already exists in map
	_, isFound := k.GetDenom(
		ctx,
		symbol,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set in map")
	}

	// Check if existing only in bankKeeper state (GGT case)
	_, found := k.bankKeeper.GetDenomMetaData(ctx, "u"+symbol)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index already set in bank")
	}

	var denom = types.Denom{
		Creator:            msg.Creator,
		Symbol:             symbol,
		Project:            msg.Project,
		MaxSupply:          msg.MaxSupply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	// Set metadata for Denom
	k.SetMetadata(ctx, symbol, msg.Name, msg.Description)

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
	err = k.ValidateProjectOwnershipOrDelegateByDenom(ctx, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	// Check if input texts meet requirements
	err = k.ValidateArgsDenom(msg.Symbol, msg.Description, msg.Name)
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

	// Getting existing metadata
	var existingMetadata, found = k.bankKeeper.GetDenomMetaData(ctx, valFound.Symbol)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "metadata for denom not found")
	}

	// If CanChangeMaxSupply we allow editing MaxSupply
	if valFound.CanChangeMaxSupply {
		currentTotalSupply := k.bankKeeper.GetSupply(ctx, valFound.Symbol)
		if currentTotalSupply.Amount.Uint64() > msg.MaxSupply {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "total supply already exceeds the new maxSupply")
		}
		valFound.MaxSupply = msg.MaxSupply
	}

	// Set metadata for Denom
	k.SetMetadata(ctx, existingMetadata.Symbol, msg.Name, msg.Description)

	k.SetDenom(ctx, valFound)

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
	err = k.ValidateProjectOwnershipOrDelegateByDenom(ctx, msg.Creator, denom.Symbol)
	if err != nil {
		return nil, err
	}

	// Minting coins with x/bank module
	coin := sdk.NewCoin(denom.Symbol, math.NewInt(int64(msg.Amount)))
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
