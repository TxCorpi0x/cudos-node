package types

const (
	// ModuleName defines the module name
	ModuleName = "mail"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mail"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MessageKey        = "Message/value/"
	MessageCountKey   = "Message/count/"
	MessageBySender   = "MessageBySender/value/"
	MessageByReceiver = "MessageByReceiver/value/"
)
