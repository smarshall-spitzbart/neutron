package types

import (
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

var (
	_ codectypes.UnpackInterfacesMessage = MsgSubmitTXQueryResult{}
)

func (msg MsgSubmitKVQueryResult) Route() string {
	return RouterKey
}

func (msg MsgSubmitKVQueryResult) Type() string {
	return "submit-kv-query-result"
}

func (msg MsgSubmitKVQueryResult) ValidateBasic() error {
	if msg.Result == nil {
		return sdkerrors.Wrap(ErrEmptyResult, "query result can't be empty")
	}

	if len(msg.Result.KvResults) == 0 {
		return sdkerrors.Wrap(ErrEmptyResult, "query result can't be empty")
	}

	if msg.QueryId == 0 {
		return sdkerrors.Wrap(ErrInvalidQueryID, "query id cannot be equal zero")
	}

	if strings.TrimSpace(msg.Sender) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "failed to parse address: %s", msg.Sender)
	}

	if strings.TrimSpace(msg.ClientId) == "" {
		return sdkerrors.Wrap(ErrInvalidClientID, "client id cannot be empty")
	}

	return nil
}

func (msg MsgSubmitKVQueryResult) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))

}

func (msg MsgSubmitKVQueryResult) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{senderAddr}
}

func (msg MsgSubmitTXQueryResult) Route() string {
	return RouterKey
}

func (msg MsgSubmitTXQueryResult) Type() string {
	return "submit-tx-query-result"
}

func (msg MsgSubmitTXQueryResult) ValidateBasic() error {
	if msg.Block == nil {
		return sdkerrors.Wrap(ErrEmptyResult, "query result can't be empty")
	}

	if msg.Block.Tx == nil {
		return sdkerrors.Wrap(ErrEmptyResult, "result tx can't be empty")
	}

	if msg.Block.Header == nil {
		return sdkerrors.Wrapf(ErrInvalidHeader, "header can't be empty")
	}

	if msg.Block.NextBlockHeader == nil {
		return sdkerrors.Wrapf(ErrInvalidHeader, "next block header can't be empty")
	}

	if msg.QueryId == 0 {
		return sdkerrors.Wrap(ErrInvalidQueryID, "query id cannot be equal zero")
	}

	if strings.TrimSpace(msg.Sender) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "failed to parse address: %s", msg.Sender)
	}

	if strings.TrimSpace(msg.ClientId) == "" {
		return sdkerrors.Wrap(ErrInvalidClientID, "client id cannot be empty")
	}

	return nil
}

func (msg MsgSubmitTXQueryResult) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))

}

func (msg MsgSubmitTXQueryResult) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{senderAddr}
}

func (msg MsgRegisterInterchainKVQuery) Route() string {
	return RouterKey
}

func (msg MsgRegisterInterchainKVQuery) Type() string {
	return "register-interchain-tx-query"
}

func (msg MsgRegisterInterchainKVQuery) ValidateBasic() error {
	if msg.UpdatePeriod == 0 {
		return sdkerrors.Wrap(ErrInvalidUpdatePeriod, "update period cannot be equal zero")
	}

	if strings.TrimSpace(msg.Sender) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "failed to parse address: %s", msg.Sender)
	}

	if strings.TrimSpace(msg.ConnectionId) == "" {
		return sdkerrors.Wrap(ErrInvalidConnectionID, "connection id cannot be empty")
	}

	if strings.TrimSpace(msg.ZoneId) == "" {
		return sdkerrors.Wrap(ErrInvalidZoneID, "zone id cannot be empty")
	}

	return nil
}

func (msg MsgRegisterInterchainKVQuery) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))

}

func (msg MsgRegisterInterchainKVQuery) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{senderAddr}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgSubmitTXQueryResult) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var header exported.Header
	if err := unpacker.UnpackAny(msg.GetBlock().GetHeader(), &header); err != nil {
		return err
	}

	if err := unpacker.UnpackAny(msg.GetBlock().GetNextBlockHeader(), &header); err != nil {
		return err
	}

	return nil
}

func (msg MsgRegisterInterchainTXQuery) Route() string {
	return RouterKey
}

func (msg MsgRegisterInterchainTXQuery) Type() string {
	return "register-interchain-tx-query"
}

func (msg MsgRegisterInterchainTXQuery) ValidateBasic() error {
	if strings.TrimSpace(msg.Sender) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "failed to parse address: %s", msg.Sender)
	}

	if strings.TrimSpace(msg.ConnectionId) == "" {
		return sdkerrors.Wrap(ErrInvalidConnectionID, "connection id cannot be empty")
	}

	if strings.TrimSpace(msg.ZoneId) == "" {
		return sdkerrors.Wrap(ErrInvalidZoneID, "zone id cannot be empty")
	}

	if IsValidJSON(msg.TransactionsFilter) {
		return sdkerrors.Wrapf(ErrInvalidTransactionsFilter, "transactions filter must be a valid JSON")
	}

	return nil
}

func (msg MsgRegisterInterchainTXQuery) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))

}

func (msg MsgRegisterInterchainTXQuery) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{senderAddr}
}
