// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: hello-server/proto/hello.proto

package proto

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

// SayHelloClient is the client API for SayHello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SayHelloClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type sayHelloClient struct {
	cc grpc.ClientConnInterface
}

func NewSayHelloClient(cc grpc.ClientConnInterface) SayHelloClient {
	return &sayHelloClient{cc}
}

func (c *sayHelloClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/SayHello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayHelloServer is the server API for SayHello service.
// All implementations must embed UnimplementedSayHelloServer
// for forward compatibility
type SayHelloServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	mustEmbedUnimplementedSayHelloServer()
}

// UnimplementedSayHelloServer must be embedded to have forward compatible implementations.
type UnimplementedSayHelloServer struct {
}

func (UnimplementedSayHelloServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedSayHelloServer) mustEmbedUnimplementedSayHelloServer() {}

// UnsafeSayHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SayHelloServer will
// result in compilation errors.
type UnsafeSayHelloServer interface {
	mustEmbedUnimplementedSayHelloServer()
}

func RegisterSayHelloServer(s grpc.ServiceRegistrar, srv SayHelloServer) {
	s.RegisterService(&SayHello_ServiceDesc, srv)
}

func _SayHello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayHelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SayHello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayHelloServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SayHello_ServiceDesc is the grpc.ServiceDesc for SayHello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SayHello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SayHello",
	HandlerType: (*SayHelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SayHello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello-server/proto/hello.proto",
}
