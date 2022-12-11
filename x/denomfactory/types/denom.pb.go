// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: g4alchain/denomfactory/denom.proto

package types

import (
	fmt "fmt"
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

type Denom struct {
	Symbol             string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Project            string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	MaxSupply          int32  `protobuf:"varint,3,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	CanChangeMaxSupply bool   `protobuf:"varint,4,opt,name=canChangeMaxSupply,proto3" json:"canChangeMaxSupply,omitempty"`
	Creator            string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Denom) Reset()         { *m = Denom{} }
func (m *Denom) String() string { return proto.CompactTextString(m) }
func (*Denom) ProtoMessage()    {}
func (*Denom) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf14a73241a9ab29, []int{0}
}
func (m *Denom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Denom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Denom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Denom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Denom.Merge(m, src)
}
func (m *Denom) XXX_Size() int {
	return m.Size()
}
func (m *Denom) XXX_DiscardUnknown() {
	xxx_messageInfo_Denom.DiscardUnknown(m)
}

var xxx_messageInfo_Denom proto.InternalMessageInfo

func (m *Denom) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Denom) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Denom) GetMaxSupply() int32 {
	if m != nil {
		return m.MaxSupply
	}
	return 0
}

func (m *Denom) GetCanChangeMaxSupply() bool {
	if m != nil {
		return m.CanChangeMaxSupply
	}
	return false
}

func (m *Denom) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Denom)(nil), "g4alentertainment.g4alchain.denomfactory.Denom")
}

func init() {
	proto.RegisterFile("g4alchain/denomfactory/denom.proto", fileDescriptor_bf14a73241a9ab29)
}

var fileDescriptor_bf14a73241a9ab29 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x37, 0x49, 0xcc,
	0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x4f, 0x49, 0xcd, 0xcb, 0xcf, 0x4d, 0x4b, 0x4c, 0x2e, 0xc9,
	0x2f, 0xaa, 0x84, 0x70, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x34, 0x40, 0x6a, 0x52, 0xf3,
	0x4a, 0x52, 0x8b, 0x4a, 0x12, 0x33, 0xf3, 0x72, 0x53, 0xf3, 0x4a, 0xf4, 0xe0, 0xba, 0xf4, 0x90,
	0x75, 0x29, 0x2d, 0x64, 0xe4, 0x62, 0x75, 0x01, 0x09, 0x08, 0x89, 0x71, 0xb1, 0x15, 0x57, 0xe6,
	0x26, 0xe5, 0xe7, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec,
	0x05, 0x45, 0xf9, 0x59, 0xa9, 0xc9, 0x25, 0x12, 0x4c, 0x60, 0x09, 0x18, 0x57, 0x48, 0x86, 0x8b,
	0x33, 0x37, 0xb1, 0x22, 0xb8, 0xb4, 0xa0, 0x20, 0xa7, 0x52, 0x82, 0x59, 0x81, 0x51, 0x83, 0x35,
	0x08, 0x21, 0x20, 0xa4, 0xc7, 0x25, 0x94, 0x9c, 0x98, 0xe7, 0x9c, 0x91, 0x98, 0x97, 0x9e, 0xea,
	0x0b, 0x57, 0xc6, 0xa2, 0xc0, 0xa8, 0xc1, 0x11, 0x84, 0x45, 0x06, 0x64, 0x4f, 0x72, 0x51, 0x6a,
	0x62, 0x49, 0x7e, 0x91, 0x04, 0x2b, 0xc4, 0x1e, 0x28, 0xd7, 0x29, 0xfc, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x6c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0xdd, 0x4d, 0x1c, 0x7d, 0x74, 0x5d, 0x91, 0xfd, 0xac, 0x0f, 0xf2, 0xb3, 0x2e, 0x24,
	0xa8, 0x2a, 0x50, 0x03, 0xab, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x5a, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xa4, 0x5d, 0xb9, 0x53, 0x01, 0x00, 0x00,
}

func (m *Denom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Denom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Denom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CanChangeMaxSupply {
		i--
		if m.CanChangeMaxSupply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.MaxSupply != 0 {
		i = encodeVarintDenom(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Project) > 0 {
		i -= len(m.Project)
		copy(dAtA[i:], m.Project)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Project)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintDenom(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDenom(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Denom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	l = len(m.Project)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	if m.MaxSupply != 0 {
		n += 1 + sovDenom(uint64(m.MaxSupply))
	}
	if m.CanChangeMaxSupply {
		n += 2
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDenom(uint64(l))
	}
	return n
}

func sovDenom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenom(x uint64) (n int) {
	return sovDenom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Denom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenom
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
			return fmt.Errorf("proto: Denom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Denom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Project", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Project = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanChangeMaxSupply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
			m.CanChangeMaxSupply = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenom
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
				return ErrInvalidLengthDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenom
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
func skipDenom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenom
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
					return 0, ErrIntOverflowDenom
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
					return 0, ErrIntOverflowDenom
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
				return 0, ErrInvalidLengthDenom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenom = fmt.Errorf("proto: unexpected end of group")
)