// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: exerciseProto.proto

package exercise

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

// ExerciseServiceClient is the client API for ExerciseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExerciseServiceClient interface {
	Delete(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*Exercise, error)
	SoftDelete(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*Exercise, error)
	DeleteOne(ctx context.Context, in *ExerciseOneRequest, opts ...grpc.CallOption) (*Exercise, error)
	SoftDeleteOne(ctx context.Context, in *ExerciseOneRequest, opts ...grpc.CallOption) (*Exercise, error)
}

type exerciseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExerciseServiceClient(cc grpc.ClientConnInterface) ExerciseServiceClient {
	return &exerciseServiceClient{cc}
}

func (c *exerciseServiceClient) Delete(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*Exercise, error) {
	out := new(Exercise)
	err := c.cc.Invoke(ctx, "/ExerciseService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseServiceClient) SoftDelete(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*Exercise, error) {
	out := new(Exercise)
	err := c.cc.Invoke(ctx, "/ExerciseService/SoftDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseServiceClient) DeleteOne(ctx context.Context, in *ExerciseOneRequest, opts ...grpc.CallOption) (*Exercise, error) {
	out := new(Exercise)
	err := c.cc.Invoke(ctx, "/ExerciseService/DeleteOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseServiceClient) SoftDeleteOne(ctx context.Context, in *ExerciseOneRequest, opts ...grpc.CallOption) (*Exercise, error) {
	out := new(Exercise)
	err := c.cc.Invoke(ctx, "/ExerciseService/SoftDeleteOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExerciseServiceServer is the server API for ExerciseService service.
// All implementations must embed UnimplementedExerciseServiceServer
// for forward compatibility
type ExerciseServiceServer interface {
	Delete(context.Context, *ExerciseRequest) (*Exercise, error)
	SoftDelete(context.Context, *ExerciseRequest) (*Exercise, error)
	DeleteOne(context.Context, *ExerciseOneRequest) (*Exercise, error)
	SoftDeleteOne(context.Context, *ExerciseOneRequest) (*Exercise, error)
	mustEmbedUnimplementedExerciseServiceServer()
}

// UnimplementedExerciseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExerciseServiceServer struct {
}

func (UnimplementedExerciseServiceServer) Delete(context.Context, *ExerciseRequest) (*Exercise, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedExerciseServiceServer) SoftDelete(context.Context, *ExerciseRequest) (*Exercise, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftDelete not implemented")
}
func (UnimplementedExerciseServiceServer) DeleteOne(context.Context, *ExerciseOneRequest) (*Exercise, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOne not implemented")
}
func (UnimplementedExerciseServiceServer) SoftDeleteOne(context.Context, *ExerciseOneRequest) (*Exercise, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftDeleteOne not implemented")
}
func (UnimplementedExerciseServiceServer) mustEmbedUnimplementedExerciseServiceServer() {}

// UnsafeExerciseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExerciseServiceServer will
// result in compilation errors.
type UnsafeExerciseServiceServer interface {
	mustEmbedUnimplementedExerciseServiceServer()
}

func RegisterExerciseServiceServer(s grpc.ServiceRegistrar, srv ExerciseServiceServer) {
	s.RegisterService(&ExerciseService_ServiceDesc, srv)
}

func _ExerciseService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExerciseService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).Delete(ctx, req.(*ExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseService_SoftDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).SoftDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExerciseService/SoftDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).SoftDelete(ctx, req.(*ExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseService_DeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).DeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExerciseService/DeleteOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).DeleteOne(ctx, req.(*ExerciseOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseService_SoftDeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).SoftDeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExerciseService/SoftDeleteOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).SoftDeleteOne(ctx, req.(*ExerciseOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExerciseService_ServiceDesc is the grpc.ServiceDesc for ExerciseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExerciseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ExerciseService",
	HandlerType: (*ExerciseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _ExerciseService_Delete_Handler,
		},
		{
			MethodName: "SoftDelete",
			Handler:    _ExerciseService_SoftDelete_Handler,
		},
		{
			MethodName: "DeleteOne",
			Handler:    _ExerciseService_DeleteOne_Handler,
		},
		{
			MethodName: "SoftDeleteOne",
			Handler:    _ExerciseService_SoftDeleteOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exerciseProto.proto",
}
