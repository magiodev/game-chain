package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestClassMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AssetfactoryKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateClass{Creator: creator,
			Symbol: strconv.Itoa(i),
		}
		_, err := srv.CreateClass(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetClass(ctx,
			expected.Symbol,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestClassMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateClass
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateClass{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateClass{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateClass{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AssetfactoryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateClass{Creator: creator,
				Symbol: strconv.Itoa(0),
			}
			_, err := srv.CreateClass(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateClass(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetClass(ctx,
					expected.Symbol,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestClassMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteClass
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteClass{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteClass{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteClass{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AssetfactoryKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateClass(wctx, &types.MsgCreateClass{Creator: creator,
				Symbol: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteClass(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetClass(ctx,
					tc.request.Symbol,
				)
				require.False(t, found)
			}
		})
	}
}
