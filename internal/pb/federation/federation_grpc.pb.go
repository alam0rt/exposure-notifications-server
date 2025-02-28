// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package federation

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

// FederationClient is the client API for Federation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FederationClient interface {
	Fetch(ctx context.Context, in *FederationFetchRequest, opts ...grpc.CallOption) (*FederationFetchResponse, error)
}

type federationClient struct {
	cc grpc.ClientConnInterface
}

func NewFederationClient(cc grpc.ClientConnInterface) FederationClient {
	return &federationClient{cc}
}

func (c *federationClient) Fetch(ctx context.Context, in *FederationFetchRequest, opts ...grpc.CallOption) (*FederationFetchResponse, error) {
	out := new(FederationFetchResponse)
	err := c.cc.Invoke(ctx, "/Federation/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederationServer is the server API for Federation service.
// All implementations must embed UnimplementedFederationServer
// for forward compatibility
type FederationServer interface {
	Fetch(context.Context, *FederationFetchRequest) (*FederationFetchResponse, error)
	mustEmbedUnimplementedFederationServer()
}

// UnimplementedFederationServer must be embedded to have forward compatible implementations.
type UnimplementedFederationServer struct {
}

func (UnimplementedFederationServer) Fetch(context.Context, *FederationFetchRequest) (*FederationFetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedFederationServer) mustEmbedUnimplementedFederationServer() {}

// UnsafeFederationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FederationServer will
// result in compilation errors.
type UnsafeFederationServer interface {
	mustEmbedUnimplementedFederationServer()
}

func RegisterFederationServer(s grpc.ServiceRegistrar, srv FederationServer) {
	s.RegisterService(&Federation_ServiceDesc, srv)
}

func _Federation_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FederationFetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Federation/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServer).Fetch(ctx, req.(*FederationFetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Federation_ServiceDesc is the grpc.ServiceDesc for Federation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Federation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Federation",
	HandlerType: (*FederationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Federation_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pb/federation/federation.proto",
}
