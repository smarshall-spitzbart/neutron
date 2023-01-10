package feeburner_test

import (
	"testing"

	"github.com/neutron-org/neutron/app"

	"github.com/neutron-org/neutron/testutil/feeburner/keeper"
	"github.com/neutron-org/neutron/testutil/feeburner/nullify"
	"github.com/neutron-org/neutron/x/feeburner"
	"github.com/neutron-org/neutron/x/feeburner/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	_ = app.GetDefaultConfig()

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keeper.FeeburnerKeeper(t)
	feeburner.InitGenesis(ctx, *k, genesisState)
	got := feeburner.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
