package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateClass = "create_class"
	TypeMsgUpdateClass = "update_class"
	TypeMsgDeleteClass = "delete_class"
)

var _ sdk.Msg = &MsgCreateClass{}

func NewMsgCreateClass(
	creator string,
	symbol string,
	project string,
	maxSupply int32,
	canChangeMaxSupply bool,

) *MsgCreateClass {
	return &MsgCreateClass{
		Creator:            creator,
		Symbol:             symbol,
		Project:            project,
		MaxSupply:          maxSupply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func (msg *MsgCreateClass) Route() string {
	return RouterKey
}

func (msg *MsgCreateClass) Type() string {
	return TypeMsgCreateClass
}

func (msg *MsgCreateClass) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateClass{}

func NewMsgUpdateClass(
	creator string,
	symbol string,
	project string,
	maxSupply int32,
	canChangeMaxSupply bool,

) *MsgUpdateClass {
	return &MsgUpdateClass{
		Creator:            creator,
		Symbol:             symbol,
		Project:            project,
		MaxSupply:          maxSupply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func (msg *MsgUpdateClass) Route() string {
	return RouterKey
}

func (msg *MsgUpdateClass) Type() string {
	return TypeMsgUpdateClass
}

func (msg *MsgUpdateClass) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteClass{}

func NewMsgDeleteClass(
	creator string,
	symbol string,

) *MsgDeleteClass {
	return &MsgDeleteClass{
		Creator: creator,
		Symbol:  symbol,
	}
}
func (msg *MsgDeleteClass) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClass) Type() string {
	return TypeMsgDeleteClass
}

func (msg *MsgDeleteClass) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
