package keeper

import (
	"context"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintNft(goCtx context.Context, msg *types.MsgMintNft) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMintNftResponse{}, nil
}

func (k msgServer) UpdateNft(goCtx context.Context, msg *types.MsgUpdateNft) (*types.MsgUpdateNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateNftResponse{}, nil
}

func (k msgServer) BurnNft(goCtx context.Context, msg *types.MsgBurnNft) (*types.MsgBurnNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBurnNftResponse{}, nil
}

func (k msgServer) TransferNft(goCtx context.Context, msg *types.MsgTransferNft) (*types.MsgTransferNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTransferNftResponse{}, nil
}
