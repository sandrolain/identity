// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: clientgrpc.proto

package clientgrpc

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

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	// User APIs
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	LoginConfirm(ctx context.Context, in *LoginConfirmRequest, opts ...grpc.CallOption) (*LoginConfirmResponse, error)
	InitValidation(ctx context.Context, in *InitValidationRequest, opts ...grpc.CallOption) (*InitValidationResponse, error)
	CompleteValidation(ctx context.Context, in *CompleteValidationRequest, opts ...grpc.CallOption) (*CompleteValidationResponse, error)
	GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
	PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*PasswordChangeResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	BeginWebauthnRegister(ctx context.Context, in *BeginWebauthnRegisterRequest, opts ...grpc.CallOption) (*BeginWebauthnRegisterResponse, error)
	FinishWebauthnRegister(ctx context.Context, in *FinishWebauthnRegisterRequest, opts ...grpc.CallOption) (*FinishWebauthnRegisterResponse, error)
	BeginWebauthnLogin(ctx context.Context, in *BeginWebauthnLoginRequest, opts ...grpc.CallOption) (*BeginWebauthnLoginResponse, error)
	FinishWebauthnLogin(ctx context.Context, in *FinishWebauthnLoginRequest, opts ...grpc.CallOption) (*FinishWebauthnLoginResponse, error)
	// Machine APIs
	AuthenticateMachine(ctx context.Context, in *AuthenticateMachineRequest, opts ...grpc.CallOption) (*AuthenticateMachineResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) LoginConfirm(ctx context.Context, in *LoginConfirmRequest, opts ...grpc.CallOption) (*LoginConfirmResponse, error) {
	out := new(LoginConfirmResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/LoginConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) InitValidation(ctx context.Context, in *InitValidationRequest, opts ...grpc.CallOption) (*InitValidationResponse, error) {
	out := new(InitValidationResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/InitValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CompleteValidation(ctx context.Context, in *CompleteValidationRequest, opts ...grpc.CallOption) (*CompleteValidationResponse, error) {
	out := new(CompleteValidationResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/CompleteValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	out := new(GetUserDetailsResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/GetUserDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) PasswordChange(ctx context.Context, in *PasswordChangeRequest, opts ...grpc.CallOption) (*PasswordChangeResponse, error) {
	out := new(PasswordChangeResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/PasswordChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) BeginWebauthnRegister(ctx context.Context, in *BeginWebauthnRegisterRequest, opts ...grpc.CallOption) (*BeginWebauthnRegisterResponse, error) {
	out := new(BeginWebauthnRegisterResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/BeginWebauthnRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) FinishWebauthnRegister(ctx context.Context, in *FinishWebauthnRegisterRequest, opts ...grpc.CallOption) (*FinishWebauthnRegisterResponse, error) {
	out := new(FinishWebauthnRegisterResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/FinishWebauthnRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) BeginWebauthnLogin(ctx context.Context, in *BeginWebauthnLoginRequest, opts ...grpc.CallOption) (*BeginWebauthnLoginResponse, error) {
	out := new(BeginWebauthnLoginResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/BeginWebauthnLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) FinishWebauthnLogin(ctx context.Context, in *FinishWebauthnLoginRequest, opts ...grpc.CallOption) (*FinishWebauthnLoginResponse, error) {
	out := new(FinishWebauthnLoginResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/FinishWebauthnLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) AuthenticateMachine(ctx context.Context, in *AuthenticateMachineRequest, opts ...grpc.CallOption) (*AuthenticateMachineResponse, error) {
	out := new(AuthenticateMachineResponse)
	err := c.cc.Invoke(ctx, "/clientgrpc.ClientService/AuthenticateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	// User APIs
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	LoginConfirm(context.Context, *LoginConfirmRequest) (*LoginConfirmResponse, error)
	InitValidation(context.Context, *InitValidationRequest) (*InitValidationResponse, error)
	CompleteValidation(context.Context, *CompleteValidationRequest) (*CompleteValidationResponse, error)
	GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error)
	PasswordChange(context.Context, *PasswordChangeRequest) (*PasswordChangeResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	BeginWebauthnRegister(context.Context, *BeginWebauthnRegisterRequest) (*BeginWebauthnRegisterResponse, error)
	FinishWebauthnRegister(context.Context, *FinishWebauthnRegisterRequest) (*FinishWebauthnRegisterResponse, error)
	BeginWebauthnLogin(context.Context, *BeginWebauthnLoginRequest) (*BeginWebauthnLoginResponse, error)
	FinishWebauthnLogin(context.Context, *FinishWebauthnLoginRequest) (*FinishWebauthnLoginResponse, error)
	// Machine APIs
	AuthenticateMachine(context.Context, *AuthenticateMachineRequest) (*AuthenticateMachineResponse, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedClientServiceServer) LoginConfirm(context.Context, *LoginConfirmRequest) (*LoginConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginConfirm not implemented")
}
func (UnimplementedClientServiceServer) InitValidation(context.Context, *InitValidationRequest) (*InitValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitValidation not implemented")
}
func (UnimplementedClientServiceServer) CompleteValidation(context.Context, *CompleteValidationRequest) (*CompleteValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteValidation not implemented")
}
func (UnimplementedClientServiceServer) GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetails not implemented")
}
func (UnimplementedClientServiceServer) PasswordChange(context.Context, *PasswordChangeRequest) (*PasswordChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordChange not implemented")
}
func (UnimplementedClientServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedClientServiceServer) BeginWebauthnRegister(context.Context, *BeginWebauthnRegisterRequest) (*BeginWebauthnRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginWebauthnRegister not implemented")
}
func (UnimplementedClientServiceServer) FinishWebauthnRegister(context.Context, *FinishWebauthnRegisterRequest) (*FinishWebauthnRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishWebauthnRegister not implemented")
}
func (UnimplementedClientServiceServer) BeginWebauthnLogin(context.Context, *BeginWebauthnLoginRequest) (*BeginWebauthnLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginWebauthnLogin not implemented")
}
func (UnimplementedClientServiceServer) FinishWebauthnLogin(context.Context, *FinishWebauthnLoginRequest) (*FinishWebauthnLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishWebauthnLogin not implemented")
}
func (UnimplementedClientServiceServer) AuthenticateMachine(context.Context, *AuthenticateMachineRequest) (*AuthenticateMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateMachine not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_LoginConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).LoginConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/LoginConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).LoginConfirm(ctx, req.(*LoginConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_InitValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).InitValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/InitValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).InitValidation(ctx, req.(*InitValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CompleteValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CompleteValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/CompleteValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CompleteValidation(ctx, req.(*CompleteValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/GetUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetUserDetails(ctx, req.(*GetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_PasswordChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).PasswordChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/PasswordChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).PasswordChange(ctx, req.(*PasswordChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_BeginWebauthnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginWebauthnRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).BeginWebauthnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/BeginWebauthnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).BeginWebauthnRegister(ctx, req.(*BeginWebauthnRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_FinishWebauthnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishWebauthnRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).FinishWebauthnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/FinishWebauthnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).FinishWebauthnRegister(ctx, req.(*FinishWebauthnRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_BeginWebauthnLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginWebauthnLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).BeginWebauthnLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/BeginWebauthnLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).BeginWebauthnLogin(ctx, req.(*BeginWebauthnLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_FinishWebauthnLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishWebauthnLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).FinishWebauthnLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/FinishWebauthnLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).FinishWebauthnLogin(ctx, req.(*FinishWebauthnLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_AuthenticateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).AuthenticateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientgrpc.ClientService/AuthenticateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).AuthenticateMachine(ctx, req.(*AuthenticateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clientgrpc.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ClientService_Login_Handler,
		},
		{
			MethodName: "LoginConfirm",
			Handler:    _ClientService_LoginConfirm_Handler,
		},
		{
			MethodName: "InitValidation",
			Handler:    _ClientService_InitValidation_Handler,
		},
		{
			MethodName: "CompleteValidation",
			Handler:    _ClientService_CompleteValidation_Handler,
		},
		{
			MethodName: "GetUserDetails",
			Handler:    _ClientService_GetUserDetails_Handler,
		},
		{
			MethodName: "PasswordChange",
			Handler:    _ClientService_PasswordChange_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ClientService_Logout_Handler,
		},
		{
			MethodName: "BeginWebauthnRegister",
			Handler:    _ClientService_BeginWebauthnRegister_Handler,
		},
		{
			MethodName: "FinishWebauthnRegister",
			Handler:    _ClientService_FinishWebauthnRegister_Handler,
		},
		{
			MethodName: "BeginWebauthnLogin",
			Handler:    _ClientService_BeginWebauthnLogin_Handler,
		},
		{
			MethodName: "FinishWebauthnLogin",
			Handler:    _ClientService_FinishWebauthnLogin_Handler,
		},
		{
			MethodName: "AuthenticateMachine",
			Handler:    _ClientService_AuthenticateMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clientgrpc.proto",
}
