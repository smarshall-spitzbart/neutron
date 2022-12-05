package keeper_test

import (
	"testing"

	testkeeper "github.com/neutron-org/neutron/testutil/tokendistribution/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/neutron-org/neutron/x/tokendistribution/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.TokendistributionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
