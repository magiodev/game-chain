package assetfactory_test

import (
	"testing"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/testutil/nullify"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AssetfactoryKeeper(t)
	assetfactory.InitGenesis(ctx, *k, genesisState)
	got := assetfactory.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
