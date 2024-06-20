package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSendMessage   = "create_message"
	TypeMsgUpdateMessage = "update_message"
	TypeMsgDeleteMessage = "delete_message"
)

var _ sdk.Msg = &MsgSendMessage{}

func NewMsgSendMessage(sender string, receiver string, body string, forwardedId uint64) *MsgSendMessage {
	return &MsgSendMessage{
		Sender:      sender,
		Receiver:    receiver,
		Body:        body,
		ForwardedId: forwardedId,
	}
}

func (msg *MsgSendMessage) Route() string {
	return RouterKey
}

func (msg *MsgSendMessage) Type() string {
	return TypeMsgSendMessage
}

func (msg *MsgSendMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender and receiver can not be a same address")
	}

	return nil
}
