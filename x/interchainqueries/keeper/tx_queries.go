package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

func (k Keeper) GetLastRegisteredTXQueryKey(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.LastRegisteredTXQueryIdKey)
	if bytes == nil {
		k.Logger(ctx).Debug("Last registered query key doesn't exist, GetLastRegisteredTXQueryKey returns 0")
		return 0
	}
	return sdk.BigEndianToUint64(bytes)
}

func (k Keeper) SetLastRegisteredTXQueryKey(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastRegisteredTXQueryIdKey, sdk.Uint64ToBigEndian(id))
}

func (k Keeper) SaveTXQuery(ctx sdk.Context, query types.RegisteredTXQuery) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := k.cdc.Marshal(&query)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrProtoMarshal, "failed to marshal registered query: %v", err)
	}

	store.Set(types.GetRegisteredTXQueryByIDKey(query.Id), bz)
	k.Logger(ctx).Debug("SaveTXQuery successful", "query", query)
	return nil
}

func (k Keeper) GetTXQueryByID(ctx sdk.Context, id uint64) (*types.RegisteredTXQuery, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetRegisteredTXQueryByIDKey(id))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "there is no query with id: %v", id)
	}

	var query types.RegisteredTXQuery
	if err := k.cdc.Unmarshal(bz, &query); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}

	return &query, nil
}

// SaveTransactionAsProcessed simply stores a key (SubmittedTxKey + bigEndianBytes(queryID) + tx_hash) with
// mock data. This key can be used to check whether a certain transaction was already submitted for a specific
// transaction query.
func (k Keeper) SaveTransactionAsProcessed(ctx sdk.Context, queryID uint64, txHash []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSubmittedTransactionIDForQueryKey(queryID, txHash)

	store.Set(key, []byte{})
}

func (k Keeper) CheckTransactionIsAlreadyProcessed(ctx sdk.Context, queryID uint64, txHash []byte) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSubmittedTransactionIDForQueryKey(queryID, txHash)

	return store.Has(key)
}

func (k Keeper) IterateRegisteredTXQueries(ctx sdk.Context, fn func(index int64, queryInfo types.RegisteredTXQuery) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RegisteredTXQueryKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		query := types.RegisteredTXQuery{}
		if err := k.cdc.Unmarshal(iterator.Value(), &query); err != nil {
			continue
		}
		stop := fn(i, query)

		if stop {
			break
		}
		i++
	}
	k.Logger(ctx).Debug("Iterated over registered queries", "quantity", i)
}
