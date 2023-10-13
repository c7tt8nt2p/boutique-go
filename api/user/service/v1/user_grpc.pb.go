// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: v1/user.proto

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
	User_GetMe_FullMethodName        = "/user.service.v1.User/GetMe"
	User_Register_FullMethodName     = "/user.service.v1.User/Register"
	User_GetUserById_FullMethodName  = "/user.service.v1.User/GetUserById"
	User_GetIdByEmail_FullMethodName = "/user.service.v1.User/GetIdByEmail"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GetMe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMeResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
	GetIdByEmail(ctx context.Context, in *GetIdByEmailReq, opts ...grpc.CallOption) (*GetIdByEmailResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetMe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMeResp, error) {
	out := new(GetMeResp)
	err := c.cc.Invoke(ctx, User_GetMe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	out := new(GetUserByIdResp)
	err := c.cc.Invoke(ctx, User_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetIdByEmail(ctx context.Context, in *GetIdByEmailReq, opts ...grpc.CallOption) (*GetIdByEmailResp, error) {
	out := new(GetIdByEmailResp)
	err := c.cc.Invoke(ctx, User_GetIdByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetMe(context.Context, *emptypb.Empty) (*GetMeResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error)
	GetIdByEmail(context.Context, *GetIdByEmailReq) (*GetIdByEmailResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetMe(context.Context, *emptypb.Empty) (*GetMeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServer) GetIdByEmail(context.Context, *GetIdByEmailReq) (*GetIdByEmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdByEmail not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetMe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetMe(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetIdByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetIdByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetIdByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetIdByEmail(ctx, req.(*GetIdByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.service.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMe",
			Handler:    _User_GetMe_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
		{
			MethodName: "GetIdByEmail",
			Handler:    _User_GetIdByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/user.proto",
}
