// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: helios/peggy/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState struct
type GenesisState struct {
	Params                     *Params                        `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	LastObservedNonce          uint64                         `protobuf:"varint,2,opt,name=last_observed_nonce,json=lastObservedNonce,proto3" json:"last_observed_nonce,omitempty"`
	Valsets                    []*Valset                      `protobuf:"bytes,3,rep,name=valsets,proto3" json:"valsets,omitempty"`
	ValsetConfirms             []*MsgValsetConfirm            `protobuf:"bytes,4,rep,name=valset_confirms,json=valsetConfirms,proto3" json:"valset_confirms,omitempty"`
	Batches                    []*OutgoingTxBatch             `protobuf:"bytes,5,rep,name=batches,proto3" json:"batches,omitempty"`
	BatchConfirms              []*MsgConfirmBatch             `protobuf:"bytes,6,rep,name=batch_confirms,json=batchConfirms,proto3" json:"batch_confirms,omitempty"`
	Attestations               []*Attestation                 `protobuf:"bytes,7,rep,name=attestations,proto3" json:"attestations,omitempty"`
	OrchestratorAddresses      []*MsgSetOrchestratorAddresses `protobuf:"bytes,8,rep,name=orchestrator_addresses,json=orchestratorAddresses,proto3" json:"orchestrator_addresses,omitempty"`
	Erc20ToDenoms              []*ERC20ToDenom                `protobuf:"bytes,9,rep,name=erc20_to_denoms,json=erc20ToDenoms,proto3" json:"erc20_to_denoms,omitempty"`
	UnbatchedTransfers         []*OutgoingTransferTx          `protobuf:"bytes,10,rep,name=unbatched_transfers,json=unbatchedTransfers,proto3" json:"unbatched_transfers,omitempty"`
	LastObservedEthereumHeight uint64                         `protobuf:"varint,11,opt,name=last_observed_ethereum_height,json=lastObservedEthereumHeight,proto3" json:"last_observed_ethereum_height,omitempty"`
	LastOutgoingBatchId        uint64                         `protobuf:"varint,12,opt,name=last_outgoing_batch_id,json=lastOutgoingBatchId,proto3" json:"last_outgoing_batch_id,omitempty"`
	LastOutgoingPoolId         uint64                         `protobuf:"varint,13,opt,name=last_outgoing_pool_id,json=lastOutgoingPoolId,proto3" json:"last_outgoing_pool_id,omitempty"`
	LastObservedValset         Valset                         `protobuf:"bytes,14,opt,name=last_observed_valset,json=lastObservedValset,proto3" json:"last_observed_valset"`
	EthereumBlacklist          []string                       `protobuf:"bytes,15,rep,name=ethereum_blacklist,json=ethereumBlacklist,proto3" json:"ethereum_blacklist,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4b552146cd7e0bf, []int{0}
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

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetLastObservedNonce() uint64 {
	if m != nil {
		return m.LastObservedNonce
	}
	return 0
}

func (m *GenesisState) GetValsets() []*Valset {
	if m != nil {
		return m.Valsets
	}
	return nil
}

func (m *GenesisState) GetValsetConfirms() []*MsgValsetConfirm {
	if m != nil {
		return m.ValsetConfirms
	}
	return nil
}

func (m *GenesisState) GetBatches() []*OutgoingTxBatch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *GenesisState) GetBatchConfirms() []*MsgConfirmBatch {
	if m != nil {
		return m.BatchConfirms
	}
	return nil
}

func (m *GenesisState) GetAttestations() []*Attestation {
	if m != nil {
		return m.Attestations
	}
	return nil
}

func (m *GenesisState) GetOrchestratorAddresses() []*MsgSetOrchestratorAddresses {
	if m != nil {
		return m.OrchestratorAddresses
	}
	return nil
}

func (m *GenesisState) GetErc20ToDenoms() []*ERC20ToDenom {
	if m != nil {
		return m.Erc20ToDenoms
	}
	return nil
}

func (m *GenesisState) GetUnbatchedTransfers() []*OutgoingTransferTx {
	if m != nil {
		return m.UnbatchedTransfers
	}
	return nil
}

func (m *GenesisState) GetLastObservedEthereumHeight() uint64 {
	if m != nil {
		return m.LastObservedEthereumHeight
	}
	return 0
}

func (m *GenesisState) GetLastOutgoingBatchId() uint64 {
	if m != nil {
		return m.LastOutgoingBatchId
	}
	return 0
}

func (m *GenesisState) GetLastOutgoingPoolId() uint64 {
	if m != nil {
		return m.LastOutgoingPoolId
	}
	return 0
}

func (m *GenesisState) GetLastObservedValset() Valset {
	if m != nil {
		return m.LastObservedValset
	}
	return Valset{}
}

func (m *GenesisState) GetEthereumBlacklist() []string {
	if m != nil {
		return m.EthereumBlacklist
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "helios.peggy.v1.GenesisState")
}

func init() { proto.RegisterFile("helios/peggy/v1/genesis.proto", fileDescriptor_c4b552146cd7e0bf) }

var fileDescriptor_c4b552146cd7e0bf = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xcb, 0x52, 0xdb, 0x3e,
	0x14, 0xc6, 0x93, 0x3f, 0xfc, 0x49, 0x11, 0x97, 0x0c, 0xe2, 0x52, 0x4d, 0x0a, 0x69, 0x68, 0x37,
	0x2c, 0x20, 0x26, 0xb0, 0xeb, 0xaa, 0x84, 0x52, 0x4a, 0x67, 0x68, 0x18, 0x93, 0xe9, 0xa2, 0x1b,
	0x8f, 0x62, 0x1f, 0x1c, 0x4f, 0x6d, 0x2b, 0xa3, 0xa3, 0x64, 0xe0, 0x2d, 0xfa, 0x58, 0x2c, 0x59,
	0x76, 0xd5, 0xe9, 0x84, 0x17, 0xe9, 0x58, 0x92, 0x21, 0x17, 0xb2, 0x93, 0xf5, 0x7d, 0xbf, 0x4f,
	0x47, 0xd2, 0xb1, 0xc8, 0x4e, 0x17, 0xe2, 0x48, 0xa0, 0xd3, 0x83, 0x30, 0xbc, 0x73, 0x06, 0x0d,
	0x27, 0x84, 0x14, 0x30, 0xc2, 0x7a, 0x4f, 0x0a, 0x25, 0x68, 0xd9, 0xc8, 0x75, 0x2d, 0xd7, 0x07,
	0x8d, 0xca, 0x46, 0x28, 0x42, 0xa1, 0x35, 0x27, 0x1b, 0x19, 0x5b, 0xe5, 0xcd, 0x64, 0x8a, 0xba,
	0xeb, 0x81, 0xcd, 0xa8, 0x54, 0x26, 0xc5, 0x04, 0x43, 0x9c, 0x05, 0x76, 0xb8, 0xf2, 0xbb, 0x56,
	0xdc, 0x9d, 0x14, 0xb9, 0x52, 0x80, 0x8a, 0xab, 0x48, 0xa4, 0xd6, 0xb2, 0x3d, 0x69, 0xe9, 0x71,
	0xc9, 0x93, 0x3c, 0xbd, 0xea, 0x0b, 0x4c, 0x04, 0x3a, 0x1d, 0x8e, 0xe0, 0x0c, 0x1a, 0x1d, 0x50,
	0xbc, 0xe1, 0xf8, 0x22, 0xb2, 0xf4, 0xbb, 0x61, 0x89, 0x2c, 0x9f, 0x9b, 0xfd, 0x5e, 0x2b, 0xae,
	0x80, 0x3a, 0x64, 0xc1, 0x04, 0xb0, 0x62, 0xad, 0xb8, 0xb7, 0x74, 0xf4, 0xba, 0x3e, 0xb1, 0xff,
	0xfa, 0x95, 0x96, 0x5d, 0x6b, 0xa3, 0x75, 0xb2, 0x1e, 0x73, 0x54, 0x9e, 0xe8, 0x20, 0xc8, 0x01,
	0x04, 0x5e, 0x2a, 0x52, 0x1f, 0xd8, 0x7f, 0xb5, 0xe2, 0xde, 0xbc, 0xbb, 0x96, 0x49, 0x2d, 0xab,
	0x7c, 0xcb, 0x04, 0xda, 0x20, 0xa5, 0x01, 0x8f, 0x11, 0x14, 0xb2, 0xb9, 0xda, 0xdc, 0x8b, 0x2b,
	0x7c, 0xd7, 0xba, 0x9b, 0xfb, 0xe8, 0x57, 0x52, 0x36, 0x43, 0xcf, 0x17, 0xe9, 0x4d, 0x24, 0x13,
	0x64, 0xf3, 0x1a, 0xdd, 0x9d, 0x42, 0x2f, 0x31, 0x34, 0xf4, 0xa9, 0x71, 0xba, 0xab, 0x83, 0xd1,
	0x4f, 0xa4, 0x1f, 0x48, 0x49, 0x1f, 0x30, 0x20, 0xfb, 0x5f, 0x67, 0xd4, 0xa6, 0x32, 0x5a, 0x7d,
	0x15, 0x8a, 0x28, 0x0d, 0xdb, 0xb7, 0xcd, 0xcc, 0xe9, 0xe6, 0x00, 0x3d, 0x27, 0xab, 0x7a, 0xf8,
	0x5c, 0xc6, 0xc2, 0x8c, 0x88, 0x4b, 0x0c, 0xed, 0x8a, 0x26, 0x62, 0x45, 0x73, 0x4f, 0x45, 0x7c,
	0x24, 0xcb, 0x23, 0x17, 0x89, 0xac, 0xa4, 0x63, 0xb6, 0xa7, 0x62, 0x4e, 0x9e, 0x4d, 0xee, 0x18,
	0x41, 0x7d, 0xb2, 0x25, 0x64, 0x56, 0x94, 0x92, 0x5c, 0x09, 0xe9, 0xf1, 0x20, 0x90, 0x80, 0x08,
	0xc8, 0x5e, 0xe9, 0xac, 0xfd, 0x97, 0x4a, 0xba, 0x06, 0xd5, 0x1a, 0x81, 0x4e, 0x72, 0xc6, 0xdd,
	0x14, 0x2f, 0x4d, 0xd3, 0x33, 0x52, 0x06, 0xe9, 0x1f, 0x1d, 0x7a, 0x4a, 0x78, 0x01, 0xa4, 0x22,
	0x41, 0xb6, 0xa8, 0xd3, 0x77, 0xa6, 0xd2, 0xcf, 0xdc, 0xd3, 0xa3, 0xc3, 0xb6, 0xf8, 0x94, 0xb9,
	0xdc, 0x15, 0x4d, 0xd9, 0x2f, 0xa4, 0x6d, 0xb2, 0xde, 0x4f, 0xcd, 0x19, 0x06, 0x9e, 0x92, 0x3c,
	0xc5, 0x1b, 0x90, 0xc8, 0x88, 0x8e, 0x7a, 0x3f, 0xfb, 0xf8, 0xad, 0xb3, 0x7d, 0xeb, 0xd2, 0x27,
	0x3e, 0x9f, 0x44, 0x7a, 0x42, 0x76, 0xc6, 0xfb, 0x0e, 0x54, 0x17, 0x24, 0xf4, 0x13, 0xaf, 0x0b,
	0x51, 0xd8, 0x55, 0x6c, 0x49, 0x77, 0x60, 0x65, 0xb4, 0x03, 0xcf, 0xac, 0xe5, 0x8b, 0x76, 0xd0,
	0x63, 0xb2, 0x65, 0x22, 0xec, 0x8a, 0x9e, 0xb9, 0xdd, 0x28, 0x60, 0xcb, 0x9a, 0xd5, 0x8d, 0x9d,
	0x97, 0xa3, 0x2f, 0xf2, 0x22, 0xa0, 0x0d, 0xb2, 0x39, 0x0e, 0xf5, 0x84, 0x88, 0x33, 0x66, 0x45,
	0x33, 0x74, 0x94, 0xb9, 0x12, 0x22, 0xbe, 0x08, 0x68, 0x8b, 0x6c, 0x8c, 0x97, 0x6a, 0x7a, 0x92,
	0xad, 0xce, 0xf8, 0xc3, 0x4c, 0x07, 0x37, 0xe7, 0xef, 0xff, 0xbc, 0x2d, 0xd8, 0x40, 0x4b, 0x1a,
	0x85, 0x1e, 0x10, 0xfa, 0xb4, 0xdb, 0x4e, 0xcc, 0xfd, 0x9f, 0x71, 0x84, 0x8a, 0x95, 0x6b, 0x73,
	0x7b, 0x8b, 0xee, 0x5a, 0xae, 0x34, 0x73, 0xa1, 0xf9, 0xf9, 0x7e, 0x58, 0x2d, 0x3e, 0x0c, 0xab,
	0xc5, 0xbf, 0xc3, 0x6a, 0xf1, 0xd7, 0x63, 0xb5, 0xf0, 0xf0, 0x58, 0x2d, 0xfc, 0x7e, 0xac, 0x16,
	0x7e, 0xec, 0x9b, 0xa5, 0x0f, 0x7c, 0x21, 0xc1, 0xc9, 0xc7, 0x5d, 0x1e, 0xa5, 0x4e, 0x22, 0x82,
	0x7e, 0x0c, 0xf9, 0xb3, 0xa2, 0x1f, 0xb3, 0xce, 0x82, 0x7e, 0x33, 0x8e, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x9a, 0xde, 0x1e, 0x43, 0x32, 0x05, 0x00, 0x00,
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
	if len(m.EthereumBlacklist) > 0 {
		for iNdEx := len(m.EthereumBlacklist) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EthereumBlacklist[iNdEx])
			copy(dAtA[i:], m.EthereumBlacklist[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.EthereumBlacklist[iNdEx])))
			i--
			dAtA[i] = 0x7a
		}
	}
	{
		size, err := m.LastObservedValset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	if m.LastOutgoingPoolId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastOutgoingPoolId))
		i--
		dAtA[i] = 0x68
	}
	if m.LastOutgoingBatchId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastOutgoingBatchId))
		i--
		dAtA[i] = 0x60
	}
	if m.LastObservedEthereumHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastObservedEthereumHeight))
		i--
		dAtA[i] = 0x58
	}
	if len(m.UnbatchedTransfers) > 0 {
		for iNdEx := len(m.UnbatchedTransfers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbatchedTransfers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Erc20ToDenoms) > 0 {
		for iNdEx := len(m.Erc20ToDenoms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Erc20ToDenoms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.OrchestratorAddresses) > 0 {
		for iNdEx := len(m.OrchestratorAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrchestratorAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Attestations) > 0 {
		for iNdEx := len(m.Attestations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attestations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.BatchConfirms) > 0 {
		for iNdEx := len(m.BatchConfirms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BatchConfirms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Batches) > 0 {
		for iNdEx := len(m.Batches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Batches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ValsetConfirms) > 0 {
		for iNdEx := len(m.ValsetConfirms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValsetConfirms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Valsets) > 0 {
		for iNdEx := len(m.Valsets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Valsets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.LastObservedNonce != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastObservedNonce))
		i--
		dAtA[i] = 0x10
	}
	if m.Params != nil {
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
	}
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
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LastObservedNonce != 0 {
		n += 1 + sovGenesis(uint64(m.LastObservedNonce))
	}
	if len(m.Valsets) > 0 {
		for _, e := range m.Valsets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValsetConfirms) > 0 {
		for _, e := range m.ValsetConfirms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Batches) > 0 {
		for _, e := range m.Batches {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BatchConfirms) > 0 {
		for _, e := range m.BatchConfirms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Attestations) > 0 {
		for _, e := range m.Attestations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OrchestratorAddresses) > 0 {
		for _, e := range m.OrchestratorAddresses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Erc20ToDenoms) > 0 {
		for _, e := range m.Erc20ToDenoms {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UnbatchedTransfers) > 0 {
		for _, e := range m.UnbatchedTransfers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastObservedEthereumHeight != 0 {
		n += 1 + sovGenesis(uint64(m.LastObservedEthereumHeight))
	}
	if m.LastOutgoingBatchId != 0 {
		n += 1 + sovGenesis(uint64(m.LastOutgoingBatchId))
	}
	if m.LastOutgoingPoolId != 0 {
		n += 1 + sovGenesis(uint64(m.LastOutgoingPoolId))
	}
	l = m.LastObservedValset.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.EthereumBlacklist) > 0 {
		for _, s := range m.EthereumBlacklist {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
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
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedNonce", wireType)
			}
			m.LastObservedNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastObservedNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valsets", wireType)
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
			m.Valsets = append(m.Valsets, &Valset{})
			if err := m.Valsets[len(m.Valsets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetConfirms", wireType)
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
			m.ValsetConfirms = append(m.ValsetConfirms, &MsgValsetConfirm{})
			if err := m.ValsetConfirms[len(m.ValsetConfirms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Batches", wireType)
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
			m.Batches = append(m.Batches, &OutgoingTxBatch{})
			if err := m.Batches[len(m.Batches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchConfirms", wireType)
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
			m.BatchConfirms = append(m.BatchConfirms, &MsgConfirmBatch{})
			if err := m.BatchConfirms[len(m.BatchConfirms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attestations", wireType)
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
			m.Attestations = append(m.Attestations, &Attestation{})
			if err := m.Attestations[len(m.Attestations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrchestratorAddresses", wireType)
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
			m.OrchestratorAddresses = append(m.OrchestratorAddresses, &MsgSetOrchestratorAddresses{})
			if err := m.OrchestratorAddresses[len(m.OrchestratorAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20ToDenoms", wireType)
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
			m.Erc20ToDenoms = append(m.Erc20ToDenoms, &ERC20ToDenom{})
			if err := m.Erc20ToDenoms[len(m.Erc20ToDenoms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbatchedTransfers", wireType)
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
			m.UnbatchedTransfers = append(m.UnbatchedTransfers, &OutgoingTransferTx{})
			if err := m.UnbatchedTransfers[len(m.UnbatchedTransfers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedEthereumHeight", wireType)
			}
			m.LastObservedEthereumHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastObservedEthereumHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastOutgoingBatchId", wireType)
			}
			m.LastOutgoingBatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastOutgoingBatchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastOutgoingPoolId", wireType)
			}
			m.LastOutgoingPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastOutgoingPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastObservedValset", wireType)
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
			if err := m.LastObservedValset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumBlacklist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumBlacklist = append(m.EthereumBlacklist, string(dAtA[iNdEx:postIndex]))
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
