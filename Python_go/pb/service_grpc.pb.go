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
	CommunicationService_SendMessage_FullMethodName = "/CommunicationService/SendMessage"
)

// CommunicationServiceClient is the client API for CommunicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicationServiceClient interface {
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type communicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationServiceClient(cc grpc.ClientConnInterface) CommunicationServiceClient {
	return &communicationServiceClient{cc}
}

func (c *communicationServiceClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, CommunicationService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationServiceServer is the server API for CommunicationService service.
// All implementations must embed UnimplementedCommunicationServiceServer
// for forward compatibility.
type CommunicationServiceServer interface {
	SendMessage(context.Context, *MessageRequest) (*MessageResponse, error)
	mustEmbedUnimplementedCommunicationServiceServer()
}

// UnimplementedCommunicationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommunicationServiceServer struct{}

func (UnimplementedCommunicationServiceServer) SendMessage(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedCommunicationServiceServer) mustEmbedUnimplementedCommunicationServiceServer() {}
func (UnimplementedCommunicationServiceServer) testEmbeddedByValue()                              {}

// UnsafeCommunicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicationServiceServer will
// result in compilation errors.
type UnsafeCommunicationServiceServer interface {
	mustEmbedUnimplementedCommunicationServiceServer()
}

func RegisterCommunicationServiceServer(s grpc.ServiceRegistrar, srv CommunicationServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommunicationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommunicationService_ServiceDesc, srv)
}

func _CommunicationService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServiceServer).SendMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunicationService_ServiceDesc is the grpc.ServiceDesc for CommunicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CommunicationService",
	HandlerType: (*CommunicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _CommunicationService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
