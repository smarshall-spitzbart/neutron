package keeper

import (
	"bytes"
	"context"
	"time"

	ics23 "github.com/confio/ics23/go"
	"github.com/cosmos/cosmos-sdk/telemetry"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	ibccommitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"

	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) RegisterInterchainKVQuery(goCtx context.Context, msg *types.MsgRegisterInterchainKVQuery) (*types.MsgRegisterInterchainQueryResponse, error) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), LabelRegisterInterchainQuery)

	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.Logger().Debug("RegisterInterchainKVQuery", "msg", msg)

	if err := msg.ValidateBasic(); err != nil {
		ctx.Logger().Debug("RegisterInterchainKVQuery: failed to validate message", "message", msg)
		return nil, sdkerrors.Wrapf(err, "invalid msg: %v", err)
	}

	if _, err := k.ibcKeeper.ConnectionKeeper.Connection(goCtx, &ibcconnectiontypes.QueryConnectionRequest{ConnectionId: msg.ConnectionId}); err != nil {
		ctx.Logger().Debug("RegisterInterchainKVQuery: failed to get connection with ID", "message", msg)
		return nil, sdkerrors.Wrapf(types.ErrInvalidConnectionID, "failed to get connection with ID '%s': %v", msg.ConnectionId, err)
	}

	qs := NewKVQueryStorage(&k.Keeper)

	lastID := qs.GetLastRegisteredQueryID(ctx)
	lastID += 1

	registeredQuery := types.RegisteredKVQuery{
		Id:                lastID,
		Owner:             msg.Sender,
		Keys:              msg.Keys,
		ZoneId:            msg.ZoneId,
		UpdatePeriod:      msg.UpdatePeriod,
		ConnectionId:      msg.ConnectionId,
		LastEmittedHeight: uint64(ctx.BlockHeight()),
	}

	qs.SetLastRegisteredQueryID(ctx, lastID)
	if err := qs.SaveQuery(ctx, &registeredQuery); err != nil {
		ctx.Logger().Debug("RegisterInterchainKVQuery: failed to save query", "message", &msg, "error", err)
		return nil, sdkerrors.Wrapf(err, "failed to save query: %v", err)
	}

	return &types.MsgRegisterInterchainQueryResponse{Id: lastID}, nil
}

func (k msgServer) RegisterInterchainTXQuery(goCtx context.Context, msg *types.MsgRegisterInterchainTXQuery) (*types.MsgRegisterInterchainQueryResponse, error) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), LabelRegisterInterchainQuery)

	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.Logger().Debug("RegisterInterchainTXQuery", "msg", msg)

	if err := msg.ValidateBasic(); err != nil {
		ctx.Logger().Debug("RegisterInterchainTXQuery: failed to validate message", "message", msg)
		return nil, sdkerrors.Wrapf(err, "invalid msg: %v", err)
	}

	if _, err := k.ibcKeeper.ConnectionKeeper.Connection(goCtx, &ibcconnectiontypes.QueryConnectionRequest{ConnectionId: msg.ConnectionId}); err != nil {
		ctx.Logger().Debug("RegisterInterchainTXQuery: failed to get connection with ID", "message", msg)
		return nil, sdkerrors.Wrapf(types.ErrInvalidConnectionID, "failed to get connection with ID '%s': %v", msg.ConnectionId, err)
	}

	qs := NewTXQueryStorage(&k.Keeper)

	lastID := qs.GetLastRegisteredQueryID(ctx)
	lastID += 1

	registeredQuery := types.RegisteredTXQuery{
		Id:                 lastID,
		Owner:              msg.Sender,
		ZoneId:             msg.ZoneId,
		ConnectionId:       msg.ConnectionId,
		TransactionsFilter: msg.TransactionsFilter,
	}

	qs.SetLastRegisteredQueryID(ctx, lastID)
	if err := qs.SaveQuery(ctx, &registeredQuery); err != nil {
		ctx.Logger().Debug("RegisterInterchainTXQuery: failed to save query", "message", &msg, "error", err)
		return nil, sdkerrors.Wrapf(err, "failed to save query: %v", err)
	}

	return &types.MsgRegisterInterchainQueryResponse{Id: lastID}, nil
}

func (k msgServer) SubmitKVQueryResult(goCtx context.Context, msg *types.MsgSubmitKVQueryResult) (*types.MsgSubmitQueryResultResponse, error) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), LabelRegisterInterchainQuery)

	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.Logger().Debug("SubmitKVQueryResult", "query_id", msg.QueryId)

	if err := msg.ValidateBasic(); err != nil {
		ctx.Logger().Debug("SubmitKVQueryResult: failed to validate message",
			"error", err, "message", msg)
		return nil, sdkerrors.Wrapf(err, "invalid msg: %v", err)
	}

	qs := NewKVQueryStorage(&k.Keeper)

	query, err := qs.GetQueryByID(ctx, msg.QueryId)
	if err != nil {
		ctx.Logger().Debug("SubmitKVQueryResult: failed to GetKVQueryByID",
			"error", err, "query_id", msg.QueryId)
		return nil, sdkerrors.Wrapf(err, "failed to get query by id: %v", err)
	}

	if len(msg.Result.KvResults) != len(query.Keys) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSubmittedResult, "KV keys length from result is not equal to registered query keys length: %v != %v", len(msg.Result.KvResults), query.Keys)
	}

	resp, err := k.ibcKeeper.ConnectionConsensusState(goCtx, &ibcconnectiontypes.QueryConnectionConsensusStateRequest{
		ConnectionId:   query.ConnectionId,
		RevisionNumber: msg.Result.Revision,
		RevisionHeight: msg.Result.Height + 1,
	})
	if err != nil {
		ctx.Logger().Debug("SubmitKVQueryResult: failed to get ConnectionConsensusState",
			"error", err, "query", query, "message", msg)
		return nil, sdkerrors.Wrapf(ibcclienttypes.ErrConsensusStateNotFound, "failed to get consensus state: %v", err)
	}

	consensusState, err := ibcclienttypes.UnpackConsensusState(resp.ConsensusState)
	if err != nil {
		ctx.Logger().Error("SubmitKVQueryResult: failed to UnpackConsensusState",
			"error", err, "query", query, "message", msg)
		return nil, sdkerrors.Wrapf(types.ErrProtoUnmarshal, "failed to unpack consesus state: %v", err)
	}

	for index, result := range msg.Result.KvResults {
		proof, err := ibccommitmenttypes.ConvertProofs(result.Proof)
		if err != nil {
			ctx.Logger().Debug("SubmitKVQueryResult: failed to ConvertProofs",
				"error", err, "query", query, "message", msg)
			return nil, sdkerrors.Wrapf(types.ErrInvalidType, "failed to convert crypto.ProofOps to MerkleProof: %v", err)
		}

		if !bytes.Equal(result.Key, query.Keys[index].Key) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidSubmittedResult, "KV key from result is not equal to registered query key: %v != %v", result.Key, query.Keys[index].Key)
		}

		if result.StoragePrefix != query.Keys[index].Path {
			return nil, sdkerrors.Wrapf(types.ErrInvalidSubmittedResult, "KV path from result is not equal to registered query storage prefix: %v != %v", result.StoragePrefix, query.Keys[index].Path)
		}

		path := ibccommitmenttypes.NewMerklePath(result.StoragePrefix, string(result.Key))

		// identify what kind proofs (non-existence proof always has *ics23.CommitmentProof_Nonexist as the first item) we got
		// and call corresponding method to verify it
		switch proof.GetProofs()[0].GetProof().(type) {
		// we can get non-existence proof if someone queried some key which is not exists in the storage on remote chain
		case *ics23.CommitmentProof_Nonexist:
			if err := proof.VerifyNonMembership(ibccommitmenttypes.GetSDKSpecs(), consensusState.GetRoot(), path); err != nil {
				ctx.Logger().Debug("SubmitKVQueryResult: failed to VerifyNonMembership",
					"error", err, "query", query, "message", msg, "path", path)
				return nil, sdkerrors.Wrapf(types.ErrInvalidProof, "failed to verify proof: %v", err)
			}
			result.Value = nil
		case *ics23.CommitmentProof_Exist:
			if err := proof.VerifyMembership(ibccommitmenttypes.GetSDKSpecs(), consensusState.GetRoot(), path, result.Value); err != nil {
				ctx.Logger().Debug("SubmitKVQueryResult: failed to VerifyMembership",
					"error", err, "query", query, "message", msg, "path", path)
				return nil, sdkerrors.Wrapf(types.ErrInvalidProof, "failed to verify proof: %v", err)
			}
		default:
			return nil, sdkerrors.Wrapf(types.ErrInvalidProof, "unknown proof type %T", proof.GetProofs()[0].GetProof())
		}
	}

	queryOwner, err := sdk.AccAddressFromBech32(query.Owner)
	if err != nil {
		ctx.Logger().Error("SubmitKVQueryResult: failed to decode AccAddressFromBech32",
			"error", err, "query", query, "message", msg)
		return nil, sdkerrors.Wrapf(types.ErrInternal, "failed to decode owner contract address: %v", err)
	}

	if err = k.SaveKVQueryResult(ctx, msg.QueryId, msg.Result); err != nil {
		ctx.Logger().Error("SubmitKVQueryResult: failed to SaveKVQueryResult",
			"error", err, "query", query, "message", msg)
		return nil, sdkerrors.Wrapf(err, "failed to SaveKVQueryResult: %v", err)
	}

	if msg.Result.GetAllowKvCallbacks() {
		// Let the query owner contract process the query result.
		if _, err := k.sudoHandler.SudoKVQueryResult(ctx, queryOwner, query.Id); err != nil {
			ctx.Logger().Debug("SubmitKVQueryResult: failed to SudoKVQueryResult",
				"error", err, "query_id", query.GetId())
			return nil, sdkerrors.Wrapf(err, "contract %s rejected KV query result (query_id: %d)",
				queryOwner, query.GetId())
		}
		return &types.MsgSubmitQueryResultResponse{}, nil
	}

	return &types.MsgSubmitQueryResultResponse{}, nil
}

func (k msgServer) SubmitTXQueryResult(goCtx context.Context, msg *types.MsgSubmitTXQueryResult) (*types.MsgSubmitQueryResultResponse, error) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), LabelRegisterInterchainQuery)

	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.Logger().Debug("SubmitTXQueryResult", "query_id", msg.QueryId)

	if err := msg.ValidateBasic(); err != nil {
		ctx.Logger().Debug("SubmitTXQueryResult: failed to validate message",
			"error", err, "message", msg)
		return nil, sdkerrors.Wrapf(err, "invalid msg: %v", err)
	}

	qs := NewKVQueryStorage(&k.Keeper)

	query, err := qs.GetQueryByID(ctx, msg.QueryId)
	if err != nil {
		ctx.Logger().Debug("SubmitTXQueryResult: failed to GetTXQueryByID",
			"error", err, "query_id", msg.QueryId)
		return nil, sdkerrors.Wrapf(err, "failed to get query by id: %v", err)
	}

	queryOwner, err := sdk.AccAddressFromBech32(query.Owner)
	if err != nil {
		ctx.Logger().Error("SubmitTXQueryResult: failed to decode AccAddressFromBech32",
			"error", err, "query", query, "message", msg)
		return nil, sdkerrors.Wrapf(err, "failed to decode owner contract address (%s)", query.Owner)
	}

	if err := k.ProcessBlock(ctx, queryOwner, msg.QueryId, msg.ClientId, msg.Block); err != nil {
		ctx.Logger().Debug("SubmitTXQueryResult: failed to ProcessBlock",
			"error", err, "query", query, "message", msg)
		return nil, sdkerrors.Wrapf(err, "failed to ProcessBlock: %v", err)
	}

	return &types.MsgSubmitQueryResultResponse{}, nil
}

var _ types.MsgServer = msgServer{}
