// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: v1/auth.proto

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
	Auth_NewAuth_FullMethodName    = "/user.service.v1.Auth/NewAuth"
	Auth_DeleteAuth_FullMethodName = "/user.service.v1.Auth/DeleteAuth"
	Auth_Login_FullMethodName      = "/user.service.v1.Auth/Login"
	Auth_Validate_FullMethodName   = "/user.service.v1.Auth/Validate"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	NewAuth(ctx context.Context, in *NewAuthReq, opts ...grpc.CallOption) (*NewAuthResp, error)
	DeleteAuth(ctx context.Context, in *DeleteAuthReq, opts ...grpc.CallOption) (*DeleteAuthResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Validate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ValidateResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) NewAuth(ctx context.Context, in *NewAuthReq, opts ...grpc.CallOption) (*NewAuthResp, error) {
	out := new(NewAuthResp)
	err := c.cc.Invoke(ctx, Auth_NewAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteAuth(ctx context.Context, in *DeleteAuthReq, opts ...grpc.CallOption) (*DeleteAuthResp, error) {
	out := new(DeleteAuthResp)
	err := c.cc.Invoke(ctx, Auth_DeleteAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Auth_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Validate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ValidateResp, error) {
	out := new(ValidateResp)
	err := c.cc.Invoke(ctx, Auth_Validate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	NewAuth(context.Context, *NewAuthReq) (*NewAuthResp, error)
	DeleteAuth(context.Context, *DeleteAuthReq) (*DeleteAuthResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Validate(context.Context, *emptypb.Empty) (*ValidateResp, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) NewAuth(context.Context, *NewAuthReq) (*NewAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAuth not implemented")
}
func (UnimplementedAuthServer) DeleteAuth(context.Context, *DeleteAuthReq) (*DeleteAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuth not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) Validate(context.Context, *emptypb.Empty) (*ValidateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_NewAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).NewAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_NewAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).NewAuth(ctx, req.(*NewAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteAuth(ctx, req.(*DeleteAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Validate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.service.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAuth",
			Handler:    _Auth_NewAuth_Handler,
		},
		{
			MethodName: "DeleteAuth",
			Handler:    _Auth_DeleteAuth_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Auth_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/auth.proto",
}
