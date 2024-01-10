// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: SagaActions.proto

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

const (
	GRPCSagaService_FindSagaObject_FullMethodName         = "/proto.gRPCSagaService/FindSagaObject"
	GRPCSagaService_FindSagaObjectChildren_FullMethodName = "/proto.gRPCSagaService/FindSagaObjectChildren"
	GRPCSagaService_DeleteObject_FullMethodName           = "/proto.gRPCSagaService/DeleteObject"
	GRPCSagaService_UnDeleteObject_FullMethodName         = "/proto.gRPCSagaService/UnDeleteObject"
)

// GRPCSagaServiceClient is the client API for GRPCSagaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCSagaServiceClient interface {
	// RPC method to get user information
	FindSagaObject(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*SagaObject, error)
	FindSagaObjectChildren(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error)
	DeleteObject(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*SagaObject, error)
	UnDeleteObject(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*SagaObject, error)
}

type gRPCSagaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCSagaServiceClient(cc grpc.ClientConnInterface) GRPCSagaServiceClient {
	return &gRPCSagaServiceClient{cc}
}

func (c *gRPCSagaServiceClient) FindSagaObject(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*SagaObject, error) {
	out := new(SagaObject)
	err := c.cc.Invoke(ctx, GRPCSagaService_FindSagaObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCSagaServiceClient) FindSagaObjectChildren(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectResponse, error) {
	out := new(ObjectResponse)
	err := c.cc.Invoke(ctx, GRPCSagaService_FindSagaObjectChildren_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCSagaServiceClient) DeleteObject(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*SagaObject, error) {
	out := new(SagaObject)
	err := c.cc.Invoke(ctx, GRPCSagaService_DeleteObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCSagaServiceClient) UnDeleteObject(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*SagaObject, error) {
	out := new(SagaObject)
	err := c.cc.Invoke(ctx, GRPCSagaService_UnDeleteObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCSagaServiceServer is the server API for GRPCSagaService service.
// All implementations must embed UnimplementedGRPCSagaServiceServer
// for forward compatibility
type GRPCSagaServiceServer interface {
	// RPC method to get user information
	FindSagaObject(context.Context, *ObjectRequest) (*SagaObject, error)
	FindSagaObjectChildren(context.Context, *ObjectRequest) (*ObjectResponse, error)
	DeleteObject(context.Context, *ObjectRequest) (*SagaObject, error)
	UnDeleteObject(context.Context, *ObjectRequest) (*SagaObject, error)
	mustEmbedUnimplementedGRPCSagaServiceServer()
}

// UnimplementedGRPCSagaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCSagaServiceServer struct {
}

func (UnimplementedGRPCSagaServiceServer) FindSagaObject(context.Context, *ObjectRequest) (*SagaObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSagaObject not implemented")
}
func (UnimplementedGRPCSagaServiceServer) FindSagaObjectChildren(context.Context, *ObjectRequest) (*ObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSagaObjectChildren not implemented")
}
func (UnimplementedGRPCSagaServiceServer) DeleteObject(context.Context, *ObjectRequest) (*SagaObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedGRPCSagaServiceServer) UnDeleteObject(context.Context, *ObjectRequest) (*SagaObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnDeleteObject not implemented")
}
func (UnimplementedGRPCSagaServiceServer) mustEmbedUnimplementedGRPCSagaServiceServer() {}

// UnsafeGRPCSagaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCSagaServiceServer will
// result in compilation errors.
type UnsafeGRPCSagaServiceServer interface {
	mustEmbedUnimplementedGRPCSagaServiceServer()
}

func RegisterGRPCSagaServiceServer(s grpc.ServiceRegistrar, srv GRPCSagaServiceServer) {
	s.RegisterService(&GRPCSagaService_ServiceDesc, srv)
}

func _GRPCSagaService_FindSagaObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCSagaServiceServer).FindSagaObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCSagaService_FindSagaObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCSagaServiceServer).FindSagaObject(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCSagaService_FindSagaObjectChildren_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCSagaServiceServer).FindSagaObjectChildren(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCSagaService_FindSagaObjectChildren_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCSagaServiceServer).FindSagaObjectChildren(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCSagaService_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCSagaServiceServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCSagaService_DeleteObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCSagaServiceServer).DeleteObject(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCSagaService_UnDeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCSagaServiceServer).UnDeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCSagaService_UnDeleteObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCSagaServiceServer).UnDeleteObject(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCSagaService_ServiceDesc is the grpc.ServiceDesc for GRPCSagaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCSagaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.gRPCSagaService",
	HandlerType: (*GRPCSagaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSagaObject",
			Handler:    _GRPCSagaService_FindSagaObject_Handler,
		},
		{
			MethodName: "FindSagaObjectChildren",
			Handler:    _GRPCSagaService_FindSagaObjectChildren_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _GRPCSagaService_DeleteObject_Handler,
		},
		{
			MethodName: "UnDeleteObject",
			Handler:    _GRPCSagaService_UnDeleteObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SagaActions.proto",
}
