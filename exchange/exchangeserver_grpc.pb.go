// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: exchangeserver.proto

package exchange

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExchangeService_CreateConversionRate_FullMethodName    = "/db.ExchangeService/CreateConversionRate"
	ExchangeService_ReadConversionRate_FullMethodName      = "/db.ExchangeService/ReadConversionRate"
	ExchangeService_UpdateConversionRate_FullMethodName    = "/db.ExchangeService/UpdateConversionRate"
	ExchangeService_DeleteConversionRate_FullMethodName    = "/db.ExchangeService/DeleteConversionRate"
	ExchangeService_CreateAccount_FullMethodName           = "/db.ExchangeService/CreateAccount"
	ExchangeService_DeleteAccount_FullMethodName           = "/db.ExchangeService/DeleteAccount"
	ExchangeService_UpdateAccount_FullMethodName           = "/db.ExchangeService/UpdateAccount"
	ExchangeService_SearchAccount_FullMethodName           = "/db.ExchangeService/SearchAccount"
	ExchangeService_CreateTransaction_FullMethodName       = "/db.ExchangeService/CreateTransaction"
	ExchangeService_DeleteTransaction_FullMethodName       = "/db.ExchangeService/DeleteTransaction"
	ExchangeService_UpdateTransaction_FullMethodName       = "/db.ExchangeService/UpdateTransaction"
	ExchangeService_GetTransaction_FullMethodName          = "/db.ExchangeService/GetTransaction"
	ExchangeService_GetTransactionByAccount_FullMethodName = "/db.ExchangeService/GetTransactionByAccount"
	ExchangeService_CreateTrade_FullMethodName             = "/db.ExchangeService/CreateTrade"
	ExchangeService_DeleteTrade_FullMethodName             = "/db.ExchangeService/DeleteTrade"
	ExchangeService_UpdateTrade_FullMethodName             = "/db.ExchangeService/UpdateTrade"
	ExchangeService_GetTrade_FullMethodName                = "/db.ExchangeService/GetTrade"
	ExchangeService_GetTradeByAccount_FullMethodName       = "/db.ExchangeService/GetTradeByAccount"
)

// ExchangeServiceClient is the client API for ExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangeServiceClient interface {
	// RATES
	CreateConversionRate(ctx context.Context, in *CreateConversionRateRequest, opts ...grpc.CallOption) (*CreateConversionRateResponse, error)
	ReadConversionRate(ctx context.Context, in *ReadConversionRateRequest, opts ...grpc.CallOption) (*ReadConversionRateResponse, error)
	UpdateConversionRate(ctx context.Context, in *UpdateConversionRateRequest, opts ...grpc.CallOption) (*UpdateConversionRateResponse, error)
	DeleteConversionRate(ctx context.Context, in *DeleteConversionRateRequest, opts ...grpc.CallOption) (*DeleteConversionRateResponse, error)
	// ACCOUNT
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error)
	SearchAccount(ctx context.Context, in *SearchAccountRequest, opts ...grpc.CallOption) (*SearchAccountResponse, error)
	// TRANSACTION
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	GetTransactionByAccount(ctx context.Context, in *GetTransactionByAccountRequest, opts ...grpc.CallOption) (*GetTransactionByAccountResponse, error)
	// TRADING
	CreateTrade(ctx context.Context, in *CreateTradeRequest, opts ...grpc.CallOption) (*CreateTradeResponse, error)
	DeleteTrade(ctx context.Context, in *DeleteTradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateTrade(ctx context.Context, in *UpdateTradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTrade(ctx context.Context, in *GetTradeRequest, opts ...grpc.CallOption) (*Trade, error)
	GetTradeByAccount(ctx context.Context, in *GetTradeByAccountRequest, opts ...grpc.CallOption) (*GetTradeByAccountResponse, error)
}

type exchangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeServiceClient(cc grpc.ClientConnInterface) ExchangeServiceClient {
	return &exchangeServiceClient{cc}
}

func (c *exchangeServiceClient) CreateConversionRate(ctx context.Context, in *CreateConversionRateRequest, opts ...grpc.CallOption) (*CreateConversionRateResponse, error) {
	out := new(CreateConversionRateResponse)
	err := c.cc.Invoke(ctx, ExchangeService_CreateConversionRate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) ReadConversionRate(ctx context.Context, in *ReadConversionRateRequest, opts ...grpc.CallOption) (*ReadConversionRateResponse, error) {
	out := new(ReadConversionRateResponse)
	err := c.cc.Invoke(ctx, ExchangeService_ReadConversionRate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) UpdateConversionRate(ctx context.Context, in *UpdateConversionRateRequest, opts ...grpc.CallOption) (*UpdateConversionRateResponse, error) {
	out := new(UpdateConversionRateResponse)
	err := c.cc.Invoke(ctx, ExchangeService_UpdateConversionRate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) DeleteConversionRate(ctx context.Context, in *DeleteConversionRateRequest, opts ...grpc.CallOption) (*DeleteConversionRateResponse, error) {
	out := new(DeleteConversionRateResponse)
	err := c.cc.Invoke(ctx, ExchangeService_DeleteConversionRate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, ExchangeService_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, ExchangeService_DeleteAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountResponse, error) {
	out := new(UpdateAccountResponse)
	err := c.cc.Invoke(ctx, ExchangeService_UpdateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) SearchAccount(ctx context.Context, in *SearchAccountRequest, opts ...grpc.CallOption) (*SearchAccountResponse, error) {
	out := new(SearchAccountResponse)
	err := c.cc.Invoke(ctx, ExchangeService_SearchAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, ExchangeService_CreateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) DeleteTransaction(ctx context.Context, in *DeleteTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExchangeService_DeleteTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExchangeService_UpdateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, ExchangeService_GetTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) GetTransactionByAccount(ctx context.Context, in *GetTransactionByAccountRequest, opts ...grpc.CallOption) (*GetTransactionByAccountResponse, error) {
	out := new(GetTransactionByAccountResponse)
	err := c.cc.Invoke(ctx, ExchangeService_GetTransactionByAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) CreateTrade(ctx context.Context, in *CreateTradeRequest, opts ...grpc.CallOption) (*CreateTradeResponse, error) {
	out := new(CreateTradeResponse)
	err := c.cc.Invoke(ctx, ExchangeService_CreateTrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) DeleteTrade(ctx context.Context, in *DeleteTradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExchangeService_DeleteTrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) UpdateTrade(ctx context.Context, in *UpdateTradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExchangeService_UpdateTrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) GetTrade(ctx context.Context, in *GetTradeRequest, opts ...grpc.CallOption) (*Trade, error) {
	out := new(Trade)
	err := c.cc.Invoke(ctx, ExchangeService_GetTrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeServiceClient) GetTradeByAccount(ctx context.Context, in *GetTradeByAccountRequest, opts ...grpc.CallOption) (*GetTradeByAccountResponse, error) {
	out := new(GetTradeByAccountResponse)
	err := c.cc.Invoke(ctx, ExchangeService_GetTradeByAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServiceServer is the server API for ExchangeService service.
// All implementations must embed UnimplementedExchangeServiceServer
// for forward compatibility
type ExchangeServiceServer interface {
	// RATES
	CreateConversionRate(context.Context, *CreateConversionRateRequest) (*CreateConversionRateResponse, error)
	ReadConversionRate(context.Context, *ReadConversionRateRequest) (*ReadConversionRateResponse, error)
	UpdateConversionRate(context.Context, *UpdateConversionRateRequest) (*UpdateConversionRateResponse, error)
	DeleteConversionRate(context.Context, *DeleteConversionRateRequest) (*DeleteConversionRateResponse, error)
	// ACCOUNT
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error)
	SearchAccount(context.Context, *SearchAccountRequest) (*SearchAccountResponse, error)
	// TRANSACTION
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	DeleteTransaction(context.Context, *DeleteTransactionRequest) (*emptypb.Empty, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*emptypb.Empty, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error)
	GetTransactionByAccount(context.Context, *GetTransactionByAccountRequest) (*GetTransactionByAccountResponse, error)
	// TRADING
	CreateTrade(context.Context, *CreateTradeRequest) (*CreateTradeResponse, error)
	DeleteTrade(context.Context, *DeleteTradeRequest) (*emptypb.Empty, error)
	UpdateTrade(context.Context, *UpdateTradeRequest) (*emptypb.Empty, error)
	GetTrade(context.Context, *GetTradeRequest) (*Trade, error)
	GetTradeByAccount(context.Context, *GetTradeByAccountRequest) (*GetTradeByAccountResponse, error)
	mustEmbedUnimplementedExchangeServiceServer()
}

// UnimplementedExchangeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExchangeServiceServer struct {
}

func (UnimplementedExchangeServiceServer) CreateConversionRate(context.Context, *CreateConversionRateRequest) (*CreateConversionRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConversionRate not implemented")
}
func (UnimplementedExchangeServiceServer) ReadConversionRate(context.Context, *ReadConversionRateRequest) (*ReadConversionRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadConversionRate not implemented")
}
func (UnimplementedExchangeServiceServer) UpdateConversionRate(context.Context, *UpdateConversionRateRequest) (*UpdateConversionRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConversionRate not implemented")
}
func (UnimplementedExchangeServiceServer) DeleteConversionRate(context.Context, *DeleteConversionRateRequest) (*DeleteConversionRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConversionRate not implemented")
}
func (UnimplementedExchangeServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedExchangeServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedExchangeServiceServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedExchangeServiceServer) SearchAccount(context.Context, *SearchAccountRequest) (*SearchAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccount not implemented")
}
func (UnimplementedExchangeServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedExchangeServiceServer) DeleteTransaction(context.Context, *DeleteTransactionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTransaction not implemented")
}
func (UnimplementedExchangeServiceServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedExchangeServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedExchangeServiceServer) GetTransactionByAccount(context.Context, *GetTransactionByAccountRequest) (*GetTransactionByAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionByAccount not implemented")
}
func (UnimplementedExchangeServiceServer) CreateTrade(context.Context, *CreateTradeRequest) (*CreateTradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrade not implemented")
}
func (UnimplementedExchangeServiceServer) DeleteTrade(context.Context, *DeleteTradeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrade not implemented")
}
func (UnimplementedExchangeServiceServer) UpdateTrade(context.Context, *UpdateTradeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrade not implemented")
}
func (UnimplementedExchangeServiceServer) GetTrade(context.Context, *GetTradeRequest) (*Trade, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrade not implemented")
}
func (UnimplementedExchangeServiceServer) GetTradeByAccount(context.Context, *GetTradeByAccountRequest) (*GetTradeByAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTradeByAccount not implemented")
}
func (UnimplementedExchangeServiceServer) mustEmbedUnimplementedExchangeServiceServer() {}

// UnsafeExchangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeServiceServer will
// result in compilation errors.
type UnsafeExchangeServiceServer interface {
	mustEmbedUnimplementedExchangeServiceServer()
}

func RegisterExchangeServiceServer(s grpc.ServiceRegistrar, srv ExchangeServiceServer) {
	s.RegisterService(&ExchangeService_ServiceDesc, srv)
}

func _ExchangeService_CreateConversionRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConversionRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).CreateConversionRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_CreateConversionRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).CreateConversionRate(ctx, req.(*CreateConversionRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_ReadConversionRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadConversionRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).ReadConversionRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_ReadConversionRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).ReadConversionRate(ctx, req.(*ReadConversionRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_UpdateConversionRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConversionRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).UpdateConversionRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_UpdateConversionRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).UpdateConversionRate(ctx, req.(*UpdateConversionRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_DeleteConversionRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConversionRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).DeleteConversionRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_DeleteConversionRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).DeleteConversionRate(ctx, req.(*DeleteConversionRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_UpdateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_SearchAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).SearchAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_SearchAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).SearchAccount(ctx, req.(*SearchAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_DeleteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).DeleteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_DeleteTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).DeleteTransaction(ctx, req.(*DeleteTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_UpdateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_GetTransactionByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetTransactionByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetTransactionByAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetTransactionByAccount(ctx, req.(*GetTransactionByAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_CreateTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).CreateTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_CreateTrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).CreateTrade(ctx, req.(*CreateTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_DeleteTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).DeleteTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_DeleteTrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).DeleteTrade(ctx, req.(*DeleteTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_UpdateTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).UpdateTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_UpdateTrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).UpdateTrade(ctx, req.(*UpdateTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_GetTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetTrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetTrade(ctx, req.(*GetTradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeService_GetTradeByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradeByAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServiceServer).GetTradeByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeService_GetTradeByAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServiceServer).GetTradeByAccount(ctx, req.(*GetTradeByAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExchangeService_ServiceDesc is the grpc.ServiceDesc for ExchangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExchangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.ExchangeService",
	HandlerType: (*ExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConversionRate",
			Handler:    _ExchangeService_CreateConversionRate_Handler,
		},
		{
			MethodName: "ReadConversionRate",
			Handler:    _ExchangeService_ReadConversionRate_Handler,
		},
		{
			MethodName: "UpdateConversionRate",
			Handler:    _ExchangeService_UpdateConversionRate_Handler,
		},
		{
			MethodName: "DeleteConversionRate",
			Handler:    _ExchangeService_DeleteConversionRate_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _ExchangeService_CreateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _ExchangeService_DeleteAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _ExchangeService_UpdateAccount_Handler,
		},
		{
			MethodName: "SearchAccount",
			Handler:    _ExchangeService_SearchAccount_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _ExchangeService_CreateTransaction_Handler,
		},
		{
			MethodName: "DeleteTransaction",
			Handler:    _ExchangeService_DeleteTransaction_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _ExchangeService_UpdateTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _ExchangeService_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactionByAccount",
			Handler:    _ExchangeService_GetTransactionByAccount_Handler,
		},
		{
			MethodName: "CreateTrade",
			Handler:    _ExchangeService_CreateTrade_Handler,
		},
		{
			MethodName: "DeleteTrade",
			Handler:    _ExchangeService_DeleteTrade_Handler,
		},
		{
			MethodName: "UpdateTrade",
			Handler:    _ExchangeService_UpdateTrade_Handler,
		},
		{
			MethodName: "GetTrade",
			Handler:    _ExchangeService_GetTrade_Handler,
		},
		{
			MethodName: "GetTradeByAccount",
			Handler:    _ExchangeService_GetTradeByAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exchangeserver.proto",
}
