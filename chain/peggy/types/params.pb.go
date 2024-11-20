// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: helios/peggy/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	PeggyId                       string                      `protobuf:"bytes,1,opt,name=peggy_id,json=peggyId,proto3" json:"peggy_id,omitempty"`
	ContractSourceHash            string                      `protobuf:"bytes,2,opt,name=contract_source_hash,json=contractSourceHash,proto3" json:"contract_source_hash,omitempty"`
	BridgeEthereumAddress         string                      `protobuf:"bytes,3,opt,name=bridge_ethereum_address,json=bridgeEthereumAddress,proto3" json:"bridge_ethereum_address,omitempty"`
	BridgeChainId                 uint64                      `protobuf:"varint,4,opt,name=bridge_chain_id,json=bridgeChainId,proto3" json:"bridge_chain_id,omitempty"`
	SignedValsetsWindow           uint64                      `protobuf:"varint,5,opt,name=signed_valsets_window,json=signedValsetsWindow,proto3" json:"signed_valsets_window,omitempty"`
	SignedBatchesWindow           uint64                      `protobuf:"varint,6,opt,name=signed_batches_window,json=signedBatchesWindow,proto3" json:"signed_batches_window,omitempty"`
	SignedClaimsWindow            uint64                      `protobuf:"varint,7,opt,name=signed_claims_window,json=signedClaimsWindow,proto3" json:"signed_claims_window,omitempty"`
	TargetBatchTimeout            uint64                      `protobuf:"varint,8,opt,name=target_batch_timeout,json=targetBatchTimeout,proto3" json:"target_batch_timeout,omitempty"`
	AverageBlockTime              uint64                      `protobuf:"varint,9,opt,name=average_block_time,json=averageBlockTime,proto3" json:"average_block_time,omitempty"`
	AverageEthereumBlockTime      uint64                      `protobuf:"varint,10,opt,name=average_ethereum_block_time,json=averageEthereumBlockTime,proto3" json:"average_ethereum_block_time,omitempty"`
	SlashFractionValset           cosmossdk_io_math.LegacyDec `protobuf:"bytes,11,opt,name=slash_fraction_valset,json=slashFractionValset,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_fraction_valset"`
	SlashFractionBatch            cosmossdk_io_math.LegacyDec `protobuf:"bytes,12,opt,name=slash_fraction_batch,json=slashFractionBatch,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_fraction_batch"`
	SlashFractionClaim            cosmossdk_io_math.LegacyDec `protobuf:"bytes,13,opt,name=slash_fraction_claim,json=slashFractionClaim,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_fraction_claim"`
	SlashFractionConflictingClaim cosmossdk_io_math.LegacyDec `protobuf:"bytes,14,opt,name=slash_fraction_conflicting_claim,json=slashFractionConflictingClaim,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_fraction_conflicting_claim"`
	UnbondSlashingValsetsWindow   uint64                      `protobuf:"varint,15,opt,name=unbond_slashing_valsets_window,json=unbondSlashingValsetsWindow,proto3" json:"unbond_slashing_valsets_window,omitempty"`
	SlashFractionBadEthSignature  cosmossdk_io_math.LegacyDec `protobuf:"bytes,16,opt,name=slash_fraction_bad_eth_signature,json=slashFractionBadEthSignature,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_fraction_bad_eth_signature"`
	CosmosCoinDenom               string                      `protobuf:"bytes,17,opt,name=cosmos_coin_denom,json=cosmosCoinDenom,proto3" json:"cosmos_coin_denom,omitempty"`
	CosmosCoinErc20Contract       string                      `protobuf:"bytes,18,opt,name=cosmos_coin_erc20_contract,json=cosmosCoinErc20Contract,proto3" json:"cosmos_coin_erc20_contract,omitempty"`
	ClaimSlashingEnabled          bool                        `protobuf:"varint,19,opt,name=claim_slashing_enabled,json=claimSlashingEnabled,proto3" json:"claim_slashing_enabled,omitempty"`
	BridgeContractStartHeight     uint64                      `protobuf:"varint,20,opt,name=bridge_contract_start_height,json=bridgeContractStartHeight,proto3" json:"bridge_contract_start_height,omitempty"`
	ValsetReward                  types.Coin                  `protobuf:"bytes,21,opt,name=valset_reward,json=valsetReward,proto3" json:"valset_reward"`
	Admins                        []string                    `protobuf:"bytes,22,rep,name=admins,proto3" json:"admins,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9072ab4399c3b996, []int{0}
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

func (m *Params) GetPeggyId() string {
	if m != nil {
		return m.PeggyId
	}
	return ""
}

func (m *Params) GetContractSourceHash() string {
	if m != nil {
		return m.ContractSourceHash
	}
	return ""
}

func (m *Params) GetBridgeEthereumAddress() string {
	if m != nil {
		return m.BridgeEthereumAddress
	}
	return ""
}

func (m *Params) GetBridgeChainId() uint64 {
	if m != nil {
		return m.BridgeChainId
	}
	return 0
}

func (m *Params) GetSignedValsetsWindow() uint64 {
	if m != nil {
		return m.SignedValsetsWindow
	}
	return 0
}

func (m *Params) GetSignedBatchesWindow() uint64 {
	if m != nil {
		return m.SignedBatchesWindow
	}
	return 0
}

func (m *Params) GetSignedClaimsWindow() uint64 {
	if m != nil {
		return m.SignedClaimsWindow
	}
	return 0
}

func (m *Params) GetTargetBatchTimeout() uint64 {
	if m != nil {
		return m.TargetBatchTimeout
	}
	return 0
}

func (m *Params) GetAverageBlockTime() uint64 {
	if m != nil {
		return m.AverageBlockTime
	}
	return 0
}

func (m *Params) GetAverageEthereumBlockTime() uint64 {
	if m != nil {
		return m.AverageEthereumBlockTime
	}
	return 0
}

func (m *Params) GetUnbondSlashingValsetsWindow() uint64 {
	if m != nil {
		return m.UnbondSlashingValsetsWindow
	}
	return 0
}

func (m *Params) GetCosmosCoinDenom() string {
	if m != nil {
		return m.CosmosCoinDenom
	}
	return ""
}

func (m *Params) GetCosmosCoinErc20Contract() string {
	if m != nil {
		return m.CosmosCoinErc20Contract
	}
	return ""
}

func (m *Params) GetClaimSlashingEnabled() bool {
	if m != nil {
		return m.ClaimSlashingEnabled
	}
	return false
}

func (m *Params) GetBridgeContractStartHeight() uint64 {
	if m != nil {
		return m.BridgeContractStartHeight
	}
	return 0
}

func (m *Params) GetValsetReward() types.Coin {
	if m != nil {
		return m.ValsetReward
	}
	return types.Coin{}
}

func (m *Params) GetAdmins() []string {
	if m != nil {
		return m.Admins
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "helios.peggy.v1.Params")
}

func init() { proto.RegisterFile("helios/peggy/v1/params.proto", fileDescriptor_9072ab4399c3b996) }

var fileDescriptor_9072ab4399c3b996 = []byte{
	// 786 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0x63, 0xb6, 0x74, 0xdb, 0xd9, 0x94, 0x6c, 0xa7, 0x49, 0x77, 0xfa, 0x07, 0xaf, 0x05,
	0x12, 0x8a, 0x56, 0x8b, 0xdd, 0x14, 0xc4, 0x01, 0x84, 0x10, 0x49, 0xb3, 0xda, 0x95, 0x38, 0xa0,
	0x14, 0x58, 0x89, 0xcb, 0x68, 0xec, 0x99, 0xb5, 0x47, 0xb5, 0x3d, 0x91, 0x67, 0x92, 0xaa, 0x37,
	0xce, 0x9c, 0xf8, 0x28, 0x7c, 0x8c, 0x3d, 0xee, 0x11, 0x21, 0xb4, 0x42, 0xed, 0x01, 0x89, 0x4f,
	0x81, 0xe6, 0x1d, 0x3b, 0x49, 0x03, 0x87, 0x8a, 0x4b, 0x64, 0xcf, 0xef, 0x79, 0xde, 0x77, 0xf2,
	0xce, 0xe4, 0x09, 0x3a, 0xce, 0x44, 0x2e, 0x95, 0x8e, 0xa6, 0x22, 0x4d, 0xaf, 0xa2, 0xf9, 0x20,
	0x9a, 0xb2, 0x8a, 0x15, 0x3a, 0x9c, 0x56, 0xca, 0x28, 0xdc, 0x71, 0x34, 0x04, 0x1a, 0xce, 0x07,
	0x87, 0xdd, 0x54, 0xa5, 0x0a, 0x58, 0x64, 0x9f, 0x9c, 0xec, 0xd0, 0x4f, 0x94, 0x2e, 0x94, 0x8e,
	0x62, 0xa6, 0x45, 0x34, 0x1f, 0xc4, 0xc2, 0xb0, 0x41, 0x94, 0x28, 0x59, 0xd6, 0x7c, 0x97, 0x15,
	0xb2, 0x54, 0x11, 0x7c, 0xba, 0xa5, 0x0f, 0xfe, 0x46, 0x68, 0xf3, 0x5b, 0x68, 0x85, 0x0f, 0xd0,
	0x16, 0xd4, 0xa7, 0x92, 0x13, 0x2f, 0xf0, 0xfa, 0xdb, 0x93, 0xfb, 0xf0, 0xfe, 0x82, 0xe3, 0x13,
	0xd4, 0x4d, 0x54, 0x69, 0x2a, 0x96, 0x18, 0xaa, 0xd5, 0xac, 0x4a, 0x04, 0xcd, 0x98, 0xce, 0xc8,
	0x3b, 0x20, 0xc3, 0x0d, 0x3b, 0x07, 0xf4, 0x9c, 0xe9, 0x0c, 0x7f, 0x86, 0x1e, 0xc5, 0x95, 0xe4,
	0xa9, 0xa0, 0xc2, 0x64, 0xa2, 0x12, 0xb3, 0x82, 0x32, 0xce, 0x2b, 0xa1, 0x35, 0xb9, 0x07, 0xa6,
	0x9e, 0xc3, 0xe3, 0x9a, 0x7e, 0xed, 0x20, 0xfe, 0x08, 0x75, 0x6a, 0x5f, 0x92, 0x31, 0x59, 0xda,
	0xbd, 0x6c, 0x04, 0x5e, 0x7f, 0x63, 0xb2, 0xe3, 0x96, 0x47, 0x76, 0xf5, 0x05, 0xc7, 0xa7, 0xa8,
	0xa7, 0x65, 0x5a, 0x0a, 0x4e, 0xe7, 0x2c, 0xd7, 0xc2, 0x68, 0x7a, 0x29, 0x4b, 0xae, 0x2e, 0xc9,
	0xbb, 0xa0, 0xde, 0x73, 0xf0, 0x07, 0xc7, 0x5e, 0x02, 0x5a, 0xf1, 0xc4, 0xcc, 0x24, 0x99, 0x58,
	0x78, 0x36, 0x57, 0x3d, 0x43, 0xc7, 0x6a, 0xcf, 0x09, 0xea, 0xd6, 0x9e, 0x24, 0x67, 0xb2, 0x58,
	0x58, 0xee, 0x83, 0x05, 0x3b, 0x36, 0x02, 0xb4, 0x74, 0x18, 0x56, 0xa5, 0xc2, 0xb8, 0x2e, 0xd4,
	0xc8, 0x42, 0xa8, 0x99, 0x21, 0x5b, 0xce, 0xe1, 0x18, 0x34, 0xf9, 0xce, 0x11, 0xfc, 0x14, 0x61,
	0x36, 0x17, 0x15, 0x4b, 0x05, 0x8d, 0x73, 0x95, 0x5c, 0x80, 0x85, 0x6c, 0x83, 0xfe, 0x61, 0x4d,
	0x86, 0x16, 0x58, 0x03, 0xfe, 0x12, 0x1d, 0x35, 0xea, 0xc5, 0x68, 0x57, 0x6c, 0x08, 0x6c, 0xa4,
	0x96, 0x34, 0xe3, 0x5d, 0xda, 0x5f, 0xa2, 0x9e, 0xce, 0x99, 0xce, 0xe8, 0x2b, 0x7b, 0x62, 0x52,
	0x95, 0xf5, 0x00, 0xc9, 0x83, 0xc0, 0xeb, 0xb7, 0x87, 0x1f, 0xbe, 0x7e, 0xfb, 0xb8, 0xf5, 0xfb,
	0xdb, 0xc7, 0x47, 0xee, 0x2a, 0x69, 0x7e, 0x11, 0x4a, 0x15, 0x15, 0xcc, 0x64, 0xe1, 0x37, 0x22,
	0x65, 0xc9, 0xd5, 0x99, 0x48, 0x26, 0x7b, 0x50, 0xe1, 0x59, 0x5d, 0xc0, 0x0d, 0x19, 0x7f, 0x8f,
	0xba, 0x6b, 0x85, 0xe1, 0xfb, 0x93, 0xf6, 0xdd, 0xeb, 0xe2, 0x5b, 0x75, 0x61, 0x46, 0xff, 0x51,
	0x16, 0x0e, 0x82, 0xec, 0xfc, 0xdf, 0xb2, 0x70, 0x58, 0x38, 0x47, 0xc1, 0x7a, 0x59, 0x55, 0xbe,
	0xca, 0x65, 0x62, 0x64, 0x99, 0xd6, 0x2d, 0xde, 0xbb, 0x7b, 0x8b, 0xf7, 0x6f, 0xb7, 0x58, 0x96,
	0x72, 0xdd, 0x46, 0xc8, 0x9f, 0x95, 0xb1, 0x2a, 0x39, 0x05, 0x9d, 0x6d, 0xb1, 0x76, 0x6d, 0x3b,
	0x70, 0x6c, 0x47, 0x4e, 0x75, 0x5e, 0x8b, 0x6e, 0x5f, 0xdf, 0x8b, 0x7f, 0x6d, 0x39, 0x66, 0xdc,
	0xde, 0x01, 0x6a, 0x6f, 0x21, 0x33, 0xb3, 0x4a, 0x90, 0x87, 0x77, 0xdf, 0xf2, 0xf1, 0xda, 0xb0,
	0xf9, 0xd8, 0x64, 0xe7, 0x4d, 0x21, 0xfc, 0x04, 0xed, 0x3a, 0x33, 0xb5, 0xf9, 0x41, 0xb9, 0x28,
	0x55, 0x41, 0x76, 0xe1, 0x97, 0xdb, 0x71, 0x60, 0xa4, 0x64, 0x79, 0x66, 0x97, 0xf1, 0x17, 0xe8,
	0x70, 0x55, 0x2b, 0xaa, 0xe4, 0xf4, 0x84, 0x36, 0x99, 0x40, 0x30, 0x98, 0x1e, 0x2d, 0x4d, 0x63,
	0xcb, 0x47, 0x35, 0xc6, 0x9f, 0xa2, 0x7d, 0x98, 0xf6, 0x72, 0x32, 0xa2, 0x64, 0x71, 0x2e, 0x38,
	0xd9, 0x0b, 0xbc, 0xfe, 0xd6, 0xa4, 0x0b, 0xb4, 0x99, 0xc8, 0xd8, 0x31, 0xfc, 0x15, 0x3a, 0x6e,
	0x62, 0x62, 0x91, 0x4b, 0x86, 0x55, 0x86, 0x66, 0x42, 0xa6, 0x99, 0x21, 0x5d, 0x18, 0xe7, 0x41,
	0x9d, 0x19, 0x4d, 0x3c, 0x59, 0xc5, 0x73, 0x10, 0xe0, 0x33, 0xb4, 0xe3, 0x4e, 0x80, 0x56, 0xe2,
	0x92, 0x55, 0x9c, 0xf4, 0x02, 0xaf, 0xff, 0xe0, 0xf4, 0x20, 0x74, 0xfb, 0x0c, 0x6d, 0x84, 0x86,
	0x75, 0x84, 0x86, 0x76, 0xd7, 0xc3, 0x0d, 0x3b, 0xd4, 0x49, 0xdb, 0xb9, 0x26, 0x60, 0xc2, 0xfb,
	0x68, 0x93, 0xf1, 0x42, 0x96, 0x9a, 0xec, 0x07, 0xf7, 0xfa, 0xdb, 0x93, 0xfa, 0xed, 0xf3, 0xde,
	0x4f, 0x7f, 0x04, 0xad, 0x9f, 0xff, 0xfa, 0xf5, 0x49, 0xdb, 0xe5, 0xb9, 0x4b, 0xd8, 0xe1, 0xb3,
	0xd7, 0xd7, 0xbe, 0xf7, 0xe6, 0xda, 0xf7, 0xfe, 0xbc, 0xf6, 0xbd, 0x5f, 0x6e, 0xfc, 0xd6, 0x9b,
	0x1b, 0xbf, 0xf5, 0xdb, 0x8d, 0xdf, 0xfa, 0xf1, 0xa9, 0x0b, 0xf8, 0x8f, 0x13, 0x55, 0x89, 0xa8,
	0x79, 0xb6, 0x51, 0x17, 0x15, 0x8a, 0xcf, 0x72, 0xd1, 0xfc, 0x31, 0x98, 0xab, 0xa9, 0xd0, 0xf1,
	0x26, 0x64, 0xf7, 0x27, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x70, 0xed, 0x39, 0x97, 0x35, 0x06,
	0x00, 0x00,
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
	if len(m.Admins) > 0 {
		for iNdEx := len(m.Admins) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Admins[iNdEx])
			copy(dAtA[i:], m.Admins[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.Admins[iNdEx])))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0xb2
		}
	}
	{
		size, err := m.ValsetReward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0xaa
	if m.BridgeContractStartHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BridgeContractStartHeight))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.ClaimSlashingEnabled {
		i--
		if m.ClaimSlashingEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if len(m.CosmosCoinErc20Contract) > 0 {
		i -= len(m.CosmosCoinErc20Contract)
		copy(dAtA[i:], m.CosmosCoinErc20Contract)
		i = encodeVarintParams(dAtA, i, uint64(len(m.CosmosCoinErc20Contract)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.CosmosCoinDenom) > 0 {
		i -= len(m.CosmosCoinDenom)
		copy(dAtA[i:], m.CosmosCoinDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.CosmosCoinDenom)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	{
		size := m.SlashFractionBadEthSignature.Size()
		i -= size
		if _, err := m.SlashFractionBadEthSignature.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x82
	if m.UnbondSlashingValsetsWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnbondSlashingValsetsWindow))
		i--
		dAtA[i] = 0x78
	}
	{
		size := m.SlashFractionConflictingClaim.Size()
		i -= size
		if _, err := m.SlashFractionConflictingClaim.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.SlashFractionClaim.Size()
		i -= size
		if _, err := m.SlashFractionClaim.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.SlashFractionBatch.Size()
		i -= size
		if _, err := m.SlashFractionBatch.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.SlashFractionValset.Size()
		i -= size
		if _, err := m.SlashFractionValset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if m.AverageEthereumBlockTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AverageEthereumBlockTime))
		i--
		dAtA[i] = 0x50
	}
	if m.AverageBlockTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AverageBlockTime))
		i--
		dAtA[i] = 0x48
	}
	if m.TargetBatchTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TargetBatchTimeout))
		i--
		dAtA[i] = 0x40
	}
	if m.SignedClaimsWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SignedClaimsWindow))
		i--
		dAtA[i] = 0x38
	}
	if m.SignedBatchesWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SignedBatchesWindow))
		i--
		dAtA[i] = 0x30
	}
	if m.SignedValsetsWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SignedValsetsWindow))
		i--
		dAtA[i] = 0x28
	}
	if m.BridgeChainId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BridgeChainId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.BridgeEthereumAddress) > 0 {
		i -= len(m.BridgeEthereumAddress)
		copy(dAtA[i:], m.BridgeEthereumAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BridgeEthereumAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractSourceHash) > 0 {
		i -= len(m.ContractSourceHash)
		copy(dAtA[i:], m.ContractSourceHash)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ContractSourceHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PeggyId) > 0 {
		i -= len(m.PeggyId)
		copy(dAtA[i:], m.PeggyId)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PeggyId)))
		i--
		dAtA[i] = 0xa
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
	l = len(m.PeggyId)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ContractSourceHash)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.BridgeEthereumAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.BridgeChainId != 0 {
		n += 1 + sovParams(uint64(m.BridgeChainId))
	}
	if m.SignedValsetsWindow != 0 {
		n += 1 + sovParams(uint64(m.SignedValsetsWindow))
	}
	if m.SignedBatchesWindow != 0 {
		n += 1 + sovParams(uint64(m.SignedBatchesWindow))
	}
	if m.SignedClaimsWindow != 0 {
		n += 1 + sovParams(uint64(m.SignedClaimsWindow))
	}
	if m.TargetBatchTimeout != 0 {
		n += 1 + sovParams(uint64(m.TargetBatchTimeout))
	}
	if m.AverageBlockTime != 0 {
		n += 1 + sovParams(uint64(m.AverageBlockTime))
	}
	if m.AverageEthereumBlockTime != 0 {
		n += 1 + sovParams(uint64(m.AverageEthereumBlockTime))
	}
	l = m.SlashFractionValset.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashFractionBatch.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashFractionClaim.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashFractionConflictingClaim.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.UnbondSlashingValsetsWindow != 0 {
		n += 1 + sovParams(uint64(m.UnbondSlashingValsetsWindow))
	}
	l = m.SlashFractionBadEthSignature.Size()
	n += 2 + l + sovParams(uint64(l))
	l = len(m.CosmosCoinDenom)
	if l > 0 {
		n += 2 + l + sovParams(uint64(l))
	}
	l = len(m.CosmosCoinErc20Contract)
	if l > 0 {
		n += 2 + l + sovParams(uint64(l))
	}
	if m.ClaimSlashingEnabled {
		n += 3
	}
	if m.BridgeContractStartHeight != 0 {
		n += 2 + sovParams(uint64(m.BridgeContractStartHeight))
	}
	l = m.ValsetReward.Size()
	n += 2 + l + sovParams(uint64(l))
	if len(m.Admins) > 0 {
		for _, s := range m.Admins {
			l = len(s)
			n += 2 + l + sovParams(uint64(l))
		}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeggyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeggyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractSourceHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractSourceHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeEthereumAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BridgeEthereumAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeChainId", wireType)
			}
			m.BridgeChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BridgeChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedValsetsWindow", wireType)
			}
			m.SignedValsetsWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedValsetsWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBatchesWindow", wireType)
			}
			m.SignedBatchesWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBatchesWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedClaimsWindow", wireType)
			}
			m.SignedClaimsWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedClaimsWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBatchTimeout", wireType)
			}
			m.TargetBatchTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetBatchTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageBlockTime", wireType)
			}
			m.AverageBlockTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AverageBlockTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AverageEthereumBlockTime", wireType)
			}
			m.AverageEthereumBlockTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AverageEthereumBlockTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionValset", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionValset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionBatch", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionBatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionClaim", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionConflictingClaim", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionConflictingClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondSlashingValsetsWindow", wireType)
			}
			m.UnbondSlashingValsetsWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondSlashingValsetsWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionBadEthSignature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionBadEthSignature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmosCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosCoinErc20Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmosCoinErc20Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimSlashingEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.ClaimSlashingEnabled = bool(v != 0)
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeContractStartHeight", wireType)
			}
			m.BridgeContractStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BridgeContractStartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValsetReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admins", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admins = append(m.Admins, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
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
