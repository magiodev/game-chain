package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintNft(goCtx context.Context, msg *types.MsgMintNft) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateMintNft()
	if err != nil {
		return nil, err
	}

	bech32, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	toMint := nft.NFT{
		ClassId: msg.Symbol,
		Id:      string(k.nftKeeper.GetTotalSupply(ctx, msg.Symbol)), // check conversion
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		//Data: msg.Data, // TODO
	}

	err = k.nftKeeper.Mint(ctx, toMint, bech32)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintNftResponse{}, nil
}

func (k msgServer) UpdateNft(goCtx context.Context, msg *types.MsgUpdateNft) (*types.MsgUpdateNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateUpdateNft()
	if err != nil {
		return nil, err
	}

	toUpdate := nft.NFT{
		ClassId: msg.Symbol,
		Id:      "TODO", // TODO
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		//Data: msg.Data, // TODO
	}

	err = k.nftKeeper.Update(ctx, toUpdate)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateNftResponse{}, nil
}

func (k msgServer) BurnNft(goCtx context.Context, msg *types.MsgBurnNft) (*types.MsgBurnNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := validateBurnNft()
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

	err := validateTransferNft()
	if err != nil {
		return nil, err
	}

	bech32, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	err = k.nftKeeper.Transfer(ctx, msg.Symbol, msg.Id, bech32)
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferNftResponse{}, nil
}

func validateMintNft() error {
	// TODO validate is developer
	// TODO validate is owner or delegate
	// TODO validate maxSupply (consider 0 as unlimited)
	return nil // TODO
}

func validateBurnNft() error {
	// TODO validate is developer
	// TODO validate is owner or delegate
	return nil // TODO
}

func validateUpdateNft() error {
	// TODO validate is developer
	// TODO validate is owner or delegate
	// TODO validate data and params
	return nil // TODO
}

func validateTransferNft() error {
	// TODO validate is owner
	// TODO validate allowance (see x/authz if can help with nfts?) (or maybe skip ownership if developer. but not delegate?)
	return nil // TODO
}
