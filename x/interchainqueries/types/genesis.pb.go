// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainqueries/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RegisteredKVQuery struct {
	// The unique id of the registered query.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The address that registered the query.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// The KV-storage keys for which we want to get values from remote chain
	Keys []*KVKey `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
	// The chain of interest identifier.
	ZoneId string `protobuf:"bytes,4,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// The IBC connection ID for getting ConsensusState to verify proofs
	ConnectionId string `protobuf:"bytes,5,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// Parameter that defines how often the query must be updated.
	UpdatePeriod uint64 `protobuf:"varint,6,opt,name=update_period,json=updatePeriod,proto3" json:"update_period,omitempty"`
	// The local height when the event to update the query result was emitted last time.
	LastEmittedHeight uint64 `protobuf:"varint,7,opt,name=last_emitted_height,json=lastEmittedHeight,proto3" json:"last_emitted_height,omitempty"`
	// The local chain last block height when the query result was updated.
	LastSubmittedResultLocalHeight uint64 `protobuf:"varint,8,opt,name=last_submitted_result_local_height,json=lastSubmittedResultLocalHeight,proto3" json:"last_submitted_result_local_height,omitempty"`
	// The remote chain last block height when the query result was updated.
	LastSubmittedResultRemoteHeight uint64 `protobuf:"varint,9,opt,name=last_submitted_result_remote_height,json=lastSubmittedResultRemoteHeight,proto3" json:"last_submitted_result_remote_height,omitempty"`
}

func (m *RegisteredKVQuery) Reset()         { *m = RegisteredKVQuery{} }
func (m *RegisteredKVQuery) String() string { return proto.CompactTextString(m) }
func (*RegisteredKVQuery) ProtoMessage()    {}
func (*RegisteredKVQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{0}
}
func (m *RegisteredKVQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteredKVQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteredKVQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteredKVQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredKVQuery.Merge(m, src)
}
func (m *RegisteredKVQuery) XXX_Size() int {
	return m.Size()
}
func (m *RegisteredKVQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredKVQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredKVQuery proto.InternalMessageInfo

func (m *RegisteredKVQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisteredKVQuery) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RegisteredKVQuery) GetKeys() []*KVKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *RegisteredKVQuery) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *RegisteredKVQuery) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *RegisteredKVQuery) GetUpdatePeriod() uint64 {
	if m != nil {
		return m.UpdatePeriod
	}
	return 0
}

func (m *RegisteredKVQuery) GetLastEmittedHeight() uint64 {
	if m != nil {
		return m.LastEmittedHeight
	}
	return 0
}

func (m *RegisteredKVQuery) GetLastSubmittedResultLocalHeight() uint64 {
	if m != nil {
		return m.LastSubmittedResultLocalHeight
	}
	return 0
}

func (m *RegisteredKVQuery) GetLastSubmittedResultRemoteHeight() uint64 {
	if m != nil {
		return m.LastSubmittedResultRemoteHeight
	}
	return 0
}

type KVKey struct {
	// Path (storage prefix) to the storage where you want to read value by key (usually name of cosmos-sdk module: 'staking', 'bank', etc.)
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Key you want to read from the storage
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *KVKey) Reset()         { *m = KVKey{} }
func (m *KVKey) String() string { return proto.CompactTextString(m) }
func (*KVKey) ProtoMessage()    {}
func (*KVKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{1}
}
func (m *KVKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KVKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KVKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVKey.Merge(m, src)
}
func (m *KVKey) XXX_Size() int {
	return m.Size()
}
func (m *KVKey) XXX_DiscardUnknown() {
	xxx_messageInfo_KVKey.DiscardUnknown(m)
}

var xxx_messageInfo_KVKey proto.InternalMessageInfo

func (m *KVKey) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *KVKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type RegisteredTXQuery struct {
	// The unique id of the registered query.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The address that registered the query.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// The filter for transaction search ICQ
	TransactionsFilter string `protobuf:"bytes,3,opt,name=transactions_filter,json=transactionsFilter,proto3" json:"transactions_filter,omitempty"`
	// The chain of interest identifier.
	ZoneId string `protobuf:"bytes,4,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// The IBC connection ID for getting ConsensusState to verify proofs
	ConnectionId string `protobuf:"bytes,5,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
}

func (m *RegisteredTXQuery) Reset()         { *m = RegisteredTXQuery{} }
func (m *RegisteredTXQuery) String() string { return proto.CompactTextString(m) }
func (*RegisteredTXQuery) ProtoMessage()    {}
func (*RegisteredTXQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{2}
}
func (m *RegisteredTXQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteredTXQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteredTXQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteredTXQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredTXQuery.Merge(m, src)
}
func (m *RegisteredTXQuery) XXX_Size() int {
	return m.Size()
}
func (m *RegisteredTXQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredTXQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredTXQuery proto.InternalMessageInfo

func (m *RegisteredTXQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisteredTXQuery) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RegisteredTXQuery) GetTransactionsFilter() string {
	if m != nil {
		return m.TransactionsFilter
	}
	return ""
}

func (m *RegisteredTXQuery) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *RegisteredTXQuery) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

// GenesisState defines the interchainadapter module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e6c14f58b92f58, []int{3}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*RegisteredKVQuery)(nil), "neutron.interchainadapter.interchainqueries.RegisteredKVQuery")
	proto.RegisterType((*KVKey)(nil), "neutron.interchainadapter.interchainqueries.KVKey")
	proto.RegisterType((*RegisteredTXQuery)(nil), "neutron.interchainadapter.interchainqueries.RegisteredTXQuery")
	proto.RegisterType((*GenesisState)(nil), "neutron.interchainadapter.interchainqueries.GenesisState")
}

func init() { proto.RegisterFile("interchainqueries/genesis.proto", fileDescriptor_68e6c14f58b92f58) }

var fileDescriptor_68e6c14f58b92f58 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6a, 0x1b, 0x31,
	0x10, 0xf6, 0xfa, 0x2f, 0xb5, 0xe2, 0x96, 0x46, 0x09, 0x74, 0xc9, 0x61, 0x6d, 0x9c, 0x8b, 0xa1,
	0x64, 0x17, 0x9c, 0x4b, 0xcf, 0x81, 0xa6, 0x3f, 0xce, 0x21, 0x51, 0x4a, 0x28, 0xbd, 0x2c, 0xb2,
	0x77, 0xba, 0x16, 0xb1, 0xa5, 0xad, 0x34, 0x4b, 0xeb, 0x3e, 0x45, 0x1f, 0xa3, 0xa7, 0x3e, 0x47,
	0x8e, 0x39, 0xf6, 0x54, 0x8a, 0xfd, 0x22, 0x65, 0x67, 0x6d, 0x12, 0x70, 0x2e, 0xa6, 0xb7, 0x91,
	0xbe, 0x9f, 0x19, 0xf4, 0x8d, 0x58, 0x47, 0x69, 0x04, 0x3b, 0x9e, 0x48, 0xa5, 0xbf, 0xe4, 0x60,
	0x15, 0xb8, 0x28, 0x05, 0x0d, 0x4e, 0xb9, 0x30, 0xb3, 0x06, 0x0d, 0x7f, 0xa9, 0x21, 0x47, 0x6b,
	0x74, 0x78, 0x4f, 0x94, 0x89, 0xcc, 0x10, 0x6c, 0xb8, 0x21, 0x3d, 0x3c, 0x48, 0x4d, 0x6a, 0x48,
	0x17, 0x15, 0x55, 0x69, 0x71, 0x18, 0x6c, 0xf6, 0xc8, 0xa4, 0x95, 0xb3, 0x55, 0x8b, 0xde, 0xaf,
	0x1a, 0xdb, 0x13, 0x90, 0x2a, 0x87, 0x60, 0x21, 0x19, 0x5e, 0x5f, 0xe6, 0x60, 0xe7, 0xfc, 0x19,
	0xab, 0xaa, 0xc4, 0xf7, 0xba, 0x5e, 0xbf, 0x2e, 0xaa, 0x2a, 0xe1, 0x07, 0xac, 0x61, 0xbe, 0x6a,
	0xb0, 0x7e, 0xb5, 0xeb, 0xf5, 0x5b, 0xa2, 0x3c, 0xf0, 0x33, 0x56, 0xbf, 0x81, 0xb9, 0xf3, 0x6b,
	0xdd, 0x5a, 0x7f, 0x77, 0x30, 0x08, 0xb7, 0x98, 0x36, 0x1c, 0x5e, 0x0f, 0x61, 0x2e, 0x48, 0xcf,
	0x5f, 0xb0, 0x9d, 0xef, 0x46, 0x43, 0xac, 0x12, 0xbf, 0x4e, 0xfe, 0xcd, 0xe2, 0xf8, 0x2e, 0xe1,
	0x47, 0xec, 0xe9, 0xd8, 0x68, 0x0d, 0x63, 0x54, 0x46, 0x17, 0x70, 0x83, 0xe0, 0xf6, 0xfd, 0x65,
	0x49, 0xca, 0xb3, 0x44, 0x22, 0xc4, 0x19, 0x58, 0x65, 0x12, 0xbf, 0x49, 0x63, 0xb7, 0xcb, 0xcb,
	0x0b, 0xba, 0xe3, 0x21, 0xdb, 0x9f, 0x4a, 0x87, 0x31, 0xcc, 0x14, 0x22, 0x24, 0xf1, 0x04, 0x54,
	0x3a, 0x41, 0x7f, 0x87, 0xa8, 0x7b, 0x05, 0xf4, 0xba, 0x44, 0xde, 0x12, 0xc0, 0xdf, 0xb3, 0x1e,
	0xf1, 0x5d, 0x3e, 0x5a, 0x29, 0x2c, 0xb8, 0x7c, 0x8a, 0xf1, 0xd4, 0x8c, 0xe5, 0x74, 0x2d, 0x7f,
	0x42, 0xf2, 0xa0, 0x60, 0x5e, 0xad, 0x89, 0x82, 0x78, 0xe7, 0x05, 0x6d, 0xe5, 0x75, 0xce, 0x8e,
	0x1e, 0xf7, 0xb2, 0x30, 0x33, 0x08, 0x6b, 0xb3, 0x16, 0x99, 0x75, 0x1e, 0x31, 0x13, 0xc4, 0x2b,
	0xdd, 0x7a, 0xc7, 0xac, 0x41, 0x6f, 0xc7, 0x39, 0xab, 0x67, 0x12, 0x27, 0x94, 0x52, 0x4b, 0x50,
	0xcd, 0x9f, 0xb3, 0xda, 0x0d, 0xcc, 0x29, 0xa5, 0xb6, 0x28, 0xca, 0xde, 0x4f, 0xef, 0x61, 0xbe,
	0x1f, 0x3e, 0x6e, 0x93, 0x6f, 0xc4, 0xf6, 0xd1, 0x4a, 0xed, 0x24, 0x3d, 0xb5, 0x8b, 0x3f, 0xab,
	0x29, 0x82, 0xf5, 0x6b, 0xc4, 0xe1, 0x0f, 0xa1, 0x33, 0x42, 0xfe, 0x2f, 0xc8, 0x9e, 0x64, 0xed,
	0x37, 0xe5, 0xfa, 0x5f, 0xa1, 0x44, 0xe0, 0x97, 0xac, 0x59, 0xae, 0x2a, 0x0d, 0xba, 0x3b, 0x38,
	0xd9, 0x6a, 0xc1, 0x2e, 0x48, 0x7a, 0x5a, 0xbf, 0xfd, 0xd3, 0xa9, 0x88, 0x95, 0xd1, 0xa9, 0xb8,
	0x5d, 0x04, 0xde, 0xdd, 0x22, 0xf0, 0xfe, 0x2e, 0x02, 0xef, 0xc7, 0x32, 0xa8, 0xdc, 0x2d, 0x83,
	0xca, 0xef, 0x65, 0x50, 0xf9, 0xf4, 0x2a, 0x55, 0x38, 0xc9, 0x47, 0xe1, 0xd8, 0xcc, 0xa2, 0x55,
	0x9b, 0x63, 0x63, 0xd3, 0x75, 0x1d, 0x7d, 0x8b, 0x36, 0x3f, 0x12, 0xce, 0x33, 0x70, 0xa3, 0x26,
	0x7d, 0xa4, 0x93, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0x7c, 0xe6, 0x1a, 0xce, 0x03, 0x00,
	0x00,
}

func (m *RegisteredKVQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteredKVQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisteredKVQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastSubmittedResultRemoteHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastSubmittedResultRemoteHeight))
		i--
		dAtA[i] = 0x48
	}
	if m.LastSubmittedResultLocalHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastSubmittedResultLocalHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.LastEmittedHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastEmittedHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.UpdatePeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.UpdatePeriod))
		i--
		dAtA[i] = 0x30
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ZoneId) > 0 {
		i -= len(m.ZoneId)
		copy(dAtA[i:], m.ZoneId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ZoneId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Keys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *KVKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KVKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisteredTXQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteredTXQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisteredTXQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ZoneId) > 0 {
		i -= len(m.ZoneId)
		copy(dAtA[i:], m.ZoneId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ZoneId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TransactionsFilter) > 0 {
		i -= len(m.TransactionsFilter)
		copy(dAtA[i:], m.TransactionsFilter)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TransactionsFilter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisteredKVQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Keys) > 0 {
		for _, e := range m.Keys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.ZoneId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.UpdatePeriod != 0 {
		n += 1 + sovGenesis(uint64(m.UpdatePeriod))
	}
	if m.LastEmittedHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastEmittedHeight))
	}
	if m.LastSubmittedResultLocalHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastSubmittedResultLocalHeight))
	}
	if m.LastSubmittedResultRemoteHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastSubmittedResultRemoteHeight))
	}
	return n
}

func (m *KVKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *RegisteredTXQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.TransactionsFilter)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ZoneId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisteredKVQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisteredKVQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteredKVQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, &KVKey{})
			if err := m.Keys[len(m.Keys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZoneId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatePeriod", wireType)
			}
			m.UpdatePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEmittedHeight", wireType)
			}
			m.LastEmittedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastEmittedHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSubmittedResultLocalHeight", wireType)
			}
			m.LastSubmittedResultLocalHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSubmittedResultLocalHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSubmittedResultRemoteHeight", wireType)
			}
			m.LastSubmittedResultRemoteHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSubmittedResultRemoteHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *KVKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KVKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisteredTXQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisteredTXQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteredTXQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionsFilter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionsFilter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZoneId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
