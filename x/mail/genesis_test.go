package mail_test

import (
	"testing"

	keepertest "github.com/CudoVentures/cudos-node/testutil/keeper"
	"github.com/CudoVentures/cudos-node/testutil/nullify"
	"github.com/CudoVentures/cudos-node/x/mail"
	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MessageList: []types.Message{
			{
				Id: 1,
			},
			{
				Id: 2,
			},
		},
		MessageCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MailKeeper(t)
	mail.InitGenesis(ctx, *k, genesisState)
	got := mail.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MessageList, got.MessageList)
	require.Equal(t, genesisState.MessageCount, got.MessageCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
