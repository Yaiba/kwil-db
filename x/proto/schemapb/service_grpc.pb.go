// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: kwil/schemasvc/service.proto

package schemapb

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

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemaServiceClient interface {
	PlanSchema(ctx context.Context, in *PlanSchemaRequest, opts ...grpc.CallOption) (*PlanSchemaResponse, error)
	ApplySchema(ctx context.Context, in *ApplySchemaRequest, opts ...grpc.CallOption) (*ApplySchemaResponse, error)
	GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error)
}

type schemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaServiceClient(cc grpc.ClientConnInterface) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) PlanSchema(ctx context.Context, in *PlanSchemaRequest, opts ...grpc.CallOption) (*PlanSchemaResponse, error) {
	out := new(PlanSchemaResponse)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/PlanSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ApplySchema(ctx context.Context, in *ApplySchemaRequest, opts ...grpc.CallOption) (*ApplySchemaResponse, error) {
	out := new(ApplySchemaResponse)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/ApplySchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error) {
	out := new(GetMetadataResponse)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
// All implementations must embed UnimplementedSchemaServiceServer
// for forward compatibility
type SchemaServiceServer interface {
	PlanSchema(context.Context, *PlanSchemaRequest) (*PlanSchemaResponse, error)
	ApplySchema(context.Context, *ApplySchemaRequest) (*ApplySchemaResponse, error)
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)
	mustEmbedUnimplementedSchemaServiceServer()
}

// UnimplementedSchemaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchemaServiceServer struct {
}

func (UnimplementedSchemaServiceServer) PlanSchema(context.Context, *PlanSchemaRequest) (*PlanSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlanSchema not implemented")
}
func (UnimplementedSchemaServiceServer) ApplySchema(context.Context, *ApplySchemaRequest) (*ApplySchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplySchema not implemented")
}
func (UnimplementedSchemaServiceServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedSchemaServiceServer) mustEmbedUnimplementedSchemaServiceServer() {}

// UnsafeSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemaServiceServer will
// result in compilation errors.
type UnsafeSchemaServiceServer interface {
	mustEmbedUnimplementedSchemaServiceServer()
}

func RegisterSchemaServiceServer(s grpc.ServiceRegistrar, srv SchemaServiceServer) {
	s.RegisterService(&SchemaService_ServiceDesc, srv)
}

func _SchemaService_PlanSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).PlanSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/PlanSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).PlanSchema(ctx, req.(*PlanSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ApplySchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplySchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ApplySchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/ApplySchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ApplySchema(ctx, req.(*ApplySchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetMetadata(ctx, req.(*GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemaService_ServiceDesc is the grpc.ServiceDesc for SchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schemasvc.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlanSchema",
			Handler:    _SchemaService_PlanSchema_Handler,
		},
		{
			MethodName: "ApplySchema",
			Handler:    _SchemaService_ApplySchema_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _SchemaService_GetMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kwil/schemasvc/service.proto",
}
