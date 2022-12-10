package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestProjectMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.GameKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateProject{Creator: creator,
			Symbol: strconv.Itoa(i),
		}
		_, err := srv.CreateProject(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetProject(ctx,
			expected.Symbol,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestProjectMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateProject
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateProject{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateProject{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateProject{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.GameKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateProject{Creator: creator,
				Symbol: strconv.Itoa(0),
			}
			_, err := srv.CreateProject(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateProject(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetProject(ctx,
					expected.Symbol,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestProjectMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteProject
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteProject{Creator: creator,
				Symbol: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteProject{Creator: "B",
				Symbol: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteProject{Creator: creator,
				Symbol: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.GameKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateProject(wctx, &types.MsgCreateProject{Creator: creator,
				Symbol: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteProject(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetProject(ctx,
					tc.request.Symbol,
				)
				require.False(t, found)
			}
		})
	}
}
