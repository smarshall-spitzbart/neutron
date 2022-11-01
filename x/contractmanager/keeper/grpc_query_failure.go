package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/neutron-org/neutron/x/contractmanager/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllFailures(c context.Context, req *types.QueryFailuresRequest) (*types.QueryFailuresResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var failures []types.Failure
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)

	var failureStore prefix.Store
	if req.Address != "" {
		if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to parse address: %s", req.Address)
		}
		failureStore = prefix.NewStore(store, types.GetFailureKeyPrefix(req.Address))
	} else {
		failureStore = prefix.NewStore(store, types.ContractFailuresKey)
	}

	pageRes, err := query.Paginate(failureStore, req.Pagination, func(key []byte, value []byte) error {
		var failure types.Failure
		if err := k.cdc.Unmarshal(value, &failure); err != nil {
			return err
		}

		failures = append(failures, failure)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFailuresResponse{Failures: failures, Pagination: pageRes}, nil
}

func (k Keeper) Failures(c context.Context, req *types.QueryFailuresRequest) (*types.QueryFailuresResponse, error) {
	return k.AllFailures(c, req)
}
