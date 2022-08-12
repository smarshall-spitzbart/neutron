package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) RegisteredKVQuery(goCtx context.Context, request *types.QueryRegisteredKVQueryRequest) (*types.QueryRegisteredKVQueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	qs := NewKVQueryStorage(&k)

	registeredQuery, err := qs.GetQueryByID(ctx, request.QueryId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "failed to get registered query by query id: %v", err)
	}

	return &types.QueryRegisteredKVQueryResponse{RegisteredQuery: registeredQuery}, nil
}

func (k Keeper) RegisteredTXQuery(goCtx context.Context, request *types.QueryRegisteredTXQueryRequest) (*types.QueryRegisteredTXQueryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	qs := NewTXQueryStorage(&k)

	registeredQuery, err := qs.GetQueryByID(ctx, request.QueryId)
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
	qs := NewKVQueryStorage(&k)
	queries := make([]types.RegisteredKVQuery, 0)

	qs.IterateRegisteredQueries(ctx, func(i int64, registeredQuery *types.RegisteredKVQuery) (stop bool) {
		queries = append(queries, *registeredQuery)
		return false
	})

	return &types.QueryRegisteredKVQueriesResponse{RegisteredQueries: queries}, nil
}

func (k Keeper) RegisteredTXQueries(goCtx context.Context, req *types.QueryRegisteredTXQueriesRequest) (*types.QueryRegisteredTXQueriesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return k.GetRegisteredTXQueries(ctx, req)
}

func (k Keeper) GetRegisteredTXQueries(ctx sdk.Context, _ *types.QueryRegisteredTXQueriesRequest) (*types.QueryRegisteredTXQueriesResponse, error) {
	qs := NewTXQueryStorage(&k)
	queries := make([]types.RegisteredTXQuery, 0)

	qs.IterateRegisteredQueries(ctx, func(i int64, registeredQuery *types.RegisteredTXQuery) (stop bool) {
		queries = append(queries, *registeredQuery)
		return false
	})

	return &types.QueryRegisteredTXQueriesResponse{RegisteredQueries: queries}, nil
}

func (k Keeper) KVQueryResult(goCtx context.Context, request *types.QueryRegisteredKVQueryResultRequest) (*types.QueryRegisteredKVQueryResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	qs := NewKVQueryStorage(&k)

	if !qs.CheckQueryWithIDExists(ctx, request.QueryId) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidQueryID, "query with id %d doesn't exist", request.QueryId)
	}

	result, err := k.GetKVQueryResultByID(ctx, request.QueryId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to get query result by query id: %v", err)
	}
	return &types.QueryRegisteredKVQueryResultResponse{Result: result}, nil
}
