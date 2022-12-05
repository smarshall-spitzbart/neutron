package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/neutron-org/neutron/testutil/tokendistribution/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/neutron-org/neutron/x/tokendistribution/keeper"
	"github.com/neutron-org/neutron/x/tokendistribution/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TokendistributionKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
