package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDenomMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.DenomfactoryKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDenom{Creator: creator,
			Symbol: strconv.Itoa(i),
		}
		_, err := srv.CreateDenom(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDenom(ctx,
			expected.Symbol,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestDenomMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDenom
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDenom{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDenom{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDenom{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.DenomfactoryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateDenom{Creator: creator,
				Symbol: strconv.Itoa(0),
			}
			_, err := srv.CreateDenom(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDenom(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDenom(ctx,
					expected.Symbol,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}
