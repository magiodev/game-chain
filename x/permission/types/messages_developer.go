package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDeveloper = "create_developer"
	TypeMsgUpdateDeveloper = "update_developer"
)

var _ sdk.Msg = &MsgCreateDeveloper{}

func NewMsgCreateDeveloper(
	creator string,
	address string,
	createdAt int32,
	updatedAt int32,
	blocked bool,

) *MsgCreateDeveloper {
	return &MsgCreateDeveloper{
		Creator:   creator,
		Address:   address,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Blocked:   blocked,
	}
}

func (msg *MsgCreateDeveloper) Route() string {
	return RouterKey
}

func (msg *MsgCreateDeveloper) Type() string {
	return TypeMsgCreateDeveloper
}

func (msg *MsgCreateDeveloper) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDeveloper) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDeveloper) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDeveloper{}

func NewMsgUpdateDeveloper(
	creator string,
	address string,
	createdAt int32,
	updatedAt int32,
	blocked bool,

) *MsgUpdateDeveloper {
	return &MsgUpdateDeveloper{
		Creator:   creator,
		Address:   address,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Blocked:   blocked,
	}
}

func (msg *MsgUpdateDeveloper) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDeveloper) Type() string {
	return TypeMsgUpdateDeveloper
}

func (msg *MsgUpdateDeveloper) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDeveloper) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDeveloper) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
