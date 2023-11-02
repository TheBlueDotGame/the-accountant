// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: gossip.proto

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

type AliveData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion    string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	ApiHeader     string `protobuf:"bytes,2,opt,name=api_header,json=apiHeader,proto3" json:"api_header,omitempty"`
	PublicAddress string `protobuf:"bytes,3,opt,name=public_address,json=publicAddress,proto3" json:"public_address,omitempty"`
}

func (x *AliveData) Reset() {
	*x = AliveData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AliveData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AliveData) ProtoMessage() {}

func (x *AliveData) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AliveData.ProtoReflect.Descriptor instead.
func (*AliveData) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{0}
}

func (x *AliveData) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *AliveData) GetApiHeader() string {
	if x != nil {
		return x.ApiHeader
	}
	return ""
}

func (x *AliveData) GetPublicAddress() string {
	if x != nil {
		return x.PublicAddress
	}
	return ""
}

type Vertex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignerPublicAddress string       `protobuf:"bytes,1,opt,name=signer_public_address,json=signerPublicAddress,proto3" json:"signer_public_address,omitempty"`
	CreaterdAt          uint64       `protobuf:"varint,2,opt,name=createrd_at,json=createrdAt,proto3" json:"createrd_at,omitempty"`
	Signature           []byte       `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Transaction         *Transaction `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Hash                []byte       `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	LeftParentHash      []byte       `protobuf:"bytes,6,opt,name=left_parent_hash,json=leftParentHash,proto3" json:"left_parent_hash,omitempty"`
	RightParentHash     []byte       `protobuf:"bytes,7,opt,name=right_parent_hash,json=rightParentHash,proto3" json:"right_parent_hash,omitempty"`
	Weight              uint64       `protobuf:"varint,8,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Vertex) Reset() {
	*x = Vertex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vertex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vertex) ProtoMessage() {}

func (x *Vertex) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vertex.ProtoReflect.Descriptor instead.
func (*Vertex) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{1}
}

func (x *Vertex) GetSignerPublicAddress() string {
	if x != nil {
		return x.SignerPublicAddress
	}
	return ""
}

func (x *Vertex) GetCreaterdAt() uint64 {
	if x != nil {
		return x.CreaterdAt
	}
	return 0
}

func (x *Vertex) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Vertex) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *Vertex) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Vertex) GetLeftParentHash() []byte {
	if x != nil {
		return x.LeftParentHash
	}
	return nil
}

func (x *Vertex) GetRightParentHash() []byte {
	if x != nil {
		return x.RightParentHash
	}
	return nil
}

func (x *Vertex) GetWeight() uint64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type VertexGossip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vertex    *Vertex  `protobuf:"bytes,1,opt,name=vertex,proto3" json:"vertex,omitempty"`
	Gossipers []string `protobuf:"bytes,2,rep,name=gossipers,proto3" json:"gossipers,omitempty"`
}

func (x *VertexGossip) Reset() {
	*x = VertexGossip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VertexGossip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VertexGossip) ProtoMessage() {}

func (x *VertexGossip) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VertexGossip.ProtoReflect.Descriptor instead.
func (*VertexGossip) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{2}
}

func (x *VertexGossip) GetVertex() *Vertex {
	if x != nil {
		return x.Vertex
	}
	return nil
}

func (x *VertexGossip) GetGossipers() []string {
	if x != nil {
		return x.Gossipers
	}
	return nil
}

type ConnectionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicAddress string `protobuf:"bytes,1,opt,name=public_address,json=publicAddress,proto3" json:"public_address,omitempty"`
	Url           string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt     uint64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Digest        []byte `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	Signature     []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ConnectionData) Reset() {
	*x = ConnectionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionData) ProtoMessage() {}

func (x *ConnectionData) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionData.ProtoReflect.Descriptor instead.
func (*ConnectionData) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionData) GetPublicAddress() string {
	if x != nil {
		return x.PublicAddress
	}
	return ""
}

func (x *ConnectionData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ConnectionData) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ConnectionData) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *ConnectionData) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ConnectedNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignerPublicAddress string            `protobuf:"bytes,1,opt,name=signer_public_address,json=signerPublicAddress,proto3" json:"signer_public_address,omitempty"`
	Connections         []*ConnectionData `protobuf:"bytes,2,rep,name=connections,proto3" json:"connections,omitempty"`
}

func (x *ConnectedNodes) Reset() {
	*x = ConnectedNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectedNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectedNodes) ProtoMessage() {}

func (x *ConnectedNodes) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectedNodes.ProtoReflect.Descriptor instead.
func (*ConnectedNodes) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectedNodes) GetSignerPublicAddress() string {
	if x != nil {
		return x.SignerPublicAddress
	}
	return ""
}

func (x *ConnectedNodes) GetConnections() []*ConnectionData {
	if x != nil {
		return x.Connections
	}
	return nil
}

var File_gossip_proto protoreflect.FileDescriptor

var file_gossip_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x1a, 0x11, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x09, 0x41,
	0x6c, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x69,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x70, 0x69, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0xb9, 0x02, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3a, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x28, 0x0a,
	0x10, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6c, 0x65, 0x66, 0x74, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x59, 0x0a, 0x0c, 0x56,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x2b, 0x0a, 0x06, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78,
	0x52, 0x06, 0x76, 0x65, 0x72, 0x74, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x65, 0x72, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xcd, 0x01,
	0x0a, 0x09, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x50, 0x49, 0x12, 0x39, 0x0a, 0x05, 0x41,
	0x6c, 0x69, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x41, 0x6c, 0x69, 0x76, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x06, 0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x74, 0x65, 0x78, 0x47, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x72, 0x74,
	0x6f, 0x73, 0x73, 0x68, 0x2f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gossip_proto_rawDescOnce sync.Once
	file_gossip_proto_rawDescData = file_gossip_proto_rawDesc
)

func file_gossip_proto_rawDescGZIP() []byte {
	file_gossip_proto_rawDescOnce.Do(func() {
		file_gossip_proto_rawDescData = protoimpl.X.CompressGZIP(file_gossip_proto_rawDescData)
	})
	return file_gossip_proto_rawDescData
}

var file_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gossip_proto_goTypes = []interface{}{
	(*AliveData)(nil),      // 0: computantis.AliveData
	(*Vertex)(nil),         // 1: computantis.Vertex
	(*VertexGossip)(nil),   // 2: computantis.VertexGossip
	(*ConnectionData)(nil), // 3: computantis.ConnectionData
	(*ConnectedNodes)(nil), // 4: computantis.ConnectedNodes
	(*Transaction)(nil),    // 5: computantis.Transaction
	(*emptypb.Empty)(nil),  // 6: google.protobuf.Empty
}
var file_gossip_proto_depIdxs = []int32{
	5, // 0: computantis.Vertex.transaction:type_name -> computantis.Transaction
	1, // 1: computantis.VertexGossip.vertex:type_name -> computantis.Vertex
	3, // 2: computantis.ConnectedNodes.connections:type_name -> computantis.ConnectionData
	6, // 3: computantis.GossipAPI.Alive:input_type -> google.protobuf.Empty
	3, // 4: computantis.GossipAPI.Discover:input_type -> computantis.ConnectionData
	2, // 5: computantis.GossipAPI.Gossip:input_type -> computantis.VertexGossip
	0, // 6: computantis.GossipAPI.Alive:output_type -> computantis.AliveData
	4, // 7: computantis.GossipAPI.Discover:output_type -> computantis.ConnectedNodes
	6, // 8: computantis.GossipAPI.Gossip:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gossip_proto_init() }
func file_gossip_proto_init() {
	if File_gossip_proto != nil {
		return
	}
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gossip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AliveData); i {
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
		file_gossip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vertex); i {
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
		file_gossip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VertexGossip); i {
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
		file_gossip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionData); i {
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
		file_gossip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectedNodes); i {
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
			RawDescriptor: file_gossip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gossip_proto_goTypes,
		DependencyIndexes: file_gossip_proto_depIdxs,
		MessageInfos:      file_gossip_proto_msgTypes,
	}.Build()
	File_gossip_proto = out.File
	file_gossip_proto_rawDesc = nil
	file_gossip_proto_goTypes = nil
	file_gossip_proto_depIdxs = nil
}
