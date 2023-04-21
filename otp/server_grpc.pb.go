// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: otp/server.proto

package otp

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
	OtpService_HealthCheck_FullMethodName      = "/otp.OtpService/HealthCheck"
	OtpService_CreateAndSendOtp_FullMethodName = "/otp.OtpService/CreateAndSendOtp"
	OtpService_VerifyOtp_FullMethodName        = "/otp.OtpService/VerifyOtp"
	OtpService_ResendOTP_FullMethodName        = "/otp.OtpService/ResendOTP"
)

// OtpServiceClient is the client API for OtpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OtpServiceClient interface {
	HealthCheck(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	CreateAndSendOtp(ctx context.Context, in *CreateAndSendOtpReq, opts ...grpc.CallOption) (*CreateAndSendOtpRes, error)
	VerifyOtp(ctx context.Context, in *VerifyOTPReq, opts ...grpc.CallOption) (*VerifyOTPRes, error)
	ResendOTP(ctx context.Context, in *ResendOTPRed, opts ...grpc.CallOption) (*ResendOTPRes, error)
}

type otpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOtpServiceClient(cc grpc.ClientConnInterface) OtpServiceClient {
	return &otpServiceClient{cc}
}

func (c *otpServiceClient) HealthCheck(ctx context.Context, in *DefaultRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, OtpService_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otpServiceClient) CreateAndSendOtp(ctx context.Context, in *CreateAndSendOtpReq, opts ...grpc.CallOption) (*CreateAndSendOtpRes, error) {
	out := new(CreateAndSendOtpRes)
	err := c.cc.Invoke(ctx, OtpService_CreateAndSendOtp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otpServiceClient) VerifyOtp(ctx context.Context, in *VerifyOTPReq, opts ...grpc.CallOption) (*VerifyOTPRes, error) {
	out := new(VerifyOTPRes)
	err := c.cc.Invoke(ctx, OtpService_VerifyOtp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otpServiceClient) ResendOTP(ctx context.Context, in *ResendOTPRed, opts ...grpc.CallOption) (*ResendOTPRes, error) {
	out := new(ResendOTPRes)
	err := c.cc.Invoke(ctx, OtpService_ResendOTP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OtpServiceServer is the server API for OtpService service.
// All implementations must embed UnimplementedOtpServiceServer
// for forward compatibility
type OtpServiceServer interface {
	HealthCheck(context.Context, *DefaultRequest) (*HealthResponse, error)
	CreateAndSendOtp(context.Context, *CreateAndSendOtpReq) (*CreateAndSendOtpRes, error)
	VerifyOtp(context.Context, *VerifyOTPReq) (*VerifyOTPRes, error)
	ResendOTP(context.Context, *ResendOTPRed) (*ResendOTPRes, error)
	mustEmbedUnimplementedOtpServiceServer()
}

// UnimplementedOtpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOtpServiceServer struct {
}

func (UnimplementedOtpServiceServer) HealthCheck(context.Context, *DefaultRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedOtpServiceServer) CreateAndSendOtp(context.Context, *CreateAndSendOtpReq) (*CreateAndSendOtpRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAndSendOtp not implemented")
}
func (UnimplementedOtpServiceServer) VerifyOtp(context.Context, *VerifyOTPReq) (*VerifyOTPRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOtp not implemented")
}
func (UnimplementedOtpServiceServer) ResendOTP(context.Context, *ResendOTPRed) (*ResendOTPRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendOTP not implemented")
}
func (UnimplementedOtpServiceServer) mustEmbedUnimplementedOtpServiceServer() {}

// UnsafeOtpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OtpServiceServer will
// result in compilation errors.
type UnsafeOtpServiceServer interface {
	mustEmbedUnimplementedOtpServiceServer()
}

func RegisterOtpServiceServer(s grpc.ServiceRegistrar, srv OtpServiceServer) {
	s.RegisterService(&OtpService_ServiceDesc, srv)
}

func _OtpService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OtpService_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpServiceServer).HealthCheck(ctx, req.(*DefaultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OtpService_CreateAndSendOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAndSendOtpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpServiceServer).CreateAndSendOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OtpService_CreateAndSendOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpServiceServer).CreateAndSendOtp(ctx, req.(*CreateAndSendOtpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OtpService_VerifyOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpServiceServer).VerifyOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OtpService_VerifyOtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpServiceServer).VerifyOtp(ctx, req.(*VerifyOTPReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OtpService_ResendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendOTPRed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpServiceServer).ResendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OtpService_ResendOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpServiceServer).ResendOTP(ctx, req.(*ResendOTPRed))
	}
	return interceptor(ctx, in, info, handler)
}

// OtpService_ServiceDesc is the grpc.ServiceDesc for OtpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OtpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "otp.OtpService",
	HandlerType: (*OtpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _OtpService_HealthCheck_Handler,
		},
		{
			MethodName: "CreateAndSendOtp",
			Handler:    _OtpService_CreateAndSendOtp_Handler,
		},
		{
			MethodName: "VerifyOtp",
			Handler:    _OtpService_VerifyOtp_Handler,
		},
		{
			MethodName: "ResendOTP",
			Handler:    _OtpService_ResendOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "otp/server.proto",
}
