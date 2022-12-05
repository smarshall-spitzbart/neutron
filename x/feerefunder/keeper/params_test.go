package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/neutron-org/neutron/testutil/keeper"
	"github.com/neutron-org/neutron/x/feerefunder/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FeeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
