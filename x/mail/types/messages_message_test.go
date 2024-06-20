package types

import (
	"testing"

	"github.com/CudoVentures/cudos-node/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSendMessage_ValidateBasic(t *testing.T) {
	senderAddr := sample.AccAddress()
	receiverAddr := sample.AccAddress()
	tests := []struct {
		name string
		msg  MsgSendMessage
		err  error
	}{
		{
			name: "invalid sender address",
			msg: MsgSendMessage{
				Sender:   "invalid_address",
				Receiver: receiverAddr,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid receiver address",
			msg: MsgSendMessage{
				Sender:   senderAddr,
				Receiver: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "same sender and receiver",
			msg: MsgSendMessage{
				Sender:   senderAddr,
				Receiver: senderAddr,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "valid address",
			msg: MsgSendMessage{
				Sender:      senderAddr,
				Receiver:    receiverAddr,
				Body:        "valid body",
				ForwardedId: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
