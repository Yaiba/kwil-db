// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: kwil/schemasvc/service.proto

package schemasvc

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
	Plan(ctx context.Context, in *PlanRequest, opts ...grpc.CallOption) (*PlanResponse, error)
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
	ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error)
	ListQueries(ctx context.Context, in *ListQueriesRequest, opts ...grpc.CallOption) (*ListQueriesResponse, error)
	GetQuery(ctx context.Context, in *GetQueryRequest, opts ...grpc.CallOption) (*Query, error)
}

type schemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaServiceClient(cc grpc.ClientConnInterface) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) Plan(ctx context.Context, in *PlanRequest, opts ...grpc.CallOption) (*PlanResponse, error) {
	out := new(PlanResponse)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/Plan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/ListRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ListQueries(ctx context.Context, in *ListQueriesRequest, opts ...grpc.CallOption) (*ListQueriesResponse, error) {
	out := new(ListQueriesResponse)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/ListQueries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetQuery(ctx context.Context, in *GetQueryRequest, opts ...grpc.CallOption) (*Query, error) {
	out := new(Query)
	err := c.cc.Invoke(ctx, "/schemasvc.SchemaService/GetQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
// All implementations must embed UnimplementedSchemaServiceServer
// for forward compatibility
type SchemaServiceServer interface {
	Plan(context.Context, *PlanRequest) (*PlanResponse, error)
	Apply(context.Context, *ApplyRequest) (*ApplyResponse, error)
	ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error)
	GetRole(context.Context, *GetRoleRequest) (*Role, error)
	ListQueries(context.Context, *ListQueriesRequest) (*ListQueriesResponse, error)
	GetQuery(context.Context, *GetQueryRequest) (*Query, error)
	mustEmbedUnimplementedSchemaServiceServer()
}

// UnimplementedSchemaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchemaServiceServer struct {
}

func (UnimplementedSchemaServiceServer) Plan(context.Context, *PlanRequest) (*PlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Plan not implemented")
}
func (UnimplementedSchemaServiceServer) Apply(context.Context, *ApplyRequest) (*ApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedSchemaServiceServer) ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (UnimplementedSchemaServiceServer) GetRole(context.Context, *GetRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedSchemaServiceServer) ListQueries(context.Context, *ListQueriesRequest) (*ListQueriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQueries not implemented")
}
func (UnimplementedSchemaServiceServer) GetQuery(context.Context, *GetQueryRequest) (*Query, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuery not implemented")
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

func _SchemaService_Plan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).Plan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/Plan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).Plan(ctx, req.(*PlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/ListRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ListRoles(ctx, req.(*ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ListQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQueriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ListQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/ListQueries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ListQueries(ctx, req.(*ListQueriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schemasvc.SchemaService/GetQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetQuery(ctx, req.(*GetQueryRequest))
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
			MethodName: "Plan",
			Handler:    _SchemaService_Plan_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _SchemaService_Apply_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _SchemaService_ListRoles_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _SchemaService_GetRole_Handler,
		},
		{
			MethodName: "ListQueries",
			Handler:    _SchemaService_ListQueries_Handler,
		},
		{
			MethodName: "GetQuery",
			Handler:    _SchemaService_GetQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kwil/schemasvc/service.proto",
}
