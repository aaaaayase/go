// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: auth.proto

package service

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
	Auth_SignIn_FullMethodName                    = "/Auth/SignIn"
	Auth_LinkIdentity_FullMethodName              = "/Auth/LinkIdentity"
	Auth_UnlinkIdentity_FullMethodName            = "/Auth/UnlinkIdentity"
	Auth_DeleteIdentities_FullMethodName          = "/Auth/DeleteIdentities"
	Auth_AuthRealName_FullMethodName              = "/Auth/AuthRealName"
	Auth_GetIdentity_FullMethodName               = "/Auth/GetIdentity"
	Auth_GetIdentitiesByExternalId_FullMethodName = "/Auth/GetIdentitiesByExternalId"
	Auth_GetIdentities_FullMethodName             = "/Auth/GetIdentities"
	Auth_GetIdentitiesByComboId_FullMethodName    = "/Auth/GetIdentitiesByComboId"
	Auth_GetClientMeta_FullMethodName             = "/Auth/GetClientMeta"
	Auth_DescribeIdentityProviders_FullMethodName = "/Auth/DescribeIdentityProviders"
	Auth_DescribeService_FullMethodName           = "/Auth/DescribeService"
	Auth_AddDenylist_FullMethodName               = "/Auth/AddDenylist"
	Auth_RemoveDenylist_FullMethodName            = "/Auth/RemoveDenylist"
	Auth_QueryDenylist_FullMethodName             = "/Auth/QueryDenylist"
	Auth_ListDenylist_FullMethodName              = "/Auth/ListDenylist"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 所有 rpc 都遵循的通用 gRPC Status Code
// https://grpc.github.io/grpc/core/md_doc_statuscodes.html
//
//	INVALID_ARGUMENT | 缺少参数或者参数错误
//	INTERNAL | 服务内部错误
type AuthClient interface {
	// 聚合登录。
	// 验证 IdP 提供的 Credential，返回聚合用户标识 Combo ID 与 IdP 中的身份信息
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	// InvalidArgument | 缺少参数或者参数错误
	// UNAUTHENTICATED | IdP 提供的 Credential 验证不通过
	// PermissionDenied | IdP 提供的 Credential 验证通过，但是该用户被封禁
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// 为已有的 Combo ID 关联更多 IdP 的身份。
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	//
	//	NOT_FOUND | 指定的 Combo ID 不存在
	//	UNAUTHENTICATED | IdP 提供的 Credential 验证不通过
	LinkIdentity(ctx context.Context, in *LinkIdentityRequest, opts ...grpc.CallOption) (*LinkIdentityResponse, error)
	// 解除关联得 IdP。
	// 该功能需要配置开关进行开放。
	UnlinkIdentity(ctx context.Context, in *UnlinkIdentityRequest, opts ...grpc.CallOption) (*UnlinkIdentityResponse, error)
	// 删除 Combo User，包括关联的 IdP 数据也会被同时删除。
	// 该功能需要配置开关进行开放。
	DeleteIdentities(ctx context.Context, in *DeleteIdentitiesRequest, opts ...grpc.CallOption) (*DeleteIdentitiesResponse, error)
	// 对 Combo ID 关联的某个 Identity 进行实名认证（真实姓名 + 身份证号）
	AuthRealName(ctx context.Context, in *AuthRealNameRequest, opts ...grpc.CallOption) (*AuthRealNameResponse, error)
	// Deprecated: Do not use.
	// 获取外部用户的聚合信息
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	//
	//	NOT_FOUND | 指定的外部用户不存在
	//
	// Deprecated: 请使用 GetIdentitiesByExternalId 代替
	GetIdentity(ctx context.Context, in *GetIdentityRequest, opts ...grpc.CallOption) (*GetIdentityResponse, error)
	// 根据 外部身份标识（external_id）获取用户的聚合信息
	GetIdentitiesByExternalId(ctx context.Context, in *GetIdentitiesByExternalIdRequest, opts ...grpc.CallOption) (*GetIdentitiesByExternalIdResponse, error)
	// Deprecated: Do not use.
	// 获取聚合用户所关联的所有身份，返回简略的身份信息。
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	//
	//	NOT_FOUND | 指定的 Combo ID 不存在
	//
	// Deprecated: 请使用 GetIdentitiesByComboId 代替
	GetIdentities(ctx context.Context, in *GetIdentitiesRequest, opts ...grpc.CallOption) (*GetIdentitiesResponse, error)
	// 根据 combo_id 获取用户的所有身份，返回完整的身份信息。
	GetIdentitiesByComboId(ctx context.Context, in *GetIdentitiesByComboIdRequest, opts ...grpc.CallOption) (*GetIdentitiesByComboIdResponse, error)
	// 获取聚合用户身份的客户端元数据。
	// 客户端元数据是在首次 SignIn 或 LinkIdentity 时写入的。
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	//
	//	NOT_FOUND | 指定的 Combo ID 不存在，或 Combo ID 未关联指定的 IdP
	GetClientMeta(ctx context.Context, in *GetClientMetaRequest, opts ...grpc.CallOption) (*GetClientMetaResponse, error)
	// 获取所有当前支持的 Identity Provider 的元数据
	DescribeIdentityProviders(ctx context.Context, in *DescribeIdentityProvidersRequest, opts ...grpc.CallOption) (*DescribeIdentityProvidersResponse, error)
	// 获取配置信息
	DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*DescribeServiceResponse, error)
	// 封禁数据管理接口
	AddDenylist(ctx context.Context, in *AddDenylistRequest, opts ...grpc.CallOption) (*AddDenylistResponse, error)
	RemoveDenylist(ctx context.Context, in *RemoveDenylistRequest, opts ...grpc.CallOption) (*RemoveDenylistResponse, error)
	QueryDenylist(ctx context.Context, in *QueryDenylistRequest, opts ...grpc.CallOption) (*QueryDenylistResponse, error)
	ListDenylist(ctx context.Context, in *ListDenylistRequest, opts ...grpc.CallOption) (*ListDenylistResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, Auth_SignIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LinkIdentity(ctx context.Context, in *LinkIdentityRequest, opts ...grpc.CallOption) (*LinkIdentityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LinkIdentityResponse)
	err := c.cc.Invoke(ctx, Auth_LinkIdentity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UnlinkIdentity(ctx context.Context, in *UnlinkIdentityRequest, opts ...grpc.CallOption) (*UnlinkIdentityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnlinkIdentityResponse)
	err := c.cc.Invoke(ctx, Auth_UnlinkIdentity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteIdentities(ctx context.Context, in *DeleteIdentitiesRequest, opts ...grpc.CallOption) (*DeleteIdentitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteIdentitiesResponse)
	err := c.cc.Invoke(ctx, Auth_DeleteIdentities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AuthRealName(ctx context.Context, in *AuthRealNameRequest, opts ...grpc.CallOption) (*AuthRealNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthRealNameResponse)
	err := c.cc.Invoke(ctx, Auth_AuthRealName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *authClient) GetIdentity(ctx context.Context, in *GetIdentityRequest, opts ...grpc.CallOption) (*GetIdentityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIdentityResponse)
	err := c.cc.Invoke(ctx, Auth_GetIdentity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetIdentitiesByExternalId(ctx context.Context, in *GetIdentitiesByExternalIdRequest, opts ...grpc.CallOption) (*GetIdentitiesByExternalIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIdentitiesByExternalIdResponse)
	err := c.cc.Invoke(ctx, Auth_GetIdentitiesByExternalId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *authClient) GetIdentities(ctx context.Context, in *GetIdentitiesRequest, opts ...grpc.CallOption) (*GetIdentitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIdentitiesResponse)
	err := c.cc.Invoke(ctx, Auth_GetIdentities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetIdentitiesByComboId(ctx context.Context, in *GetIdentitiesByComboIdRequest, opts ...grpc.CallOption) (*GetIdentitiesByComboIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIdentitiesByComboIdResponse)
	err := c.cc.Invoke(ctx, Auth_GetIdentitiesByComboId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetClientMeta(ctx context.Context, in *GetClientMetaRequest, opts ...grpc.CallOption) (*GetClientMetaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetClientMetaResponse)
	err := c.cc.Invoke(ctx, Auth_GetClientMeta_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DescribeIdentityProviders(ctx context.Context, in *DescribeIdentityProvidersRequest, opts ...grpc.CallOption) (*DescribeIdentityProvidersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeIdentityProvidersResponse)
	err := c.cc.Invoke(ctx, Auth_DescribeIdentityProviders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*DescribeServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeServiceResponse)
	err := c.cc.Invoke(ctx, Auth_DescribeService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddDenylist(ctx context.Context, in *AddDenylistRequest, opts ...grpc.CallOption) (*AddDenylistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDenylistResponse)
	err := c.cc.Invoke(ctx, Auth_AddDenylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RemoveDenylist(ctx context.Context, in *RemoveDenylistRequest, opts ...grpc.CallOption) (*RemoveDenylistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveDenylistResponse)
	err := c.cc.Invoke(ctx, Auth_RemoveDenylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) QueryDenylist(ctx context.Context, in *QueryDenylistRequest, opts ...grpc.CallOption) (*QueryDenylistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryDenylistResponse)
	err := c.cc.Invoke(ctx, Auth_QueryDenylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ListDenylist(ctx context.Context, in *ListDenylistRequest, opts ...grpc.CallOption) (*ListDenylistResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDenylistResponse)
	err := c.cc.Invoke(ctx, Auth_ListDenylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility.
//
// 所有 rpc 都遵循的通用 gRPC Status Code
// https://grpc.github.io/grpc/core/md_doc_statuscodes.html
//
//	INVALID_ARGUMENT | 缺少参数或者参数错误
//	INTERNAL | 服务内部错误
type AuthServer interface {
	// 聚合登录。
	// 验证 IdP 提供的 Credential，返回聚合用户标识 Combo ID 与 IdP 中的身份信息
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	// InvalidArgument | 缺少参数或者参数错误
	// UNAUTHENTICATED | IdP 提供的 Credential 验证不通过
	// PermissionDenied | IdP 提供的 Credential 验证通过，但是该用户被封禁
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	// 为已有的 Combo ID 关联更多 IdP 的身份。
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	//
	//	NOT_FOUND | 指定的 Combo ID 不存在
	//	UNAUTHENTICATED | IdP 提供的 Credential 验证不通过
	LinkIdentity(context.Context, *LinkIdentityRequest) (*LinkIdentityResponse, error)
	// 解除关联得 IdP。
	// 该功能需要配置开关进行开放。
	UnlinkIdentity(context.Context, *UnlinkIdentityRequest) (*UnlinkIdentityResponse, error)
	// 删除 Combo User，包括关联的 IdP 数据也会被同时删除。
	// 该功能需要配置开关进行开放。
	DeleteIdentities(context.Context, *DeleteIdentitiesRequest) (*DeleteIdentitiesResponse, error)
	// 对 Combo ID 关联的某个 Identity 进行实名认证（真实姓名 + 身份证号）
	AuthRealName(context.Context, *AuthRealNameRequest) (*AuthRealNameResponse, error)
	// Deprecated: Do not use.
	// 获取外部用户的聚合信息
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	//
	//	NOT_FOUND | 指定的外部用户不存在
	//
	// Deprecated: 请使用 GetIdentitiesByExternalId 代替
	GetIdentity(context.Context, *GetIdentityRequest) (*GetIdentityResponse, error)
	// 根据 外部身份标识（external_id）获取用户的聚合信息
	GetIdentitiesByExternalId(context.Context, *GetIdentitiesByExternalIdRequest) (*GetIdentitiesByExternalIdResponse, error)
	// Deprecated: Do not use.
	// 获取聚合用户所关联的所有身份，返回简略的身份信息。
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	//
	//	NOT_FOUND | 指定的 Combo ID 不存在
	//
	// Deprecated: 请使用 GetIdentitiesByComboId 代替
	GetIdentities(context.Context, *GetIdentitiesRequest) (*GetIdentitiesResponse, error)
	// 根据 combo_id 获取用户的所有身份，返回完整的身份信息。
	GetIdentitiesByComboId(context.Context, *GetIdentitiesByComboIdRequest) (*GetIdentitiesByComboIdResponse, error)
	// 获取聚合用户身份的客户端元数据。
	// 客户端元数据是在首次 SignIn 或 LinkIdentity 时写入的。
	// 调用方需要处理的，特定业务逻辑的 gRPC Status Code
	//
	//	NOT_FOUND | 指定的 Combo ID 不存在，或 Combo ID 未关联指定的 IdP
	GetClientMeta(context.Context, *GetClientMetaRequest) (*GetClientMetaResponse, error)
	// 获取所有当前支持的 Identity Provider 的元数据
	DescribeIdentityProviders(context.Context, *DescribeIdentityProvidersRequest) (*DescribeIdentityProvidersResponse, error)
	// 获取配置信息
	DescribeService(context.Context, *DescribeServiceRequest) (*DescribeServiceResponse, error)
	// 封禁数据管理接口
	AddDenylist(context.Context, *AddDenylistRequest) (*AddDenylistResponse, error)
	RemoveDenylist(context.Context, *RemoveDenylistRequest) (*RemoveDenylistResponse, error)
	QueryDenylist(context.Context, *QueryDenylistRequest) (*QueryDenylistResponse, error)
	ListDenylist(context.Context, *ListDenylistRequest) (*ListDenylistResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServer) LinkIdentity(context.Context, *LinkIdentityRequest) (*LinkIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkIdentity not implemented")
}
func (UnimplementedAuthServer) UnlinkIdentity(context.Context, *UnlinkIdentityRequest) (*UnlinkIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkIdentity not implemented")
}
func (UnimplementedAuthServer) DeleteIdentities(context.Context, *DeleteIdentitiesRequest) (*DeleteIdentitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIdentities not implemented")
}
func (UnimplementedAuthServer) AuthRealName(context.Context, *AuthRealNameRequest) (*AuthRealNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthRealName not implemented")
}
func (UnimplementedAuthServer) GetIdentity(context.Context, *GetIdentityRequest) (*GetIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentity not implemented")
}
func (UnimplementedAuthServer) GetIdentitiesByExternalId(context.Context, *GetIdentitiesByExternalIdRequest) (*GetIdentitiesByExternalIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentitiesByExternalId not implemented")
}
func (UnimplementedAuthServer) GetIdentities(context.Context, *GetIdentitiesRequest) (*GetIdentitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentities not implemented")
}
func (UnimplementedAuthServer) GetIdentitiesByComboId(context.Context, *GetIdentitiesByComboIdRequest) (*GetIdentitiesByComboIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentitiesByComboId not implemented")
}
func (UnimplementedAuthServer) GetClientMeta(context.Context, *GetClientMetaRequest) (*GetClientMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientMeta not implemented")
}
func (UnimplementedAuthServer) DescribeIdentityProviders(context.Context, *DescribeIdentityProvidersRequest) (*DescribeIdentityProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeIdentityProviders not implemented")
}
func (UnimplementedAuthServer) DescribeService(context.Context, *DescribeServiceRequest) (*DescribeServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}
func (UnimplementedAuthServer) AddDenylist(context.Context, *AddDenylistRequest) (*AddDenylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDenylist not implemented")
}
func (UnimplementedAuthServer) RemoveDenylist(context.Context, *RemoveDenylistRequest) (*RemoveDenylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDenylist not implemented")
}
func (UnimplementedAuthServer) QueryDenylist(context.Context, *QueryDenylistRequest) (*QueryDenylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDenylist not implemented")
}
func (UnimplementedAuthServer) ListDenylist(context.Context, *ListDenylistRequest) (*ListDenylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDenylist not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}
func (UnimplementedAuthServer) testEmbeddedByValue()              {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_LinkIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LinkIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_LinkIdentity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LinkIdentity(ctx, req.(*LinkIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UnlinkIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UnlinkIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_UnlinkIdentity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UnlinkIdentity(ctx, req.(*UnlinkIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteIdentities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteIdentities(ctx, req.(*DeleteIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AuthRealName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRealNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthRealName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AuthRealName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthRealName(ctx, req.(*AuthRealNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetIdentity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetIdentity(ctx, req.(*GetIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetIdentitiesByExternalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentitiesByExternalIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetIdentitiesByExternalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetIdentitiesByExternalId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetIdentitiesByExternalId(ctx, req.(*GetIdentitiesByExternalIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetIdentities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetIdentities(ctx, req.(*GetIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetIdentitiesByComboId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentitiesByComboIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetIdentitiesByComboId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetIdentitiesByComboId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetIdentitiesByComboId(ctx, req.(*GetIdentitiesByComboIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetClientMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetClientMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetClientMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetClientMeta(ctx, req.(*GetClientMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DescribeIdentityProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeIdentityProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DescribeIdentityProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DescribeIdentityProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DescribeIdentityProviders(ctx, req.(*DescribeIdentityProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DescribeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DescribeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DescribeService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DescribeService(ctx, req.(*DescribeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddDenylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDenylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddDenylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddDenylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddDenylist(ctx, req.(*AddDenylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RemoveDenylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDenylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RemoveDenylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RemoveDenylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RemoveDenylist(ctx, req.(*RemoveDenylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_QueryDenylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).QueryDenylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_QueryDenylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).QueryDenylist(ctx, req.(*QueryDenylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ListDenylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDenylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListDenylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ListDenylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListDenylist(ctx, req.(*ListDenylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _Auth_SignIn_Handler,
		},
		{
			MethodName: "LinkIdentity",
			Handler:    _Auth_LinkIdentity_Handler,
		},
		{
			MethodName: "UnlinkIdentity",
			Handler:    _Auth_UnlinkIdentity_Handler,
		},
		{
			MethodName: "DeleteIdentities",
			Handler:    _Auth_DeleteIdentities_Handler,
		},
		{
			MethodName: "AuthRealName",
			Handler:    _Auth_AuthRealName_Handler,
		},
		{
			MethodName: "GetIdentity",
			Handler:    _Auth_GetIdentity_Handler,
		},
		{
			MethodName: "GetIdentitiesByExternalId",
			Handler:    _Auth_GetIdentitiesByExternalId_Handler,
		},
		{
			MethodName: "GetIdentities",
			Handler:    _Auth_GetIdentities_Handler,
		},
		{
			MethodName: "GetIdentitiesByComboId",
			Handler:    _Auth_GetIdentitiesByComboId_Handler,
		},
		{
			MethodName: "GetClientMeta",
			Handler:    _Auth_GetClientMeta_Handler,
		},
		{
			MethodName: "DescribeIdentityProviders",
			Handler:    _Auth_DescribeIdentityProviders_Handler,
		},
		{
			MethodName: "DescribeService",
			Handler:    _Auth_DescribeService_Handler,
		},
		{
			MethodName: "AddDenylist",
			Handler:    _Auth_AddDenylist_Handler,
		},
		{
			MethodName: "RemoveDenylist",
			Handler:    _Auth_RemoveDenylist_Handler,
		},
		{
			MethodName: "QueryDenylist",
			Handler:    _Auth_QueryDenylist_Handler,
		},
		{
			MethodName: "ListDenylist",
			Handler:    _Auth_ListDenylist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
