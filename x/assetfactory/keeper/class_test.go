package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/testutil/nullify"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNClass(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Class {
	items := make([]types.Class, n)
	for i := range items {
		items[i].Symbol = strconv.Itoa(i)

		keeper.SetClass(ctx, items[i])
	}
	return items
}

func TestClassGet(t *testing.T) {
	keeper, ctx := keepertest.AssetfactoryKeeper(t)
	items := createNClass(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClass(ctx,
			item.Symbol,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClassRemove(t *testing.T) {
	keeper, ctx := keepertest.AssetfactoryKeeper(t)
	items := createNClass(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClass(ctx,
			item.Symbol,
		)
		_, found := keeper.GetClass(ctx,
			item.Symbol,
		)
		require.False(t, found)
	}
}

func TestClassGetAll(t *testing.T) {
	keeper, ctx := keepertest.AssetfactoryKeeper(t)
	items := createNClass(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClass(ctx)),
	)
}
