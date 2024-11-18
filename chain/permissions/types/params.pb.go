// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/permissions/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the permissions module.
type Params struct {
	WasmHookQueryMaxGas uint64 `protobuf:"varint,1,opt,name=wasm_hook_query_max_gas,json=wasmHookQueryMaxGas,proto3" json:"wasm_hook_query_max_gas,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a7ea0496163621f, []int{0}
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

func (m *Params) GetWasmHookQueryMaxGas() uint64 {
	if m != nil {
		return m.WasmHookQueryMaxGas
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "helios.permissions.v1beta1.Params")
}

func init() {
	proto.RegisterFile("injective/permissions/v1beta1/params.proto", fileDescriptor_4a7ea0496163621f)
}

var fileDescriptor_4a7ea0496163621f = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x6b, 0x09, 0x75, 0xc8, 0x46, 0x40, 0x02, 0x8a, 0x30, 0x88, 0x09, 0x55, 0x22, 0x56,
	0x05, 0x13, 0x23, 0x0b, 0x20, 0x81, 0x44, 0x19, 0x61, 0x88, 0x5e, 0x82, 0x95, 0x98, 0xe0, 0xbc,
	0x90, 0xe7, 0x94, 0xf6, 0x0a, 0x4c, 0x1c, 0x81, 0x23, 0x70, 0x0c, 0xc6, 0x8e, 0x8c, 0x28, 0x19,
	0xe0, 0x18, 0x28, 0x71, 0x88, 0xc2, 0x62, 0xf9, 0xf7, 0xf7, 0xdb, 0x7e, 0xff, 0xef, 0x8c, 0x55,
	0xfa, 0x20, 0x43, 0xa3, 0x66, 0x52, 0x64, 0x32, 0xd7, 0x8a, 0x48, 0x61, 0x4a, 0x62, 0x36, 0x09,
	0xa4, 0x81, 0x89, 0xc8, 0x20, 0x07, 0x4d, 0x5e, 0x96, 0xa3, 0x41, 0x77, 0xa7, 0xf3, 0x7a, 0x3d,
	0xaf, 0xd7, 0x7a, 0x47, 0xeb, 0x11, 0x46, 0xd8, 0x38, 0x45, 0xbd, 0xb3, 0x97, 0x46, 0x5b, 0x21,
	0x92, 0x46, 0xf2, 0x2d, 0xb0, 0xa2, 0x45, 0xdc, 0x2a, 0x11, 0x00, 0xc9, 0xee, 0xc7, 0x10, 0x55,
	0xda, 0xf2, 0x55, 0xd0, 0x2a, 0x45, 0xd1, 0xac, 0xf6, 0x68, 0xff, 0xce, 0x19, 0x5e, 0x37, 0x23,
	0xb9, 0xc7, 0xce, 0xc6, 0x33, 0x90, 0xf6, 0x63, 0xc4, 0xc4, 0x7f, 0x2a, 0x64, 0xbe, 0xf0, 0x35,
	0xcc, 0xfd, 0x08, 0x68, 0x93, 0xed, 0xb1, 0x83, 0x95, 0x9b, 0xb5, 0x1a, 0x9f, 0x23, 0x26, 0xd3,
	0x1a, 0x5e, 0xc1, 0xfc, 0x0c, 0xe8, 0x64, 0xfb, 0xe7, 0x6d, 0x97, 0xbd, 0x7c, 0xbf, 0x8f, 0xdd,
	0x7e, 0x5a, 0xfb, 0xe4, 0x69, 0xf2, 0x51, 0x72, 0xb6, 0x2c, 0x39, 0xfb, 0x2a, 0x39, 0x7b, 0xad,
	0xf8, 0x60, 0x59, 0xf1, 0xc1, 0x67, 0xc5, 0x07, 0xb7, 0xd3, 0x48, 0x99, 0xb8, 0x08, 0xbc, 0x10,
	0xb5, 0xb8, 0xf8, 0x2b, 0xe1, 0x12, 0x02, 0x12, 0x5d, 0x25, 0x87, 0x21, 0xe6, 0xb2, 0x2f, 0x63,
	0x50, 0xa9, 0xd0, 0x78, 0x5f, 0x3c, 0x4a, 0xfa, 0xd7, 0xad, 0x59, 0x64, 0x92, 0x82, 0x61, 0x13,
	0xe8, 0xe8, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x42, 0x3f, 0x45, 0x81, 0x01, 0x00, 0x00,
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
	if this.WasmHookQueryMaxGas != that1.WasmHookQueryMaxGas {
		return false
	}
	return true
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
	if m.WasmHookQueryMaxGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.WasmHookQueryMaxGas))
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
	if m.WasmHookQueryMaxGas != 0 {
		n += 1 + sovParams(uint64(m.WasmHookQueryMaxGas))
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
				return fmt.Errorf("proto: wrong wireType = %d for field WasmHookQueryMaxGas", wireType)
			}
			m.WasmHookQueryMaxGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WasmHookQueryMaxGas |= uint64(b&0x7F) << shift
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
