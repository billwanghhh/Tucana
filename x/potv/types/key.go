package types

// todo: 3. add keys value
const (
	// ModuleName defines the module name
	ModuleName = "potv"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for potv
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the potv module.
	QuerierRoute = StoreKey
)

// KeyPrefixPOTV defines prefix key for storing potv-related data
var KeyPrefixPOTV = []byte{0x01} // 示例前缀字节，实际应根据需要设计
