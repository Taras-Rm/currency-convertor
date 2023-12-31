// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/currency.proto

package api

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

// CurrencyConvertorClient is the client API for CurrencyConvertor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyConvertorClient interface {
	ExchangeRate(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (*ExchangeRateResponse, error)
}

type currencyConvertorClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyConvertorClient(cc grpc.ClientConnInterface) CurrencyConvertorClient {
	return &currencyConvertorClient{cc}
}

func (c *currencyConvertorClient) ExchangeRate(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (*ExchangeRateResponse, error) {
	out := new(ExchangeRateResponse)
	err := c.cc.Invoke(ctx, "/main.CurrencyConvertor/ExchangeRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyConvertorServer is the server API for CurrencyConvertor service.
// All implementations must embed UnimplementedCurrencyConvertorServer
// for forward compatibility
type CurrencyConvertorServer interface {
	ExchangeRate(context.Context, *ExchangeRateRequest) (*ExchangeRateResponse, error)
	mustEmbedUnimplementedCurrencyConvertorServer()
}

// UnimplementedCurrencyConvertorServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyConvertorServer struct {
}

func (UnimplementedCurrencyConvertorServer) ExchangeRate(context.Context, *ExchangeRateRequest) (*ExchangeRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeRate not implemented")
}
func (UnimplementedCurrencyConvertorServer) mustEmbedUnimplementedCurrencyConvertorServer() {}

// UnsafeCurrencyConvertorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyConvertorServer will
// result in compilation errors.
type UnsafeCurrencyConvertorServer interface {
	mustEmbedUnimplementedCurrencyConvertorServer()
}

func RegisterCurrencyConvertorServer(s grpc.ServiceRegistrar, srv CurrencyConvertorServer) {
	s.RegisterService(&CurrencyConvertor_ServiceDesc, srv)
}

func _CurrencyConvertor_ExchangeRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyConvertorServer).ExchangeRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CurrencyConvertor/ExchangeRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyConvertorServer).ExchangeRate(ctx, req.(*ExchangeRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrencyConvertor_ServiceDesc is the grpc.ServiceDesc for CurrencyConvertor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyConvertor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.CurrencyConvertor",
	HandlerType: (*CurrencyConvertorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExchangeRate",
			Handler:    _CurrencyConvertor_ExchangeRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/currency.proto",
}
