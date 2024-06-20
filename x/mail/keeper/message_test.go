package keeper_test

import (
	"testing"

	keepertest "github.com/CudoVentures/cudos-node/testutil/keeper"
	"github.com/CudoVentures/cudos-node/testutil/nullify"
	"github.com/CudoVentures/cudos-node/x/mail/keeper"
	"github.com/CudoVentures/cudos-node/x/mail/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNMessage(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Message {
	items := make([]types.Message, n)
	for i := range items {
		items[i].Id = keeper.AppendMessage(ctx, items[i])
	}
	return items
}

func TestMessageGet(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNMessage(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMessage(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMessageRemove(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNMessage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMessage(ctx, item.Id)
		_, found := keeper.GetMessage(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMessageGetAll(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNMessage(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMessage(ctx)),
	)
}

func TestMessageCount(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNMessage(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMessageCount(ctx))
}
