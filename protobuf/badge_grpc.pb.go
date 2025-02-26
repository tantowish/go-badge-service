// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// BadgeServiceClient is the client API for BadgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BadgeServiceClient interface {
	GetBadge(ctx context.Context, in *GetBadgeRequest, opts ...grpc.CallOption) (*Badge, error)
	CreateBadge(ctx context.Context, in *CreateBadgeRequest, opts ...grpc.CallOption) (*Badge, error)
	UpdateBadge(ctx context.Context, in *UpdateBadgeRequest, opts ...grpc.CallOption) (*Badge, error)
	DeleteBadge(ctx context.Context, in *DeleteBadgeRequest, opts ...grpc.CallOption) (*DeleteBadgeResponse, error)
	GetShopBadge(ctx context.Context, in *GetShopBadgeRequest, opts ...grpc.CallOption) (*ShopBadge, error)
	InvokeNsq(ctx context.Context, in *GetShopBadgeRequest, opts ...grpc.CallOption) (*InvokeNsqResponse, error)
}

type badgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBadgeServiceClient(cc grpc.ClientConnInterface) BadgeServiceClient {
	return &badgeServiceClient{cc}
}

func (c *badgeServiceClient) GetBadge(ctx context.Context, in *GetBadgeRequest, opts ...grpc.CallOption) (*Badge, error) {
	out := new(Badge)
	err := c.cc.Invoke(ctx, "/badge.BadgeService/GetBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) CreateBadge(ctx context.Context, in *CreateBadgeRequest, opts ...grpc.CallOption) (*Badge, error) {
	out := new(Badge)
	err := c.cc.Invoke(ctx, "/badge.BadgeService/CreateBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) UpdateBadge(ctx context.Context, in *UpdateBadgeRequest, opts ...grpc.CallOption) (*Badge, error) {
	out := new(Badge)
	err := c.cc.Invoke(ctx, "/badge.BadgeService/UpdateBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) DeleteBadge(ctx context.Context, in *DeleteBadgeRequest, opts ...grpc.CallOption) (*DeleteBadgeResponse, error) {
	out := new(DeleteBadgeResponse)
	err := c.cc.Invoke(ctx, "/badge.BadgeService/DeleteBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) GetShopBadge(ctx context.Context, in *GetShopBadgeRequest, opts ...grpc.CallOption) (*ShopBadge, error) {
	out := new(ShopBadge)
	err := c.cc.Invoke(ctx, "/badge.BadgeService/GetShopBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgeServiceClient) InvokeNsq(ctx context.Context, in *GetShopBadgeRequest, opts ...grpc.CallOption) (*InvokeNsqResponse, error) {
	out := new(InvokeNsqResponse)
	err := c.cc.Invoke(ctx, "/badge.BadgeService/InvokeNsq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BadgeServiceServer is the server API for BadgeService service.
// All implementations must embed UnimplementedBadgeServiceServer
// for forward compatibility
type BadgeServiceServer interface {
	GetBadge(context.Context, *GetBadgeRequest) (*Badge, error)
	CreateBadge(context.Context, *CreateBadgeRequest) (*Badge, error)
	UpdateBadge(context.Context, *UpdateBadgeRequest) (*Badge, error)
	DeleteBadge(context.Context, *DeleteBadgeRequest) (*DeleteBadgeResponse, error)
	GetShopBadge(context.Context, *GetShopBadgeRequest) (*ShopBadge, error)
	InvokeNsq(context.Context, *GetShopBadgeRequest) (*InvokeNsqResponse, error)
	mustEmbedUnimplementedBadgeServiceServer()
}

// UnimplementedBadgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBadgeServiceServer struct {
}

func (UnimplementedBadgeServiceServer) GetBadge(context.Context, *GetBadgeRequest) (*Badge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBadge not implemented")
}
func (UnimplementedBadgeServiceServer) CreateBadge(context.Context, *CreateBadgeRequest) (*Badge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBadge not implemented")
}
func (UnimplementedBadgeServiceServer) UpdateBadge(context.Context, *UpdateBadgeRequest) (*Badge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBadge not implemented")
}
func (UnimplementedBadgeServiceServer) DeleteBadge(context.Context, *DeleteBadgeRequest) (*DeleteBadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBadge not implemented")
}
func (UnimplementedBadgeServiceServer) GetShopBadge(context.Context, *GetShopBadgeRequest) (*ShopBadge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopBadge not implemented")
}
func (UnimplementedBadgeServiceServer) InvokeNsq(context.Context, *GetShopBadgeRequest) (*InvokeNsqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeNsq not implemented")
}
func (UnimplementedBadgeServiceServer) mustEmbedUnimplementedBadgeServiceServer() {}

// UnsafeBadgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BadgeServiceServer will
// result in compilation errors.
type UnsafeBadgeServiceServer interface {
	mustEmbedUnimplementedBadgeServiceServer()
}

func RegisterBadgeServiceServer(s grpc.ServiceRegistrar, srv BadgeServiceServer) {
	s.RegisterService(&BadgeService_ServiceDesc, srv)
}

func _BadgeService_GetBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).GetBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badge.BadgeService/GetBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).GetBadge(ctx, req.(*GetBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_CreateBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).CreateBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badge.BadgeService/CreateBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).CreateBadge(ctx, req.(*CreateBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_UpdateBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).UpdateBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badge.BadgeService/UpdateBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).UpdateBadge(ctx, req.(*UpdateBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_DeleteBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).DeleteBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badge.BadgeService/DeleteBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).DeleteBadge(ctx, req.(*DeleteBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_GetShopBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).GetShopBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badge.BadgeService/GetShopBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).GetShopBadge(ctx, req.(*GetShopBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BadgeService_InvokeNsq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopBadgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgeServiceServer).InvokeNsq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badge.BadgeService/InvokeNsq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgeServiceServer).InvokeNsq(ctx, req.(*GetShopBadgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BadgeService_ServiceDesc is the grpc.ServiceDesc for BadgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BadgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "badge.BadgeService",
	HandlerType: (*BadgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBadge",
			Handler:    _BadgeService_GetBadge_Handler,
		},
		{
			MethodName: "CreateBadge",
			Handler:    _BadgeService_CreateBadge_Handler,
		},
		{
			MethodName: "UpdateBadge",
			Handler:    _BadgeService_UpdateBadge_Handler,
		},
		{
			MethodName: "DeleteBadge",
			Handler:    _BadgeService_DeleteBadge_Handler,
		},
		{
			MethodName: "GetShopBadge",
			Handler:    _BadgeService_GetShopBadge_Handler,
		},
		{
			MethodName: "InvokeNsq",
			Handler:    _BadgeService_InvokeNsq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "badge.proto",
}
