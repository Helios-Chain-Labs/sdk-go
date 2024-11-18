// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/auction/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the auction module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to auction.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// current auction round
	AuctionRound uint64 `protobuf:"varint,2,opt,name=auction_round,json=auctionRound,proto3" json:"auction_round,omitempty"`
	// current highest bid
	HighestBid *Bid `protobuf:"bytes,3,opt,name=highest_bid,json=highestBid,proto3" json:"highest_bid,omitempty"`
	// auction ending timestamp
	AuctionEndingTimestamp int64 `protobuf:"varint,4,opt,name=auction_ending_timestamp,json=auctionEndingTimestamp,proto3" json:"auction_ending_timestamp,omitempty"`
	// last auction result
	LastAuctionResult *LastAuctionResult `protobuf:"bytes,5,opt,name=last_auction_result,json=lastAuctionResult,proto3" json:"last_auction_result,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f095f457353f49, []int{0}
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

func (m *GenesisState) GetAuctionRound() uint64 {
	if m != nil {
		return m.AuctionRound
	}
	return 0
}

func (m *GenesisState) GetHighestBid() *Bid {
	if m != nil {
		return m.HighestBid
	}
	return nil
}

func (m *GenesisState) GetAuctionEndingTimestamp() int64 {
	if m != nil {
		return m.AuctionEndingTimestamp
	}
	return 0
}

func (m *GenesisState) GetLastAuctionResult() *LastAuctionResult {
	if m != nil {
		return m.LastAuctionResult
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "helios.auction.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("injective/auction/v1beta1/genesis.proto", fileDescriptor_56f095f457353f49)
}

var fileDescriptor_56f095f457353f49 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x93, 0xb6, 0xb7, 0x8b, 0x69, 0xef, 0xe2, 0xe6, 0x8a, 0xc4, 0x2e, 0x62, 0xd5, 0x85,
	0x5d, 0x68, 0x42, 0x75, 0xe3, 0xae, 0x18, 0x10, 0x11, 0x2a, 0x48, 0x74, 0x25, 0x42, 0x98, 0x24,
	0xc3, 0x64, 0x24, 0x99, 0x09, 0x99, 0x93, 0x82, 0x6f, 0xe1, 0x63, 0x75, 0x23, 0x74, 0xe9, 0x4a,
	0xa4, 0x7d, 0x11, 0xe9, 0x74, 0x12, 0x44, 0xb0, 0xbb, 0x33, 0xff, 0xf9, 0xcf, 0xf9, 0xfe, 0xe1,
	0xa0, 0x63, 0xc6, 0x9f, 0x49, 0x0c, 0x6c, 0x46, 0x3c, 0x5c, 0xc5, 0xc0, 0x04, 0xf7, 0x66, 0xe3,
	0x88, 0x00, 0x1e, 0x7b, 0x94, 0x70, 0x22, 0x99, 0x74, 0x8b, 0x52, 0x80, 0xb0, 0xf6, 0x1a, 0xa3,
	0xab, 0x8d, 0xae, 0x36, 0x0e, 0xb6, 0xec, 0xa8, 0xad, 0x6a, 0xc7, 0x60, 0x87, 0x0a, 0x2a, 0x54,
	0xe9, 0xad, 0xab, 0x8d, 0x7a, 0xf8, 0xd6, 0x42, 0xfd, 0xeb, 0x0d, 0xeb, 0x1e, 0x30, 0x10, 0x6b,
	0x82, 0xba, 0x05, 0x2e, 0x71, 0x2e, 0x6d, 0x73, 0x68, 0x8e, 0x7a, 0x67, 0x07, 0xee, 0xaf, 0x6c,
	0xf7, 0x4e, 0x19, 0xfd, 0xce, 0xfc, 0x63, 0xdf, 0x08, 0xf4, 0x98, 0x75, 0x84, 0xfe, 0x6a, 0x5f,
	0x58, 0x8a, 0x8a, 0x27, 0x76, 0x6b, 0x68, 0x8e, 0x3a, 0x41, 0x5f, 0x8b, 0xc1, 0x5a, 0xb3, 0x26,
	0xa8, 0x97, 0x32, 0x9a, 0x12, 0x09, 0x61, 0xc4, 0x12, 0xbb, 0xad, 0x50, 0xce, 0x16, 0x94, 0xcf,
	0x92, 0x00, 0xe9, 0x11, 0x9f, 0x25, 0xd6, 0x05, 0xb2, 0x6b, 0x0a, 0xe1, 0x09, 0xe3, 0x34, 0x04,
	0x96, 0x13, 0x09, 0x38, 0x2f, 0xec, 0xce, 0xd0, 0x1c, 0xb5, 0x83, 0x5d, 0xdd, 0xbf, 0x52, 0xed,
	0x87, 0xba, 0x6b, 0x3d, 0xa1, 0xff, 0x19, 0x96, 0x10, 0x36, 0x21, 0x89, 0xac, 0x32, 0xb0, 0xff,
	0xa8, 0x08, 0x27, 0x5b, 0x22, 0x4c, 0xb1, 0x84, 0x4b, 0xfd, 0x09, 0x35, 0x13, 0xfc, 0xcb, 0x7e,
	0x4a, 0x3e, 0x9d, 0x2f, 0x1d, 0x73, 0xb1, 0x74, 0xcc, 0xcf, 0xa5, 0x63, 0xbe, 0xae, 0x1c, 0x63,
	0xb1, 0x72, 0x8c, 0xf7, 0x95, 0x63, 0x3c, 0xde, 0x52, 0x06, 0x69, 0x15, 0xb9, 0xb1, 0xc8, 0xbd,
	0x9b, 0x1a, 0x32, 0xc5, 0x91, 0xf4, 0x1a, 0xe4, 0x69, 0x2c, 0x4a, 0xf2, 0xfd, 0x99, 0x62, 0xc6,
	0xbd, 0x5c, 0x24, 0x55, 0x46, 0x64, 0x73, 0x5e, 0x78, 0x29, 0x88, 0x8c, 0xba, 0xea, 0x7e, 0xe7,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x91, 0x86, 0x72, 0xfe, 0x44, 0x02, 0x00, 0x00,
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
	if m.LastAuctionResult != nil {
		{
			size, err := m.LastAuctionResult.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.AuctionEndingTimestamp != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AuctionEndingTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if m.HighestBid != nil {
		{
			size, err := m.HighestBid.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.AuctionRound != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AuctionRound))
		i--
		dAtA[i] = 0x10
	}
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.AuctionRound != 0 {
		n += 1 + sovGenesis(uint64(m.AuctionRound))
	}
	if m.HighestBid != nil {
		l = m.HighestBid.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.AuctionEndingTimestamp != 0 {
		n += 1 + sovGenesis(uint64(m.AuctionEndingTimestamp))
	}
	if m.LastAuctionResult != nil {
		l = m.LastAuctionResult.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionRound", wireType)
			}
			m.AuctionRound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionRound |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestBid", wireType)
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
			if m.HighestBid == nil {
				m.HighestBid = &Bid{}
			}
			if err := m.HighestBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionEndingTimestamp", wireType)
			}
			m.AuctionEndingTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionEndingTimestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastAuctionResult", wireType)
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
			if m.LastAuctionResult == nil {
				m.LastAuctionResult = &LastAuctionResult{}
			}
			if err := m.LastAuctionResult.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
