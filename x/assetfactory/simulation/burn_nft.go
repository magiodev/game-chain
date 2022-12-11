package simulation

import (
	"math/rand"

	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/keeper"
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgBurnNft(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBurnNft{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the BurnNft simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "BurnNft simulation not implemented"), nil, nil
	}
}
