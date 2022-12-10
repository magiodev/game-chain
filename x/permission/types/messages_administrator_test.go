package types

import (
	"testing"

	"github.com/G4AL-Entertainment/g4al-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateAdministrator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAdministrator
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAdministrator{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateAdministrator{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateAdministrator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAdministrator
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAdministrator{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAdministrator{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteAdministrator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAdministrator
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteAdministrator{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteAdministrator{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
