// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.3
// source: auditProto.proto

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

// AuditInterfaceClient is the client API for AuditInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditInterfaceClient interface {
	ObtainProtoAudit(ctx context.Context, in *AuditProtoRequest, opts ...grpc.CallOption) (*AuditProtoResponse, error)
	EventInfo(ctx context.Context, in *AuditProtoRequest, opts ...grpc.CallOption) (*AuditProtoResponse, error)
	EventInfoStream(ctx context.Context, opts ...grpc.CallOption) (AuditInterface_EventInfoStreamClient, error)
	FlowAudit(ctx context.Context, in *AuditProtoRequest, opts ...grpc.CallOption) (*AuditProtoResponse, error)
	AllowList(ctx context.Context, in *AuditProtoRequest, opts ...grpc.CallOption) (*AuditProtoResponse, error)
}

type auditInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditInterfaceClient(cc grpc.ClientConnInterface) AuditInterfaceClient {
	return &auditInterfaceClient{cc}
}

func (c *auditInterfaceClient) ObtainProtoAudit(ctx context.Context, in *AuditProtoRequest, opts ...grpc.CallOption) (*AuditProtoResponse, error) {
	out := new(AuditProtoResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditInterface/ObtainProtoAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditInterfaceClient) EventInfo(ctx context.Context, in *AuditProtoRequest, opts ...grpc.CallOption) (*AuditProtoResponse, error) {
	out := new(AuditProtoResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditInterface/EventInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditInterfaceClient) EventInfoStream(ctx context.Context, opts ...grpc.CallOption) (AuditInterface_EventInfoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &AuditInterface_ServiceDesc.Streams[0], "/audit.AuditInterface/EventInfoStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &auditInterfaceEventInfoStreamClient{stream}
	return x, nil
}

type AuditInterface_EventInfoStreamClient interface {
	Send(*AuditProtoRequest) error
	CloseAndRecv() (*AuditProtoResponse, error)
	grpc.ClientStream
}

type auditInterfaceEventInfoStreamClient struct {
	grpc.ClientStream
}

func (x *auditInterfaceEventInfoStreamClient) Send(m *AuditProtoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *auditInterfaceEventInfoStreamClient) CloseAndRecv() (*AuditProtoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AuditProtoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *auditInterfaceClient) FlowAudit(ctx context.Context, in *AuditProtoRequest, opts ...grpc.CallOption) (*AuditProtoResponse, error) {
	out := new(AuditProtoResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditInterface/FlowAudit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditInterfaceClient) AllowList(ctx context.Context, in *AuditProtoRequest, opts ...grpc.CallOption) (*AuditProtoResponse, error) {
	out := new(AuditProtoResponse)
	err := c.cc.Invoke(ctx, "/audit.AuditInterface/AllowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditInterfaceServer is the server API for AuditInterface service.
// All implementations must embed UnimplementedAuditInterfaceServer
// for forward compatibility
type AuditInterfaceServer interface {
	ObtainProtoAudit(context.Context, *AuditProtoRequest) (*AuditProtoResponse, error)
	EventInfo(context.Context, *AuditProtoRequest) (*AuditProtoResponse, error)
	EventInfoStream(AuditInterface_EventInfoStreamServer) error
	FlowAudit(context.Context, *AuditProtoRequest) (*AuditProtoResponse, error)
	AllowList(context.Context, *AuditProtoRequest) (*AuditProtoResponse, error)
	mustEmbedUnimplementedAuditInterfaceServer()
}

// UnimplementedAuditInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedAuditInterfaceServer struct {
}

func (UnimplementedAuditInterfaceServer) ObtainProtoAudit(context.Context, *AuditProtoRequest) (*AuditProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtainProtoAudit not implemented")
}
func (UnimplementedAuditInterfaceServer) EventInfo(context.Context, *AuditProtoRequest) (*AuditProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventInfo not implemented")
}
func (UnimplementedAuditInterfaceServer) EventInfoStream(AuditInterface_EventInfoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EventInfoStream not implemented")
}
func (UnimplementedAuditInterfaceServer) FlowAudit(context.Context, *AuditProtoRequest) (*AuditProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowAudit not implemented")
}
func (UnimplementedAuditInterfaceServer) AllowList(context.Context, *AuditProtoRequest) (*AuditProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowList not implemented")
}
func (UnimplementedAuditInterfaceServer) mustEmbedUnimplementedAuditInterfaceServer() {}

// UnsafeAuditInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditInterfaceServer will
// result in compilation errors.
type UnsafeAuditInterfaceServer interface {
	mustEmbedUnimplementedAuditInterfaceServer()
}

func RegisterAuditInterfaceServer(s grpc.ServiceRegistrar, srv AuditInterfaceServer) {
	s.RegisterService(&AuditInterface_ServiceDesc, srv)
}

func _AuditInterface_ObtainProtoAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditProtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditInterfaceServer).ObtainProtoAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditInterface/ObtainProtoAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditInterfaceServer).ObtainProtoAudit(ctx, req.(*AuditProtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditInterface_EventInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditProtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditInterfaceServer).EventInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditInterface/EventInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditInterfaceServer).EventInfo(ctx, req.(*AuditProtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditInterface_EventInfoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuditInterfaceServer).EventInfoStream(&auditInterfaceEventInfoStreamServer{stream})
}

type AuditInterface_EventInfoStreamServer interface {
	SendAndClose(*AuditProtoResponse) error
	Recv() (*AuditProtoRequest, error)
	grpc.ServerStream
}

type auditInterfaceEventInfoStreamServer struct {
	grpc.ServerStream
}

func (x *auditInterfaceEventInfoStreamServer) SendAndClose(m *AuditProtoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *auditInterfaceEventInfoStreamServer) Recv() (*AuditProtoRequest, error) {
	m := new(AuditProtoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AuditInterface_FlowAudit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditProtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditInterfaceServer).FlowAudit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditInterface/FlowAudit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditInterfaceServer).FlowAudit(ctx, req.(*AuditProtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditInterface_AllowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditProtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditInterfaceServer).AllowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/audit.AuditInterface/AllowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditInterfaceServer).AllowList(ctx, req.(*AuditProtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuditInterface_ServiceDesc is the grpc.ServiceDesc for AuditInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuditInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "audit.AuditInterface",
	HandlerType: (*AuditInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ObtainProtoAudit",
			Handler:    _AuditInterface_ObtainProtoAudit_Handler,
		},
		{
			MethodName: "EventInfo",
			Handler:    _AuditInterface_EventInfo_Handler,
		},
		{
			MethodName: "FlowAudit",
			Handler:    _AuditInterface_FlowAudit_Handler,
		},
		{
			MethodName: "AllowList",
			Handler:    _AuditInterface_AllowList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventInfoStream",
			Handler:       _AuditInterface_EventInfoStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "auditProto.proto",
}
