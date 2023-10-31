// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: wallet_client_api.proto

package protobufcompiled

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AliveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	ApiHeader  string `protobuf:"bytes,2,opt,name=api_header,json=apiHeader,proto3" json:"api_header,omitempty"`
	Alive      bool   `protobuf:"varint,3,opt,name=alive,proto3" json:"alive,omitempty"`
}

func (x *AliveInfo) Reset() {
	*x = AliveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AliveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AliveInfo) ProtoMessage() {}

func (x *AliveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AliveInfo.ProtoReflect.Descriptor instead.
func (*AliveInfo) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{0}
}

func (x *AliveInfo) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *AliveInfo) GetApiHeader() string {
	if x != nil {
		return x.ApiHeader
	}
	return ""
}

func (x *AliveInfo) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

type IssueTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject         string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Data            []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ReceiverAddress string `protobuf:"bytes,3,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
}

func (x *IssueTransaction) Reset() {
	*x = IssueTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTransaction) ProtoMessage() {}

func (x *IssueTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTransaction.ProtoReflect.Descriptor instead.
func (*IssueTransaction) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{1}
}

func (x *IssueTransaction) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *IssueTransaction) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *IssueTransaction) GetReceiverAddress() string {
	if x != nil {
		return x.ReceiverAddress
	}
	return ""
}

type Reject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Reject) Reset() {
	*x = Reject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reject) ProtoMessage() {}

func (x *Reject) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reject.ProtoReflect.Descriptor instead.
func (*Reject) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{2}
}

func (x *Reject) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{3}
}

func (x *ServerInfo) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ServerInfo) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Transactions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info         *ServerInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Transactions) Reset() {
	*x = Transactions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transactions) ProtoMessage() {}

func (x *Transactions) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transactions.ProtoReflect.Descriptor instead.
func (*Transactions) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{4}
}

func (x *Transactions) GetInfo() *ServerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Transactions) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type CreateWallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateWallet) Reset() {
	*x = CreateWallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWallet) ProtoMessage() {}

func (x *CreateWallet) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWallet.ProtoReflect.Descriptor instead.
func (*CreateWallet) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{5}
}

func (x *CreateWallet) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CreateWebHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateWebHook) Reset() {
	*x = CreateWebHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWebHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebHook) ProtoMessage() {}

func (x *CreateWebHook) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWebHook.ProtoReflect.Descriptor instead.
func (*CreateWebHook) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{6}
}

func (x *CreateWebHook) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type WalletPublicAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *ServerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Address string      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *WalletPublicAddress) Reset() {
	*x = WalletPublicAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletPublicAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletPublicAddress) ProtoMessage() {}

func (x *WalletPublicAddress) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletPublicAddress.ProtoReflect.Descriptor instead.
func (*WalletPublicAddress) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{7}
}

func (x *WalletPublicAddress) GetInfo() *ServerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *WalletPublicAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Paggination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Paggination) Reset() {
	*x = Paggination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paggination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paggination) ProtoMessage() {}

func (x *Paggination) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paggination.ProtoReflect.Descriptor instead.
func (*Paggination) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{8}
}

func (x *Paggination) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Paggination) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type NotaryNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *NotaryNode) Reset() {
	*x = NotaryNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_client_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotaryNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotaryNode) ProtoMessage() {}

func (x *NotaryNode) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_client_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotaryNode.ProtoReflect.Descriptor instead.
func (*NotaryNode) Descriptor() ([]byte, []int) {
	return file_wallet_client_api_proto_rawDescGZIP(), []int{9}
}

func (x *NotaryNode) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_wallet_client_api_proto protoreflect.FileDescriptor

var file_wallet_client_api_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x1a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x09, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x6b, 0x0a, 0x10, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2e,
	0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x79,
	0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x0c, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x21, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x5c, 0x0a, 0x13, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x3b, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x1e, 0x0a,
	0x0a, 0x4e, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xcd, 0x04,
	0x0a, 0x0f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x50,
	0x49, 0x12, 0x39, 0x0a, 0x05, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x2e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x4e, 0x6f,
	0x64, 0x65, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12,
	0x4c, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65,
	0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a,
	0x14, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x14,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74,
	0x69, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x57,
	0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f,
	0x6f, 0x6b, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x72, 0x74,
	0x6f, 0x73, 0x73, 0x68, 0x2f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_client_api_proto_rawDescOnce sync.Once
	file_wallet_client_api_proto_rawDescData = file_wallet_client_api_proto_rawDesc
)

func file_wallet_client_api_proto_rawDescGZIP() []byte {
	file_wallet_client_api_proto_rawDescOnce.Do(func() {
		file_wallet_client_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_client_api_proto_rawDescData)
	})
	return file_wallet_client_api_proto_rawDescData
}

var file_wallet_client_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_wallet_client_api_proto_goTypes = []interface{}{
	(*AliveInfo)(nil),           // 0: computantis.AliveInfo
	(*IssueTransaction)(nil),    // 1: computantis.IssueTransaction
	(*Reject)(nil),              // 2: computantis.Reject
	(*ServerInfo)(nil),          // 3: computantis.ServerInfo
	(*Transactions)(nil),        // 4: computantis.Transactions
	(*CreateWallet)(nil),        // 5: computantis.CreateWallet
	(*CreateWebHook)(nil),       // 6: computantis.CreateWebHook
	(*WalletPublicAddress)(nil), // 7: computantis.WalletPublicAddress
	(*Paggination)(nil),         // 8: computantis.Paggination
	(*NotaryNode)(nil),          // 9: computantis.NotaryNode
	(*Transaction)(nil),         // 10: computantis.Transaction
	(*emptypb.Empty)(nil),       // 11: google.protobuf.Empty
}
var file_wallet_client_api_proto_depIdxs = []int32{
	10, // 0: computantis.Reject.transactions:type_name -> computantis.Transaction
	3,  // 1: computantis.Transactions.info:type_name -> computantis.ServerInfo
	10, // 2: computantis.Transactions.transactions:type_name -> computantis.Transaction
	3,  // 3: computantis.WalletPublicAddress.info:type_name -> computantis.ServerInfo
	11, // 4: computantis.WalletClientAPI.Alive:input_type -> google.protobuf.Empty
	11, // 5: computantis.WalletClientAPI.Address:input_type -> google.protobuf.Empty
	9,  // 6: computantis.WalletClientAPI.IssuedTransactions:input_type -> computantis.NotaryNode
	9,  // 7: computantis.WalletClientAPI.ReceivedTransactions:input_type -> computantis.NotaryNode
	8,  // 8: computantis.WalletClientAPI.ApprovedTransactions:input_type -> computantis.Paggination
	8,  // 9: computantis.WalletClientAPI.RejectedTransactions:input_type -> computantis.Paggination
	5,  // 10: computantis.WalletClientAPI.Wallet:input_type -> computantis.CreateWallet
	6,  // 11: computantis.WalletClientAPI.WebHook:input_type -> computantis.CreateWebHook
	0,  // 12: computantis.WalletClientAPI.Alive:output_type -> computantis.AliveInfo
	7,  // 13: computantis.WalletClientAPI.Address:output_type -> computantis.WalletPublicAddress
	4,  // 14: computantis.WalletClientAPI.IssuedTransactions:output_type -> computantis.Transactions
	4,  // 15: computantis.WalletClientAPI.ReceivedTransactions:output_type -> computantis.Transactions
	4,  // 16: computantis.WalletClientAPI.ApprovedTransactions:output_type -> computantis.Transactions
	4,  // 17: computantis.WalletClientAPI.RejectedTransactions:output_type -> computantis.Transactions
	3,  // 18: computantis.WalletClientAPI.Wallet:output_type -> computantis.ServerInfo
	3,  // 19: computantis.WalletClientAPI.WebHook:output_type -> computantis.ServerInfo
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_wallet_client_api_proto_init() }
func file_wallet_client_api_proto_init() {
	if File_wallet_client_api_proto != nil {
		return
	}
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wallet_client_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AliveInfo); i {
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
		file_wallet_client_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueTransaction); i {
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
		file_wallet_client_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reject); i {
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
		file_wallet_client_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_wallet_client_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transactions); i {
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
		file_wallet_client_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWallet); i {
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
		file_wallet_client_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWebHook); i {
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
		file_wallet_client_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletPublicAddress); i {
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
		file_wallet_client_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paggination); i {
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
		file_wallet_client_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotaryNode); i {
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
			RawDescriptor: file_wallet_client_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_client_api_proto_goTypes,
		DependencyIndexes: file_wallet_client_api_proto_depIdxs,
		MessageInfos:      file_wallet_client_api_proto_msgTypes,
	}.Build()
	File_wallet_client_api_proto = out.File
	file_wallet_client_api_proto_rawDesc = nil
	file_wallet_client_api_proto_goTypes = nil
	file_wallet_client_api_proto_depIdxs = nil
}
