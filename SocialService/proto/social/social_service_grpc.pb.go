// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: social_service.proto

package social

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

// SocialServiceClient is the client API for SocialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocialServiceClient interface {
	RegUser(ctx context.Context, in *RegUserRequest, opts ...grpc.CallOption) (*RegUserResponse, error)
	RequestToFollow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowIntentResponse, error)
}

type socialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialServiceClient(cc grpc.ClientConnInterface) SocialServiceClient {
	return &socialServiceClient{cc}
}

func (c *socialServiceClient) RegUser(ctx context.Context, in *RegUserRequest, opts ...grpc.CallOption) (*RegUserResponse, error) {
	out := new(RegUserResponse)
	err := c.cc.Invoke(ctx, "/SocialService/RegUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialServiceClient) RequestToFollow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowIntentResponse, error) {
	out := new(FollowIntentResponse)
	err := c.cc.Invoke(ctx, "/SocialService/RequestToFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialServiceServer is the server API for SocialService service.
// All implementations must embed UnimplementedSocialServiceServer
// for forward compatibility
type SocialServiceServer interface {
	RegUser(context.Context, *RegUserRequest) (*RegUserResponse, error)
	RequestToFollow(context.Context, *FollowRequest) (*FollowIntentResponse, error)
	mustEmbedUnimplementedSocialServiceServer()
}

// UnimplementedSocialServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSocialServiceServer struct {
}

func (UnimplementedSocialServiceServer) RegUser(context.Context, *RegUserRequest) (*RegUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegUser not implemented")
}
func (UnimplementedSocialServiceServer) RequestToFollow(context.Context, *FollowRequest) (*FollowIntentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToFollow not implemented")
}
func (UnimplementedSocialServiceServer) mustEmbedUnimplementedSocialServiceServer() {}

// UnsafeSocialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialServiceServer will
// result in compilation errors.
type UnsafeSocialServiceServer interface {
	mustEmbedUnimplementedSocialServiceServer()
}

func RegisterSocialServiceServer(s grpc.ServiceRegistrar, srv SocialServiceServer) {
	s.RegisterService(&SocialService_ServiceDesc, srv)
}

func _SocialService_RegUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).RegUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SocialService/RegUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).RegUser(ctx, req.(*RegUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocialService_RequestToFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialServiceServer).RequestToFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SocialService/RequestToFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialServiceServer).RequestToFollow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SocialService_ServiceDesc is the grpc.ServiceDesc for SocialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SocialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SocialService",
	HandlerType: (*SocialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegUser",
			Handler:    _SocialService_RegUser_Handler,
		},
		{
			MethodName: "RequestToFollow",
			Handler:    _SocialService_RequestToFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "social_service.proto",
}
