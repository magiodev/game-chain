// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: g4alchain/permission/developer.proto

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

type Developer struct {
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Blocked   bool   `protobuf:"varint,4,opt,name=blocked,proto3" json:"blocked,omitempty"`
	Creator   string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Developer) Reset()         { *m = Developer{} }
func (m *Developer) String() string { return proto.CompactTextString(m) }
func (*Developer) ProtoMessage()    {}
func (*Developer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5cf969a4df33bf9, []int{0}
}
func (m *Developer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Developer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Developer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Developer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Developer.Merge(m, src)
}
func (m *Developer) XXX_Size() int {
	return m.Size()
}
func (m *Developer) XXX_DiscardUnknown() {
	xxx_messageInfo_Developer.DiscardUnknown(m)
}

var xxx_messageInfo_Developer proto.InternalMessageInfo

func (m *Developer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Developer) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Developer) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Developer) GetBlocked() bool {
	if m != nil {
		return m.Blocked
	}
	return false
}

func (m *Developer) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Developer)(nil), "g4alentertainment.g4alchain.permission.Developer")
}

func init() {
	proto.RegisterFile("g4alchain/permission/developer.proto", fileDescriptor_a5cf969a4df33bf9)
}

var fileDescriptor_a5cf969a4df33bf9 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x37, 0x49, 0xcc,
	0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x48, 0x2d, 0xca, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf,
	0xd3, 0x4f, 0x49, 0x2d, 0x4b, 0xcd, 0xc9, 0x2f, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x03, 0xa9, 0x4a, 0xcd, 0x2b, 0x49, 0x2d, 0x2a, 0x49, 0xcc, 0xcc, 0xcb, 0x4d, 0xcd,
	0x2b, 0xd1, 0x83, 0xeb, 0xd3, 0x43, 0xe8, 0x53, 0x9a, 0xca, 0xc8, 0xc5, 0xe9, 0x02, 0xd3, 0x2b,
	0x24, 0xc1, 0xc5, 0x9e, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0xe3, 0x0a, 0xc9, 0x70, 0x71, 0x26, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0xa6, 0x38, 0x96,
	0x48, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07, 0x21, 0x04, 0x40, 0xb2, 0xa5, 0x05, 0x29, 0x50, 0x59,
	0x66, 0x88, 0x2c, 0x5c, 0x00, 0x64, 0x6a, 0x52, 0x4e, 0x7e, 0x72, 0x76, 0x6a, 0x8a, 0x04, 0x8b,
	0x02, 0xa3, 0x06, 0x47, 0x10, 0x8c, 0x0b, 0x92, 0x01, 0x1b, 0x92, 0x5f, 0x24, 0xc1, 0x0a, 0xb1,
	0x0f, 0xca, 0x75, 0x0a, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xeb,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x77, 0x13, 0x47, 0x1f, 0x5d,
	0x57, 0x64, 0x5f, 0xea, 0x83, 0x7c, 0xa9, 0x0b, 0x09, 0x9e, 0x0a, 0xe4, 0x00, 0x2a, 0xa9, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0x87, 0x8e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xf5, 0x62,
	0x9b, 0x45, 0x01, 0x00, 0x00,
}

func (m *Developer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Developer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Developer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDeveloper(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Blocked {
		i--
		if m.Blocked {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.UpdatedAt != 0 {
		i = encodeVarintDeveloper(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x18
	}
	if m.CreatedAt != 0 {
		i = encodeVarintDeveloper(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintDeveloper(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeveloper(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeveloper(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Developer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovDeveloper(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovDeveloper(uint64(m.CreatedAt))
	}
	if m.UpdatedAt != 0 {
		n += 1 + sovDeveloper(uint64(m.UpdatedAt))
	}
	if m.Blocked {
		n += 2
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDeveloper(uint64(l))
	}
	return n
}

func sovDeveloper(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeveloper(x uint64) (n int) {
	return sovDeveloper(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Developer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeveloper
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
			return fmt.Errorf("proto: Developer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Developer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeveloper
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
				return ErrInvalidLengthDeveloper
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeveloper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeveloper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeveloper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocked", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeveloper
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
			m.Blocked = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeveloper
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
				return ErrInvalidLengthDeveloper
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeveloper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeveloper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeveloper
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
func skipDeveloper(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeveloper
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
					return 0, ErrIntOverflowDeveloper
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
					return 0, ErrIntOverflowDeveloper
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
				return 0, ErrInvalidLengthDeveloper
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeveloper
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeveloper
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeveloper        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeveloper          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeveloper = fmt.Errorf("proto: unexpected end of group")
)
