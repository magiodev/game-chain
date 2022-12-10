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

func createNDeveloper(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Developer {
	items := make([]types.Developer, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetDeveloper(ctx, items[i])
	}
	return items
}

func TestDeveloperGet(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNDeveloper(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDeveloper(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeveloperRemove(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNDeveloper(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeveloper(ctx,
			item.Address,
		)
		_, found := keeper.GetDeveloper(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestDeveloperGetAll(t *testing.T) {
	keeper, ctx := keepertest.PermissionKeeper(t)
	items := createNDeveloper(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDeveloper(ctx)),
	)
}
