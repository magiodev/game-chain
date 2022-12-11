package game_test

import (
	"testing"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/testutil/nullify"
	"github.com/G4AL-Entertainment/g4al-chain/x/game"
	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		ProjectList: []types.Project{
			{
				Symbol: "0",
			},
			{
				Symbol: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GameKeeper(t)
	game.InitGenesis(ctx, *k, genesisState)
	got := game.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.ProjectList, got.ProjectList)
	// this line is used by starport scaffolding # genesis/test/assert
}
