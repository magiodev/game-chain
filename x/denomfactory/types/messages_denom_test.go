package types

import (
	"testing"

	"github.com/G4AL-Entertainment/g4al-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDenom{
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

func TestMsgUpdateDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDenom{
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

func TestMsgMintDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMintDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMintDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMintDenom{
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

func TestMsgBurnDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgBurnDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgBurnDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgBurnDenom{
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

func TestMsgTransferDenom_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgTransferDenom
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgTransferDenom{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgTransferDenom{
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
