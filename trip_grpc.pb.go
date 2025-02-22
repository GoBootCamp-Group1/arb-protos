// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: trip.proto

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
	TripService_CreateTrip_FullMethodName     = "/transportation.TripService/CreateTrip"
	TripService_GetTrip_FullMethodName        = "/transportation.TripService/GetTrip"
	TripService_UpdateTrip_FullMethodName     = "/transportation.TripService/UpdateTrip"
	TripService_DeleteTrip_FullMethodName     = "/transportation.TripService/DeleteTrip"
	TripService_SearchTrips_FullMethodName    = "/transportation.TripService/SearchTrips"
	TripService_GetTrips_FullMethodName       = "/transportation.TripService/GetTrips"
	TripService_GetTripsAgency_FullMethodName = "/transportation.TripService/GetTripsAgency"
	TripService_ActivateTrip_FullMethodName   = "/transportation.TripService/ActivateTrip"
	TripService_FinishTrip_FullMethodName     = "/transportation.TripService/FinishTrip"
)

// TripServiceClient is the client API for TripService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TripServiceClient interface {
	CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*CreateTripResponse, error)
	GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error)
	UpdateTrip(ctx context.Context, in *UpdateTripRequest, opts ...grpc.CallOption) (*UpdateTripResponse, error)
	DeleteTrip(ctx context.Context, in *DeleteTripRequest, opts ...grpc.CallOption) (*DeleteTripResponse, error)
	SearchTrips(ctx context.Context, in *SearchTripsRequest, opts ...grpc.CallOption) (*SearchTripsResponse, error)
	GetTrips(ctx context.Context, in *GetTripsRequest, opts ...grpc.CallOption) (*SearchTripsResponse, error)
	GetTripsAgency(ctx context.Context, in *GetTripsRequest, opts ...grpc.CallOption) (*SearchTripsResponse, error)
	ActivateTrip(ctx context.Context, in *ActivateTripRequest, opts ...grpc.CallOption) (*ActivateTripResponse, error)
	FinishTrip(ctx context.Context, in *FinishTripRequest, opts ...grpc.CallOption) (*FinishTripResponse, error)
}

type tripServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTripServiceClient(cc grpc.ClientConnInterface) TripServiceClient {
	return &tripServiceClient{cc}
}

func (c *tripServiceClient) CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*CreateTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTripResponse)
	err := c.cc.Invoke(ctx, TripService_CreateTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTripResponse)
	err := c.cc.Invoke(ctx, TripService_GetTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) UpdateTrip(ctx context.Context, in *UpdateTripRequest, opts ...grpc.CallOption) (*UpdateTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTripResponse)
	err := c.cc.Invoke(ctx, TripService_UpdateTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) DeleteTrip(ctx context.Context, in *DeleteTripRequest, opts ...grpc.CallOption) (*DeleteTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTripResponse)
	err := c.cc.Invoke(ctx, TripService_DeleteTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) SearchTrips(ctx context.Context, in *SearchTripsRequest, opts ...grpc.CallOption) (*SearchTripsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTripsResponse)
	err := c.cc.Invoke(ctx, TripService_SearchTrips_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTrips(ctx context.Context, in *GetTripsRequest, opts ...grpc.CallOption) (*SearchTripsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTripsResponse)
	err := c.cc.Invoke(ctx, TripService_GetTrips_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTripsAgency(ctx context.Context, in *GetTripsRequest, opts ...grpc.CallOption) (*SearchTripsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTripsResponse)
	err := c.cc.Invoke(ctx, TripService_GetTripsAgency_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) ActivateTrip(ctx context.Context, in *ActivateTripRequest, opts ...grpc.CallOption) (*ActivateTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivateTripResponse)
	err := c.cc.Invoke(ctx, TripService_ActivateTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) FinishTrip(ctx context.Context, in *FinishTripRequest, opts ...grpc.CallOption) (*FinishTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FinishTripResponse)
	err := c.cc.Invoke(ctx, TripService_FinishTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TripServiceServer is the server API for TripService service.
// All implementations must embed UnimplementedTripServiceServer
// for forward compatibility
type TripServiceServer interface {
	CreateTrip(context.Context, *CreateTripRequest) (*CreateTripResponse, error)
	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
	UpdateTrip(context.Context, *UpdateTripRequest) (*UpdateTripResponse, error)
	DeleteTrip(context.Context, *DeleteTripRequest) (*DeleteTripResponse, error)
	SearchTrips(context.Context, *SearchTripsRequest) (*SearchTripsResponse, error)
	GetTrips(context.Context, *GetTripsRequest) (*SearchTripsResponse, error)
	GetTripsAgency(context.Context, *GetTripsRequest) (*SearchTripsResponse, error)
	ActivateTrip(context.Context, *ActivateTripRequest) (*ActivateTripResponse, error)
	FinishTrip(context.Context, *FinishTripRequest) (*FinishTripResponse, error)
	mustEmbedUnimplementedTripServiceServer()
}

// UnimplementedTripServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTripServiceServer struct {
}

func (UnimplementedTripServiceServer) CreateTrip(context.Context, *CreateTripRequest) (*CreateTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}
func (UnimplementedTripServiceServer) GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (UnimplementedTripServiceServer) UpdateTrip(context.Context, *UpdateTripRequest) (*UpdateTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}
func (UnimplementedTripServiceServer) DeleteTrip(context.Context, *DeleteTripRequest) (*DeleteTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrip not implemented")
}
func (UnimplementedTripServiceServer) SearchTrips(context.Context, *SearchTripsRequest) (*SearchTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTrips not implemented")
}
func (UnimplementedTripServiceServer) GetTrips(context.Context, *GetTripsRequest) (*SearchTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrips not implemented")
}
func (UnimplementedTripServiceServer) GetTripsAgency(context.Context, *GetTripsRequest) (*SearchTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTripsAgency not implemented")
}
func (UnimplementedTripServiceServer) ActivateTrip(context.Context, *ActivateTripRequest) (*ActivateTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateTrip not implemented")
}
func (UnimplementedTripServiceServer) FinishTrip(context.Context, *FinishTripRequest) (*FinishTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishTrip not implemented")
}
func (UnimplementedTripServiceServer) mustEmbedUnimplementedTripServiceServer() {}

// UnsafeTripServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TripServiceServer will
// result in compilation errors.
type UnsafeTripServiceServer interface {
	mustEmbedUnimplementedTripServiceServer()
}

func RegisterTripServiceServer(s grpc.ServiceRegistrar, srv TripServiceServer) {
	s.RegisterService(&TripService_ServiceDesc, srv)
}

func _TripService_CreateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).CreateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_CreateTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).CreateTrip(ctx, req.(*CreateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_GetTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTrip(ctx, req.(*GetTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_UpdateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).UpdateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_UpdateTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).UpdateTrip(ctx, req.(*UpdateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_DeleteTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).DeleteTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_DeleteTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).DeleteTrip(ctx, req.(*DeleteTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_SearchTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).SearchTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_SearchTrips_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).SearchTrips(ctx, req.(*SearchTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_GetTrips_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTrips(ctx, req.(*GetTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetTripsAgency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTripsAgency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_GetTripsAgency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTripsAgency(ctx, req.(*GetTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_ActivateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).ActivateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_ActivateTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).ActivateTrip(ctx, req.(*ActivateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_FinishTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).FinishTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TripService_FinishTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).FinishTrip(ctx, req.(*FinishTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TripService_ServiceDesc is the grpc.ServiceDesc for TripService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TripService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transportation.TripService",
	HandlerType: (*TripServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrip",
			Handler:    _TripService_CreateTrip_Handler,
		},
		{
			MethodName: "GetTrip",
			Handler:    _TripService_GetTrip_Handler,
		},
		{
			MethodName: "UpdateTrip",
			Handler:    _TripService_UpdateTrip_Handler,
		},
		{
			MethodName: "DeleteTrip",
			Handler:    _TripService_DeleteTrip_Handler,
		},
		{
			MethodName: "SearchTrips",
			Handler:    _TripService_SearchTrips_Handler,
		},
		{
			MethodName: "GetTrips",
			Handler:    _TripService_GetTrips_Handler,
		},
		{
			MethodName: "GetTripsAgency",
			Handler:    _TripService_GetTripsAgency_Handler,
		},
		{
			MethodName: "ActivateTrip",
			Handler:    _TripService_ActivateTrip_Handler,
		},
		{
			MethodName: "FinishTrip",
			Handler:    _TripService_FinishTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trip.proto",
}
