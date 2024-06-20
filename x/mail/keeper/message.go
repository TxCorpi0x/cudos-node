package keeper

import (
	"encoding/binary"

	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
)

// GetMessageCount get the total number of message
func (k Keeper) GetMessageCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MessageCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMessageCount set the total number of message
func (k Keeper) SetMessageCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MessageCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// SetMessageCount set the total number of message
func (k Keeper) SetMessageIDSender(ctx sdk.Context, sender string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MessageBySender + sender + cast.ToString(id))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(byteKey, bz)
}

// SetMessageCount set the total number of message
func (k Keeper) SetMessageIDReceiver(ctx sdk.Context, receiver string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MessageByReceiver + receiver + cast.ToString(id))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(byteKey, bz)
}

// AppendMessage appends a message in the store with a new id and update the count
func (k Keeper) AppendMessage(
	ctx sdk.Context,
	message types.Message,
) uint64 {
	// Create the message
	count := k.GetMessageCount(ctx)

	// Set the ID of the appended value
	message.Id = count + 1

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	appendedValue := k.cdc.MustMarshal(&message)
	store.Set(GetMessageIDBytes(message.Id), appendedValue)

	k.SetMessageIDSender(ctx, message.Sender, message.Id)
	k.SetMessageIDReceiver(ctx, message.Receiver, message.Id)

	// Update message count
	k.SetMessageCount(ctx, count+1)

	return message.Id
}

// SetMessage set a specific message in the store
func (k Keeper) SetMessage(ctx sdk.Context, message types.Message) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	b := k.cdc.MustMarshal(&message)
	store.Set(GetMessageIDBytes(message.Id), b)
}

// GetMessage returns a message from its id
func (k Keeper) GetMessage(ctx sdk.Context, id uint64) (val types.Message, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	b := store.Get(GetMessageIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMessage removes a message from the store
func (k Keeper) RemoveMessage(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	store.Delete(GetMessageIDBytes(id))
}

// GetAllMessage returns all message
func (k Keeper) GetAllMessage(ctx sdk.Context) (list []types.Message) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Message
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllSentMessage returns all sent messages
func (k Keeper) GetAllSentMessage(ctx sdk.Context) (list []uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageBySender))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val := binary.BigEndian.Uint64(iterator.Value())
		list = append(list, val)
	}

	return
}

// GetAllReceivedMessage returns all received messages
func (k Keeper) GetAllReceivedMessage(ctx sdk.Context) (list []uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MessageByReceiver))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val := binary.BigEndian.Uint64(iterator.Value())
		list = append(list, val)
	}

	return
}

// GetMessageIDBytes returns the byte representation of the ID
func GetMessageIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMessageIDFromBytes returns ID in uint64 format from a byte array
func GetMessageIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
