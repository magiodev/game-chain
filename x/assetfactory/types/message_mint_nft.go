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
