package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAdministrator = "create_administrator"
	TypeMsgUpdateAdministrator = "update_administrator"
)

var _ sdk.Msg = &MsgCreateAdministrator{}

func NewMsgCreateAdministrator(
	creator string,
	address string,

) *MsgCreateAdministrator {
	return &MsgCreateAdministrator{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateAdministrator) Route() string {
	return RouterKey
}

func (msg *MsgCreateAdministrator) Type() string {
	return TypeMsgCreateAdministrator
}

func (msg *MsgCreateAdministrator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAdministrator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAdministrator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAdministrator{}

func NewMsgUpdateAdministrator(
	creator string,
	address string,
	blocked bool,

) *MsgUpdateAdministrator {
	return &MsgUpdateAdministrator{
		Creator: creator,
		Address: address,
		Blocked: blocked,
	}
}

func (msg *MsgUpdateAdministrator) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAdministrator) Type() string {
	return TypeMsgUpdateAdministrator
}

func (msg *MsgUpdateAdministrator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAdministrator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAdministrator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
