package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

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
	qs := NewKVQueryStorage(&k)

	query, err := qs.GetQueryByID(ctx, queryID)
	if err != nil {
		return sdkerrors.Wrapf(err, "failed to get query by id: %d", queryID)
	}

	query.LastSubmittedResultLocalHeight = newLocalHeight

	return qs.SaveQuery(ctx, query)
}

func (k Keeper) UpdateLastRemoteHeight(ctx sdk.Context, queryID uint64, newRemoteHeight uint64) error {
	qs := NewKVQueryStorage(&k)

	query, err := qs.GetQueryByID(ctx, queryID)
	if err != nil {
		return sdkerrors.Wrapf(err, "failed to get query by id: %d", queryID)
	}

	if query.LastSubmittedResultRemoteHeight >= newRemoteHeight {
		return sdkerrors.Wrapf(types.ErrInvalidHeight, "can't save query result for height %d: result height can't be less or equal then last submitted query result height %d", newRemoteHeight, query.LastSubmittedResultRemoteHeight)
	}

	query.LastSubmittedResultRemoteHeight = newRemoteHeight
	k.Logger(ctx).Debug("Updated last remote height on given query", "queryID", queryID, "new remote height", newRemoteHeight)
	return qs.SaveQuery(ctx, query)
}
