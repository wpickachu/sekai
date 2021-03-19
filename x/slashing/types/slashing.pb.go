// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slashing.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ValidatorSigningInfo defines a validator's signing info for monitoring their
// liveness activity.
type ValidatorSigningInfo struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// height at which validator was first a candidate OR was activated
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty" yaml:"start_height"`
	// index offset into signed block bit array
	IndexOffset int64 `protobuf:"varint,3,opt,name=index_offset,json=indexOffset,proto3" json:"index_offset,omitempty" yaml:"index_offset"`
	// timestamp validator cannot be activated until
	InactiveUntil time.Time `protobuf:"bytes,4,opt,name=inactive_until,json=inactiveUntil,proto3,stdtime" json:"inactive_until" yaml:"inactive_until"`
	// whether or not a validator has been tombstoned (killed out of validator
	// set)
	Tombstoned bool `protobuf:"varint,5,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	// missed blocks counter (to avoid scanning the array every time)
	MissedBlocksCounter int64 `protobuf:"varint,6,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty" yaml:"missed_blocks_counter"`
}

func (m *ValidatorSigningInfo) Reset()      { *m = ValidatorSigningInfo{} }
func (*ValidatorSigningInfo) ProtoMessage() {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f622956ca78100, []int{0}
}
func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetIndexOffset() int64 {
	if m != nil {
		return m.IndexOffset
	}
	return 0
}

func (m *ValidatorSigningInfo) GetInactiveUntil() time.Time {
	if m != nil {
		return m.InactiveUntil
	}
	return time.Time{}
}

func (m *ValidatorSigningInfo) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

// Params represents the parameters used for by the slashing module.
type Params struct {
	SignedBlocksWindow       int64                                  `protobuf:"varint,1,opt,name=signed_blocks_window,json=signedBlocksWindow,proto3" json:"signed_blocks_window,omitempty" yaml:"signed_blocks_window"`
	MinSignedPerWindow       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=min_signed_per_window,json=minSignedPerWindow,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_signed_per_window" yaml:"min_signed_per_window"`
	DowntimeInactiveDuration time.Duration                          `protobuf:"bytes,3,opt,name=downtime_inactive_duration,json=downtimeInactiveDuration,proto3,stdduration" json:"downtime_inactive_duration" yaml:"downtime_inactive_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_31f622956ca78100, []int{1}
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

func (m *Params) GetSignedBlocksWindow() int64 {
	if m != nil {
		return m.SignedBlocksWindow
	}
	return 0
}

func (m *Params) GetDowntimeInactiveDuration() time.Duration {
	if m != nil {
		return m.DowntimeInactiveDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidatorSigningInfo)(nil), "kira.slashing.ValidatorSigningInfo")
	proto.RegisterType((*Params)(nil), "kira.slashing.Params")
}

func init() { proto.RegisterFile("slashing.proto", fileDescriptor_31f622956ca78100) }

var fileDescriptor_31f622956ca78100 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0xcd, 0x35, 0x50, 0x8a, 0xd3, 0x76, 0x70, 0x5b, 0x61, 0x02, 0xf2, 0xa5, 0x1e, 0x50, 0x84,
	0x54, 0x5b, 0x2a, 0x5b, 0x47, 0xb7, 0x43, 0x2b, 0x24, 0x28, 0x6e, 0x01, 0x89, 0xc5, 0xba, 0xf8,
	0x2e, 0xce, 0x29, 0xf6, 0x5d, 0x74, 0x77, 0x21, 0xed, 0xc8, 0xc2, 0x88, 0x3a, 0x76, 0xec, 0xc8,
	0x4f, 0xe9, 0xd8, 0x11, 0x31, 0x18, 0x94, 0x2c, 0x9d, 0xf3, 0x0b, 0x90, 0xcf, 0x76, 0xb1, 0x4a,
	0x98, 0x92, 0xef, 0xbd, 0xf7, 0xbd, 0xbb, 0xef, 0x7b, 0x67, 0x63, 0x5d, 0x26, 0x48, 0x0e, 0x28,
	0x8b, 0xdd, 0x91, 0xe0, 0x8a, 0x9b, 0x6b, 0x43, 0x2a, 0x90, 0x5b, 0x81, 0xed, 0xcd, 0x98, 0xc7,
	0x5c, 0x33, 0x5e, 0xfe, 0xaf, 0x10, 0xb5, 0xed, 0x98, 0xf3, 0x38, 0x21, 0x9e, 0xae, 0x7a, 0xe3,
	0xbe, 0x87, 0xc7, 0x02, 0x29, 0xca, 0x59, 0xc9, 0xc3, 0xfb, 0xbc, 0xa2, 0x29, 0x91, 0x0a, 0xa5,
	0xa3, 0x42, 0xe0, 0x7c, 0x6b, 0x1a, 0x9b, 0x1f, 0x50, 0x42, 0x31, 0x52, 0x5c, 0x9c, 0xd0, 0x98,
	0x51, 0x16, 0x1f, 0xb1, 0x3e, 0x37, 0x2d, 0xe3, 0x11, 0xc2, 0x58, 0x10, 0x29, 0x2d, 0xd0, 0x01,
	0xdd, 0xc7, 0x41, 0x55, 0x9a, 0x7b, 0xc6, 0xaa, 0x54, 0x48, 0xa8, 0x70, 0x40, 0x68, 0x3c, 0x50,
	0xd6, 0x52, 0x07, 0x74, 0x9b, 0xfe, 0x93, 0x79, 0x06, 0x37, 0xce, 0x51, 0x9a, 0xec, 0x39, 0x75,
	0xd6, 0x09, 0x5a, 0xba, 0x3c, 0xd4, 0x55, 0xde, 0x4b, 0x19, 0x26, 0x67, 0x21, 0xef, 0xf7, 0x25,
	0x51, 0x56, 0xf3, 0x7e, 0x6f, 0x9d, 0x75, 0x82, 0x96, 0x2e, 0xdf, 0xea, 0xca, 0xc4, 0xc6, 0x3a,
	0x65, 0x28, 0x52, 0xf4, 0x33, 0x09, 0xc7, 0x4c, 0xd1, 0xc4, 0x7a, 0xd0, 0x01, 0xdd, 0xd6, 0x6e,
	0xdb, 0x2d, 0x86, 0x74, 0xab, 0x21, 0xdd, 0xd3, 0x6a, 0x48, 0x7f, 0xfb, 0x3a, 0x83, 0x8d, 0x79,
	0x06, 0xb7, 0x2a, 0xf7, 0x7a, 0xbf, 0x73, 0xf1, 0x0b, 0x82, 0x60, 0xad, 0x02, 0xdf, 0xe7, 0x98,
	0x69, 0x1b, 0x86, 0xe2, 0x69, 0x4f, 0x2a, 0xce, 0x08, 0xb6, 0x1e, 0x76, 0x40, 0x77, 0x25, 0xa8,
	0x21, 0xe6, 0xa9, 0xb1, 0x95, 0x52, 0x29, 0x09, 0x0e, 0x7b, 0x09, 0x8f, 0x86, 0x32, 0x8c, 0xf8,
	0x98, 0x29, 0x22, 0xac, 0x65, 0x3d, 0x4a, 0x67, 0x9e, 0xc1, 0xe7, 0xc5, 0x61, 0x0b, 0x65, 0x4e,
	0xb0, 0x51, 0xe0, 0xbe, 0x86, 0xf7, 0x0b, 0x74, 0x6f, 0xe5, 0xf2, 0x0a, 0x36, 0x6e, 0xaf, 0x20,
	0x70, 0x6e, 0x97, 0x8c, 0xe5, 0x63, 0x24, 0x50, 0x2a, 0xcd, 0x77, 0xc6, 0xa6, 0xa4, 0x31, 0xfb,
	0xeb, 0x31, 0xa1, 0x0c, 0xf3, 0x89, 0xce, 0xa3, 0xe9, 0xc3, 0x79, 0x06, 0x9f, 0x95, 0x0b, 0x5f,
	0xa0, 0x72, 0x02, 0xb3, 0x80, 0x8b, 0x83, 0x3e, 0x6a, 0xd0, 0xfc, 0x02, 0xf2, 0xeb, 0xb3, 0xb0,
	0xec, 0x18, 0x11, 0x51, 0x99, 0xe6, 0x29, 0xae, 0xfa, 0x6f, 0xf2, 0x7d, 0xfd, 0xcc, 0xe0, 0x8b,
	0x98, 0xaa, 0xc1, 0xb8, 0xe7, 0x46, 0x3c, 0xf5, 0x22, 0x2e, 0x53, 0x2e, 0xcb, 0x9f, 0x1d, 0x89,
	0x87, 0x9e, 0x3a, 0x1f, 0x11, 0xe9, 0x1e, 0x90, 0xa8, 0x3e, 0xec, 0x02, 0x53, 0x27, 0x30, 0x53,
	0xca, 0x4e, 0x34, 0x7c, 0x4c, 0x44, 0x79, 0x87, 0xaf, 0xc0, 0x68, 0x63, 0x3e, 0x61, 0xf9, 0x53,
	0x0c, 0xef, 0x12, 0xa9, 0x1e, 0xae, 0x7e, 0x12, 0xad, 0xdd, 0xa7, 0xff, 0x84, 0x7a, 0x50, 0x0a,
	0xfc, 0x9d, 0x32, 0xd3, 0xed, 0xe2, 0xe4, 0xff, 0x5b, 0x39, 0x97, 0x79, 0xbe, 0x56, 0x25, 0x38,
	0x2a, 0xf9, 0x3b, 0xa3, 0xc3, 0xef, 0x53, 0x1b, 0x5c, 0x4f, 0x6d, 0x70, 0x33, 0xb5, 0xc1, 0xef,
	0xa9, 0x0d, 0x2e, 0x66, 0x76, 0xe3, 0x66, 0x66, 0x37, 0x7e, 0xcc, 0xec, 0xc6, 0xa7, 0x97, 0xb5,
	0x15, 0xbc, 0xa6, 0x02, 0xed, 0x73, 0x41, 0x3c, 0x49, 0x86, 0x88, 0x7a, 0x67, 0x5e, 0xf5, 0x59,
	0x16, 0xab, 0xe8, 0x2d, 0xeb, 0x5b, 0xbe, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x74, 0x9c, 0x6c,
	0x07, 0xc4, 0x03, 0x00, 0x00,
}

func (this *ValidatorSigningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfo)
	if !ok {
		that2, ok := that.(ValidatorSigningInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if this.IndexOffset != that1.IndexOffset {
		return false
	}
	if !this.InactiveUntil.Equal(that1.InactiveUntil) {
		return false
	}
	if this.Tombstoned != that1.Tombstoned {
		return false
	}
	if this.MissedBlocksCounter != that1.MissedBlocksCounter {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SignedBlocksWindow != that1.SignedBlocksWindow {
		return false
	}
	if !this.MinSignedPerWindow.Equal(that1.MinSignedPerWindow) {
		return false
	}
	if this.DowntimeInactiveDuration != that1.DowntimeInactiveDuration {
		return false
	}
	return true
}
func (m *ValidatorSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x30
	}
	if m.Tombstoned {
		i--
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.InactiveUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.InactiveUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSlashing(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.IndexOffset != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.IndexOffset))
		i--
		dAtA[i] = 0x18
	}
	if m.StartHeight != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSlashing(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DowntimeInactiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeInactiveDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSlashing(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinSignedPerWindow.Size()
		i -= size
		if _, err := m.MinSignedPerWindow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSlashing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SignedBlocksWindow != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.SignedBlocksWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSlashing(dAtA []byte, offset int, v uint64) int {
	offset -= sovSlashing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSlashing(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovSlashing(uint64(m.StartHeight))
	}
	if m.IndexOffset != 0 {
		n += 1 + sovSlashing(uint64(m.IndexOffset))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.InactiveUntil)
	n += 1 + l + sovSlashing(uint64(l))
	if m.Tombstoned {
		n += 2
	}
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovSlashing(uint64(m.MissedBlocksCounter))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedBlocksWindow != 0 {
		n += 1 + sovSlashing(uint64(m.SignedBlocksWindow))
	}
	l = m.MinSignedPerWindow.Size()
	n += 1 + l + sovSlashing(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeInactiveDuration)
	n += 1 + l + sovSlashing(uint64(l))
	return n
}

func sovSlashing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSlashing(x uint64) (n int) {
	return sovSlashing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorSigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexOffset", wireType)
			}
			m.IndexOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexOffset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.InactiveUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlashing
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBlocksWindow", wireType)
			}
			m.SignedBlocksWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBlocksWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSignedPerWindow", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSignedPerWindow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DowntimeInactiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DowntimeInactiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlashing
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
func skipSlashing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
				return 0, ErrInvalidLengthSlashing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSlashing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSlashing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSlashing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlashing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSlashing = fmt.Errorf("proto: unexpected end of group")
)
