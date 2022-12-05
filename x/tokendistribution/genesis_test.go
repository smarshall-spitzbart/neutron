package tokendistribution_test

import (
	"testing"

	keepertest "github.com/neutron-org/neutron/testutil/tokendistribution/keeper"

	"github.com/stretchr/testify/require"

	"github.com/neutron-org/neutron/testutil/tokendistribution/nullify"
	"github.com/neutron-org/neutron/x/tokendistribution"
	"github.com/neutron-org/neutron/x/tokendistribution/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokendistributionKeeper(t)
	tokendistribution.InitGenesis(ctx, *k, genesisState)
	got := tokendistribution.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
