// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: helios/insurance/v1beta1/insurance.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	types "github.com/Helios-Chain-Labs/sdk-go/chain/oracle/types"
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

type Params struct {
	// default_redemption_notice_period_duration defines the default minimum
	// notice period duration that must pass after an underwriter sends a
	// redemption request before the underwriter can claim his tokens
	DefaultRedemptionNoticePeriodDuration time.Duration `protobuf:"bytes,1,opt,name=default_redemption_notice_period_duration,json=defaultRedemptionNoticePeriodDuration,proto3,stdduration" json:"default_redemption_notice_period_duration" yaml:"default_redemption_notice_period_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b7246d78b44910b, []int{0}
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

func (m *Params) GetDefaultRedemptionNoticePeriodDuration() time.Duration {
	if m != nil {
		return m.DefaultRedemptionNoticePeriodDuration
	}
	return 0
}

type InsuranceFund struct {
	// deposit denomination for the given insurance fund
	DepositDenom string `protobuf:"bytes,1,opt,name=deposit_denom,json=depositDenom,proto3" json:"deposit_denom,omitempty"`
	// insurance fund pool token denomination for the given insurance fund
	InsurancePoolTokenDenom string `protobuf:"bytes,2,opt,name=insurance_pool_token_denom,json=insurancePoolTokenDenom,proto3" json:"insurance_pool_token_denom,omitempty"`
	// redemption_notice_period_duration defines the minimum notice period
	// duration that must pass after an underwriter sends a redemption request
	// before the underwriter can claim his tokens
	RedemptionNoticePeriodDuration time.Duration `protobuf:"bytes,3,opt,name=redemption_notice_period_duration,json=redemptionNoticePeriodDuration,proto3,stdduration" json:"redemption_notice_period_duration" yaml:"redemption_notice_period_duration"`
	// balance of fund
	Balance cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=balance,proto3,customtype=cosmossdk.io/math.Int" json:"balance"`
	// total share tokens minted
	TotalShare cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=total_share,json=totalShare,proto3,customtype=cosmossdk.io/math.Int" json:"total_share"`
	// marketID of the derivative market
	MarketId string `protobuf:"bytes,6,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// ticker of the derivative market
	MarketTicker string `protobuf:"bytes,7,opt,name=market_ticker,json=marketTicker,proto3" json:"market_ticker,omitempty"`
	// Oracle base currency of the derivative market OR the oracle symbol for the
	// binary options market.
	OracleBase string `protobuf:"bytes,8,opt,name=oracle_base,json=oracleBase,proto3" json:"oracle_base,omitempty"`
	// Oracle quote currency of the derivative market OR the oracle provider for
	// the binary options market.
	OracleQuote string `protobuf:"bytes,9,opt,name=oracle_quote,json=oracleQuote,proto3" json:"oracle_quote,omitempty"`
	// Oracle type of the binary options or derivative market
	OracleType types.OracleType `protobuf:"varint,10,opt,name=oracle_type,json=oracleType,proto3,enum=helios.oracle.v1beta1.OracleType" json:"oracle_type,omitempty"`
	// Expiration time of the derivative market. Should be -1 for perpetual or -2
	// for binary options markets.
	Expiry int64 `protobuf:"varint,11,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (m *InsuranceFund) Reset()         { *m = InsuranceFund{} }
func (m *InsuranceFund) String() string { return proto.CompactTextString(m) }
func (*InsuranceFund) ProtoMessage()    {}
func (*InsuranceFund) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b7246d78b44910b, []int{1}
}
func (m *InsuranceFund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InsuranceFund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InsuranceFund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InsuranceFund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsuranceFund.Merge(m, src)
}
func (m *InsuranceFund) XXX_Size() int {
	return m.Size()
}
func (m *InsuranceFund) XXX_DiscardUnknown() {
	xxx_messageInfo_InsuranceFund.DiscardUnknown(m)
}

var xxx_messageInfo_InsuranceFund proto.InternalMessageInfo

func (m *InsuranceFund) GetDepositDenom() string {
	if m != nil {
		return m.DepositDenom
	}
	return ""
}

func (m *InsuranceFund) GetInsurancePoolTokenDenom() string {
	if m != nil {
		return m.InsurancePoolTokenDenom
	}
	return ""
}

func (m *InsuranceFund) GetRedemptionNoticePeriodDuration() time.Duration {
	if m != nil {
		return m.RedemptionNoticePeriodDuration
	}
	return 0
}

func (m *InsuranceFund) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *InsuranceFund) GetMarketTicker() string {
	if m != nil {
		return m.MarketTicker
	}
	return ""
}

func (m *InsuranceFund) GetOracleBase() string {
	if m != nil {
		return m.OracleBase
	}
	return ""
}

func (m *InsuranceFund) GetOracleQuote() string {
	if m != nil {
		return m.OracleQuote
	}
	return ""
}

func (m *InsuranceFund) GetOracleType() types.OracleType {
	if m != nil {
		return m.OracleType
	}
	return types.OracleType_Unspecified
}

func (m *InsuranceFund) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type RedemptionSchedule struct {
	// id of redemption schedule
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// marketId of insurance fund for the redemption
	MarketId string `protobuf:"bytes,2,opt,name=marketId,proto3" json:"marketId,omitempty"`
	// address of the redeemer
	Redeemer string `protobuf:"bytes,3,opt,name=redeemer,proto3" json:"redeemer,omitempty"`
	// the time after which the redemption can be claimed
	ClaimableRedemptionTime time.Time `protobuf:"bytes,4,opt,name=claimable_redemption_time,json=claimableRedemptionTime,proto3,stdtime" json:"claimable_redemption_time" yaml:"claimable_redemption_time"`
	// the insurance_pool_token amount to redeem
	RedemptionAmount types1.Coin `protobuf:"bytes,5,opt,name=redemption_amount,json=redemptionAmount,proto3" json:"redemption_amount"`
}

func (m *RedemptionSchedule) Reset()         { *m = RedemptionSchedule{} }
func (m *RedemptionSchedule) String() string { return proto.CompactTextString(m) }
func (*RedemptionSchedule) ProtoMessage()    {}
func (*RedemptionSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b7246d78b44910b, []int{2}
}
func (m *RedemptionSchedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RedemptionSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RedemptionSchedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RedemptionSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedemptionSchedule.Merge(m, src)
}
func (m *RedemptionSchedule) XXX_Size() int {
	return m.Size()
}
func (m *RedemptionSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_RedemptionSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_RedemptionSchedule proto.InternalMessageInfo

func (m *RedemptionSchedule) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RedemptionSchedule) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *RedemptionSchedule) GetRedeemer() string {
	if m != nil {
		return m.Redeemer
	}
	return ""
}

func (m *RedemptionSchedule) GetClaimableRedemptionTime() time.Time {
	if m != nil {
		return m.ClaimableRedemptionTime
	}
	return time.Time{}
}

func (m *RedemptionSchedule) GetRedemptionAmount() types1.Coin {
	if m != nil {
		return m.RedemptionAmount
	}
	return types1.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "helios.insurance.v1beta1.Params")
	proto.RegisterType((*InsuranceFund)(nil), "helios.insurance.v1beta1.InsuranceFund")
	proto.RegisterType((*RedemptionSchedule)(nil), "helios.insurance.v1beta1.RedemptionSchedule")
}

func init() {
	proto.RegisterFile("helios/insurance/v1beta1/insurance.proto", fileDescriptor_6b7246d78b44910b)
}

var fileDescriptor_6b7246d78b44910b = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x8d, 0xd3, 0xbc, 0xb4, 0x99, 0xb4, 0x55, 0x3b, 0x7a, 0xef, 0xd5, 0x0d, 0xc2, 0x49, 0x83,
	0x90, 0x02, 0x02, 0xbb, 0x2d, 0x48, 0x48, 0x05, 0x21, 0x61, 0x2a, 0xa4, 0x22, 0x04, 0xc5, 0xcd,
	0x8a, 0x8d, 0x35, 0xb1, 0xa7, 0xc9, 0x28, 0xb6, 0xc7, 0xd8, 0x63, 0x44, 0xf6, 0xac, 0x58, 0x75,
	0x89, 0x58, 0xf1, 0x13, 0xe8, 0xbf, 0xe8, 0xb2, 0x3b, 0x10, 0x8b, 0x82, 0xda, 0x05, 0xb0, 0xe5,
	0x17, 0xa0, 0xf9, 0xb0, 0x5d, 0x81, 0xa0, 0xdd, 0x58, 0x9e, 0x33, 0xe7, 0xde, 0xb9, 0xf7, 0xdc,
	0x33, 0x03, 0x7a, 0x23, 0x1c, 0x10, 0x9a, 0x5a, 0x24, 0x4a, 0xb3, 0x04, 0x45, 0x1e, 0xb6, 0x5e,
	0xac, 0x0d, 0x30, 0x43, 0x6b, 0x25, 0x62, 0xc6, 0x09, 0x65, 0x14, 0xea, 0x92, 0x69, 0x96, 0xb8,
	0x62, 0xb6, 0xfe, 0x1d, 0xd2, 0x21, 0x15, 0x24, 0x8b, 0xff, 0x49, 0x7e, 0xcb, 0x18, 0x52, 0x3a,
	0x0c, 0xb0, 0x25, 0x56, 0x83, 0x6c, 0xd7, 0xf2, 0xb3, 0x04, 0x31, 0x42, 0x23, 0xb5, 0xdf, 0xfe,
	0x75, 0x9f, 0x91, 0x10, 0xa7, 0x0c, 0x85, 0x71, 0x9e, 0xc0, 0xa3, 0x69, 0x48, 0x53, 0x6b, 0x80,
	0xd2, 0xb2, 0x2a, 0x8f, 0x92, 0x3c, 0x41, 0x57, 0x95, 0x4e, 0x13, 0xe4, 0x05, 0x25, 0x43, 0x2e,
	0x15, 0x67, 0x11, 0x85, 0x24, 0xa2, 0x96, 0xf8, 0x4a, 0xa8, 0xfb, 0x41, 0x03, 0xf5, 0x6d, 0x94,
	0xa0, 0x30, 0x85, 0xfb, 0x1a, 0xb8, 0xe2, 0xe3, 0x5d, 0x94, 0x05, 0xcc, 0x4d, 0xb0, 0x8f, 0xc3,
	0x98, 0xd7, 0xe7, 0x46, 0x94, 0x11, 0x0f, 0xbb, 0x31, 0x4e, 0x08, 0xf5, 0xdd, 0xbc, 0x6c, 0x5d,
	0xeb, 0x68, 0xbd, 0xe6, 0xfa, 0xb2, 0x29, 0xeb, 0x36, 0xf3, 0xba, 0xcd, 0x4d, 0x45, 0xb0, 0xef,
	0x1c, 0x1c, 0xb5, 0x2b, 0x3f, 0x8e, 0xda, 0xab, 0x13, 0x14, 0x06, 0x1b, 0xdd, 0x73, 0x67, 0xee,
	0xbe, 0xf9, 0xdc, 0xd6, 0x9c, 0xcb, 0x8a, 0xef, 0x14, 0xf4, 0xc7, 0x82, 0xbd, 0x2d, 0xc8, 0xf9,
	0x21, 0x1b, 0xcb, 0xdf, 0xde, 0xb5, 0xb5, 0xd7, 0x5f, 0xdf, 0x5f, 0x5d, 0x28, 0x47, 0x26, 0xdb,
	0xe9, 0x7e, 0xaf, 0x81, 0xb9, 0xad, 0x1c, 0x7c, 0x90, 0x45, 0x3e, 0xbc, 0x04, 0xe6, 0x7c, 0x1c,
	0xd3, 0x94, 0x30, 0xd7, 0xc7, 0x11, 0x0d, 0x45, 0x0f, 0x0d, 0x67, 0x56, 0x81, 0x9b, 0x1c, 0x83,
	0xb7, 0x41, 0xab, 0x48, 0xe5, 0xc6, 0x94, 0x06, 0x2e, 0xa3, 0x63, 0x1c, 0xa9, 0x88, 0xaa, 0x88,
	0x58, 0x2a, 0x18, 0xdb, 0x94, 0x06, 0x7d, 0xbe, 0x2f, 0x83, 0xdf, 0x6a, 0x60, 0xe5, 0x6c, 0xe9,
	0xa6, 0xce, 0x92, 0xee, 0xa6, 0x92, 0xae, 0x27, 0xa5, 0x3b, 0xa7, 0x64, 0x46, 0xf2, 0x57, 0xad,
	0xe0, 0x2d, 0x30, 0x3d, 0x40, 0x01, 0xaf, 0x5a, 0xaf, 0xf1, 0x36, 0xec, 0x8b, 0xfc, 0x98, 0x4f,
	0x47, 0xed, 0xff, 0xa4, 0xb5, 0x52, 0x7f, 0x6c, 0x12, 0x6a, 0x85, 0x88, 0x8d, 0xcc, 0xad, 0x88,
	0x39, 0x39, 0x1b, 0xde, 0x05, 0x4d, 0x46, 0x19, 0x0a, 0xdc, 0x74, 0x84, 0x12, 0xac, 0xff, 0x73,
	0x9e, 0x60, 0x20, 0x22, 0x76, 0x78, 0x00, 0xbc, 0x00, 0x1a, 0x21, 0x4a, 0xc6, 0x98, 0xb9, 0xc4,
	0xd7, 0xeb, 0x42, 0xc1, 0x19, 0x09, 0x6c, 0x89, 0xa1, 0xa8, 0x4d, 0x46, 0xbc, 0x31, 0x4e, 0xf4,
	0x69, 0x39, 0x14, 0x09, 0xf6, 0x05, 0x06, 0xdb, 0xa0, 0x29, 0x8d, 0xec, 0x72, 0xfb, 0xeb, 0x33,
	0x82, 0x02, 0x24, 0x64, 0xa3, 0x14, 0xc3, 0x15, 0x30, 0xab, 0x08, 0xcf, 0x33, 0xca, 0xb0, 0xde,
	0x10, 0x0c, 0x15, 0xf4, 0x94, 0x43, 0xd0, 0x2e, 0x72, 0xb0, 0x49, 0x8c, 0x75, 0xd0, 0xd1, 0x7a,
	0xf3, 0xeb, 0x2b, 0xa6, 0xba, 0xc7, 0xea, 0x9e, 0xa8, 0x6b, 0x63, 0x3e, 0x11, 0xcb, 0xfe, 0x24,
	0xc6, 0xf9, 0x31, 0xfc, 0x1f, 0xfe, 0x0f, 0xea, 0xf8, 0x65, 0x4c, 0x92, 0x89, 0xde, 0xec, 0x68,
	0xbd, 0x29, 0x47, 0xad, 0xba, 0xfb, 0x55, 0x00, 0x4b, 0xa7, 0xee, 0x78, 0x23, 0xec, 0x67, 0x01,
	0x86, 0xf3, 0xa0, 0x4a, 0x7c, 0xe1, 0xb2, 0x9a, 0x53, 0x25, 0x3e, 0x6c, 0x81, 0xa2, 0x6f, 0xe5,
	0xa4, 0x52, 0x87, 0x16, 0x98, 0xe1, 0xf3, 0xc3, 0x21, 0x4e, 0x84, 0x41, 0x1a, 0x4e, 0xb1, 0x86,
	0xaf, 0x34, 0xb0, 0xec, 0x05, 0x88, 0x84, 0x68, 0x10, 0xe0, 0xd3, 0x37, 0x88, 0x3f, 0x12, 0x62,
	0x98, 0xcd, 0xf5, 0xd6, 0x6f, 0x76, 0xea, 0xe7, 0x2f, 0x88, 0x7d, 0x4d, 0xf9, 0xa9, 0x23, 0xfd,
	0xf4, 0xc7, 0x54, 0xdd, 0x3d, 0xee, 0xa3, 0xa5, 0x62, 0xbf, 0x6c, 0x89, 0xe7, 0x82, 0x8f, 0xc0,
	0xe2, 0xa9, 0x00, 0x14, 0xd2, 0x2c, 0x62, 0xc2, 0x0d, 0xdc, 0xcc, 0xd2, 0x06, 0x26, 0x9f, 0x4f,
	0xa1, 0xe2, 0x7d, 0x4a, 0x22, 0xbb, 0xc6, 0x0f, 0x77, 0x16, 0xca, 0xc8, 0x7b, 0x22, 0xd0, 0x7e,
	0x78, 0x70, 0x6c, 0x68, 0x87, 0xc7, 0x86, 0xf6, 0xe5, 0xd8, 0xd0, 0xf6, 0x4e, 0x8c, 0xca, 0xe1,
	0x89, 0x51, 0xf9, 0x78, 0x62, 0x54, 0x9e, 0xad, 0xca, 0x99, 0x5c, 0xf7, 0x68, 0x82, 0xad, 0xfc,
	0x7f, 0x84, 0x48, 0x64, 0x85, 0x94, 0xeb, 0x7a, 0xfa, 0x7d, 0xe6, 0xb3, 0x4c, 0x07, 0x75, 0xd1,
	0xf4, 0x8d, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20, 0x03, 0x37, 0xf3, 0xc0, 0x05, 0x00, 0x00,
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
	if this.DefaultRedemptionNoticePeriodDuration != that1.DefaultRedemptionNoticePeriodDuration {
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
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.DefaultRedemptionNoticePeriodDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.DefaultRedemptionNoticePeriodDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintInsurance(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *InsuranceFund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InsuranceFund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InsuranceFund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiry != 0 {
		i = encodeVarintInsurance(dAtA, i, uint64(m.Expiry))
		i--
		dAtA[i] = 0x58
	}
	if m.OracleType != 0 {
		i = encodeVarintInsurance(dAtA, i, uint64(m.OracleType))
		i--
		dAtA[i] = 0x50
	}
	if len(m.OracleQuote) > 0 {
		i -= len(m.OracleQuote)
		copy(dAtA[i:], m.OracleQuote)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.OracleQuote)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.OracleBase) > 0 {
		i -= len(m.OracleBase)
		copy(dAtA[i:], m.OracleBase)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.OracleBase)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.MarketTicker) > 0 {
		i -= len(m.MarketTicker)
		copy(dAtA[i:], m.MarketTicker)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.MarketTicker)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.MarketId) > 0 {
		i -= len(m.MarketId)
		copy(dAtA[i:], m.MarketId)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.MarketId)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.TotalShare.Size()
		i -= size
		if _, err := m.TotalShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInsurance(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInsurance(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.RedemptionNoticePeriodDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.RedemptionNoticePeriodDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintInsurance(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.InsurancePoolTokenDenom) > 0 {
		i -= len(m.InsurancePoolTokenDenom)
		copy(dAtA[i:], m.InsurancePoolTokenDenom)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.InsurancePoolTokenDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DepositDenom) > 0 {
		i -= len(m.DepositDenom)
		copy(dAtA[i:], m.DepositDenom)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.DepositDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RedemptionSchedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedemptionSchedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedemptionSchedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RedemptionAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintInsurance(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n4, err4 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.ClaimableRedemptionTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.ClaimableRedemptionTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintInsurance(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x22
	if len(m.Redeemer) > 0 {
		i -= len(m.Redeemer)
		copy(dAtA[i:], m.Redeemer)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.Redeemer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MarketId) > 0 {
		i -= len(m.MarketId)
		copy(dAtA[i:], m.MarketId)
		i = encodeVarintInsurance(dAtA, i, uint64(len(m.MarketId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintInsurance(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintInsurance(dAtA []byte, offset int, v uint64) int {
	offset -= sovInsurance(v)
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
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.DefaultRedemptionNoticePeriodDuration)
	n += 1 + l + sovInsurance(uint64(l))
	return n
}

func (m *InsuranceFund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DepositDenom)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	l = len(m.InsurancePoolTokenDenom)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.RedemptionNoticePeriodDuration)
	n += 1 + l + sovInsurance(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovInsurance(uint64(l))
	l = m.TotalShare.Size()
	n += 1 + l + sovInsurance(uint64(l))
	l = len(m.MarketId)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	l = len(m.MarketTicker)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	l = len(m.OracleBase)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	l = len(m.OracleQuote)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	if m.OracleType != 0 {
		n += 1 + sovInsurance(uint64(m.OracleType))
	}
	if m.Expiry != 0 {
		n += 1 + sovInsurance(uint64(m.Expiry))
	}
	return n
}

func (m *RedemptionSchedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovInsurance(uint64(m.Id))
	}
	l = len(m.MarketId)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	l = len(m.Redeemer)
	if l > 0 {
		n += 1 + l + sovInsurance(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.ClaimableRedemptionTime)
	n += 1 + l + sovInsurance(uint64(l))
	l = m.RedemptionAmount.Size()
	n += 1 + l + sovInsurance(uint64(l))
	return n
}

func sovInsurance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInsurance(x uint64) (n int) {
	return sovInsurance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInsurance
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
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultRedemptionNoticePeriodDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.DefaultRedemptionNoticePeriodDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInsurance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInsurance
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
func (m *InsuranceFund) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInsurance
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
			return fmt.Errorf("proto: InsuranceFund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InsuranceFund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsurancePoolTokenDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InsurancePoolTokenDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionNoticePeriodDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.RedemptionNoticePeriodDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketTicker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketTicker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleQuote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleQuote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleType", wireType)
			}
			m.OracleType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OracleType |= types.OracleType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			m.Expiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiry |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInsurance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInsurance
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
func (m *RedemptionSchedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInsurance
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
			return fmt.Errorf("proto: RedemptionSchedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedemptionSchedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redeemer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Redeemer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableRedemptionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.ClaimableRedemptionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedemptionAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInsurance
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
				return ErrInvalidLengthInsurance
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInsurance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RedemptionAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInsurance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInsurance
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
func skipInsurance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInsurance
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
					return 0, ErrIntOverflowInsurance
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
					return 0, ErrIntOverflowInsurance
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
				return 0, ErrInvalidLengthInsurance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInsurance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInsurance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInsurance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInsurance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInsurance = fmt.Errorf("proto: unexpected end of group")
)
