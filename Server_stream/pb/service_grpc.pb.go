// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CountDown_Start_FullMethodName = "/CountDown/Start"
)

// CountDownClient is the client API for CountDown service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountDownClient interface {
	Start(ctx context.Context, in *CountdownRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CountDownResponse], error)
}

type countDownClient struct {
	cc grpc.ClientConnInterface
}

func NewCountDownClient(cc grpc.ClientConnInterface) CountDownClient {
	return &countDownClient{cc}
}

func (c *countDownClient) Start(ctx context.Context, in *CountdownRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CountDownResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CountDown_ServiceDesc.Streams[0], CountDown_Start_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CountdownRequest, CountDownResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CountDown_StartClient = grpc.ServerStreamingClient[CountDownResponse]

// CountDownServer is the server API for CountDown service.
// All implementations must embed UnimplementedCountDownServer
// for forward compatibility.
type CountDownServer interface {
	Start(*CountdownRequest, grpc.ServerStreamingServer[CountDownResponse]) error
	mustEmbedUnimplementedCountDownServer()
}

// UnimplementedCountDownServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCountDownServer struct{}

func (UnimplementedCountDownServer) Start(*CountdownRequest, grpc.ServerStreamingServer[CountDownResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedCountDownServer) mustEmbedUnimplementedCountDownServer() {}
func (UnimplementedCountDownServer) testEmbeddedByValue()                   {}

// UnsafeCountDownServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountDownServer will
// result in compilation errors.
type UnsafeCountDownServer interface {
	mustEmbedUnimplementedCountDownServer()
}

func RegisterCountDownServer(s grpc.ServiceRegistrar, srv CountDownServer) {
	// If the following call pancis, it indicates UnimplementedCountDownServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CountDown_ServiceDesc, srv)
}

func _CountDown_Start_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CountdownRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CountDownServer).Start(m, &grpc.GenericServerStream[CountdownRequest, CountDownResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CountDown_StartServer = grpc.ServerStreamingServer[CountDownResponse]

// CountDown_ServiceDesc is the grpc.ServiceDesc for CountDown service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountDown_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CountDown",
	HandlerType: (*CountDownServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Start",
			Handler:       _CountDown_Start_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
