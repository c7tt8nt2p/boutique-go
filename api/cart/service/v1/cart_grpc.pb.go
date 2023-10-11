// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: v1/cart.proto

package v1

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
	Cart_NewCart_FullMethodName       = "/cart.service.v1.Cart/NewCart"
	Cart_AddItemToCart_FullMethodName = "/cart.service.v1.Cart/AddItemToCart"
	Cart_ViewCart_FullMethodName      = "/cart.service.v1.Cart/ViewCart"
)

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	NewCart(ctx context.Context, in *NewCartReq, opts ...grpc.CallOption) (*NewCartResp, error)
	AddItemToCart(ctx context.Context, in *AddItemToCartReq, opts ...grpc.CallOption) (*AddItemToCartResp, error)
	ViewCart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ViewCartResp, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) NewCart(ctx context.Context, in *NewCartReq, opts ...grpc.CallOption) (*NewCartResp, error) {
	out := new(NewCartResp)
	err := c.cc.Invoke(ctx, Cart_NewCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) AddItemToCart(ctx context.Context, in *AddItemToCartReq, opts ...grpc.CallOption) (*AddItemToCartResp, error) {
	out := new(AddItemToCartResp)
	err := c.cc.Invoke(ctx, Cart_AddItemToCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) ViewCart(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ViewCartResp, error) {
	out := new(ViewCartResp)
	err := c.cc.Invoke(ctx, Cart_ViewCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility
type CartServer interface {
	NewCart(context.Context, *NewCartReq) (*NewCartResp, error)
	AddItemToCart(context.Context, *AddItemToCartReq) (*AddItemToCartResp, error)
	ViewCart(context.Context, *emptypb.Empty) (*ViewCartResp, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (UnimplementedCartServer) NewCart(context.Context, *NewCartReq) (*NewCartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewCart not implemented")
}
func (UnimplementedCartServer) AddItemToCart(context.Context, *AddItemToCartReq) (*AddItemToCartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItemToCart not implemented")
}
func (UnimplementedCartServer) ViewCart(context.Context, *emptypb.Empty) (*ViewCartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewCart not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_NewCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).NewCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_NewCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).NewCart(ctx, req.(*NewCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_AddItemToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemToCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).AddItemToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_AddItemToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).AddItemToCart(ctx, req.(*AddItemToCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_ViewCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).ViewCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cart_ViewCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).ViewCart(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.service.v1.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewCart",
			Handler:    _Cart_NewCart_Handler,
		},
		{
			MethodName: "AddItemToCart",
			Handler:    _Cart_AddItemToCart_Handler,
		},
		{
			MethodName: "ViewCart",
			Handler:    _Cart_ViewCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/cart.proto",
}
