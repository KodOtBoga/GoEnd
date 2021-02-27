package greetpb

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (DivideService_DivideClient, error)
}

type divideServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &divideServiceClient{cc}
}

func (c *divideServiceClient) Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (DivideService_DivideClient, error) {
	stream, err := c.cc.NewStream(ctx, &DivideService_ServiceDesc.Streams[0], "/greet.GreetService/Divide", opts...)
	if err != nil {
		return nil, err
	}
	x := &divideServiceDivideClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DivideService_DivideClient interface {
	Recv() (*DivideResponse, error)
	grpc.ClientStream
}

type divideServiceDivideClient struct {
	grpc.ClientStream
}

func (x *divideServiceDivideClient) Recv() (*DivideResponse, error) {
	m := new(DivideResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type GreetServiceServer interface {
	Divide(*DivideRequest, DivideService_DivideServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedDivideServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDivideServiceServer struct {
}

func (UnimplementedDivideServiceServer) Divide(*DivideRequest, DivideService_DivideServer) error {
	return status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedDivideServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&DivideService_ServiceDesc, srv)
}

func _DivideService_Divide_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DivideRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).Divide(m, &greetServiceDivideServer{stream})
}

type DivideService_DivideServer interface {
	Send(*DivideResponse) error
	grpc.ServerStream
}

type greetServiceDivideServer struct {
	grpc.ServerStream
}

func (x *greetServiceDivideServer) Send(m *DivideResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DivideService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DivideService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Divide",
			Handler:       _DivideService_Divide_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}