package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/G4AL-Entertainment/g4al-chain/testutil/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AssetfactoryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
