// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.3
// source: allowListProto.proto

package test

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

// AllowListInterfaceClient is the client API for AllowListInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AllowListInterfaceClient interface {
	AllowListAudit(ctx context.Context, in *AllowListProtoRequest, opts ...grpc.CallOption) (*AllowListProtoResponse, error)
}

type allowListInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewAllowListInterfaceClient(cc grpc.ClientConnInterface) AllowListInterfaceClient {
	return &allowListInterfaceClient{cc}
}

func (c *allowListInterfaceClient) AllowListAudit(ctx context.Context, in *AllowListProtoRequest, opts ...grpc.CallOption) (*AllowListProtoResponse, error) {
	out := new(AllowListProtoResponse)
	err := c.cc.Invoke(ctx, "/audit.AllowListInterface/AllowListAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllowListInterfaceServer is the server API for AllowListInterface service.
// All implementations must embed UnimplementedAllowListInterfaceServer
// for forward compatibility
type AllowListInterfaceServer interface {
	AllowListAudit(context.Context, *AllowListProtoRequest) (*AllowListProtoResponse, error)
	mustEmbedUnimplementedAllowListInterfaceServer()
}

// UnimplementedAllowListInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedAllowListInterfaceServer struct {
}

func (UnimplementedAllowListInterfaceServer) AllowListAudit(context.Context, *AllowListProtoRequest) (*AllowListProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowListAudit not implemented")
}
func (UnimplementedAllowListInterfaceServer) mustEmbedUnimplementedAllowListInterfaceServer() {}

// UnsafeAllowListInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AllowListInterfaceServer will
// result in compilation errors.
type UnsafeAllowListInterfaceServer interface {
	mustEmbedUnimplementedAllowListInterfaceServer()
}

func RegisterAllowListInterfaceServer(s grpc.ServiceRegistrar, srv AllowListInterfaceServer) {
	s.RegisterService(&AllowListInterface_ServiceDesc, srv)
}

func _AllowListInterface_AllowListAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllowListProtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllowListInterfaceServer).AllowListAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AllowListInterface/AllowListAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllowListInterfaceServer).AllowListAudit(ctx, req.(*AllowListProtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AllowListInterface_ServiceDesc is the grpc.ServiceDesc for AllowListInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AllowListInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audit.AllowListInterface",
	HandlerType: (*AllowListInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllowListAudit",
			Handler:    _AllowListInterface_AllowListAudit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "allowListProto.proto",
}
