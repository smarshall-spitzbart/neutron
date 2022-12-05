package keeper

import (
	"github.com/neutron-org/neutron/x/tokendistribution/types"
)

var _ types.QueryServer = Keeper{}
