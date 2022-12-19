package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateProject = "create_project"
	TypeMsgUpdateProject = "update_project"
)

var _ sdk.Msg = &MsgCreateProject{}

func NewMsgCreateProject(
	creator string,
	symbol string,
	name string,
	description string,

) *MsgCreateProject {
	return &MsgCreateProject{
		Creator:     creator,
		Symbol:      symbol,
		Name:        name,
		Description: description,
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
) *MsgUpdateProject {
	return &MsgUpdateProject{
		Creator:     creator,
		Symbol:      symbol,
		Name:        name,
		Description: description,
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

const TypeMsgAddDelegate = "add_delegate"

var _ sdk.Msg = &MsgAddDelegate{}

func NewMsgAddDelegate(creator string, symbol string, delegate []string) *MsgAddDelegate {
	return &MsgAddDelegate{
		Creator:  creator,
		Symbol:   symbol,
		Delegate: delegate,
	}
}

func (msg *MsgAddDelegate) Route() string {
	return RouterKey
}

func (msg *MsgAddDelegate) Type() string {
	return TypeMsgAddDelegate
}

func (msg *MsgAddDelegate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddDelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgRemoveDelegate = "remove_delegate"

var _ sdk.Msg = &MsgRemoveDelegate{}

func NewMsgRemoveDelegate(creator string, symbol string, delegate []string) *MsgRemoveDelegate {
	return &MsgRemoveDelegate{
		Creator:  creator,
		Symbol:   symbol,
		Delegate: delegate,
	}
}

func (msg *MsgRemoveDelegate) Route() string {
	return RouterKey
}

func (msg *MsgRemoveDelegate) Type() string {
	return TypeMsgRemoveDelegate
}

func (msg *MsgRemoveDelegate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveDelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
