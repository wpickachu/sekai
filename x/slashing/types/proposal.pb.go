// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proposal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type MsgProposalUnjailValidator struct {
	Symbol   string                                        `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name     string                                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon     string                                        `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Decimals uint32                                        `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Denoms   []string                                      `protobuf:"bytes,5,rep,name=denoms,proto3" json:"denoms,omitempty"`
	Proposer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=proposer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"proposer,omitempty" yaml:"proposer"`
}

func (m *MsgProposalUnjailValidator) Reset()         { *m = MsgProposalUnjailValidator{} }
func (m *MsgProposalUnjailValidator) String() string { return proto.CompactTextString(m) }
func (*MsgProposalUnjailValidator) ProtoMessage()    {}
func (*MsgProposalUnjailValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3ac5ce23bf32d05, []int{0}
}
func (m *MsgProposalUnjailValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalUnjailValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalUnjailValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalUnjailValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalUnjailValidator.Merge(m, src)
}
func (m *MsgProposalUnjailValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalUnjailValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalUnjailValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalUnjailValidator proto.InternalMessageInfo

func (m *MsgProposalUnjailValidator) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MsgProposalUnjailValidator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgProposalUnjailValidator) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *MsgProposalUnjailValidator) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *MsgProposalUnjailValidator) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

func (m *MsgProposalUnjailValidator) GetProposer() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Proposer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgProposalUnjailValidator)(nil), "kira.slashing.MsgProposalUnjailValidator")
}

func init() { proto.RegisterFile("proposal.proto", fileDescriptor_c3ac5ce23bf32d05) }

var fileDescriptor_c3ac5ce23bf32d05 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4b, 0xfb, 0x40,
	0x18, 0xc7, 0x7b, 0xbf, 0xf6, 0x57, 0x6a, 0xb0, 0x0a, 0x41, 0x24, 0x66, 0xb8, 0x86, 0x4e, 0x45,
	0x68, 0x32, 0xb8, 0xb9, 0xb5, 0x76, 0x13, 0x41, 0x0a, 0x3a, 0xb8, 0xe8, 0xf5, 0xee, 0x48, 0xcf,
	0xde, 0xe5, 0x09, 0xf7, 0x44, 0xb0, 0xef, 0xc2, 0x97, 0xe5, 0xd8, 0xd1, 0xa9, 0x48, 0xfb, 0x0e,
	0x3a, 0x3a, 0x49, 0x72, 0x69, 0x71, 0xba, 0xef, 0x9f, 0xe3, 0x03, 0xcf, 0xd7, 0x3b, 0xc9, 0x2d,
	0xe4, 0x80, 0x4c, 0xc7, 0xb9, 0x85, 0x02, 0xfc, 0xee, 0x42, 0x59, 0x16, 0xa3, 0x66, 0x38, 0x57,
	0x59, 0x1a, 0x9e, 0xa5, 0x90, 0x42, 0xd5, 0x24, 0xa5, 0x72, 0x9f, 0xc2, 0x0b, 0x0e, 0x68, 0x00,
	0x9f, 0x5d, 0xe1, 0x8c, 0xab, 0xfa, 0x3b, 0xe2, 0x85, 0x77, 0x98, 0xde, 0xd7, 0xd4, 0x87, 0xec,
	0x95, 0x29, 0xfd, 0xc8, 0xb4, 0x12, 0xac, 0x00, 0xeb, 0x9f, 0x7b, 0x6d, 0x5c, 0x9a, 0x19, 0xe8,
	0x80, 0x44, 0x64, 0x70, 0x34, 0xad, 0x9d, 0xef, 0x7b, 0xad, 0x8c, 0x19, 0x19, 0xfc, 0xab, 0xd2,
	0x4a, 0x97, 0x99, 0xe2, 0x90, 0x05, 0x4d, 0x97, 0x95, 0xda, 0x0f, 0xbd, 0x8e, 0x90, 0x5c, 0x19,
	0xa6, 0x31, 0x68, 0x45, 0x64, 0xd0, 0x9d, 0x1e, 0x7c, 0xc9, 0x16, 0x32, 0x03, 0x83, 0xc1, 0xff,
	0xa8, 0x59, 0xb2, 0x9d, 0xf3, 0x5f, 0xbc, 0x8e, 0x3b, 0x52, 0xda, 0xa0, 0x1d, 0x91, 0xc1, 0xf1,
	0x78, 0xb2, 0x5b, 0xf7, 0x4e, 0x97, 0xcc, 0xe8, 0xeb, 0xfe, 0xbe, 0xe9, 0xff, 0xac, 0x7b, 0xc3,
	0x54, 0x15, 0xf3, 0xb7, 0x59, 0xcc, 0xc1, 0xd4, 0x47, 0xd5, 0xcf, 0x10, 0xc5, 0x22, 0x29, 0x96,
	0xb9, 0xc4, 0x78, 0xc4, 0xf9, 0x48, 0x08, 0x2b, 0x11, 0xa7, 0x07, 0xea, 0x78, 0xf2, 0xb9, 0xa1,
	0x64, 0xb5, 0xa1, 0xe4, 0x7b, 0x43, 0xc9, 0xc7, 0x96, 0x36, 0x56, 0x5b, 0xda, 0xf8, 0xda, 0xd2,
	0xc6, 0xd3, 0xe5, 0x1f, 0xe4, 0xad, 0xb2, 0xec, 0x06, 0xac, 0x4c, 0x50, 0x2e, 0x98, 0x4a, 0xde,
	0x93, 0xfd, 0xca, 0x0e, 0x3d, 0x6b, 0x57, 0x0b, 0x5e, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x91,
	0xce, 0x77, 0xb6, 0x93, 0x01, 0x00, 0x00,
}

func (m *MsgProposalUnjailValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalUnjailValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalUnjailValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Decimals != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Icon) > 0 {
		i -= len(m.Icon)
		copy(dAtA[i:], m.Icon)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Icon)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgProposalUnjailValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovProposal(uint64(m.Decimals))
	}
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgProposalUnjailValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: MsgProposalUnjailValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalUnjailValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Icon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = append(m.Proposer[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposer == nil {
				m.Proposer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProposal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)