// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: wallet_client_api.proto

package protobufcompiled

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	WalletClientAPI_Alive_FullMethodName                = "/computantis.WalletClientAPI/Alive"
	WalletClientAPI_Address_FullMethodName              = "/computantis.WalletClientAPI/Address"
	WalletClientAPI_IssuedTransactions_FullMethodName   = "/computantis.WalletClientAPI/IssuedTransactions"
	WalletClientAPI_ReceivedTransactions_FullMethodName = "/computantis.WalletClientAPI/ReceivedTransactions"
	WalletClientAPI_ApprovedTransactions_FullMethodName = "/computantis.WalletClientAPI/ApprovedTransactions"
	WalletClientAPI_RejectedTransactions_FullMethodName = "/computantis.WalletClientAPI/RejectedTransactions"
	WalletClientAPI_Wallet_FullMethodName               = "/computantis.WalletClientAPI/Wallet"
	WalletClientAPI_WebHook_FullMethodName              = "/computantis.WalletClientAPI/WebHook"
)

// WalletClientAPIClient is the client API for WalletClientAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClientAPIClient interface {
	Alive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AliveInfo, error)
	Address(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WalletPublicAddress, error)
	IssuedTransactions(ctx context.Context, in *NotaryNode, opts ...grpc.CallOption) (*Transactions, error)
	ReceivedTransactions(ctx context.Context, in *NotaryNode, opts ...grpc.CallOption) (*Transactions, error)
	ApprovedTransactions(ctx context.Context, in *Paggination, opts ...grpc.CallOption) (*Transactions, error)
	RejectedTransactions(ctx context.Context, in *Paggination, opts ...grpc.CallOption) (*Transactions, error)
	Wallet(ctx context.Context, in *CreateWallet, opts ...grpc.CallOption) (*ServerInfo, error)
	WebHook(ctx context.Context, in *CreateWebHook, opts ...grpc.CallOption) (*ServerInfo, error)
}

type walletClientAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClientAPIClient(cc grpc.ClientConnInterface) WalletClientAPIClient {
	return &walletClientAPIClient{cc}
}

func (c *walletClientAPIClient) Alive(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AliveInfo, error) {
	out := new(AliveInfo)
	err := c.cc.Invoke(ctx, WalletClientAPI_Alive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClientAPIClient) Address(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WalletPublicAddress, error) {
	out := new(WalletPublicAddress)
	err := c.cc.Invoke(ctx, WalletClientAPI_Address_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClientAPIClient) IssuedTransactions(ctx context.Context, in *NotaryNode, opts ...grpc.CallOption) (*Transactions, error) {
	out := new(Transactions)
	err := c.cc.Invoke(ctx, WalletClientAPI_IssuedTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClientAPIClient) ReceivedTransactions(ctx context.Context, in *NotaryNode, opts ...grpc.CallOption) (*Transactions, error) {
	out := new(Transactions)
	err := c.cc.Invoke(ctx, WalletClientAPI_ReceivedTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClientAPIClient) ApprovedTransactions(ctx context.Context, in *Paggination, opts ...grpc.CallOption) (*Transactions, error) {
	out := new(Transactions)
	err := c.cc.Invoke(ctx, WalletClientAPI_ApprovedTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClientAPIClient) RejectedTransactions(ctx context.Context, in *Paggination, opts ...grpc.CallOption) (*Transactions, error) {
	out := new(Transactions)
	err := c.cc.Invoke(ctx, WalletClientAPI_RejectedTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClientAPIClient) Wallet(ctx context.Context, in *CreateWallet, opts ...grpc.CallOption) (*ServerInfo, error) {
	out := new(ServerInfo)
	err := c.cc.Invoke(ctx, WalletClientAPI_Wallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClientAPIClient) WebHook(ctx context.Context, in *CreateWebHook, opts ...grpc.CallOption) (*ServerInfo, error) {
	out := new(ServerInfo)
	err := c.cc.Invoke(ctx, WalletClientAPI_WebHook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletClientAPIServer is the server API for WalletClientAPI service.
// All implementations must embed UnimplementedWalletClientAPIServer
// for forward compatibility
type WalletClientAPIServer interface {
	Alive(context.Context, *emptypb.Empty) (*AliveInfo, error)
	Address(context.Context, *emptypb.Empty) (*WalletPublicAddress, error)
	IssuedTransactions(context.Context, *NotaryNode) (*Transactions, error)
	ReceivedTransactions(context.Context, *NotaryNode) (*Transactions, error)
	ApprovedTransactions(context.Context, *Paggination) (*Transactions, error)
	RejectedTransactions(context.Context, *Paggination) (*Transactions, error)
	Wallet(context.Context, *CreateWallet) (*ServerInfo, error)
	WebHook(context.Context, *CreateWebHook) (*ServerInfo, error)
	mustEmbedUnimplementedWalletClientAPIServer()
}

// UnimplementedWalletClientAPIServer must be embedded to have forward compatible implementations.
type UnimplementedWalletClientAPIServer struct {
}

func (UnimplementedWalletClientAPIServer) Alive(context.Context, *emptypb.Empty) (*AliveInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}
func (UnimplementedWalletClientAPIServer) Address(context.Context, *emptypb.Empty) (*WalletPublicAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Address not implemented")
}
func (UnimplementedWalletClientAPIServer) IssuedTransactions(context.Context, *NotaryNode) (*Transactions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssuedTransactions not implemented")
}
func (UnimplementedWalletClientAPIServer) ReceivedTransactions(context.Context, *NotaryNode) (*Transactions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceivedTransactions not implemented")
}
func (UnimplementedWalletClientAPIServer) ApprovedTransactions(context.Context, *Paggination) (*Transactions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApprovedTransactions not implemented")
}
func (UnimplementedWalletClientAPIServer) RejectedTransactions(context.Context, *Paggination) (*Transactions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectedTransactions not implemented")
}
func (UnimplementedWalletClientAPIServer) Wallet(context.Context, *CreateWallet) (*ServerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Wallet not implemented")
}
func (UnimplementedWalletClientAPIServer) WebHook(context.Context, *CreateWebHook) (*ServerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebHook not implemented")
}
func (UnimplementedWalletClientAPIServer) mustEmbedUnimplementedWalletClientAPIServer() {}

// UnsafeWalletClientAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletClientAPIServer will
// result in compilation errors.
type UnsafeWalletClientAPIServer interface {
	mustEmbedUnimplementedWalletClientAPIServer()
}

func RegisterWalletClientAPIServer(s grpc.ServiceRegistrar, srv WalletClientAPIServer) {
	s.RegisterService(&WalletClientAPI_ServiceDesc, srv)
}

func _WalletClientAPI_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletClientAPIServer).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletClientAPI_Alive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletClientAPIServer).Alive(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletClientAPI_Address_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletClientAPIServer).Address(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletClientAPI_Address_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletClientAPIServer).Address(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletClientAPI_IssuedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotaryNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletClientAPIServer).IssuedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletClientAPI_IssuedTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletClientAPIServer).IssuedTransactions(ctx, req.(*NotaryNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletClientAPI_ReceivedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotaryNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletClientAPIServer).ReceivedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletClientAPI_ReceivedTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletClientAPIServer).ReceivedTransactions(ctx, req.(*NotaryNode))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletClientAPI_ApprovedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paggination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletClientAPIServer).ApprovedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletClientAPI_ApprovedTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletClientAPIServer).ApprovedTransactions(ctx, req.(*Paggination))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletClientAPI_RejectedTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paggination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletClientAPIServer).RejectedTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletClientAPI_RejectedTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletClientAPIServer).RejectedTransactions(ctx, req.(*Paggination))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletClientAPI_Wallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWallet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletClientAPIServer).Wallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletClientAPI_Wallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletClientAPIServer).Wallet(ctx, req.(*CreateWallet))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletClientAPI_WebHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWebHook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletClientAPIServer).WebHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletClientAPI_WebHook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletClientAPIServer).WebHook(ctx, req.(*CreateWebHook))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletClientAPI_ServiceDesc is the grpc.ServiceDesc for WalletClientAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletClientAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "computantis.WalletClientAPI",
	HandlerType: (*WalletClientAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Alive",
			Handler:    _WalletClientAPI_Alive_Handler,
		},
		{
			MethodName: "Address",
			Handler:    _WalletClientAPI_Address_Handler,
		},
		{
			MethodName: "IssuedTransactions",
			Handler:    _WalletClientAPI_IssuedTransactions_Handler,
		},
		{
			MethodName: "ReceivedTransactions",
			Handler:    _WalletClientAPI_ReceivedTransactions_Handler,
		},
		{
			MethodName: "ApprovedTransactions",
			Handler:    _WalletClientAPI_ApprovedTransactions_Handler,
		},
		{
			MethodName: "RejectedTransactions",
			Handler:    _WalletClientAPI_RejectedTransactions_Handler,
		},
		{
			MethodName: "Wallet",
			Handler:    _WalletClientAPI_Wallet_Handler,
		},
		{
			MethodName: "WebHook",
			Handler:    _WalletClientAPI_WebHook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet_client_api.proto",
}
