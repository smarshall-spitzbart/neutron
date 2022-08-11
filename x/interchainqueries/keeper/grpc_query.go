package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) RegisteredKVQuery(goCtx context.Context, request *types.QueryRegisteredKVQueryRequest) (*types.QueryRegisteredKVQueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	registeredQuery, err := k.GetKVQueryByID(ctx, request.QueryId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "failed to get registered query by query id: %v", err)
	}

	return &types.QueryRegisteredKVQueryResponse{RegisteredQuery: registeredQuery}, nil
}

func (k Keeper) RegisteredTXQuery(goCtx context.Context, request *types.QueryRegisteredTXQueryRequest) (*types.QueryRegisteredTXQueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	registeredQuery, err := k.GetTXQueryByID(ctx, request.QueryId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "failed to get registered query by query id: %v", err)
	}

	return &types.QueryRegisteredTXQueryResponse{RegisteredQuery: registeredQuery}, nil
}

func (k Keeper) RegisteredKVQueries(goCtx context.Context, req *types.QueryRegisteredKVQueriesRequest) (*types.QueryRegisteredKVQueriesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return k.GetRegisteredKVQueries(ctx, req)
}

func (k Keeper) GetRegisteredKVQueries(ctx sdk.Context, _ *types.QueryRegisteredKVQueriesRequest) (*types.QueryRegisteredKVQueriesResponse, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RegisteredKVQueryKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	queries := make([]types.RegisteredKVQuery, 0)
	for ; iterator.Valid(); iterator.Next() {
		query := types.RegisteredKVQuery{}
		k.cdc.MustUnmarshal(iterator.Value(), &query)
		queries = append(queries, query)
	}

	return &types.QueryRegisteredKVQueriesResponse{RegisteredQueries: queries}, nil
}

func (k Keeper) RegisteredTXQueries(goCtx context.Context, req *types.QueryRegisteredTXQueriesRequest) (*types.QueryRegisteredTXQueriesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return k.GetRegisteredTXQueries(ctx, req)
}

func (k Keeper) GetRegisteredTXQueries(ctx sdk.Context, _ *types.QueryRegisteredTXQueriesRequest) (*types.QueryRegisteredTXQueriesResponse, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RegisteredTXQueryKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	queries := make([]types.RegisteredTXQuery, 0)
	for ; iterator.Valid(); iterator.Next() {
		query := types.RegisteredTXQuery{}
		k.cdc.MustUnmarshal(iterator.Value(), &query)
		queries = append(queries, query)
	}

	return &types.QueryRegisteredTXQueriesResponse{RegisteredQueries: queries}, nil
}

func (k Keeper) KVQueryResult(goCtx context.Context, request *types.QueryRegisteredKVQueryResultRequest) (*types.QueryRegisteredKVQueryResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.checkRegisteredKVQueryExists(ctx, request.QueryId) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "query with id %d doesn't exist", request.QueryId)
	}

	result, err := k.GetKVQueryResultByID(ctx, request.QueryId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to get query result by query id: %v", err)
	}
	return &types.QueryRegisteredKVQueryResultResponse{Result: result}, nil
}
