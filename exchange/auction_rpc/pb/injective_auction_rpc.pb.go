// Code generated with goa v3.5.2, DO NOT EDIT.
//
// HeliosAuctionRPC protocol buffer definition
//
// Command:
// $ goa gen github.com/Helios-Chain-Labs/-indexer/api/design -o ../

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: _auction_rpc.proto

package _auction_rpcpb

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

type AuctionEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The auction round number. -1 for latest.
	Round int64 `protobuf:"zigzag64,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *AuctionEndpointRequest) Reset() {
	*x = AuctionEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionEndpointRequest) ProtoMessage() {}

func (x *AuctionEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionEndpointRequest.ProtoReflect.Descriptor instead.
func (*AuctionEndpointRequest) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *AuctionEndpointRequest) GetRound() int64 {
	if x != nil {
		return x.Round
	}
	return 0
}

type AuctionEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The auction
	Auction *Auction `protobuf:"bytes,1,opt,name=auction,proto3" json:"auction,omitempty"`
	// Bids of the auction
	Bids []*Bid `protobuf:"bytes,2,rep,name=bids,proto3" json:"bids,omitempty"`
}

func (x *AuctionEndpointResponse) Reset() {
	*x = AuctionEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionEndpointResponse) ProtoMessage() {}

func (x *AuctionEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionEndpointResponse.ProtoReflect.Descriptor instead.
func (*AuctionEndpointResponse) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *AuctionEndpointResponse) GetAuction() *Auction {
	if x != nil {
		return x.Auction
	}
	return nil
}

func (x *AuctionEndpointResponse) GetBids() []*Bid {
	if x != nil {
		return x.Bids
	}
	return nil
}

type Auction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account address of the auction winner
	Winner string `protobuf:"bytes,1,opt,name=winner,proto3" json:"winner,omitempty"`
	// Coins in the basket
	Basket           []*Coin `protobuf:"bytes,2,rep,name=basket,proto3" json:"basket,omitempty"`
	WinningBidAmount string  `protobuf:"bytes,3,opt,name=winning_bid_amount,json=winningBidAmount,proto3" json:"winning_bid_amount,omitempty"`
	Round            uint64  `protobuf:"varint,4,opt,name=round,proto3" json:"round,omitempty"`
	// Auction end timestamp in UNIX millis.
	EndTimestamp int64 `protobuf:"zigzag64,5,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
	// UpdatedAt timestamp in UNIX millis.
	UpdatedAt int64 `protobuf:"zigzag64,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Auction) Reset() {
	*x = Auction{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction) ProtoMessage() {}

func (x *Auction) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auction.ProtoReflect.Descriptor instead.
func (*Auction) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *Auction) GetWinner() string {
	if x != nil {
		return x.Winner
	}
	return ""
}

func (x *Auction) GetBasket() []*Coin {
	if x != nil {
		return x.Basket
	}
	return nil
}

func (x *Auction) GetWinningBidAmount() string {
	if x != nil {
		return x.WinningBidAmount
	}
	return ""
}

func (x *Auction) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *Auction) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

func (x *Auction) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Denom of the coin
	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *Coin) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Coin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account address of the bidder
	Bidder string `protobuf:"bytes,1,opt,name=bidder,proto3" json:"bidder,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// Bid timestamp in UNIX millis.
	Timestamp int64 `protobuf:"zigzag64,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *Bid) GetBidder() string {
	if x != nil {
		return x.Bidder
	}
	return ""
}

func (x *Bid) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Bid) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type AuctionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuctionsRequest) Reset() {
	*x = AuctionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionsRequest) ProtoMessage() {}

func (x *AuctionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionsRequest.ProtoReflect.Descriptor instead.
func (*AuctionsRequest) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{5}
}

type AuctionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The historical auctions
	Auctions []*Auction `protobuf:"bytes,1,rep,name=auctions,proto3" json:"auctions,omitempty"`
}

func (x *AuctionsResponse) Reset() {
	*x = AuctionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuctionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionsResponse) ProtoMessage() {}

func (x *AuctionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionsResponse.ProtoReflect.Descriptor instead.
func (*AuctionsResponse) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *AuctionsResponse) GetAuctions() []*Auction {
	if x != nil {
		return x.Auctions
	}
	return nil
}

type StreamBidsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamBidsRequest) Reset() {
	*x = StreamBidsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamBidsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamBidsRequest) ProtoMessage() {}

func (x *StreamBidsRequest) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamBidsRequest.ProtoReflect.Descriptor instead.
func (*StreamBidsRequest) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{7}
}

type StreamBidsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Account address of the bidder
	Bidder    string `protobuf:"bytes,1,opt,name=bidder,proto3" json:"bidder,omitempty"`
	BidAmount string `protobuf:"bytes,2,opt,name=bid_amount,json=bidAmount,proto3" json:"bid_amount,omitempty"`
	Round     uint64 `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	// Operation timestamp in UNIX millis.
	Timestamp int64 `protobuf:"zigzag64,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StreamBidsResponse) Reset() {
	*x = StreamBidsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file__auction_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamBidsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamBidsResponse) ProtoMessage() {}

func (x *StreamBidsResponse) ProtoReflect() protoreflect.Message {
	mi := &file__auction_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamBidsResponse.ProtoReflect.Descriptor instead.
func (*StreamBidsResponse) Descriptor() ([]byte, []int) {
	return file__auction_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *StreamBidsResponse) GetBidder() string {
	if x != nil {
		return x.Bidder
	}
	return ""
}

func (x *StreamBidsResponse) GetBidAmount() string {
	if x != nil {
		return x.BidAmount
	}
	return ""
}

func (x *StreamBidsResponse) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *StreamBidsResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File__auction_rpc_proto protoreflect.FileDescriptor

var file__auction_rpc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x70, 0x63, 0x22, 0x2e, 0x0a, 0x16, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x17, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x62, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x42, 0x69, 0x64, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x07, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x33,
	0x0a, 0x06, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x06, 0x62, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x62,
	0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x77, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0c,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x34, 0x0a, 0x04, 0x43,
	0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x53, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x64, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x10, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x08, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x42, 0x69, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7f,
	0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x69, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x62, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32,
	0xc9, 0x02, 0x0a, 0x13, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x50, 0x43, 0x12, 0x70, 0x0a, 0x0f, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x69, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x69, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x08, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x42, 0x69, 0x64, 0x73, 0x12, 0x28, 0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x42, 0x69, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x69, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x2f,
	0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x70, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__auction_rpc_proto_rawDescOnce sync.Once
	file__auction_rpc_proto_rawDescData = file__auction_rpc_proto_rawDesc
)

func file__auction_rpc_proto_rawDescGZIP() []byte {
	file__auction_rpc_proto_rawDescOnce.Do(func() {
		file__auction_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file__auction_rpc_proto_rawDescData)
	})
	return file__auction_rpc_proto_rawDescData
}

var file__auction_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file__auction_rpc_proto_goTypes = []interface{}{
	(*AuctionEndpointRequest)(nil),  // 0: _auction_rpc.AuctionEndpointRequest
	(*AuctionEndpointResponse)(nil), // 1: _auction_rpc.AuctionEndpointResponse
	(*Auction)(nil),                 // 2: _auction_rpc.Auction
	(*Coin)(nil),                    // 3: _auction_rpc.Coin
	(*Bid)(nil),                     // 4: _auction_rpc.Bid
	(*AuctionsRequest)(nil),         // 5: _auction_rpc.AuctionsRequest
	(*AuctionsResponse)(nil),        // 6: _auction_rpc.AuctionsResponse
	(*StreamBidsRequest)(nil),       // 7: _auction_rpc.StreamBidsRequest
	(*StreamBidsResponse)(nil),      // 8: _auction_rpc.StreamBidsResponse
}
var file__auction_rpc_proto_depIdxs = []int32{
	2, // 0: _auction_rpc.AuctionEndpointResponse.auction:type_name -> _auction_rpc.Auction
	4, // 1: _auction_rpc.AuctionEndpointResponse.bids:type_name -> _auction_rpc.Bid
	3, // 2: _auction_rpc.Auction.basket:type_name -> _auction_rpc.Coin
	2, // 3: _auction_rpc.AuctionsResponse.auctions:type_name -> _auction_rpc.Auction
	0, // 4: _auction_rpc.HeliosAuctionRPC.AuctionEndpoint:input_type -> _auction_rpc.AuctionEndpointRequest
	5, // 5: _auction_rpc.HeliosAuctionRPC.Auctions:input_type -> _auction_rpc.AuctionsRequest
	7, // 6: _auction_rpc.HeliosAuctionRPC.StreamBids:input_type -> _auction_rpc.StreamBidsRequest
	1, // 7: _auction_rpc.HeliosAuctionRPC.AuctionEndpoint:output_type -> _auction_rpc.AuctionEndpointResponse
	6, // 8: _auction_rpc.HeliosAuctionRPC.Auctions:output_type -> _auction_rpc.AuctionsResponse
	8, // 9: _auction_rpc.HeliosAuctionRPC.StreamBids:output_type -> _auction_rpc.StreamBidsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file__auction_rpc_proto_init() }
func file__auction_rpc_proto_init() {
	if File__auction_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__auction_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionEndpointRequest); i {
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
		file__auction_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionEndpointResponse); i {
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
		file__auction_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auction); i {
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
		file__auction_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
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
		file__auction_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
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
		file__auction_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionsRequest); i {
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
		file__auction_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuctionsResponse); i {
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
		file__auction_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamBidsRequest); i {
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
		file__auction_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamBidsResponse); i {
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
			RawDescriptor: file__auction_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file__auction_rpc_proto_goTypes,
		DependencyIndexes: file__auction_rpc_proto_depIdxs,
		MessageInfos:      file__auction_rpc_proto_msgTypes,
	}.Build()
	File__auction_rpc_proto = out.File
	file__auction_rpc_proto_rawDesc = nil
	file__auction_rpc_proto_goTypes = nil
	file__auction_rpc_proto_depIdxs = nil
}
