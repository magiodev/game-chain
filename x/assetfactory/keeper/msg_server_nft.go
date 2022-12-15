package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintNft(goCtx context.Context, msg *types.MsgMintNft) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate
	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	err = k.ValidateProjectOwnershipOrDelegateByClassId(ctx, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}
	// Max supply validation
	err = k.ValidateMaxSupply(ctx, msg.Symbol)
	if err != nil {
		return nil, err
	}

	// Treating msg.Data any value
	//msgData, err := StringToAny(msg.Data)
	//if err != nil {
	//	return nil, err
	//}

	toMint := nft.NFT{
		ClassId: msg.Symbol,
		Id:      string(k.nftKeeper.GetTotalSupply(ctx, msg.Symbol)), // check conversion
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		//Data:    msgData,
	}

	// Validating receiver account as Bech32 address
	err = k.nftKeeper.Mint(ctx, toMint, sdk.AccAddress(msg.Receiver))
	if err != nil {
		return nil, err
	}

	return &types.MsgMintNftResponse{}, nil
}

func (k msgServer) UpdateNft(goCtx context.Context, msg *types.MsgUpdateNft) (*types.MsgUpdateNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.ValidateProjectOwnershipOrDelegateByClassId(ctx, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	toUpdate, found := k.nftKeeper.GetNFT(ctx, msg.Symbol, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "nftID not found (%s)", msg.Id)
	}

	// Treating msg.Data any value
	//msgData, err := utils.StringToAny(msg.Data)
	//if err != nil {
	//	return nil, err
	//}

	toUpdate.Uri = msg.Uri
	toUpdate.UriHash = msg.UriHash
	//toUpdate.Data = msgData

	err = k.nftKeeper.Update(ctx, toUpdate)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateNftResponse{}, nil
}

func (k msgServer) BurnNft(goCtx context.Context, msg *types.MsgBurnNft) (*types.MsgBurnNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.permissionKeeper.ValidateDeveloper(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.ValidateProjectOwnershipOrDelegateByClassId(ctx, msg.Creator, msg.Symbol)
	if err != nil {
		return nil, err
	}

	err = k.nftKeeper.Burn(ctx, msg.Symbol, msg.Id)
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnNftResponse{}, nil
}

func (k msgServer) TransferNft(goCtx context.Context, msg *types.MsgTransferNft) (*types.MsgTransferNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateNftOwnershipOrAllowance(ctx, msg)
	if err != nil {
		return nil, err
	}

	err = k.nftKeeper.Transfer(ctx, msg.Symbol, msg.Id, sdk.AccAddress(msg.Receiver))
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferNftResponse{}, nil
}
