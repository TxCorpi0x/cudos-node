package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CudoVentures/cudos-node/x/mail/types"
)

func TestMessageMsgServerSend(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	sender := "A"
	receiver := "B"
	for i := 1; i < 6; i++ {
		resp, err := srv.SendMessage(ctx, &types.MsgSendMessage{Sender: sender, Receiver: receiver, Body: "sample body"})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestLargeMessageMsgServerSend(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	sender := "A"
	receiver := "B"
	_, err := srv.SendMessage(ctx, &types.MsgSendMessage{Sender: sender, Receiver: receiver, Body: `a large text to check if the validate basic prevents addition of out of allowed range
		of characters according to the parameters. it should return invalid request error when there are
 		lots of characters in the mail body`})
	require.Error(t, err)
}
