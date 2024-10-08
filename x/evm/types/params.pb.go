// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/evm/v1beta1/params.proto

package types

import (
	fmt "fmt"
	utils "github.com/scalarorg/xchains-indexer/util"
	exported "github.com/scalarorg/xchains-indexer/x/nexus/exported"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/scalarorg/xchains-indexer/x/nexus/exported"
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

// Params is the parameter set for this module
type Params struct {
	Chain               github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	ConfirmationHeight  uint64                                                          `protobuf:"varint,2,opt,name=confirmation_height,json=confirmationHeight,proto3" json:"confirmation_height,omitempty"`
	Network             string                                                          `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	TokenCode           []byte                                                          `protobuf:"bytes,5,opt,name=token_code,json=tokenCode,proto3" json:"token_code,omitempty"`
	Burnable            []byte                                                          `protobuf:"bytes,6,opt,name=burnable,proto3" json:"burnable,omitempty"`
	RevoteLockingPeriod int64                                                           `protobuf:"varint,7,opt,name=revote_locking_period,json=revoteLockingPeriod,proto3" json:"revote_locking_period,omitempty"`
	Networks            []NetworkInfo                                                   `protobuf:"bytes,8,rep,name=networks,proto3" json:"networks"`
	VotingThreshold     utils.Threshold                                                 `protobuf:"bytes,9,opt,name=voting_threshold,json=votingThreshold,proto3" json:"voting_threshold"`
	MinVoterCount       int64                                                           `protobuf:"varint,10,opt,name=min_voter_count,json=minVoterCount,proto3" json:"min_voter_count,omitempty"`
	CommandsGasLimit    uint32                                                          `protobuf:"varint,11,opt,name=commands_gas_limit,json=commandsGasLimit,proto3" json:"commands_gas_limit,omitempty"`
	VotingGracePeriod   int64                                                           `protobuf:"varint,13,opt,name=voting_grace_period,json=votingGracePeriod,proto3" json:"voting_grace_period,omitempty"`
	EndBlockerLimit     int64                                                           `protobuf:"varint,14,opt,name=end_blocker_limit,json=endBlockerLimit,proto3" json:"end_blocker_limit,omitempty"`
	TransferLimit       uint64                                                          `protobuf:"varint,15,opt,name=transfer_limit,json=transferLimit,proto3" json:"transfer_limit,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8097febe234aa9c, []int{0}
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

type PendingChain struct {
	Params Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Chain  exported.Chain `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain"`
}

func (m *PendingChain) Reset()         { *m = PendingChain{} }
func (m *PendingChain) String() string { return proto.CompactTextString(m) }
func (*PendingChain) ProtoMessage()    {}
func (*PendingChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8097febe234aa9c, []int{1}
}
func (m *PendingChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingChain.Merge(m, src)
}
func (m *PendingChain) XXX_Size() int {
	return m.Size()
}
func (m *PendingChain) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingChain.DiscardUnknown(m)
}

var xxx_messageInfo_PendingChain proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "axelar.evm.v1beta1.Params")
	proto.RegisterType((*PendingChain)(nil), "axelar.evm.v1beta1.PendingChain")
}

func init() { proto.RegisterFile("axelar/evm/v1beta1/params.proto", fileDescriptor_c8097febe234aa9c) }

var fileDescriptor_c8097febe234aa9c = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x6f, 0xb6, 0xb6, 0xeb, 0xbc, 0x75, 0xed, 0x3c, 0x90, 0xa2, 0x4a, 0xa4, 0xd1, 0x34, 0x50,
	0x41, 0x90, 0xb0, 0x71, 0xe1, 0x06, 0xb4, 0x87, 0xc1, 0x34, 0x4d, 0x55, 0x84, 0x90, 0xe0, 0x12,
	0xb9, 0xc9, 0x5b, 0x6a, 0x2d, 0xb1, 0x2b, 0xc7, 0x2d, 0xe5, 0x2b, 0x70, 0xe2, 0xc6, 0x57, 0xda,
	0x71, 0x47, 0x4e, 0x13, 0x6c, 0xdf, 0x82, 0x13, 0x8a, 0xed, 0x44, 0x13, 0xf4, 0xc0, 0x2d, 0xfe,
	0xfd, 0x7b, 0x8e, 0xdf, 0x7b, 0xa8, 0x4f, 0x96, 0x90, 0x12, 0xe1, 0xc3, 0x22, 0xf3, 0x17, 0x87,
	0x13, 0x90, 0xe4, 0xd0, 0x9f, 0x11, 0x41, 0xb2, 0xdc, 0x9b, 0x09, 0x2e, 0x39, 0xc6, 0x5a, 0xe0,
	0xc1, 0x22, 0xf3, 0x8c, 0xa0, 0x77, 0x60, 0x4c, 0x73, 0x49, 0xd3, 0xbc, 0xb2, 0xc9, 0xa9, 0x80,
	0x7c, 0xca, 0xd3, 0x58, 0x3b, 0x7b, 0xce, 0x8a, 0x68, 0xf9, 0x65, 0x06, 0x26, 0xb9, 0x77, 0x2f,
	0xe1, 0x09, 0x57, 0x9f, 0x7e, 0xf1, 0x65, 0xd0, 0xc7, 0xc6, 0xc5, 0x60, 0x39, 0xcf, 0x7d, 0x58,
	0xce, 0xb8, 0x90, 0x10, 0xaf, 0x0a, 0xd8, 0xff, 0xde, 0x40, 0xcd, 0xb1, 0xba, 0x2b, 0xfe, 0x88,
	0x1a, 0xd1, 0x94, 0x50, 0x66, 0x5b, 0xae, 0x35, 0xd8, 0x1c, 0x8e, 0x7e, 0x5f, 0xf7, 0x5f, 0x25,
	0x54, 0x4e, 0xe7, 0x13, 0x2f, 0xe2, 0x99, 0xaf, 0x33, 0x19, 0xc8, 0xcf, 0x5c, 0x5c, 0x98, 0xd3,
	0xb3, 0x88, 0x0b, 0xf0, 0x97, 0x7f, 0x15, 0xf2, 0x46, 0x45, 0xcc, 0x19, 0xc9, 0x20, 0xd0, 0x89,
	0xd8, 0x47, 0x7b, 0x11, 0x67, 0xe7, 0x54, 0x64, 0x44, 0x52, 0xce, 0xc2, 0x29, 0xd0, 0x64, 0x2a,
	0xed, 0x35, 0xd7, 0x1a, 0xd4, 0x03, 0x7c, 0x97, 0x7a, 0xab, 0x18, 0x6c, 0xa3, 0x0d, 0x53, 0xc9,
	0x5e, 0x2f, 0x6e, 0x13, 0x94, 0x47, 0xfc, 0x00, 0x21, 0xc9, 0x2f, 0x80, 0x85, 0x11, 0x8f, 0xc1,
	0x6e, 0xb8, 0xd6, 0x60, 0x3b, 0xd8, 0x54, 0xc8, 0x88, 0xc7, 0x80, 0x7b, 0xa8, 0x35, 0x99, 0x0b,
	0x46, 0x26, 0x29, 0xd8, 0x4d, 0x45, 0x56, 0x67, 0x7c, 0x84, 0xee, 0x0b, 0x58, 0x70, 0x09, 0x61,
	0xca, 0xa3, 0x0b, 0xca, 0x92, 0x70, 0x06, 0x82, 0xf2, 0xd8, 0xde, 0x70, 0xad, 0xc1, 0x7a, 0xb0,
	0xa7, 0xc9, 0x53, 0xcd, 0x8d, 0x15, 0x85, 0xdf, 0xa0, 0x96, 0xa9, 0x9c, 0xdb, 0x2d, 0x77, 0x7d,
	0xb0, 0x75, 0xd4, 0xf7, 0xfe, 0xed, 0xa6, 0x77, 0xa6, 0x35, 0xef, 0xd8, 0x39, 0x1f, 0xd6, 0x2f,
	0xaf, 0xfb, 0xb5, 0xa0, 0xb2, 0xe1, 0x31, 0xea, 0x2e, 0xb8, 0x2c, 0xca, 0x55, 0xdd, 0xb5, 0x37,
	0x5d, 0xeb, 0x6e, 0x94, 0x1a, 0x82, 0x2a, 0xec, 0x7d, 0x29, 0x33, 0x51, 0x1d, 0x6d, 0xaf, 0x60,
	0xfc, 0x08, 0x75, 0x32, 0xca, 0xc2, 0xe2, 0xb6, 0x22, 0x8c, 0xf8, 0x9c, 0x49, 0x1b, 0xa9, 0x5f,
	0x68, 0x67, 0x94, 0x7d, 0x28, 0xd0, 0x51, 0x01, 0xe2, 0xa7, 0x08, 0x47, 0x3c, 0xcb, 0x08, 0x8b,
	0xf3, 0x30, 0x21, 0x79, 0x98, 0xd2, 0x8c, 0x4a, 0x7b, 0xcb, 0xb5, 0x06, 0xed, 0xa0, 0x5b, 0x32,
	0xc7, 0x24, 0x3f, 0x2d, 0x70, 0xec, 0xa1, 0x3d, 0x73, 0xcf, 0x44, 0x90, 0x08, 0xca, 0xc7, 0x69,
	0xab, 0xe4, 0x5d, 0x4d, 0x1d, 0x17, 0x8c, 0x79, 0x9a, 0x27, 0x68, 0x17, 0x58, 0x1c, 0x4e, 0x8a,
	0xc7, 0x04, 0x61, 0xc2, 0x77, 0x94, 0xba, 0x03, 0x2c, 0x1e, 0x6a, 0x5c, 0x67, 0x3f, 0x44, 0x3b,
	0x52, 0x10, 0x96, 0x9f, 0x57, 0xc2, 0x8e, 0xea, 0x7d, 0xbb, 0x44, 0x95, 0xec, 0xa4, 0xde, 0xaa,
	0x77, 0x1b, 0x27, 0xf5, 0xd6, 0x76, 0xb7, 0xbd, 0xff, 0xd5, 0x42, 0xdb, 0x63, 0x60, 0x31, 0x65,
	0x89, 0x9a, 0x27, 0xfc, 0x12, 0x35, 0xf5, 0x56, 0xa9, 0x01, 0xdd, 0x3a, 0xea, 0xad, 0x6a, 0x84,
	0x9e, 0x65, 0xf3, 0x70, 0x46, 0x8f, 0x5f, 0x97, 0x93, 0xbd, 0xa6, 0x8c, 0x07, 0xa5, 0x51, 0x8d,
	0xad, 0x57, 0x8d, 0x6d, 0x99, 0xa1, 0xca, 0x99, 0x08, 0x6d, 0x1c, 0x9e, 0x5d, 0xfe, 0x72, 0x6a,
	0x97, 0x37, 0x8e, 0x75, 0x75, 0xe3, 0x58, 0x3f, 0x6f, 0x1c, 0xeb, 0xdb, 0xad, 0x53, 0xbb, 0xba,
	0x75, 0x6a, 0x3f, 0x6e, 0x9d, 0xda, 0xa7, 0xe7, 0xff, 0xb9, 0x26, 0xc5, 0x16, 0xab, 0xe5, 0x9b,
	0x34, 0xd5, 0xf6, 0xbd, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xec, 0x46, 0xee, 0x1d, 0x3b, 0x04,
	0x00, 0x00,
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
	if m.TransferLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TransferLimit))
		i--
		dAtA[i] = 0x78
	}
	if m.EndBlockerLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockerLimit))
		i--
		dAtA[i] = 0x70
	}
	if m.VotingGracePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VotingGracePeriod))
		i--
		dAtA[i] = 0x68
	}
	if m.CommandsGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CommandsGasLimit))
		i--
		dAtA[i] = 0x58
	}
	if m.MinVoterCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinVoterCount))
		i--
		dAtA[i] = 0x50
	}
	{
		size, err := m.VotingThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Networks) > 0 {
		for iNdEx := len(m.Networks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Networks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.RevoteLockingPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RevoteLockingPeriod))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Burnable) > 0 {
		i -= len(m.Burnable)
		copy(dAtA[i:], m.Burnable)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Burnable)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenCode) > 0 {
		i -= len(m.TokenCode)
		copy(dAtA[i:], m.TokenCode)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TokenCode)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ConfirmationHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConfirmationHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PendingChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Chain.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.ConfirmationHeight != 0 {
		n += 1 + sovParams(uint64(m.ConfirmationHeight))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TokenCode)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Burnable)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.RevoteLockingPeriod != 0 {
		n += 1 + sovParams(uint64(m.RevoteLockingPeriod))
	}
	if len(m.Networks) > 0 {
		for _, e := range m.Networks {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.VotingThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MinVoterCount != 0 {
		n += 1 + sovParams(uint64(m.MinVoterCount))
	}
	if m.CommandsGasLimit != 0 {
		n += 1 + sovParams(uint64(m.CommandsGasLimit))
	}
	if m.VotingGracePeriod != 0 {
		n += 1 + sovParams(uint64(m.VotingGracePeriod))
	}
	if m.EndBlockerLimit != 0 {
		n += 1 + sovParams(uint64(m.EndBlockerLimit))
	}
	if m.TransferLimit != 0 {
		n += 1 + sovParams(uint64(m.TransferLimit))
	}
	return n
}

func (m *PendingChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Chain.Size()
	n += 1 + l + sovParams(uint64(l))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationHeight", wireType)
			}
			m.ConfirmationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfirmationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenCode = append(m.TokenCode[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenCode == nil {
				m.TokenCode = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burnable", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Burnable = append(m.Burnable[:0], dAtA[iNdEx:postIndex]...)
			if m.Burnable == nil {
				m.Burnable = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevoteLockingPeriod", wireType)
			}
			m.RevoteLockingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevoteLockingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Networks", wireType)
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
			m.Networks = append(m.Networks, NetworkInfo{})
			if err := m.Networks[len(m.Networks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingThreshold", wireType)
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
			if err := m.VotingThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinVoterCount", wireType)
			}
			m.MinVoterCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinVoterCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandsGasLimit", wireType)
			}
			m.CommandsGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommandsGasLimit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingGracePeriod", wireType)
			}
			m.VotingGracePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingGracePeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockerLimit", wireType)
			}
			m.EndBlockerLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockerLimit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferLimit", wireType)
			}
			m.TransferLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *PendingChain) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PendingChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
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
			if err := m.Chain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
