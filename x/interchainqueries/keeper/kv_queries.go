package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

func (k Keeper) GetLastRegisteredKVQueryKey(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types.LastRegisteredKVQueryIdKey)
	if bytes == nil {
		k.Logger(ctx).Debug("Last registered query key doesn't exist, GetLastRegisteredKVQueryKey returns 0")
		return 0
	}
	return sdk.BigEndianToUint64(bytes)
}

func (k Keeper) SetLastRegisteredKVQueryKey(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastRegisteredKVQueryIdKey, sdk.Uint64ToBigEndian(id))
}

func (k Keeper) SaveKVQuery(ctx sdk.Context, query types.RegisteredKVQuery) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := k.cdc.Marshal(&query)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrProtoMarshal, "failed to marshal registered query: %v", err)
	}

	store.Set(types.GetRegisteredKVQueryByIDKey(query.Id), bz)
	k.Logger(ctx).Debug("SaveKVQuery successful", "query", query)
	return nil
}

func (k Keeper) GetKVQueryByID(ctx sdk.Context, id uint64) (*types.RegisteredKVQuery, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetRegisteredKVQueryByIDKey(id))
	if bz == nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "there is no query with id: %v", id)
	}

	var query types.RegisteredKVQuery
	if err := k.cdc.Unmarshal(bz, &query); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}

	return &query, nil
}

func (k Keeper) SaveKVQueryResult(ctx sdk.Context, id uint64, result *types.QueryResult) error {
	store := ctx.KVStore(k.storeKey)

	if result.KvResults != nil {
		cleanResult := clearQueryResult(result)
		bz, err := k.cdc.Marshal(&cleanResult)
		if err != nil {
			return sdkerrors.Wrapf(types.ErrProtoMarshal, "failed to marshal result result: %v", err)
		}

		store.Set(types.GetRegisteredQueryResultByIDKey(id), bz)

		if err = k.UpdateLastRemoteHeight(ctx, id, result.Height); err != nil {
			return sdkerrors.Wrapf(err, "failed to update last remote height for a result with id %d: %v", id, err)
		}

		if err = k.UpdateLastLocalHeight(ctx, id, uint64(ctx.BlockHeight())); err != nil {
			return sdkerrors.Wrapf(err, "failed to update last local height for a result with id %d: %v", id, err)
		}
	}
	k.Logger(ctx).Debug("Successfully saved query result", "result", &result)
	return nil
}

// We don't need to store proofs or transactions, so we just remove unnecessary fields
func clearQueryResult(result *types.QueryResult) types.QueryResult {
	storageValues := make([]*types.StorageValue, 0, len(result.KvResults))
	for _, v := range result.KvResults {
		storageValues = append(storageValues, &types.StorageValue{
			StoragePrefix: v.StoragePrefix,
			Key:           v.Key,
			Value:         v.Value,
			Proof:         nil,
		})
	}

	cleanResult := types.QueryResult{
		KvResults: storageValues,
		Height:    result.Height,
	}

	return cleanResult
}

// GetKVQueryResultByID returns a QueryResult for query with id
func (k Keeper) GetKVQueryResultByID(ctx sdk.Context, id uint64) (*types.QueryResult, error) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetRegisteredQueryResultByIDKey(id))
	if bz == nil {
		return nil, types.ErrNoQueryResult
	}

	var query types.QueryResult
	if err := k.cdc.Unmarshal(bz, &query); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}

	return &query, nil
}

func (k Keeper) UpdateLastLocalHeight(ctx sdk.Context, queryID uint64, newLocalHeight uint64) error {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetRegisteredKVQueryByIDKey(queryID))
	if bz == nil {
		return sdkerrors.Wrapf(types.ErrInvalidQueryID, "query with ID %d not found", queryID)
	}

	var query types.RegisteredKVQuery
	if err := k.cdc.Unmarshal(bz, &query); err != nil {
		return sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}

	query.LastSubmittedResultLocalHeight = newLocalHeight

	return k.SaveKVQuery(ctx, query)
}

func (k Keeper) UpdateLastRemoteHeight(ctx sdk.Context, queryID uint64, newRemoteHeight uint64) error {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetRegisteredKVQueryByIDKey(queryID))
	if bz == nil {
		return sdkerrors.Wrapf(types.ErrInvalidQueryID, "query with ID %d not found", queryID)
	}

	var query types.RegisteredKVQuery
	if err := k.cdc.Unmarshal(bz, &query); err != nil {
		return sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}

	if query.LastSubmittedResultRemoteHeight >= newRemoteHeight {
		return sdkerrors.Wrapf(types.ErrInvalidHeight, "can't save query result for height %d: result height can't be less or equal then last submitted query result height %d", newRemoteHeight, query.LastSubmittedResultRemoteHeight)
	}

	query.LastSubmittedResultRemoteHeight = newRemoteHeight
	k.Logger(ctx).Debug("Updated last remote height on given query", "queryID", queryID, "new remote height", newRemoteHeight)
	return k.SaveKVQuery(ctx, query)
}

func (k Keeper) IterateRegisteredKVQueries(ctx sdk.Context, fn func(index int64, queryInfo types.RegisteredKVQuery) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RegisteredKVQueryKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		query := types.RegisteredKVQuery{}
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

func (k Keeper) checkRegisteredKVQueryExists(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetRegisteredKVQueryByIDKey(id))
}
