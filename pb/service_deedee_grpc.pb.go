// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: service_deedee.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DeeDee_CreateUser_FullMethodName       = "/pb.DeeDee/CreateUser"
	DeeDee_UpdateUser_FullMethodName       = "/pb.DeeDee/UpdateUser"
	DeeDee_LoginUser_FullMethodName        = "/pb.DeeDee/LoginUser"
	DeeDee_CreateRecruiter_FullMethodName  = "/pb.DeeDee/CreateRecruiter"
	DeeDee_GetRecruiters_FullMethodName    = "/pb.DeeDee/GetRecruiters"
	DeeDee_GetRecruiter_FullMethodName     = "/pb.DeeDee/GetRecruiter"
	DeeDee_GetResume_FullMethodName        = "/pb.DeeDee/GetResume"
	DeeDee_CreateResume_FullMethodName     = "/pb.DeeDee/CreateResume"
	DeeDee_UpdateResume_FullMethodName     = "/pb.DeeDee/UpdateResume"
	DeeDee_VerifyEmail_FullMethodName      = "/pb.DeeDee/VerifyEmail"
	DeeDee_RenewAccessToken_FullMethodName = "/pb.DeeDee/RenewAccessToken"
)

// DeeDeeClient is the client API for DeeDee service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeeDeeClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	CreateRecruiter(ctx context.Context, in *CreateRecruiterRequest, opts ...grpc.CallOption) (*CreateRecruiterResponse, error)
	GetRecruiters(ctx context.Context, in *GetRecruitersRequest, opts ...grpc.CallOption) (*GetRecruitersResponse, error)
	GetRecruiter(ctx context.Context, in *GetRecruiterRequest, opts ...grpc.CallOption) (*GetRecruiterResponse, error)
	GetResume(ctx context.Context, in *GetResumeRequest, opts ...grpc.CallOption) (*GetResumeResponse, error)
	CreateResume(ctx context.Context, in *CreateResumeRequest, opts ...grpc.CallOption) (*CreateResumeResponse, error)
	UpdateResume(ctx context.Context, in *UpdateResumeRequest, opts ...grpc.CallOption) (*UpdateResumeResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error)
}

type deeDeeClient struct {
	cc grpc.ClientConnInterface
}

func NewDeeDeeClient(cc grpc.ClientConnInterface) DeeDeeClient {
	return &deeDeeClient{cc}
}

func (c *deeDeeClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, DeeDee_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, DeeDee_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, DeeDee_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) CreateRecruiter(ctx context.Context, in *CreateRecruiterRequest, opts ...grpc.CallOption) (*CreateRecruiterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRecruiterResponse)
	err := c.cc.Invoke(ctx, DeeDee_CreateRecruiter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) GetRecruiters(ctx context.Context, in *GetRecruitersRequest, opts ...grpc.CallOption) (*GetRecruitersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecruitersResponse)
	err := c.cc.Invoke(ctx, DeeDee_GetRecruiters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) GetRecruiter(ctx context.Context, in *GetRecruiterRequest, opts ...grpc.CallOption) (*GetRecruiterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecruiterResponse)
	err := c.cc.Invoke(ctx, DeeDee_GetRecruiter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) GetResume(ctx context.Context, in *GetResumeRequest, opts ...grpc.CallOption) (*GetResumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResumeResponse)
	err := c.cc.Invoke(ctx, DeeDee_GetResume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) CreateResume(ctx context.Context, in *CreateResumeRequest, opts ...grpc.CallOption) (*CreateResumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResumeResponse)
	err := c.cc.Invoke(ctx, DeeDee_CreateResume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) UpdateResume(ctx context.Context, in *UpdateResumeRequest, opts ...grpc.CallOption) (*UpdateResumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResumeResponse)
	err := c.cc.Invoke(ctx, DeeDee_UpdateResume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, DeeDee_VerifyEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deeDeeClient) RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RenewAccessTokenResponse)
	err := c.cc.Invoke(ctx, DeeDee_RenewAccessToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeeDeeServer is the server API for DeeDee service.
// All implementations must embed UnimplementedDeeDeeServer
// for forward compatibility.
type DeeDeeServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	CreateRecruiter(context.Context, *CreateRecruiterRequest) (*CreateRecruiterResponse, error)
	GetRecruiters(context.Context, *GetRecruitersRequest) (*GetRecruitersResponse, error)
	GetRecruiter(context.Context, *GetRecruiterRequest) (*GetRecruiterResponse, error)
	GetResume(context.Context, *GetResumeRequest) (*GetResumeResponse, error)
	CreateResume(context.Context, *CreateResumeRequest) (*CreateResumeResponse, error)
	UpdateResume(context.Context, *UpdateResumeRequest) (*UpdateResumeResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error)
	mustEmbedUnimplementedDeeDeeServer()
}

// UnimplementedDeeDeeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeeDeeServer struct{}

func (UnimplementedDeeDeeServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedDeeDeeServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedDeeDeeServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedDeeDeeServer) CreateRecruiter(context.Context, *CreateRecruiterRequest) (*CreateRecruiterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecruiter not implemented")
}
func (UnimplementedDeeDeeServer) GetRecruiters(context.Context, *GetRecruitersRequest) (*GetRecruitersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecruiters not implemented")
}
func (UnimplementedDeeDeeServer) GetRecruiter(context.Context, *GetRecruiterRequest) (*GetRecruiterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecruiter not implemented")
}
func (UnimplementedDeeDeeServer) GetResume(context.Context, *GetResumeRequest) (*GetResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResume not implemented")
}
func (UnimplementedDeeDeeServer) CreateResume(context.Context, *CreateResumeRequest) (*CreateResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResume not implemented")
}
func (UnimplementedDeeDeeServer) UpdateResume(context.Context, *UpdateResumeRequest) (*UpdateResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResume not implemented")
}
func (UnimplementedDeeDeeServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedDeeDeeServer) RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewAccessToken not implemented")
}
func (UnimplementedDeeDeeServer) mustEmbedUnimplementedDeeDeeServer() {}
func (UnimplementedDeeDeeServer) testEmbeddedByValue()                {}

// UnsafeDeeDeeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeeDeeServer will
// result in compilation errors.
type UnsafeDeeDeeServer interface {
	mustEmbedUnimplementedDeeDeeServer()
}

func RegisterDeeDeeServer(s grpc.ServiceRegistrar, srv DeeDeeServer) {
	// If the following call pancis, it indicates UnimplementedDeeDeeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeeDee_ServiceDesc, srv)
}

func _DeeDee_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_CreateRecruiter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecruiterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).CreateRecruiter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_CreateRecruiter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).CreateRecruiter(ctx, req.(*CreateRecruiterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_GetRecruiters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecruitersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).GetRecruiters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_GetRecruiters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).GetRecruiters(ctx, req.(*GetRecruitersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_GetRecruiter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecruiterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).GetRecruiter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_GetRecruiter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).GetRecruiter(ctx, req.(*GetRecruiterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_GetResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).GetResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_GetResume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).GetResume(ctx, req.(*GetResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_CreateResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).CreateResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_CreateResume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).CreateResume(ctx, req.(*CreateResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_UpdateResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).UpdateResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_UpdateResume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).UpdateResume(ctx, req.(*UpdateResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeeDee_RenewAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeeDeeServer).RenewAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeeDee_RenewAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeeDeeServer).RenewAccessToken(ctx, req.(*RenewAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeeDee_ServiceDesc is the grpc.ServiceDesc for DeeDee service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeeDee_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DeeDee",
	HandlerType: (*DeeDeeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _DeeDee_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _DeeDee_UpdateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _DeeDee_LoginUser_Handler,
		},
		{
			MethodName: "CreateRecruiter",
			Handler:    _DeeDee_CreateRecruiter_Handler,
		},
		{
			MethodName: "GetRecruiters",
			Handler:    _DeeDee_GetRecruiters_Handler,
		},
		{
			MethodName: "GetRecruiter",
			Handler:    _DeeDee_GetRecruiter_Handler,
		},
		{
			MethodName: "GetResume",
			Handler:    _DeeDee_GetResume_Handler,
		},
		{
			MethodName: "CreateResume",
			Handler:    _DeeDee_CreateResume_Handler,
		},
		{
			MethodName: "UpdateResume",
			Handler:    _DeeDee_UpdateResume_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _DeeDee_VerifyEmail_Handler,
		},
		{
			MethodName: "RenewAccessToken",
			Handler:    _DeeDee_RenewAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_deedee.proto",
}
