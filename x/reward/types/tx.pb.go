// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/reward/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type RefundMsgRequest struct {
	Sender       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	InnerMessage *types.Any                                    `protobuf:"bytes,2,opt,name=inner_message,json=innerMessage,proto3" json:"inner_message,omitempty"`
}

func (m *RefundMsgRequest) Reset()         { *m = RefundMsgRequest{} }
func (m *RefundMsgRequest) String() string { return proto.CompactTextString(m) }
func (*RefundMsgRequest) ProtoMessage()    {}
func (*RefundMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b096b8b34fa76c5c, []int{0}
}
func (m *RefundMsgRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RefundMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RefundMsgRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RefundMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundMsgRequest.Merge(m, src)
}
func (m *RefundMsgRequest) XXX_Size() int {
	return m.Size()
}
func (m *RefundMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefundMsgRequest proto.InternalMessageInfo

type RefundMsgResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Log  string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *RefundMsgResponse) Reset()         { *m = RefundMsgResponse{} }
func (m *RefundMsgResponse) String() string { return proto.CompactTextString(m) }
func (*RefundMsgResponse) ProtoMessage()    {}
func (*RefundMsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b096b8b34fa76c5c, []int{1}
}
func (m *RefundMsgResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RefundMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RefundMsgResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RefundMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundMsgResponse.Merge(m, src)
}
func (m *RefundMsgResponse) XXX_Size() int {
	return m.Size()
}
func (m *RefundMsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundMsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefundMsgResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RefundMsgRequest)(nil), "axelar.reward.v1beta1.RefundMsgRequest")
	proto.RegisterType((*RefundMsgResponse)(nil), "axelar.reward.v1beta1.RefundMsgResponse")
}

func init() { proto.RegisterFile("axelar/reward/v1beta1/tx.proto", fileDescriptor_b096b8b34fa76c5c) }

var fileDescriptor_b096b8b34fa76c5c = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x8d, 0xe1, 0xa9, 0x12, 0xa6, 0xa0, 0x12, 0x15, 0xa9, 0xed, 0x60, 0xaa, 0x4e, 0x5d, 0x62,
	0xab, 0x74, 0x82, 0xad, 0xdd, 0x3a, 0x74, 0x20, 0x23, 0x4b, 0xe5, 0x24, 0xb7, 0x26, 0x6a, 0x62,
	0x07, 0xdb, 0xa1, 0xe9, 0xc6, 0x27, 0xf0, 0x25, 0x4c, 0xf0, 0x0f, 0x15, 0x53, 0x47, 0x26, 0x04,
	0xed, 0x5f, 0x30, 0x3d, 0x35, 0x76, 0xa5, 0x4e, 0x3e, 0xd7, 0xe7, 0x9e, 0x7b, 0xcf, 0xb1, 0x31,
	0xe1, 0x0d, 0x14, 0x5c, 0x33, 0x0d, 0x7b, 0xae, 0x33, 0xf6, 0x65, 0x96, 0x80, 0xe5, 0x33, 0x66,
	0x1b, 0x5a, 0x69, 0x65, 0x55, 0xf8, 0xda, 0xf1, 0xd4, 0xf1, 0xd4, 0xf3, 0xa3, 0xa1, 0x50, 0x4a,
	0x14, 0xc0, 0xda, 0xa6, 0xa4, 0xde, 0x32, 0x2e, 0x0f, 0x4e, 0x31, 0xea, 0x0b, 0x25, 0x54, 0x0b,
	0xd9, 0x15, 0xf9, 0xdb, 0x61, 0xaa, 0x4c, 0xa9, 0xcc, 0xc6, 0x11, 0xae, 0xf0, 0x14, 0xf5, 0x16,
	0x2a, 0xd0, 0x65, 0x6e, 0x4c, 0xae, 0x24, 0x83, 0xa6, 0x52, 0xda, 0xc2, 0x9d, 0x9f, 0x43, 0x05,
	0xbe, 0x7f, 0xf2, 0x1d, 0xe1, 0x5e, 0x0c, 0xdb, 0x5a, 0x66, 0x6b, 0x23, 0x62, 0xf8, 0x5c, 0x83,
	0xb1, 0xe1, 0x0a, 0x77, 0x0c, 0xc8, 0x0c, 0xf4, 0x00, 0x8d, 0xd1, 0xb4, 0xbb, 0x9c, 0xfd, 0xff,
	0xf3, 0x26, 0x12, 0xb9, 0xfd, 0x54, 0x27, 0x34, 0x55, 0xa5, 0xdf, 0xe8, 0x8f, 0xc8, 0x64, 0x3b,
	0x3f, 0x72, 0x91, 0xa6, 0x8b, 0x2c, 0xd3, 0x60, 0x4c, 0xec, 0x07, 0x84, 0x2b, 0xfc, 0x22, 0x97,
	0x12, 0xf4, 0xa6, 0x04, 0x63, 0xb8, 0x80, 0xc1, 0x93, 0x31, 0x9a, 0x3e, 0x7f, 0xdb, 0xa7, 0x2e,
	0x33, 0xbd, 0x65, 0xa6, 0x0b, 0x79, 0x58, 0xbe, 0xfc, 0xf5, 0x23, 0xc2, 0xce, 0x0d, 0x4f, 0x0a,
	0x88, 0xbb, 0xad, 0x74, 0xed, 0x94, 0xef, 0x1f, 0xbe, 0xfe, 0x1c, 0xa0, 0xc9, 0x3b, 0xfc, 0xea,
	0xce, 0xaf, 0xa9, 0x94, 0x34, 0x10, 0x86, 0xf8, 0x21, 0xe3, 0x96, 0x3b, 0xbb, 0x71, 0x8b, 0xc3,
	0x1e, 0x7e, 0x5a, 0x28, 0xd1, 0xee, 0x7b, 0x16, 0x5f, 0xe1, 0xf2, 0xc3, 0xf1, 0x1f, 0x09, 0x8e,
	0x67, 0x82, 0x4e, 0x67, 0x82, 0xfe, 0x9e, 0x09, 0xfa, 0x76, 0x21, 0xc1, 0xe9, 0x42, 0x82, 0xdf,
	0x17, 0x12, 0x7c, 0x9c, 0xdf, 0x05, 0x74, 0x8f, 0x28, 0xc1, 0xee, 0x95, 0xde, 0xf9, 0x2a, 0x4a,
	0x95, 0x06, 0xd6, 0xdc, 0x3e, 0xb7, 0x4d, 0x9c, 0x74, 0x5a, 0xff, 0xf3, 0xc7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc0, 0xef, 0x86, 0x70, 0xfa, 0x01, 0x00, 0x00,
}

func (m *RefundMsgRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RefundMsgRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefundMsgRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InnerMessage != nil {
		{
			size, err := m.InnerMessage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RefundMsgResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RefundMsgResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefundMsgResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RefundMsgRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.InnerMessage != nil {
		l = m.InnerMessage.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *RefundMsgResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RefundMsgRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RefundMsgRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefundMsgRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InnerMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InnerMessage == nil {
				m.InnerMessage = &types.Any{}
			}
			if err := m.InnerMessage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *RefundMsgResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RefundMsgResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefundMsgResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
