// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: protos/calculator.proto

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Addition(ctx context.Context, in *AdditionCalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
	Subtraction(ctx context.Context, in *SubtractionCalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
	Division(ctx context.Context, in *DivisionCalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
	Multiplication(ctx context.Context, in *MultiplicationCalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error)
	AllCalculations(ctx context.Context, in *AllCalculationRequest, opts ...grpc.CallOption) (*AllCalculationsResponse, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Addition(ctx context.Context, in *AdditionCalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Addition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Subtraction(ctx context.Context, in *SubtractionCalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Subtraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Division(ctx context.Context, in *DivisionCalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Division", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiplication(ctx context.Context, in *MultiplicationCalculationRequest, opts ...grpc.CallOption) (*CalculationResponse, error) {
	out := new(CalculationResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Multiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) AllCalculations(ctx context.Context, in *AllCalculationRequest, opts ...grpc.CallOption) (*AllCalculationsResponse, error) {
	out := new(AllCalculationsResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/AllCalculations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Addition(context.Context, *AdditionCalculationRequest) (*CalculationResponse, error)
	Subtraction(context.Context, *SubtractionCalculationRequest) (*CalculationResponse, error)
	Division(context.Context, *DivisionCalculationRequest) (*CalculationResponse, error)
	Multiplication(context.Context, *MultiplicationCalculationRequest) (*CalculationResponse, error)
	AllCalculations(context.Context, *AllCalculationRequest) (*AllCalculationsResponse, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Addition(context.Context, *AdditionCalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addition not implemented")
}
func (UnimplementedCalculatorServiceServer) Subtraction(context.Context, *SubtractionCalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtraction not implemented")
}
func (UnimplementedCalculatorServiceServer) Division(context.Context, *DivisionCalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Division not implemented")
}
func (UnimplementedCalculatorServiceServer) Multiplication(context.Context, *MultiplicationCalculationRequest) (*CalculationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiplication not implemented")
}
func (UnimplementedCalculatorServiceServer) AllCalculations(context.Context, *AllCalculationRequest) (*AllCalculationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllCalculations not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Addition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdditionCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Addition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Addition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Addition(ctx, req.(*AdditionCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Subtraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractionCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Subtraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Subtraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Subtraction(ctx, req.(*SubtractionCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Division_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivisionCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Division(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Division",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Division(ctx, req.(*DivisionCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplicationCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Multiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiplication(ctx, req.(*MultiplicationCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_AllCalculations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).AllCalculations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/AllCalculations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).AllCalculations(ctx, req.(*AllCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Addition",
			Handler:    _CalculatorService_Addition_Handler,
		},
		{
			MethodName: "Subtraction",
			Handler:    _CalculatorService_Subtraction_Handler,
		},
		{
			MethodName: "Division",
			Handler:    _CalculatorService_Division_Handler,
		},
		{
			MethodName: "Multiplication",
			Handler:    _CalculatorService_Multiplication_Handler,
		},
		{
			MethodName: "AllCalculations",
			Handler:    _CalculatorService_AllCalculations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/calculator.proto",
}
