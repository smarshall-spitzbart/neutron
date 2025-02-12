// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchainqueries/params.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type Params struct {
	// Defines amount of blocks required before query becomes available for
	// removal by anybody
	QuerySubmitTimeout uint64 `protobuf:"varint,1,opt,name=query_submit_timeout,json=querySubmitTimeout,proto3" json:"query_submit_timeout,omitempty"`
	// Amount of coins deposited for the query.
	QueryDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=query_deposit,json=queryDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"query_deposit"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1421c1e223ed164f, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetQuerySubmitTimeout() uint64 {
	if m != nil {
		return m.QuerySubmitTimeout
	}
	return 0
}

func (m *Params) GetQueryDeposit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.QueryDeposit
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "neutron.interchainadapter.interchainqueries.Params")
}

func init() { proto.RegisterFile("interchainqueries/params.proto", fileDescriptor_1421c1e223ed164f) }

var fileDescriptor_1421c1e223ed164f = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x1c, 0xc4, 0xe3, 0xef, 0xab, 0x3a, 0x04, 0x58, 0xa2, 0x0e, 0xa5, 0x83, 0x5b, 0x31, 0x45, 0x42,
	0xb5, 0x5b, 0x58, 0x10, 0x63, 0xe1, 0x01, 0x50, 0x60, 0x62, 0xa9, 0x9c, 0xc4, 0x4a, 0x2d, 0x14,
	0xff, 0x8d, 0xed, 0x20, 0xfa, 0x16, 0x8c, 0x8c, 0xcc, 0x8c, 0x3c, 0x45, 0xc7, 0x8e, 0x4c, 0x80,
	0x92, 0x17, 0x41, 0xb1, 0x83, 0x40, 0xea, 0xe4, 0x93, 0xce, 0x77, 0xbf, 0xbf, 0x2e, 0xc4, 0x42,
	0x5a, 0xae, 0xb3, 0x15, 0x13, 0xf2, 0xbe, 0xe2, 0x5a, 0x70, 0x43, 0x15, 0xd3, 0xac, 0x34, 0x44,
	0x69, 0xb0, 0x10, 0x1d, 0x4b, 0x5e, 0x59, 0x0d, 0x92, 0xfc, 0xfe, 0x63, 0x39, 0x53, 0x96, 0x6b,
	0xb2, 0x93, 0x1c, 0x0d, 0x0a, 0x28, 0xc0, 0xe5, 0x68, 0xab, 0x7c, 0xc5, 0x08, 0x67, 0x60, 0x4a,
	0x30, 0x34, 0x65, 0x86, 0xd3, 0x87, 0x79, 0xca, 0x2d, 0x9b, 0xd3, 0x0c, 0x84, 0xf4, 0xfe, 0xd1,
	0x1b, 0x0a, 0xfb, 0x57, 0x8e, 0x19, 0xcd, 0xc2, 0x41, 0xdb, 0xb5, 0x5e, 0x9a, 0x2a, 0x2d, 0x85,
	0x5d, 0x5a, 0x51, 0x72, 0xa8, 0xec, 0x10, 0x4d, 0x50, 0xdc, 0x4b, 0x22, 0xe7, 0x5d, 0x3b, 0xeb,
	0xc6, 0x3b, 0x91, 0x0a, 0x0f, 0x7c, 0x22, 0xe7, 0x0a, 0x8c, 0xb0, 0xc3, 0x7f, 0x93, 0xff, 0xf1,
	0xde, 0xc9, 0x21, 0xf1, 0x50, 0xd2, 0x42, 0x49, 0x07, 0x25, 0x17, 0x20, 0xe4, 0x62, 0xb6, 0xf9,
	0x18, 0x07, 0xaf, 0x9f, 0xe3, 0xb8, 0x10, 0x76, 0x55, 0xa5, 0x24, 0x83, 0x92, 0x76, 0x17, 0xfa,
	0x67, 0x6a, 0xf2, 0x3b, 0x6a, 0xd7, 0x8a, 0x1b, 0x17, 0x30, 0xc9, 0xbe, 0x23, 0x5c, 0x7a, 0xc0,
	0x79, 0xef, 0xf9, 0x65, 0x1c, 0x2c, 0x92, 0x4d, 0x8d, 0xd1, 0xb6, 0xc6, 0xe8, 0xab, 0xc6, 0xe8,
	0xa9, 0xc1, 0xc1, 0xb6, 0xc1, 0xc1, 0x7b, 0x83, 0x83, 0xdb, 0xb3, 0x3f, 0xbd, 0xdd, 0x78, 0x53,
	0xd0, 0xc5, 0x8f, 0xa6, 0x8f, 0x74, 0x77, 0x72, 0x47, 0x4b, 0xfb, 0x6e, 0x8f, 0xd3, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0xea, 0x48, 0x87, 0x94, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueryDeposit) > 0 {
		for iNdEx := len(m.QueryDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.QueryDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.QuerySubmitTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.QuerySubmitTimeout))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuerySubmitTimeout != 0 {
		n += 1 + sovParams(uint64(m.QuerySubmitTimeout))
	}
	if len(m.QueryDeposit) > 0 {
		for _, e := range m.QueryDeposit {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuerySubmitTimeout", wireType)
			}
			m.QuerySubmitTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuerySubmitTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryDeposit = append(m.QueryDeposit, types.Coin{})
			if err := m.QueryDeposit[len(m.QueryDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
