// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibb/repay.proto

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

type Repay struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id          uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	BlockHeight int32  `protobuf:"varint,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Asset       string `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	Amount      int32  `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom       string `protobuf:"bytes,6,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *Repay) Reset()         { *m = Repay{} }
func (m *Repay) String() string { return proto.CompactTextString(m) }
func (*Repay) ProtoMessage()    {}
func (*Repay) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ac6f0e0c17a9c4a, []int{0}
}
func (m *Repay) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Repay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Repay.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Repay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repay.Merge(m, src)
}
func (m *Repay) XXX_Size() int {
	return m.Size()
}
func (m *Repay) XXX_DiscardUnknown() {
	xxx_messageInfo_Repay.DiscardUnknown(m)
}

var xxx_messageInfo_Repay proto.InternalMessageInfo

func (m *Repay) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Repay) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Repay) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Repay) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *Repay) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Repay) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*Repay)(nil), "sapienscosmos.ibb.ibb.Repay")
}

func init() { proto.RegisterFile("ibb/repay.proto", fileDescriptor_2ac6f0e0c17a9c4a) }

var fileDescriptor_2ac6f0e0c17a9c4a = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0x7b, 0xea, 0x6d, 0xc5, 0x08, 0x0a, 0xe1, 0x2a, 0xc1, 0x21, 0x14, 0xa7, 0x3a, 0xd8,
	0x0e, 0xbe, 0x81, 0x2e, 0xce, 0x1d, 0xdd, 0x92, 0x36, 0xf4, 0x06, 0x6d, 0x4f, 0x49, 0x72, 0xc1,
	0xfb, 0x16, 0x6e, 0xbe, 0x92, 0xe3, 0x1d, 0x1d, 0xa5, 0x7d, 0x11, 0x49, 0x52, 0xc1, 0x21, 0x21,
	0x5f, 0xf8, 0xfe, 0x03, 0xff, 0x21, 0x97, 0x5a, 0xca, 0xda, 0xa8, 0x49, 0x1c, 0xaa, 0xc9, 0xa0,
	0x43, 0x7a, 0x65, 0xc5, 0xa4, 0xd5, 0x68, 0x5b, 0xb4, 0x03, 0xda, 0x4a, 0x4b, 0xe9, 0xcf, 0xcd,
	0xb6, 0xc7, 0x1e, 0x83, 0x51, 0xfb, 0x57, 0x94, 0x6f, 0x3f, 0x81, 0x64, 0x8d, 0x0f, 0x53, 0x46,
	0x4e, 0x5b, 0xa3, 0x84, 0x43, 0xc3, 0xa0, 0x80, 0xf2, 0xac, 0xf9, 0x43, 0x7a, 0x41, 0x52, 0xdd,
	0xb1, 0xb4, 0x80, 0x72, 0xd3, 0xa4, 0xba, 0xa3, 0x05, 0x39, 0x97, 0x6f, 0xd8, 0xbe, 0x3e, 0x2b,
	0xdd, 0xef, 0x1c, 0x3b, 0x29, 0xa0, 0xcc, 0x9a, 0xff, 0x5f, 0x74, 0x4b, 0x32, 0x61, 0xad, 0x72,
	0x6c, 0x13, 0x26, 0x45, 0xa0, 0xd7, 0x24, 0x17, 0x03, 0xee, 0x47, 0xc7, 0xb2, 0x10, 0x59, 0xc9,
	0xdb, 0x9d, 0x1a, 0x71, 0x60, 0x79, 0xb4, 0x03, 0x3c, 0x3e, 0x7d, 0xcd, 0x1c, 0x8e, 0x33, 0x87,
	0x9f, 0x99, 0xc3, 0xc7, 0xc2, 0x93, 0xe3, 0xc2, 0x93, 0xef, 0x85, 0x27, 0x2f, 0x77, 0xbd, 0x76,
	0xbb, 0xbd, 0xac, 0x5a, 0x1c, 0xea, 0xb5, 0xeb, 0x7d, 0x2c, 0x5b, 0xfb, 0x5d, 0xbc, 0x87, 0xdb,
	0x1d, 0x26, 0x65, 0x65, 0x1e, 0x5a, 0x3e, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x95, 0xc6,
	0x51, 0x25, 0x01, 0x00, 0x00,
}

func (m *Repay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Repay) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Repay) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRepay(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x32
	}
	if m.Amount != 0 {
		i = encodeVarintRepay(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintRepay(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlockHeight != 0 {
		i = encodeVarintRepay(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintRepay(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRepay(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRepay(dAtA []byte, offset int, v uint64) int {
	offset -= sovRepay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Repay) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRepay(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovRepay(uint64(m.Id))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovRepay(uint64(m.BlockHeight))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovRepay(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovRepay(uint64(m.Amount))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRepay(uint64(l))
	}
	return n
}

func sovRepay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRepay(x uint64) (n int) {
	return sovRepay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Repay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRepay
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
			return fmt.Errorf("proto: Repay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Repay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepay
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
				return ErrInvalidLengthRepay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRepay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepay
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepay
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
				return ErrInvalidLengthRepay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRepay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRepay
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
				return ErrInvalidLengthRepay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRepay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRepay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRepay
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
func skipRepay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRepay
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
					return 0, ErrIntOverflowRepay
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
					return 0, ErrIntOverflowRepay
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
				return 0, ErrInvalidLengthRepay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRepay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRepay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRepay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRepay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRepay = fmt.Errorf("proto: unexpected end of group")
)