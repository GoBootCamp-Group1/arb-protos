// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: technical_team.proto

package arb_protos

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
	TechnicalTeamService_CreateTechnicalTeam_FullMethodName = "/transportation.TechnicalTeamService/CreateTechnicalTeam"
	TechnicalTeamService_GetTechnicalTeam_FullMethodName    = "/transportation.TechnicalTeamService/GetTechnicalTeam"
	TechnicalTeamService_UpdateTechnicalTeam_FullMethodName = "/transportation.TechnicalTeamService/UpdateTechnicalTeam"
	TechnicalTeamService_DeleteTechnicalTeam_FullMethodName = "/transportation.TechnicalTeamService/DeleteTechnicalTeam"
)

// TechnicalTeamServiceClient is the client API for TechnicalTeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TechnicalTeamServiceClient interface {
	CreateTechnicalTeam(ctx context.Context, in *CreateTechnicalTeamRequest, opts ...grpc.CallOption) (*CreateTechnicalTeamResponse, error)
	GetTechnicalTeam(ctx context.Context, in *GetTechnicalTeamRequest, opts ...grpc.CallOption) (*GetTechnicalTeamResponse, error)
	UpdateTechnicalTeam(ctx context.Context, in *UpdateTechnicalTeamRequest, opts ...grpc.CallOption) (*UpdateTechnicalTeamResponse, error)
	DeleteTechnicalTeam(ctx context.Context, in *DeleteTechnicalTeamRequest, opts ...grpc.CallOption) (*DeleteTechnicalTeamResponse, error)
}

type technicalTeamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTechnicalTeamServiceClient(cc grpc.ClientConnInterface) TechnicalTeamServiceClient {
	return &technicalTeamServiceClient{cc}
}

func (c *technicalTeamServiceClient) CreateTechnicalTeam(ctx context.Context, in *CreateTechnicalTeamRequest, opts ...grpc.CallOption) (*CreateTechnicalTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTechnicalTeamResponse)
	err := c.cc.Invoke(ctx, TechnicalTeamService_CreateTechnicalTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *technicalTeamServiceClient) GetTechnicalTeam(ctx context.Context, in *GetTechnicalTeamRequest, opts ...grpc.CallOption) (*GetTechnicalTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTechnicalTeamResponse)
	err := c.cc.Invoke(ctx, TechnicalTeamService_GetTechnicalTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *technicalTeamServiceClient) UpdateTechnicalTeam(ctx context.Context, in *UpdateTechnicalTeamRequest, opts ...grpc.CallOption) (*UpdateTechnicalTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTechnicalTeamResponse)
	err := c.cc.Invoke(ctx, TechnicalTeamService_UpdateTechnicalTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *technicalTeamServiceClient) DeleteTechnicalTeam(ctx context.Context, in *DeleteTechnicalTeamRequest, opts ...grpc.CallOption) (*DeleteTechnicalTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTechnicalTeamResponse)
	err := c.cc.Invoke(ctx, TechnicalTeamService_DeleteTechnicalTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TechnicalTeamServiceServer is the server API for TechnicalTeamService service.
// All implementations must embed UnimplementedTechnicalTeamServiceServer
// for forward compatibility
type TechnicalTeamServiceServer interface {
	CreateTechnicalTeam(context.Context, *CreateTechnicalTeamRequest) (*CreateTechnicalTeamResponse, error)
	GetTechnicalTeam(context.Context, *GetTechnicalTeamRequest) (*GetTechnicalTeamResponse, error)
	UpdateTechnicalTeam(context.Context, *UpdateTechnicalTeamRequest) (*UpdateTechnicalTeamResponse, error)
	DeleteTechnicalTeam(context.Context, *DeleteTechnicalTeamRequest) (*DeleteTechnicalTeamResponse, error)
	mustEmbedUnimplementedTechnicalTeamServiceServer()
}

// UnimplementedTechnicalTeamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTechnicalTeamServiceServer struct {
}

func (UnimplementedTechnicalTeamServiceServer) CreateTechnicalTeam(context.Context, *CreateTechnicalTeamRequest) (*CreateTechnicalTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTechnicalTeam not implemented")
}
func (UnimplementedTechnicalTeamServiceServer) GetTechnicalTeam(context.Context, *GetTechnicalTeamRequest) (*GetTechnicalTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTechnicalTeam not implemented")
}
func (UnimplementedTechnicalTeamServiceServer) UpdateTechnicalTeam(context.Context, *UpdateTechnicalTeamRequest) (*UpdateTechnicalTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTechnicalTeam not implemented")
}
func (UnimplementedTechnicalTeamServiceServer) DeleteTechnicalTeam(context.Context, *DeleteTechnicalTeamRequest) (*DeleteTechnicalTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTechnicalTeam not implemented")
}
func (UnimplementedTechnicalTeamServiceServer) mustEmbedUnimplementedTechnicalTeamServiceServer() {}

// UnsafeTechnicalTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TechnicalTeamServiceServer will
// result in compilation errors.
type UnsafeTechnicalTeamServiceServer interface {
	mustEmbedUnimplementedTechnicalTeamServiceServer()
}

func RegisterTechnicalTeamServiceServer(s grpc.ServiceRegistrar, srv TechnicalTeamServiceServer) {
	s.RegisterService(&TechnicalTeamService_ServiceDesc, srv)
}

func _TechnicalTeamService_CreateTechnicalTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTechnicalTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnicalTeamServiceServer).CreateTechnicalTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TechnicalTeamService_CreateTechnicalTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnicalTeamServiceServer).CreateTechnicalTeam(ctx, req.(*CreateTechnicalTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TechnicalTeamService_GetTechnicalTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTechnicalTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnicalTeamServiceServer).GetTechnicalTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TechnicalTeamService_GetTechnicalTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnicalTeamServiceServer).GetTechnicalTeam(ctx, req.(*GetTechnicalTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TechnicalTeamService_UpdateTechnicalTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTechnicalTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnicalTeamServiceServer).UpdateTechnicalTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TechnicalTeamService_UpdateTechnicalTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnicalTeamServiceServer).UpdateTechnicalTeam(ctx, req.(*UpdateTechnicalTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TechnicalTeamService_DeleteTechnicalTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTechnicalTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TechnicalTeamServiceServer).DeleteTechnicalTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TechnicalTeamService_DeleteTechnicalTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TechnicalTeamServiceServer).DeleteTechnicalTeam(ctx, req.(*DeleteTechnicalTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TechnicalTeamService_ServiceDesc is the grpc.ServiceDesc for TechnicalTeamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TechnicalTeamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transportation.TechnicalTeamService",
	HandlerType: (*TechnicalTeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTechnicalTeam",
			Handler:    _TechnicalTeamService_CreateTechnicalTeam_Handler,
		},
		{
			MethodName: "GetTechnicalTeam",
			Handler:    _TechnicalTeamService_GetTechnicalTeam_Handler,
		},
		{
			MethodName: "UpdateTechnicalTeam",
			Handler:    _TechnicalTeamService_UpdateTechnicalTeam_Handler,
		},
		{
			MethodName: "DeleteTechnicalTeam",
			Handler:    _TechnicalTeamService_DeleteTechnicalTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "technical_team.proto",
}
