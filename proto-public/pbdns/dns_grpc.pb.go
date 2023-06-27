// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pbdns/dns.proto

package pbdns

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

// DNSServiceClient is the client API for DNSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSServiceClient interface {
	// Query sends a DNS request over to Consul server and returns a DNS reply message.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type dNSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSServiceClient(cc grpc.ClientConnInterface) DNSServiceClient {
	return &dNSServiceClient{cc}
}

func (c *dNSServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.consul.dns.DNSService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSServiceServer is the server API for DNSService service.
// All implementations should embed UnimplementedDNSServiceServer
// for forward compatibility
type DNSServiceServer interface {
	// Query sends a DNS request over to Consul server and returns a DNS reply message.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedDNSServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDNSServiceServer struct {
}

func (UnimplementedDNSServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

// UnsafeDNSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DNSServiceServer will
// result in compilation errors.
type UnsafeDNSServiceServer interface {
	mustEmbedUnimplementedDNSServiceServer()
}

func RegisterDNSServiceServer(s grpc.ServiceRegistrar, srv DNSServiceServer) {
	s.RegisterService(&DNSService_ServiceDesc, srv)
}

func _DNSService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.consul.dns.DNSService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DNSService_ServiceDesc is the grpc.ServiceDesc for DNSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.consul.dns.DNSService",
	HandlerType: (*DNSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DNSService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbdns/dns.proto",
}
