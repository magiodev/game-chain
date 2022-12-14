package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDenom = "create_denom"
	TypeMsgUpdateDenom = "update_denom"
)

var _ sdk.Msg = &MsgCreateDenom{}

func NewMsgCreateDenom(
	creator string,
	symbol string,
	project string,
	maxSupply uint64,
	canChangeMaxSupply bool,
	name string,
	description string,
	precision uint32,
	uri string,
	uri_hash string,

) *MsgCreateDenom {
	return &MsgCreateDenom{
		Creator:            creator,
		Symbol:             symbol,
		Project:            project,
		MaxSupply:          maxSupply,
		CanChangeMaxSupply: canChangeMaxSupply,
		Name:               name,
		Description:        description,
		Precision:          precision,
		Uri:                uri,
		UriHash:            uri_hash,
	}
}

func (msg *MsgCreateDenom) Route() string {
	return RouterKey
}

func (msg *MsgCreateDenom) Type() string {
	return TypeMsgCreateDenom
}

func (msg *MsgCreateDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDenom{}

func NewMsgUpdateDenom(
	creator string,
	symbol string,
	project string,
	maxSupply uint64,
	name string,
	description string,
	uri string,
	uri_hash string,

) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Creator:     creator,
		Symbol:      symbol,
		Project:     project,
		MaxSupply:   maxSupply,
		Name:        name,
		Description: description,
		Uri:         uri,
		UriHash:     uri_hash,
	}
}

func (msg *MsgUpdateDenom) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDenom) Type() string {
	return TypeMsgUpdateDenom
}

func (msg *MsgUpdateDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
