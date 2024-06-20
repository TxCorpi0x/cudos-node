package keeper

import (
	"context"

	"github.com/CudoVentures/cudos-node/x/mail/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mrz1836/go-sanitize"
)

func (k msgServer) SendMessage(goCtx context.Context, msg *types.MsgSendMessage) (*types.MsgSendMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bodyLen := uint32(len(msg.Body))
	if bodyLen > k.GetParams(ctx).MaxBodyCharacters {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "body is too large: %d", bodyLen)
	}

	if msg.ForwardedId > 0 {
		_, found := k.GetMessage(ctx, msg.ForwardedId)
		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no message found with id equal to forwarded id: %d", msg.ForwardedId)
		}
	}

	message := types.Message{
		Sender:      msg.Sender,
		Receiver:    msg.Receiver,
		Body:        sanitize.XSS(msg.Body),
		ForwardedId: msg.ForwardedId,
		Ts:          uint64(ctx.BlockTime().Unix()),
	}

	id := k.AppendMessage(
		ctx,
		message,
	)

	return &types.MsgSendMessageResponse{
		Id: id,
	}, nil
}
