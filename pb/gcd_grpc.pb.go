// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// GCDServiceClient is the client API for GCDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GCDServiceClient interface {
	Compute(ctx context.Context, in *GCDRequest, opts ...grpc.CallOption) (*GCDResponse, error)
}

type gCDServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGCDServiceClient(cc grpc.ClientConnInterface) GCDServiceClient {
	return &gCDServiceClient{cc}
}

func (c *gCDServiceClient) Compute(ctx context.Context, in *GCDRequest, opts ...grpc.CallOption) (*GCDResponse, error) {
	out := new(GCDResponse)
	err := c.cc.Invoke(ctx, "/pb.GCDService/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GCDServiceServer is the server API for GCDService service.
// All implementations must embed UnimplementedGCDServiceServer
// for forward compatibility
type GCDServiceServer interface {
	Compute(context.Context, *GCDRequest) (*GCDResponse, error)
	mustEmbedUnimplementedGCDServiceServer()
}

// UnimplementedGCDServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGCDServiceServer struct {
}

func (UnimplementedGCDServiceServer) Compute(context.Context, *GCDRequest) (*GCDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
}
func (UnimplementedGCDServiceServer) mustEmbedUnimplementedGCDServiceServer() {}

// UnsafeGCDServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GCDServiceServer will
// result in compilation errors.
type UnsafeGCDServiceServer interface {
	mustEmbedUnimplementedGCDServiceServer()
}

func RegisterGCDServiceServer(s grpc.ServiceRegistrar, srv GCDServiceServer) {
	s.RegisterService(&GCDService_ServiceDesc, srv)
}

func _GCDService_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GCDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GCDServiceServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GCDService/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GCDServiceServer).Compute(ctx, req.(*GCDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GCDService_ServiceDesc is the grpc.ServiceDesc for GCDService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GCDService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GCDService",
	HandlerType: (*GCDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _GCDService_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gcd.proto",
}
