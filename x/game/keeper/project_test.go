package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/testutil/nullify"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNProject(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Project {
	items := make([]types.Project, n)
	for i := range items {
		items[i].Symbol = strconv.Itoa(i)

		keeper.SetProject(ctx, items[i])
	}
	return items
}

func TestProjectGet(t *testing.T) {
	keeper, ctx := keepertest.GameKeeper(t)
	items := createNProject(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProject(ctx,
			item.Symbol,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestProjectRemove(t *testing.T) {
	keeper, ctx := keepertest.GameKeeper(t)
	items := createNProject(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProject(ctx,
			item.Symbol,
		)
		_, found := keeper.GetProject(ctx,
			item.Symbol,
		)
		require.False(t, found)
	}
}

func TestProjectGetAll(t *testing.T) {
	keeper, ctx := keepertest.GameKeeper(t)
	items := createNProject(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProject(ctx)),
	)
}
