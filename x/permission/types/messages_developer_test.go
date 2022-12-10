package types

import (
	"testing"

	"github.com/G4AL-Entertainment/g4al-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateDeveloper_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDeveloper
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDeveloper{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDeveloper{
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

func TestMsgUpdateDeveloper_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDeveloper
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDeveloper{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDeveloper{
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
