package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/testutil/nullify"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAdministratorQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAdministrator(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAdministratorRequest
		response *types.QueryGetAdministratorResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAdministratorRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetAdministratorResponse{Administrator: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAdministratorRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetAdministratorResponse{Administrator: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAdministratorRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Administrator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAdministratorQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAdministrator(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAdministratorRequest {
		return &types.QueryAllAdministratorRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AdministratorAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Administrator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Administrator),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AdministratorAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Administrator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Administrator),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AdministratorAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Administrator),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AdministratorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
