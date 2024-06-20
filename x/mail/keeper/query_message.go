package keeper

import (
	"context"
	"encoding/binary"

	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MessageAll(goCtx context.Context, req *types.QueryAllMessageRequest) (*types.QueryAllMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var messages []types.Message
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	messageStore := prefix.NewStore(store, types.KeyPrefix(types.MessageKey))

	pageRes, err := query.Paginate(messageStore, req.Pagination, func(key []byte, value []byte) error {
		var message types.Message
		if err := k.cdc.Unmarshal(value, &message); err != nil {
			return err
		}

		messages = append(messages, message)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMessageResponse{Message: messages, Pagination: pageRes}, nil
}

func (k Keeper) MessageAllBySender(goCtx context.Context, req *types.QueryAllMessageBySenderRequest) (*types.QueryAllMessageBySenderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var messages []types.Message
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	messageStore := prefix.NewStore(store, types.KeyPrefix(types.MessageBySender+req.Sender))

	pageRes, err := query.Paginate(messageStore, req.Pagination, func(key []byte, value []byte) error {
		id := binary.BigEndian.Uint64(value)

		if message, found := k.GetMessage(ctx, id); found {
			messages = append(messages, message)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMessageBySenderResponse{Message: messages, Pagination: pageRes}, nil
}

func (k Keeper) MessageAllByReceiver(goCtx context.Context, req *types.QueryAllMessageByReceiverRequest) (*types.QueryAllMessageByReceiverResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var messages []types.Message
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	messageStore := prefix.NewStore(store, types.KeyPrefix(types.MessageByReceiver+req.Receiver))

	pageRes, err := query.Paginate(messageStore, req.Pagination, func(key []byte, value []byte) error {
		id := binary.BigEndian.Uint64(value)

		if message, found := k.GetMessage(ctx, id); found {
			messages = append(messages, message)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMessageByReceiverResponse{Message: messages, Pagination: pageRes}, nil
}

func (k Keeper) Message(goCtx context.Context, req *types.QueryGetMessageRequest) (*types.QueryGetMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	message, found := k.GetMessage(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	var forwards []types.Message
	if message.ForwardedId != 0 {
		forwards = k.getAllForwards(ctx, message, forwards)
	}

	return &types.QueryGetMessageResponse{Message: message, Forwards: forwards}, nil
}

func (k Keeper) getAllForwards(ctx sdk.Context, message types.Message, forwards []types.Message) []types.Message {
	if message.ForwardedId != 0 {
		forwarded, found := k.GetMessage(ctx, message.ForwardedId)
		forwards = append(forwards, forwarded)
		if found {
			return k.getAllForwards(ctx, forwarded, forwards)
		}
	}
	return forwards
}
