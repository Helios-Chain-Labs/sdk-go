// Code generated with goa v3.5.2, DO NOT EDIT.
//
// HeliosInsuranceRPC protocol buffer definition
//
// Command:
// $ goa gen github.com/Helios-Chain-Labs/-indexer/api/design -o ../

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: _insurance_rpc.proto

package _insurance_rpcpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FundsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FundsRequest) Reset() {
	*x = FundsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__insurance_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundsRequest) ProtoMessage() {}

func (x *FundsRequest) ProtoReflect() protoreflect.Message {
	mi := &file__insurance_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundsRequest.ProtoReflect.Descriptor instead.
func (*FundsRequest) Descriptor() ([]byte, []int) {
	return file__insurance_rpc_proto_rawDescGZIP(), []int{0}
}

type FundsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Funds []*InsuranceFund `protobuf:"bytes,1,rep,name=funds,proto3" json:"funds,omitempty"`
}

func (x *FundsResponse) Reset() {
	*x = FundsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__insurance_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundsResponse) ProtoMessage() {}

func (x *FundsResponse) ProtoReflect() protoreflect.Message {
	mi := &file__insurance_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundsResponse.ProtoReflect.Descriptor instead.
func (*FundsResponse) Descriptor() ([]byte, []int) {
	return file__insurance_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *FundsResponse) GetFunds() []*InsuranceFund {
	if x != nil {
		return x.Funds
	}
	return nil
}

type InsuranceFund struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ticker of the derivative market.
	MarketTicker string `protobuf:"bytes,1,opt,name=market_ticker,json=marketTicker,proto3" json:"market_ticker,omitempty"`
	// Derivative Market ID
	MarketId string `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Coin denom used for the underwriting of the insurance fund.
	DepositDenom string `protobuf:"bytes,3,opt,name=deposit_denom,json=depositDenom,proto3" json:"deposit_denom,omitempty"`
	// Pool token denom
	PoolTokenDenom string `protobuf:"bytes,4,opt,name=pool_token_denom,json=poolTokenDenom,proto3" json:"pool_token_denom,omitempty"`
	// Redemption notice period duration in seconds.
	RedemptionNoticePeriodDuration int64  `protobuf:"zigzag64,5,opt,name=redemption_notice_period_duration,json=redemptionNoticePeriodDuration,proto3" json:"redemption_notice_period_duration,omitempty"`
	Balance                        string `protobuf:"bytes,6,opt,name=balance,proto3" json:"balance,omitempty"`
	TotalShare                     string `protobuf:"bytes,7,opt,name=total_share,json=totalShare,proto3" json:"total_share,omitempty"`
	// Oracle base currency
	OracleBase string `protobuf:"bytes,8,opt,name=oracle_base,json=oracleBase,proto3" json:"oracle_base,omitempty"`
	// Oracle quote currency
	OracleQuote string `protobuf:"bytes,9,opt,name=oracle_quote,json=oracleQuote,proto3" json:"oracle_quote,omitempty"`
	// Oracle Type
	OracleType string `protobuf:"bytes,10,opt,name=oracle_type,json=oracleType,proto3" json:"oracle_type,omitempty"`
	// Defines the expiry, if any
	Expiry int64 `protobuf:"zigzag64,11,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// Token metadata for the deposit asset, only for Ethereum-based assets
	DepositTokenMeta *TokenMeta `protobuf:"bytes,12,opt,name=deposit_token_meta,json=depositTokenMeta,proto3" json:"deposit_token_meta,omitempty"`
}

func (x *InsuranceFund) Reset() {
	*x = InsuranceFund{}
	if protoimpl.UnsafeEnabled {
		mi := &file__insurance_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsuranceFund) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsuranceFund) ProtoMessage() {}

func (x *InsuranceFund) ProtoReflect() protoreflect.Message {
	mi := &file__insurance_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsuranceFund.ProtoReflect.Descriptor instead.
func (*InsuranceFund) Descriptor() ([]byte, []int) {
	return file__insurance_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *InsuranceFund) GetMarketTicker() string {
	if x != nil {
		return x.MarketTicker
	}
	return ""
}

func (x *InsuranceFund) GetMarketId() string {
	if x != nil {
		return x.MarketId
	}
	return ""
}

func (x *InsuranceFund) GetDepositDenom() string {
	if x != nil {
		return x.DepositDenom
	}
	return ""
}

func (x *InsuranceFund) GetPoolTokenDenom() string {
	if x != nil {
		return x.PoolTokenDenom
	}
	return ""
}

func (x *InsuranceFund) GetRedemptionNoticePeriodDuration() int64 {
	if x != nil {
		return x.RedemptionNoticePeriodDuration
	}
	return 0
}

func (x *InsuranceFund) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *InsuranceFund) GetTotalShare() string {
	if x != nil {
		return x.TotalShare
	}
	return ""
}

func (x *InsuranceFund) GetOracleBase() string {
	if x != nil {
		return x.OracleBase
	}
	return ""
}

func (x *InsuranceFund) GetOracleQuote() string {
	if x != nil {
		return x.OracleQuote
	}
	return ""
}

func (x *InsuranceFund) GetOracleType() string {
	if x != nil {
		return x.OracleType
	}
	return ""
}

func (x *InsuranceFund) GetExpiry() int64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

func (x *InsuranceFund) GetDepositTokenMeta() *TokenMeta {
	if x != nil {
		return x.DepositTokenMeta
	}
	return nil
}

type TokenMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token full name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Token Ethereum contract address
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Token symbol short name
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// URL to the logo image
	Logo string `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
	// Token decimals
	Decimals int32 `protobuf:"zigzag32,5,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// Token metadata fetched timestamp in UNIX millis.
	UpdatedAt int64 `protobuf:"zigzag64,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TokenMeta) Reset() {
	*x = TokenMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file__insurance_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMeta) ProtoMessage() {}

func (x *TokenMeta) ProtoReflect() protoreflect.Message {
	mi := &file__insurance_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMeta.ProtoReflect.Descriptor instead.
func (*TokenMeta) Descriptor() ([]byte, []int) {
	return file__insurance_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *TokenMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenMeta) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TokenMeta) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenMeta) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *TokenMeta) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TokenMeta) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type RedemptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account address of the redemption owner
	Redeemer string `protobuf:"bytes,1,opt,name=redeemer,proto3" json:"redeemer,omitempty"`
	// Denom of the insurance pool token.
	RedemptionDenom string `protobuf:"bytes,2,opt,name=redemption_denom,json=redemptionDenom,proto3" json:"redemption_denom,omitempty"`
	// Status of the redemption. Either pending or disbursed.
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RedemptionsRequest) Reset() {
	*x = RedemptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__insurance_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedemptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedemptionsRequest) ProtoMessage() {}

func (x *RedemptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file__insurance_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedemptionsRequest.ProtoReflect.Descriptor instead.
func (*RedemptionsRequest) Descriptor() ([]byte, []int) {
	return file__insurance_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *RedemptionsRequest) GetRedeemer() string {
	if x != nil {
		return x.Redeemer
	}
	return ""
}

func (x *RedemptionsRequest) GetRedemptionDenom() string {
	if x != nil {
		return x.RedemptionDenom
	}
	return ""
}

func (x *RedemptionsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RedemptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedemptionSchedules []*RedemptionSchedule `protobuf:"bytes,1,rep,name=redemption_schedules,json=redemptionSchedules,proto3" json:"redemption_schedules,omitempty"`
}

func (x *RedemptionsResponse) Reset() {
	*x = RedemptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__insurance_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedemptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedemptionsResponse) ProtoMessage() {}

func (x *RedemptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file__insurance_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedemptionsResponse.ProtoReflect.Descriptor instead.
func (*RedemptionsResponse) Descriptor() ([]byte, []int) {
	return file__insurance_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *RedemptionsResponse) GetRedemptionSchedules() []*RedemptionSchedule {
	if x != nil {
		return x.RedemptionSchedules
	}
	return nil
}

type RedemptionSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Redemption ID.
	RedemptionId uint64 `protobuf:"varint,1,opt,name=redemption_id,json=redemptionId,proto3" json:"redemption_id,omitempty"`
	// Status of the redemption. Either pending or disbursed.
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Account address of the redemption owner
	Redeemer string `protobuf:"bytes,3,opt,name=redeemer,proto3" json:"redeemer,omitempty"`
	// Claimable redemption time in seconds
	ClaimableRedemptionTime int64 `protobuf:"zigzag64,4,opt,name=claimable_redemption_time,json=claimableRedemptionTime,proto3" json:"claimable_redemption_time,omitempty"`
	// Amount of pool tokens being redeemed.
	RedemptionAmount string `protobuf:"bytes,5,opt,name=redemption_amount,json=redemptionAmount,proto3" json:"redemption_amount,omitempty"`
	// Pool token denom being redeemed.
	RedemptionDenom string `protobuf:"bytes,6,opt,name=redemption_denom,json=redemptionDenom,proto3" json:"redemption_denom,omitempty"`
	// Redemption request time in unix milliseconds.
	RequestedAt int64 `protobuf:"zigzag64,7,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
	// Amount of quote tokens disbursed
	DisbursedAmount string `protobuf:"bytes,8,opt,name=disbursed_amount,json=disbursedAmount,proto3" json:"disbursed_amount,omitempty"`
	// Denom of the quote tokens disbursed
	DisbursedDenom string `protobuf:"bytes,9,opt,name=disbursed_denom,json=disbursedDenom,proto3" json:"disbursed_denom,omitempty"`
	// Redemption disbursement time in unix milliseconds.
	DisbursedAt int64 `protobuf:"zigzag64,10,opt,name=disbursed_at,json=disbursedAt,proto3" json:"disbursed_at,omitempty"`
}

func (x *RedemptionSchedule) Reset() {
	*x = RedemptionSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file__insurance_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedemptionSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedemptionSchedule) ProtoMessage() {}

func (x *RedemptionSchedule) ProtoReflect() protoreflect.Message {
	mi := &file__insurance_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedemptionSchedule.ProtoReflect.Descriptor instead.
func (*RedemptionSchedule) Descriptor() ([]byte, []int) {
	return file__insurance_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *RedemptionSchedule) GetRedemptionId() uint64 {
	if x != nil {
		return x.RedemptionId
	}
	return 0
}

func (x *RedemptionSchedule) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RedemptionSchedule) GetRedeemer() string {
	if x != nil {
		return x.Redeemer
	}
	return ""
}

func (x *RedemptionSchedule) GetClaimableRedemptionTime() int64 {
	if x != nil {
		return x.ClaimableRedemptionTime
	}
	return 0
}

func (x *RedemptionSchedule) GetRedemptionAmount() string {
	if x != nil {
		return x.RedemptionAmount
	}
	return ""
}

func (x *RedemptionSchedule) GetRedemptionDenom() string {
	if x != nil {
		return x.RedemptionDenom
	}
	return ""
}

func (x *RedemptionSchedule) GetRequestedAt() int64 {
	if x != nil {
		return x.RequestedAt
	}
	return 0
}

func (x *RedemptionSchedule) GetDisbursedAmount() string {
	if x != nil {
		return x.DisbursedAmount
	}
	return ""
}

func (x *RedemptionSchedule) GetDisbursedDenom() string {
	if x != nil {
		return x.DisbursedDenom
	}
	return ""
}

func (x *RedemptionSchedule) GetDisbursedAt() int64 {
	if x != nil {
		return x.DisbursedAt
	}
	return 0
}

var File__insurance_rpc_proto protoreflect.FileDescriptor

var file__insurance_rpc_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x0e, 0x0a, 0x0c, 0x46, 0x75, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x0d, 0x46, 0x75, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x66, 0x75, 0x6e,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6e, 0x64,
	0x52, 0x05, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x22, 0xf5, 0x03, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x75,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d,
	0x12, 0x28, 0x0a, 0x10, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6f, 0x6f, 0x6c,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x49, 0x0a, 0x21, 0x72, 0x65,
	0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x12, 0x52, 0x1e, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x50, 0x0a,
	0x12, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x10, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x22,
	0xa0, 0x01, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x73, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x65,
	0x65, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x65,
	0x65, 0x6d, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x75, 0x0a, 0x13, 0x52, 0x65, 0x64, 0x65, 0x6d,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e,
	0x0a, 0x14, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x13, 0x72, 0x65, 0x64, 0x65, 0x6d,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x9b,
	0x03, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x65,
	0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x65, 0x72, 0x12, 0x3a,
	0x0a, 0x19, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x64, 0x65,
	0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x17, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x64, 0x65,
	0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65,
	0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x64, 0x65, 0x6d,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x72, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e,
	0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73,
	0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x62, 0x75,
	0x72, 0x73, 0x65, 0x64, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x62, 0x75, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x62, 0x75, 0x72, 0x73, 0x65, 0x64, 0x41, 0x74, 0x32, 0xd9, 0x01, 0x0a,
	0x15, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x50, 0x43, 0x12, 0x56, 0x0a, 0x05, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x12,
	0x25, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68,
	0x0a, 0x0b, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x2e,
	0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x2f, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x72, 0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__insurance_rpc_proto_rawDescOnce sync.Once
	file__insurance_rpc_proto_rawDescData = file__insurance_rpc_proto_rawDesc
)

func file__insurance_rpc_proto_rawDescGZIP() []byte {
	file__insurance_rpc_proto_rawDescOnce.Do(func() {
		file__insurance_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file__insurance_rpc_proto_rawDescData)
	})
	return file__insurance_rpc_proto_rawDescData
}

var file__insurance_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file__insurance_rpc_proto_goTypes = []interface{}{
	(*FundsRequest)(nil),        // 0: _insurance_rpc.FundsRequest
	(*FundsResponse)(nil),       // 1: _insurance_rpc.FundsResponse
	(*InsuranceFund)(nil),       // 2: _insurance_rpc.InsuranceFund
	(*TokenMeta)(nil),           // 3: _insurance_rpc.TokenMeta
	(*RedemptionsRequest)(nil),  // 4: _insurance_rpc.RedemptionsRequest
	(*RedemptionsResponse)(nil), // 5: _insurance_rpc.RedemptionsResponse
	(*RedemptionSchedule)(nil),  // 6: _insurance_rpc.RedemptionSchedule
}
var file__insurance_rpc_proto_depIdxs = []int32{
	2, // 0: _insurance_rpc.FundsResponse.funds:type_name -> _insurance_rpc.InsuranceFund
	3, // 1: _insurance_rpc.InsuranceFund.deposit_token_meta:type_name -> _insurance_rpc.TokenMeta
	6, // 2: _insurance_rpc.RedemptionsResponse.redemption_schedules:type_name -> _insurance_rpc.RedemptionSchedule
	0, // 3: _insurance_rpc.HeliosInsuranceRPC.Funds:input_type -> _insurance_rpc.FundsRequest
	4, // 4: _insurance_rpc.HeliosInsuranceRPC.Redemptions:input_type -> _insurance_rpc.RedemptionsRequest
	1, // 5: _insurance_rpc.HeliosInsuranceRPC.Funds:output_type -> _insurance_rpc.FundsResponse
	5, // 6: _insurance_rpc.HeliosInsuranceRPC.Redemptions:output_type -> _insurance_rpc.RedemptionsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file__insurance_rpc_proto_init() }
func file__insurance_rpc_proto_init() {
	if File__insurance_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__insurance_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file__insurance_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file__insurance_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsuranceFund); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file__insurance_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file__insurance_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedemptionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file__insurance_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedemptionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file__insurance_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedemptionSchedule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file__insurance_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file__insurance_rpc_proto_goTypes,
		DependencyIndexes: file__insurance_rpc_proto_depIdxs,
		MessageInfos:      file__insurance_rpc_proto_msgTypes,
	}.Build()
	File__insurance_rpc_proto = out.File
	file__insurance_rpc_proto_rawDesc = nil
	file__insurance_rpc_proto_goTypes = nil
	file__insurance_rpc_proto_depIdxs = nil
}
