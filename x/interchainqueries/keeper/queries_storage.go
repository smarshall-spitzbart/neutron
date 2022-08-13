package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

type Query interface {
	codec.ProtoMarshaler
	GetId() uint64
}

type QueriesStorage[Q any, T interface {
	*Q
	Query
}] struct {
	k        *Keeper
	queryKey []byte
}

func NewKVQueryStorage(k *Keeper) *QueriesStorage[types.RegisteredKVQuery, *types.RegisteredKVQuery] {
	return &QueriesStorage[types.RegisteredKVQuery, *types.RegisteredKVQuery]{
		k:        k,
		queryKey: types.RegisteredKVQueryKey,
	}
}

func NewTXQueryStorage(k *Keeper) *QueriesStorage[types.RegisteredTXQuery, *types.RegisteredTXQuery] {
	return &QueriesStorage[types.RegisteredTXQuery, *types.RegisteredTXQuery]{
		k:        k,
		queryKey: types.RegisteredTXQueryKey,
	}
}

func (qs QueriesStorage[Q, T]) GetLastRegisteredQueryID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(qs.k.storeKey)
	bytes := store.Get(types.GetLastRegisteredQueryIdKey(qs.queryKey))
	if bytes == nil {
		qs.k.Logger(ctx).Debug("Last registered query key doesn't exist, GetLastRegisteredQueryKey returns 0")
		return 0
	}
	return sdk.BigEndianToUint64(bytes)
}

func (qs QueriesStorage[Q, T]) SetLastRegisteredQueryID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(qs.k.storeKey)
	store.Set(types.GetLastRegisteredQueryIdKey(qs.queryKey), sdk.Uint64ToBigEndian(id))
}

func (qs QueriesStorage[Q, T]) SaveQuery(ctx sdk.Context, query Q) error {
	store := ctx.KVStore(qs.k.storeKey)

	bz, err := qs.k.cdc.Marshal(T(&query))
	if err != nil {
		return sdkerrors.Wrapf(types.ErrProtoMarshal, "failed to marshal registered query: %v", err)
	}

	store.Set(types.GetRegisteredQueryByIDKey(qs.queryKey, T(&query).GetId()), bz)
	qs.k.Logger(ctx).Debug("SaveTXQuery successful", "query", query)
	return nil
}

func (qs QueriesStorage[Q, T]) GetQueryByID(ctx sdk.Context, id uint64) (Q, error) {
	store := ctx.KVStore(qs.k.storeKey)

	var query Q

	bz := store.Get(types.GetRegisteredQueryByIDKey(qs.queryKey, id))
	if bz == nil {
		return query, sdkerrors.Wrapf(types.ErrInvalidQueryID, "there is no query with id: %v", id)
	}

	if err := qs.k.cdc.Unmarshal(bz, T(&query)); err != nil {
		return query, sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unmarshal registered query: %v", err)
	}

	return query, nil
}

func (qs QueriesStorage[Q, T]) IterateRegisteredQueries(ctx sdk.Context, fn func(index int64, queryInfo Q) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(qs.k.storeKey), qs.queryKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		var query Q
		if err := qs.k.cdc.Unmarshal(iterator.Value(), T(&query)); err != nil {
			continue
		}
		stop := fn(i, query)

		if stop {
			break
		}
		i++
	}
	qs.k.Logger(ctx).Debug("Iterated over registered queries", "quantity", i)
}

func (qs QueriesStorage[Q, T]) CheckQueryWithIDExists(ctx sdk.Context, queryID uint64) bool {
	store := ctx.KVStore(qs.k.storeKey)

	return store.Has(types.GetRegisteredQueryByIDKey(qs.queryKey, queryID))
}
