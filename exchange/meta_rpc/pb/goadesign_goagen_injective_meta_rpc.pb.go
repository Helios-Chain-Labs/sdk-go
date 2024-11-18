// Code generated with goa v3.7.0, DO NOT EDIT.
//
// InjectiveMetaRPC protocol buffer definition
//
// Command:
// $ goa gen github.com/Helios-Chain-Labs/helios-indexer/api/design -o ../

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: goadesign_goagen_helios_meta_rpc.proto

package helios_meta_rpcpb

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{1}
}

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{2}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// helios-exchange code version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Additional build meta info.
	Build map[string]string `protobuf:"bytes,2,rep,name=build,proto3" json:"build,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionResponse) GetBuild() map[string]string {
	if x != nil {
		return x.Build
	}
	return nil
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provide current system UNIX timestamp in millis
	Timestamp int64 `protobuf:"zigzag64,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *InfoRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The original timestamp value in millis.
	Timestamp int64 `protobuf:"zigzag64,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// UNIX time on the server in millis.
	ServerTime int64 `protobuf:"zigzag64,2,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
	// helios-exchange code version.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Additional build meta info.
	Build map[string]string `protobuf:"bytes,4,rep,name=build,proto3" json:"build,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Server's location region
	Region string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *InfoResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *InfoResponse) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *InfoResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *InfoResponse) GetBuild() map[string]string {
	if x != nil {
		return x.Build
	}
	return nil
}

func (x *InfoResponse) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type StreamKeepaliveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamKeepaliveRequest) Reset() {
	*x = StreamKeepaliveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamKeepaliveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamKeepaliveRequest) ProtoMessage() {}

func (x *StreamKeepaliveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamKeepaliveRequest.ProtoReflect.Descriptor instead.
func (*StreamKeepaliveRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{6}
}

type StreamKeepaliveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Server event
	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// New conection endpoint for the gRPC API
	NewEndpoint string `protobuf:"bytes,2,opt,name=new_endpoint,json=newEndpoint,proto3" json:"new_endpoint,omitempty"`
	// Operation timestamp in UNIX millis.
	Timestamp int64 `protobuf:"zigzag64,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StreamKeepaliveResponse) Reset() {
	*x = StreamKeepaliveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamKeepaliveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamKeepaliveResponse) ProtoMessage() {}

func (x *StreamKeepaliveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamKeepaliveResponse.ProtoReflect.Descriptor instead.
func (*StreamKeepaliveResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *StreamKeepaliveResponse) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *StreamKeepaliveResponse) GetNewEndpoint() string {
	if x != nil {
		return x.NewEndpoint
	}
	return ""
}

func (x *StreamKeepaliveResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type TokenMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denoms []string `protobuf:"bytes,1,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (x *TokenMetadataRequest) Reset() {
	*x = TokenMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMetadataRequest) ProtoMessage() {}

func (x *TokenMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMetadataRequest.ProtoReflect.Descriptor instead.
func (*TokenMetadataRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *TokenMetadataRequest) GetDenoms() []string {
	if x != nil {
		return x.Denoms
	}
	return nil
}

type TokenMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tokens and their metadata list
	Tokens []*TokenMetadataElement `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *TokenMetadataResponse) Reset() {
	*x = TokenMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMetadataResponse) ProtoMessage() {}

func (x *TokenMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMetadataResponse.ProtoReflect.Descriptor instead.
func (*TokenMetadataResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{9}
}

func (x *TokenMetadataResponse) GetTokens() []*TokenMetadataElement {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type TokenMetadataElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token's Ethereum address, not all token have this information
	EthereumAddress string `protobuf:"bytes,1,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
	// Token's CoinGecko id for price references
	CoingeckoId string `protobuf:"bytes,2,opt,name=coingecko_id,json=coingeckoId,proto3" json:"coingecko_id,omitempty"`
	// Token's denom on helios chain
	Denom string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	// Token name
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Token symbol
	Symbol string `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Number of decimal places used to represent the token's smallest unit
	Decimals int32 `protobuf:"zigzag32,6,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// Token logo URL
	Logo string `protobuf:"bytes,7,opt,name=logo,proto3" json:"logo,omitempty"`
}

func (x *TokenMetadataElement) Reset() {
	*x = TokenMetadataElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMetadataElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMetadataElement) ProtoMessage() {}

func (x *TokenMetadataElement) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMetadataElement.ProtoReflect.Descriptor instead.
func (*TokenMetadataElement) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP(), []int{10}
}

func (x *TokenMetadataElement) GetEthereumAddress() string {
	if x != nil {
		return x.EthereumAddress
	}
	return ""
}

func (x *TokenMetadataElement) GetCoingeckoId() string {
	if x != nil {
		return x.CoingeckoId
	}
	return ""
}

func (x *TokenMetadataElement) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *TokenMetadataElement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenMetadataElement) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenMetadataElement) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TokenMetadataElement) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

var File_goadesign_goagen_helios_meta_rpc_proto protoreflect.FileDescriptor

var file_goadesign_goagen_helios_meta_rpc_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x22,
	0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e,
	0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10,
	0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xab, 0x01, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44,
	0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x1a, 0x38, 0x0a, 0x0a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b,
	0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xfc, 0x01, 0x0a, 0x0c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x1a, 0x38, 0x0a, 0x0a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x70, 0x0a, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4b, 0x65,
	0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x2e, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x73, 0x22, 0x59, 0x0a, 0x15, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x22, 0xd6, 0x01, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63,
	0x6b, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x69,
	0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x32, 0xd0, 0x03, 0x0a, 0x10, 0x49,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x50, 0x43, 0x12,
	0x49, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x2a, 0x2e, 0x69,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x64, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17, 0x5a,
	0x15, 0x2f, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x72, 0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_helios_meta_rpc_proto_rawDescOnce sync.Once
	file_goadesign_goagen_helios_meta_rpc_proto_rawDescData = file_goadesign_goagen_helios_meta_rpc_proto_rawDesc
)

func file_goadesign_goagen_helios_meta_rpc_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_helios_meta_rpc_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_helios_meta_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_helios_meta_rpc_proto_rawDescData)
	})
	return file_goadesign_goagen_helios_meta_rpc_proto_rawDescData
}

var file_goadesign_goagen_helios_meta_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_goadesign_goagen_helios_meta_rpc_proto_goTypes = []interface{}{
	(*PingRequest)(nil),             // 0: helios_meta_rpc.PingRequest
	(*PingResponse)(nil),            // 1: helios_meta_rpc.PingResponse
	(*VersionRequest)(nil),          // 2: helios_meta_rpc.VersionRequest
	(*VersionResponse)(nil),         // 3: helios_meta_rpc.VersionResponse
	(*InfoRequest)(nil),             // 4: helios_meta_rpc.InfoRequest
	(*InfoResponse)(nil),            // 5: helios_meta_rpc.InfoResponse
	(*StreamKeepaliveRequest)(nil),  // 6: helios_meta_rpc.StreamKeepaliveRequest
	(*StreamKeepaliveResponse)(nil), // 7: helios_meta_rpc.StreamKeepaliveResponse
	(*TokenMetadataRequest)(nil),    // 8: helios_meta_rpc.TokenMetadataRequest
	(*TokenMetadataResponse)(nil),   // 9: helios_meta_rpc.TokenMetadataResponse
	(*TokenMetadataElement)(nil),    // 10: helios_meta_rpc.TokenMetadataElement
	nil,                             // 11: helios_meta_rpc.VersionResponse.BuildEntry
	nil,                             // 12: helios_meta_rpc.InfoResponse.BuildEntry
}
var file_goadesign_goagen_helios_meta_rpc_proto_depIdxs = []int32{
	11, // 0: helios_meta_rpc.VersionResponse.build:type_name -> helios_meta_rpc.VersionResponse.BuildEntry
	12, // 1: helios_meta_rpc.InfoResponse.build:type_name -> helios_meta_rpc.InfoResponse.BuildEntry
	10, // 2: helios_meta_rpc.TokenMetadataResponse.tokens:type_name -> helios_meta_rpc.TokenMetadataElement
	0,  // 3: helios_meta_rpc.InjectiveMetaRPC.Ping:input_type -> helios_meta_rpc.PingRequest
	2,  // 4: helios_meta_rpc.InjectiveMetaRPC.Version:input_type -> helios_meta_rpc.VersionRequest
	4,  // 5: helios_meta_rpc.InjectiveMetaRPC.Info:input_type -> helios_meta_rpc.InfoRequest
	6,  // 6: helios_meta_rpc.InjectiveMetaRPC.StreamKeepalive:input_type -> helios_meta_rpc.StreamKeepaliveRequest
	8,  // 7: helios_meta_rpc.InjectiveMetaRPC.TokenMetadata:input_type -> helios_meta_rpc.TokenMetadataRequest
	1,  // 8: helios_meta_rpc.InjectiveMetaRPC.Ping:output_type -> helios_meta_rpc.PingResponse
	3,  // 9: helios_meta_rpc.InjectiveMetaRPC.Version:output_type -> helios_meta_rpc.VersionResponse
	5,  // 10: helios_meta_rpc.InjectiveMetaRPC.Info:output_type -> helios_meta_rpc.InfoResponse
	7,  // 11: helios_meta_rpc.InjectiveMetaRPC.StreamKeepalive:output_type -> helios_meta_rpc.StreamKeepaliveResponse
	9,  // 12: helios_meta_rpc.InjectiveMetaRPC.TokenMetadata:output_type -> helios_meta_rpc.TokenMetadataResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_helios_meta_rpc_proto_init() }
func file_goadesign_goagen_helios_meta_rpc_proto_init() {
	if File_goadesign_goagen_helios_meta_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamKeepaliveRequest); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamKeepaliveResponse); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMetadataRequest); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMetadataResponse); i {
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
		file_goadesign_goagen_helios_meta_rpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMetadataElement); i {
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
			RawDescriptor: file_goadesign_goagen_helios_meta_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_helios_meta_rpc_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_helios_meta_rpc_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_helios_meta_rpc_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_helios_meta_rpc_proto = out.File
	file_goadesign_goagen_helios_meta_rpc_proto_rawDesc = nil
	file_goadesign_goagen_helios_meta_rpc_proto_goTypes = nil
	file_goadesign_goagen_helios_meta_rpc_proto_depIdxs = nil
}
