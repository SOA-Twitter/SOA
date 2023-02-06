// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: profile_service.proto

package profile

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	Register(ctx context.Context, in *ProfileRegisterRequest, opts ...grpc.CallOption) (*ProfileRegisterResponse, error)
	GetUserProfile(ctx context.Context, in *UserProfRequest, opts ...grpc.CallOption) (*UserProfResponse, error)
	ManagePrivacy(ctx context.Context, in *ManagePrivacyRequest, opts ...grpc.CallOption) (*ManagePrivacyResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) Register(ctx context.Context, in *ProfileRegisterRequest, opts ...grpc.CallOption) (*ProfileRegisterResponse, error) {
	out := new(ProfileRegisterResponse)
	err := c.cc.Invoke(ctx, "/ProfileService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetUserProfile(ctx context.Context, in *UserProfRequest, opts ...grpc.CallOption) (*UserProfResponse, error) {
	out := new(UserProfResponse)
	err := c.cc.Invoke(ctx, "/ProfileService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) ManagePrivacy(ctx context.Context, in *ManagePrivacyRequest, opts ...grpc.CallOption) (*ManagePrivacyResponse, error) {
	out := new(ManagePrivacyResponse)
	err := c.cc.Invoke(ctx, "/ProfileService/ManagePrivacy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	Register(context.Context, *ProfileRegisterRequest) (*ProfileRegisterResponse, error)
	GetUserProfile(context.Context, *UserProfRequest) (*UserProfResponse, error)
	ManagePrivacy(context.Context, *ManagePrivacyRequest) (*ManagePrivacyResponse, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) Register(context.Context, *ProfileRegisterRequest) (*ProfileRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedProfileServiceServer) GetUserProfile(context.Context, *UserProfRequest) (*UserProfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedProfileServiceServer) ManagePrivacy(context.Context, *ManagePrivacyRequest) (*ManagePrivacyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManagePrivacy not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Register(ctx, req.(*ProfileRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetUserProfile(ctx, req.(*UserProfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_ManagePrivacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagePrivacyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).ManagePrivacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileService/ManagePrivacy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).ManagePrivacy(ctx, req.(*ManagePrivacyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ProfileService_Register_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _ProfileService_GetUserProfile_Handler,
		},
		{
			MethodName: "ManagePrivacy",
			Handler:    _ProfileService_ManagePrivacy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile_service.proto",
}
