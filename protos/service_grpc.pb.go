// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package props

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

// YourServiceClient is the client API for YourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YourServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type yourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYourServiceClient(cc grpc.ClientConnInterface) YourServiceClient {
	return &yourServiceClient{cc}
}

func (c *yourServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/service.YourService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YourServiceServer is the server API for YourService service.
// All implementations must embed UnimplementedYourServiceServer
// for forward compatibility
type YourServiceServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	mustEmbedUnimplementedYourServiceServer()
}

// UnimplementedYourServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYourServiceServer struct {
}

func (UnimplementedYourServiceServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedYourServiceServer) mustEmbedUnimplementedYourServiceServer() {}

// UnsafeYourServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YourServiceServer will
// result in compilation errors.
type UnsafeYourServiceServer interface {
	mustEmbedUnimplementedYourServiceServer()
}

func RegisterYourServiceServer(s grpc.ServiceRegistrar, srv YourServiceServer) {
	s.RegisterService(&YourService_ServiceDesc, srv)
}

func _YourService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.YourService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// YourService_ServiceDesc is the grpc.ServiceDesc for YourService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YourService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.YourService",
	HandlerType: (*YourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _YourService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/service.proto",
}
