package mail

import (
	"github.com/CudoVentures/cudos-node/x/mail/keeper"
	"github.com/CudoVentures/cudos-node/x/mail/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the message
	for _, elem := range genState.MessageList {
		k.SetMessage(ctx, elem)
		k.SetMessageIDSender(ctx, elem.Sender, elem.Id)
		k.SetMessageIDReceiver(ctx, elem.Receiver, elem.Id)
	}

	// Set message count
	k.SetMessageCount(ctx, genState.MessageCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MessageList = k.GetAllMessage(ctx)
	genesis.MessageSentList = k.GetAllSentMessage(ctx)
	genesis.MessageReceivedList = k.GetAllReceivedMessage(ctx)
	genesis.MessageCount = k.GetMessageCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
