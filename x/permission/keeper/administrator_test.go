package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/testutil/nullify"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAdministrator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Administrator {
	items := make([]types.Administrator, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetAdministrator(ctx, items[i])
	}
	return items
}

func TestAdministratorGet(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNAdministrator(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAdministrator(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAdministratorRemove(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNAdministrator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAdministrator(ctx,
			item.Address,
		)
		_, found := keeper.GetAdministrator(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestAdministratorGetAll(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNAdministrator(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAdministrator(ctx)),
	)
}
