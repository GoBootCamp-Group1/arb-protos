// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: TranslateService/translations.proto

package arb_test

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
	Translation_Translate_FullMethodName = "/Translation/Translate"
)

// TranslationClient is the client API for Translation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslationClient interface {
	Translate(ctx context.Context, in *TranslationInput, opts ...grpc.CallOption) (*TranslationOutput, error)
}

type translationClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslationClient(cc grpc.ClientConnInterface) TranslationClient {
	return &translationClient{cc}
}

func (c *translationClient) Translate(ctx context.Context, in *TranslationInput, opts ...grpc.CallOption) (*TranslationOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TranslationOutput)
	err := c.cc.Invoke(ctx, Translation_Translate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslationServer is the server API for Translation service.
// All implementations must embed UnimplementedTranslationServer
// for forward compatibility
type TranslationServer interface {
	Translate(context.Context, *TranslationInput) (*TranslationOutput, error)
	mustEmbedUnimplementedTranslationServer()
}

// UnimplementedTranslationServer must be embedded to have forward compatible implementations.
type UnimplementedTranslationServer struct {
}

func (UnimplementedTranslationServer) Translate(context.Context, *TranslationInput) (*TranslationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedTranslationServer) mustEmbedUnimplementedTranslationServer() {}

// UnsafeTranslationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslationServer will
// result in compilation errors.
type UnsafeTranslationServer interface {
	mustEmbedUnimplementedTranslationServer()
}

func RegisterTranslationServer(s grpc.ServiceRegistrar, srv TranslationServer) {
	s.RegisterService(&Translation_ServiceDesc, srv)
}

func _Translation_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Translation_Translate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServer).Translate(ctx, req.(*TranslationInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Translation_ServiceDesc is the grpc.ServiceDesc for Translation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Translation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Translation",
	HandlerType: (*TranslationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _Translation_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TranslateService/translations.proto",
}
