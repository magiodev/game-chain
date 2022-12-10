package permission_test

import (
	"testing"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/testutil/nullify"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission"
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PermissionKeeper(t)
	permission.InitGenesis(ctx, *k, genesisState)
	got := permission.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
