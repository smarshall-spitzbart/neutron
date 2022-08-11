package keeper

import (
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"

	"github.com/neutron-org/neutron/internal/sudo"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

const (
	LabelRegisterInterchainQuery = "register_interchain_query"
	LabelSubmitQueryResult       = "submit_query_result"
)

type (
	Keeper struct {
		cdc         codec.BinaryCodec
		storeKey    storetypes.StoreKey
		memKey      storetypes.StoreKey
		paramstore  paramtypes.Subspace
		ibcKeeper   *ibckeeper.Keeper
		wasmKeeper  *wasm.Keeper
		sudoHandler sudo.SudoHandler
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	ibcKeeper *ibckeeper.Keeper,
	wasmKeeper *wasm.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		memKey:      memKey,
		paramstore:  ps,
		ibcKeeper:   ibcKeeper,
		wasmKeeper:  wasmKeeper,
		sudoHandler: sudo.NewSudoHandler(wasmKeeper, types.ModuleName),
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
