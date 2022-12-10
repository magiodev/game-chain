package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintNft = "mint_nft"

var _ sdk.Msg = &MsgMintNft{}

func NewMsgMintNft(creator string, project string, symbol string, uri string, uriHash string, data string, receiver string) *MsgMintNft {
	return &MsgMintNft{
		Creator:  creator,
		Project:  project,
		Symbol:   symbol,
		Uri:      uri,
		UriHash:  uriHash,
		Data:     data,
		Receiver: receiver,
	}
}

func (msg *MsgMintNft) Route() string {
	return RouterKey
}

func (msg *MsgMintNft) Type() string {
	return TypeMsgMintNft
}

func (msg *MsgMintNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgUpdateNft = "update_nft"

var _ sdk.Msg = &MsgUpdateNft{}

func NewMsgUpdateNft(creator string, symbol string, id string, uri string, uriHash string, data string) *MsgUpdateNft {
	return &MsgUpdateNft{
		Creator: creator,
		Symbol:  symbol,
		Id:      id,
		Uri:     uri,
		UriHash: uriHash,
		Data:    data,
	}
}

func (msg *MsgUpdateNft) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNft) Type() string {
	return TypeMsgUpdateNft
}

func (msg *MsgUpdateNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgBurnNft = "burn_nft"

var _ sdk.Msg = &MsgBurnNft{}

func NewMsgBurnNft(creator string, symbol string, id string) *MsgBurnNft {
	return &MsgBurnNft{
		Creator: creator,
		Symbol:  symbol,
		Id:      id,
	}
}

func (msg *MsgBurnNft) Route() string {
	return RouterKey
}

func (msg *MsgBurnNft) Type() string {
	return TypeMsgBurnNft
}

func (msg *MsgBurnNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgTransferNft = "transfer_nft"

var _ sdk.Msg = &MsgTransferNft{}

func NewMsgTransferNft(creator string, symbol string, id string, receiver string) *MsgTransferNft {
	return &MsgTransferNft{
		Creator:  creator,
		Symbol:   symbol,
		Id:       id,
		Receiver: receiver,
	}
}

func (msg *MsgTransferNft) Route() string {
	return RouterKey
}

func (msg *MsgTransferNft) Type() string {
	return TypeMsgTransferNft
}

func (msg *MsgTransferNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
