// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: wallet.proto

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

type IssueTrx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject         string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	ReceiverAddress string `protobuf:"bytes,2,opt,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	Data            []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Spice           *Spice `protobuf:"bytes,4,opt,name=spice,proto3" json:"spice,omitempty"`
}

func (x *IssueTrx) Reset() {
	*x = IssueTrx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueTrx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTrx) ProtoMessage() {}

func (x *IssueTrx) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTrx.ProtoReflect.Descriptor instead.
func (*IssueTrx) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *IssueTrx) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *IssueTrx) GetReceiverAddress() string {
	if x != nil {
		return x.ReceiverAddress
	}
	return ""
}

func (x *IssueTrx) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *IssueTrx) GetSpice() *Spice {
	if x != nil {
		return x.Spice
	}
	return nil
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
		mi := &file_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWebHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebHook) ProtoMessage() {}

func (x *CreateWebHook) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[1]
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
	return file_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWebHook) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
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
		mi := &file_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotaryNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotaryNode) ProtoMessage() {}

func (x *NotaryNode) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[2]
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
	return file_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *NotaryNode) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TrxHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *TrxHash) Reset() {
	*x = TrxHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrxHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrxHash) ProtoMessage() {}

func (x *TrxHash) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrxHash.ProtoReflect.Descriptor instead.
func (*TrxHash) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *TrxHash) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

var File_wallet_proto protoreflect.FileDescriptor

var file_wallet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x1a, 0x16, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x01, 0x0a, 0x08, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x72, 0x78, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x70, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x73, 0x2e, 0x53, 0x70, 0x69, 0x63, 0x65, 0x52, 0x05, 0x73, 0x70, 0x69, 0x63, 0x65,
	0x22, 0x21, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f,
	0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x1e, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x1d, 0x0a, 0x07, 0x54, 0x72, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x32, 0xff, 0x03, 0x0a, 0x0f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x50, 0x49, 0x12, 0x39, 0x0a, 0x05, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x13, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x05, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x72, 0x78, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54, 0x72, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x06, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54, 0x72, 0x78, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x57, 0x61, 0x69, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x2e, 0x4e, 0x6f, 0x74, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x53, 0x61, 0x76, 0x65,
	0x64, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e,
	0x54, 0x72, 0x78, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x72, 0x74, 0x6f, 0x73, 0x73, 0x68, 0x2f, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_proto_rawDescOnce sync.Once
	file_wallet_proto_rawDescData = file_wallet_proto_rawDesc
)

func file_wallet_proto_rawDescGZIP() []byte {
	file_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_proto_rawDescData)
	})
	return file_wallet_proto_rawDescData
}

var file_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wallet_proto_goTypes = []interface{}{
	(*IssueTrx)(nil),      // 0: computantis.IssueTrx
	(*CreateWebHook)(nil), // 1: computantis.CreateWebHook
	(*NotaryNode)(nil),    // 2: computantis.NotaryNode
	(*TrxHash)(nil),       // 3: computantis.TrxHash
	(*Spice)(nil),         // 4: computantis.Spice
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
	(*AliveData)(nil),     // 6: computantis.AliveData
	(*Address)(nil),       // 7: computantis.Address
	(*Transactions)(nil),  // 8: computantis.Transactions
	(*Transaction)(nil),   // 9: computantis.Transaction
}
var file_wallet_proto_depIdxs = []int32{
	4, // 0: computantis.IssueTrx.spice:type_name -> computantis.Spice
	5, // 1: computantis.WalletClientAPI.Alive:input_type -> google.protobuf.Empty
	5, // 2: computantis.WalletClientAPI.WalletPublicAddress:input_type -> google.protobuf.Empty
	0, // 3: computantis.WalletClientAPI.Issue:input_type -> computantis.IssueTrx
	3, // 4: computantis.WalletClientAPI.Approve:input_type -> computantis.TrxHash
	3, // 5: computantis.WalletClientAPI.Reject:input_type -> computantis.TrxHash
	2, // 6: computantis.WalletClientAPI.Waiting:input_type -> computantis.NotaryNode
	3, // 7: computantis.WalletClientAPI.Saved:input_type -> computantis.TrxHash
	1, // 8: computantis.WalletClientAPI.WebHook:input_type -> computantis.CreateWebHook
	6, // 9: computantis.WalletClientAPI.Alive:output_type -> computantis.AliveData
	7, // 10: computantis.WalletClientAPI.WalletPublicAddress:output_type -> computantis.Address
	5, // 11: computantis.WalletClientAPI.Issue:output_type -> google.protobuf.Empty
	5, // 12: computantis.WalletClientAPI.Approve:output_type -> google.protobuf.Empty
	5, // 13: computantis.WalletClientAPI.Reject:output_type -> google.protobuf.Empty
	8, // 14: computantis.WalletClientAPI.Waiting:output_type -> computantis.Transactions
	9, // 15: computantis.WalletClientAPI.Saved:output_type -> computantis.Transaction
	5, // 16: computantis.WalletClientAPI.WebHook:output_type -> google.protobuf.Empty
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wallet_proto_init() }
func file_wallet_proto_init() {
	if File_wallet_proto != nil {
		return
	}
	file_computantistypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueTrx); i {
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
		file_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrxHash); i {
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
			RawDescriptor: file_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_proto_msgTypes,
	}.Build()
	File_wallet_proto = out.File
	file_wallet_proto_rawDesc = nil
	file_wallet_proto_goTypes = nil
	file_wallet_proto_depIdxs = nil
}
