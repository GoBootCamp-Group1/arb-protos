// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: UserService/perm.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PermissionService_CreatePermission_FullMethodName    = "/role.PermissionService/CreatePermission"
	PermissionService_UpdatePermission_FullMethodName    = "/role.PermissionService/UpdatePermission"
	PermissionService_GetPermissionById_FullMethodName   = "/role.PermissionService/GetPermissionById"
	PermissionService_GetPermissionByName_FullMethodName = "/role.PermissionService/GetPermissionByName"
	PermissionService_DeletePermission_FullMethodName    = "/role.PermissionService/DeletePermission"
	PermissionService_GetAllPermissions_FullMethodName   = "/role.PermissionService/GetAllPermissions"
)

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error)
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error)
	GetPermissionById(ctx context.Context, in *GetPermissionByIdRequest, opts ...grpc.CallOption) (*GetPermissionByIdResponse, error)
	GetPermissionByName(ctx context.Context, in *GetPermissionByNameRequest, opts ...grpc.CallOption) (*GetPermissionByNameResponse, error)
	DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*DeletePermissionResponse, error)
	GetAllPermissions(ctx context.Context, in *GetAllPermissionsRequest, opts ...grpc.CallOption) (*GetAllPermissionsResponse, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePermissionResponse)
	err := c.cc.Invoke(ctx, PermissionService_CreatePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*UpdatePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePermissionResponse)
	err := c.cc.Invoke(ctx, PermissionService_UpdatePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetPermissionById(ctx context.Context, in *GetPermissionByIdRequest, opts ...grpc.CallOption) (*GetPermissionByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPermissionByIdResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetPermissionById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetPermissionByName(ctx context.Context, in *GetPermissionByNameRequest, opts ...grpc.CallOption) (*GetPermissionByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPermissionByNameResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetPermissionByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*DeletePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePermissionResponse)
	err := c.cc.Invoke(ctx, PermissionService_DeletePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetAllPermissions(ctx context.Context, in *GetAllPermissionsRequest, opts ...grpc.CallOption) (*GetAllPermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllPermissionsResponse)
	err := c.cc.Invoke(ctx, PermissionService_GetAllPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations must embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	CreatePermission(context.Context, *CreatePermissionRequest) (*CreatePermissionResponse, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*UpdatePermissionResponse, error)
	GetPermissionById(context.Context, *GetPermissionByIdRequest) (*GetPermissionByIdResponse, error)
	GetPermissionByName(context.Context, *GetPermissionByNameRequest) (*GetPermissionByNameResponse, error)
	DeletePermission(context.Context, *DeletePermissionRequest) (*DeletePermissionResponse, error)
	GetAllPermissions(context.Context, *GetAllPermissionsRequest) (*GetAllPermissionsResponse, error)
	mustEmbedUnimplementedPermissionServiceServer()
}

// UnimplementedPermissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) CreatePermission(context.Context, *CreatePermissionRequest) (*CreatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (UnimplementedPermissionServiceServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*UpdatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}
func (UnimplementedPermissionServiceServer) GetPermissionById(context.Context, *GetPermissionByIdRequest) (*GetPermissionByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionById not implemented")
}
func (UnimplementedPermissionServiceServer) GetPermissionByName(context.Context, *GetPermissionByNameRequest) (*GetPermissionByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionByName not implemented")
}
func (UnimplementedPermissionServiceServer) DeletePermission(context.Context, *DeletePermissionRequest) (*DeletePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePermission not implemented")
}
func (UnimplementedPermissionServiceServer) GetAllPermissions(context.Context, *GetAllPermissionsRequest) (*GetAllPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPermissions not implemented")
}
func (UnimplementedPermissionServiceServer) mustEmbedUnimplementedPermissionServiceServer() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_CreatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_UpdatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetPermissionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetPermissionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetPermissionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetPermissionById(ctx, req.(*GetPermissionByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetPermissionByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetPermissionByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetPermissionByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetPermissionByName(ctx, req.(*GetPermissionByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_DeletePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).DeletePermission(ctx, req.(*DeletePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetAllPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetAllPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionService_GetAllPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetAllPermissions(ctx, req.(*GetAllPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "role.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePermission",
			Handler:    _PermissionService_CreatePermission_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _PermissionService_UpdatePermission_Handler,
		},
		{
			MethodName: "GetPermissionById",
			Handler:    _PermissionService_GetPermissionById_Handler,
		},
		{
			MethodName: "GetPermissionByName",
			Handler:    _PermissionService_GetPermissionByName_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _PermissionService_DeletePermission_Handler,
		},
		{
			MethodName: "GetAllPermissions",
			Handler:    _PermissionService_GetAllPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UserService/perm.proto",
}
