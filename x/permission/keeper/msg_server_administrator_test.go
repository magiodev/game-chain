package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAdministratorMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.PermissionKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateAdministrator{Creator: creator,
			Address: strconv.Itoa(i),
		}
		_, err := srv.CreateAdministrator(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetAdministrator(ctx,
			expected.Address,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestAdministratorMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateAdministrator
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateAdministrator{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateAdministrator{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateAdministrator{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PermissionKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateAdministrator{Creator: creator,
				Address: strconv.Itoa(0),
			}
			_, err := srv.CreateAdministrator(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateAdministrator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetAdministrator(ctx,
					expected.Address,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestAdministratorMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteAdministrator
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteAdministrator{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteAdministrator{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteAdministrator{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PermissionKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateAdministrator(wctx, &types.MsgCreateAdministrator{Creator: creator,
				Address: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteAdministrator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetAdministrator(ctx,
					tc.request.Address,
				)
				require.False(t, found)
			}
		})
	}
}
