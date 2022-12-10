package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateProject = "create_project"
	TypeMsgUpdateProject = "update_project"
	TypeMsgDeleteProject = "delete_project"
)

var _ sdk.Msg = &MsgCreateProject{}

func NewMsgCreateProject(
	creator string,
	symbol string,
	name string,
	description string,
	delegate []string,

) *MsgCreateProject {
	return &MsgCreateProject{
		Creator:     creator,
		Symbol:      symbol,
		Name:        name,
		Description: description,
		Delegate:    delegate,
	}
}

func (msg *MsgCreateProject) Route() string {
	return RouterKey
}

func (msg *MsgCreateProject) Type() string {
	return TypeMsgCreateProject
}

func (msg *MsgCreateProject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProject{}

func NewMsgUpdateProject(
	creator string,
	symbol string,
	name string,
	description string,
	delegate []string,

) *MsgUpdateProject {
	return &MsgUpdateProject{
		Creator:     creator,
		Symbol:      symbol,
		Name:        name,
		Description: description,
		Delegate:    delegate,
	}
}

func (msg *MsgUpdateProject) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProject) Type() string {
	return TypeMsgUpdateProject
}

func (msg *MsgUpdateProject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteProject{}

func NewMsgDeleteProject(
	creator string,
	symbol string,

) *MsgDeleteProject {
	return &MsgDeleteProject{
		Creator: creator,
		Symbol:  symbol,
	}
}
func (msg *MsgDeleteProject) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProject) Type() string {
	return TypeMsgDeleteProject
}

func (msg *MsgDeleteProject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
