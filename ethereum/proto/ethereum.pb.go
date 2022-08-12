// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/ethereum.proto

package ethereum

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

// Get the balance of an ethereum wallet
type BalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address of wallet
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *BalanceRequest) Reset() {
	*x = BalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ethereum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest) ProtoMessage() {}

func (x *BalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest.ProtoReflect.Descriptor instead.
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_ethereum_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type BalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the account balance (in wei)
	Balance int64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BalanceResponse) Reset() {
	*x = BalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ethereum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResponse) ProtoMessage() {}

func (x *BalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResponse.ProtoReflect.Descriptor instead.
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_ethereum_proto_rawDescGZIP(), []int{1}
}

func (x *BalanceResponse) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

// Get transaction details by hash
type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tx hash
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ethereum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_proto_ethereum_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tx hash
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// the block hash
	BlockHash string `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// the block number
	BlockNumber string `protobuf:"bytes,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	// sent from
	FromAddress string `protobuf:"bytes,4,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// gas
	Gas string `protobuf:"bytes,5,opt,name=gas,proto3" json:"gas,omitempty"`
	// gas price
	GasPrice string `protobuf:"bytes,6,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// max fee per gas
	MaxFeePerGas string `protobuf:"bytes,7,opt,name=max_fee_per_gas,json=maxFeePerGas,proto3" json:"max_fee_per_gas,omitempty"`
	// max priority fee per gas
	MaxPriorityFeePerGas string `protobuf:"bytes,8,opt,name=max_priority_fee_per_gas,json=maxPriorityFeePerGas,proto3" json:"max_priority_fee_per_gas,omitempty"`
	// input
	Input string `protobuf:"bytes,9,opt,name=input,proto3" json:"input,omitempty"`
	// the nonce
	Nonce string `protobuf:"bytes,10,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// to address
	ToAddress string `protobuf:"bytes,11,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// transaction index
	TxIndex string `protobuf:"bytes,12,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	// value of transaction
	Value string `protobuf:"bytes,13,opt,name=value,proto3" json:"value,omitempty"`
	// type of transaction
	Type string `protobuf:"bytes,14,opt,name=type,proto3" json:"type,omitempty"`
	// chain id
	ChainId string `protobuf:"bytes,15,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	V       string `protobuf:"bytes,16,opt,name=v,proto3" json:"v,omitempty"`
	R       string `protobuf:"bytes,17,opt,name=r,proto3" json:"r,omitempty"`
	S       string `protobuf:"bytes,18,opt,name=s,proto3" json:"s,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ethereum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_ethereum_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *TransactionResponse) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *TransactionResponse) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *TransactionResponse) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *TransactionResponse) GetGas() string {
	if x != nil {
		return x.Gas
	}
	return ""
}

func (x *TransactionResponse) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

func (x *TransactionResponse) GetMaxFeePerGas() string {
	if x != nil {
		return x.MaxFeePerGas
	}
	return ""
}

func (x *TransactionResponse) GetMaxPriorityFeePerGas() string {
	if x != nil {
		return x.MaxPriorityFeePerGas
	}
	return ""
}

func (x *TransactionResponse) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TransactionResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *TransactionResponse) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TransactionResponse) GetTxIndex() string {
	if x != nil {
		return x.TxIndex
	}
	return ""
}

func (x *TransactionResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TransactionResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransactionResponse) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *TransactionResponse) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

func (x *TransactionResponse) GetR() string {
	if x != nil {
		return x.R
	}
	return ""
}

func (x *TransactionResponse) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

var File_proto_ethereum_proto protoreflect.FileDescriptor

var file_proto_ethereum_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x22, 0x2a, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2b, 0x0a, 0x0f,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x22, 0xf1, 0x03, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47, 0x61, 0x73, 0x12, 0x36, 0x0a, 0x18, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x67, 0x61, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x61,
	0x78, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x50, 0x65, 0x72, 0x47,
	0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x0c, 0x0a,
	0x01, 0x76, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x76, 0x12, 0x0c, 0x0a, 0x01, 0x72,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x08, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x12, 0x40, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x18, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ethereum_proto_rawDescOnce sync.Once
	file_proto_ethereum_proto_rawDescData = file_proto_ethereum_proto_rawDesc
)

func file_proto_ethereum_proto_rawDescGZIP() []byte {
	file_proto_ethereum_proto_rawDescOnce.Do(func() {
		file_proto_ethereum_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ethereum_proto_rawDescData)
	})
	return file_proto_ethereum_proto_rawDescData
}

var file_proto_ethereum_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_ethereum_proto_goTypes = []interface{}{
	(*BalanceRequest)(nil),      // 0: ethereum.BalanceRequest
	(*BalanceResponse)(nil),     // 1: ethereum.BalanceResponse
	(*TransactionRequest)(nil),  // 2: ethereum.TransactionRequest
	(*TransactionResponse)(nil), // 3: ethereum.TransactionResponse
}
var file_proto_ethereum_proto_depIdxs = []int32{
	0, // 0: ethereum.Ethereum.Balance:input_type -> ethereum.BalanceRequest
	2, // 1: ethereum.Ethereum.Transaction:input_type -> ethereum.TransactionRequest
	1, // 2: ethereum.Ethereum.Balance:output_type -> ethereum.BalanceResponse
	3, // 3: ethereum.Ethereum.Transaction:output_type -> ethereum.TransactionResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ethereum_proto_init() }
func file_proto_ethereum_proto_init() {
	if File_proto_ethereum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ethereum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceRequest); i {
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
		file_proto_ethereum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceResponse); i {
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
		file_proto_ethereum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_proto_ethereum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
			RawDescriptor: file_proto_ethereum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ethereum_proto_goTypes,
		DependencyIndexes: file_proto_ethereum_proto_depIdxs,
		MessageInfos:      file_proto_ethereum_proto_msgTypes,
	}.Build()
	File_proto_ethereum_proto = out.File
	file_proto_ethereum_proto_rawDesc = nil
	file_proto_ethereum_proto_goTypes = nil
	file_proto_ethereum_proto_depIdxs = nil
}
