package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "interchainqueries"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_interchainqueries"
)

const (
	prefixRegisteredKVQuery = iota + 1
	prefixRegisteredTXQuery
	prefixRegisteredQueryResult

	prefixSubmittedTx
)

var (
	RegisteredKVQueryKey       = []byte{prefixRegisteredKVQuery}
	RegisteredKVQueryResultKey = []byte{prefixRegisteredQueryResult}

	RegisteredTXQueryKey = []byte{prefixRegisteredKVQuery}

	SubmittedTxKey = []byte{prefixSubmittedTx}

	LastRegisteredQueryIdKey = []byte{0x64}
)

func GetLastRegisteredQueryIdKey(queryKey []byte) []byte {
	return append(LastRegisteredQueryIdKey, queryKey...)
}

func GetRegisteredQueryByIDKey(queryKey []byte, id uint64) []byte {
	return append(queryKey, sdk.Uint64ToBigEndian(id)...)
}

func GetSubmittedTransactionIDForQueryKeyPrefix(queryID uint64) []byte {
	return append(SubmittedTxKey, sdk.Uint64ToBigEndian(queryID)...)
}

func GetSubmittedTransactionIDForQueryKey(queryID uint64, txHash []byte) []byte {
	return append(GetSubmittedTransactionIDForQueryKeyPrefix(queryID), txHash...)
}

func GetRegisteredQueryResultByIDKey(id uint64) []byte {
	return append(RegisteredKVQueryResultKey, sdk.Uint64ToBigEndian(id)...)
}
