// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package routeguide

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

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {
	SaveUsers(ctx context.Context, in *Users, opts ...grpc.CallOption) (*SaveReply, error)
	SavePosts(ctx context.Context, in *Posts, opts ...grpc.CallOption) (*SaveReply, error)
	FollowUsers(ctx context.Context, in *Search, opts ...grpc.CallOption) (*FollowReply, error)
	UnfollowUsers(ctx context.Context, in *Search, opts ...grpc.CallOption) (*FollowReply, error)
	LoadUsers(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Users, error)
	LoadPosts(ctx context.Context, in *Followings, opts ...grpc.CallOption) (*FollowingsPosts, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) SaveUsers(ctx context.Context, in *Users, opts ...grpc.CallOption) (*SaveReply, error) {
	out := new(SaveReply)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/SaveUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) SavePosts(ctx context.Context, in *Posts, opts ...grpc.CallOption) (*SaveReply, error) {
	out := new(SaveReply)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/SavePosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) FollowUsers(ctx context.Context, in *Search, opts ...grpc.CallOption) (*FollowReply, error) {
	out := new(FollowReply)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/FollowUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) UnfollowUsers(ctx context.Context, in *Search, opts ...grpc.CallOption) (*FollowReply, error) {
	out := new(FollowReply)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/UnfollowUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) LoadUsers(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/LoadUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) LoadPosts(ctx context.Context, in *Followings, opts ...grpc.CallOption) (*FollowingsPosts, error) {
	out := new(FollowingsPosts)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/LoadPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteGuideServer is the server API for RouteGuide service.
// All implementations must embed UnimplementedRouteGuideServer
// for forward compatibility
type RouteGuideServer interface {
	SaveUsers(context.Context, *Users) (*SaveReply, error)
	SavePosts(context.Context, *Posts) (*SaveReply, error)
	FollowUsers(context.Context, *Search) (*FollowReply, error)
	UnfollowUsers(context.Context, *Search) (*FollowReply, error)
	LoadUsers(context.Context, *Search) (*Users, error)
	LoadPosts(context.Context, *Followings) (*FollowingsPosts, error)
	mustEmbedUnimplementedRouteGuideServer()
}

// UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (UnimplementedRouteGuideServer) SaveUsers(context.Context, *Users) (*SaveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUsers not implemented")
}
func (UnimplementedRouteGuideServer) SavePosts(context.Context, *Posts) (*SaveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePosts not implemented")
}
func (UnimplementedRouteGuideServer) FollowUsers(context.Context, *Search) (*FollowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUsers not implemented")
}
func (UnimplementedRouteGuideServer) UnfollowUsers(context.Context, *Search) (*FollowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUsers not implemented")
}
func (UnimplementedRouteGuideServer) LoadUsers(context.Context, *Search) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadUsers not implemented")
}
func (UnimplementedRouteGuideServer) LoadPosts(context.Context, *Followings) (*FollowingsPosts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadPosts not implemented")
}
func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {}

// UnsafeRouteGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServer will
// result in compilation errors.
type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	s.RegisterService(&RouteGuide_ServiceDesc, srv)
}

func _RouteGuide_SaveUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Users)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).SaveUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/SaveUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).SaveUsers(ctx, req.(*Users))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_SavePosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Posts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).SavePosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/SavePosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).SavePosts(ctx, req.(*Posts))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_FollowUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).FollowUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/FollowUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).FollowUsers(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_UnfollowUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).UnfollowUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/UnfollowUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).UnfollowUsers(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_LoadUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).LoadUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/LoadUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).LoadUsers(ctx, req.(*Search))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_LoadPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Followings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).LoadPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/LoadPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).LoadPosts(ctx, req.(*Followings))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteGuide_ServiceDesc is the grpc.ServiceDesc for RouteGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUsers",
			Handler:    _RouteGuide_SaveUsers_Handler,
		},
		{
			MethodName: "SavePosts",
			Handler:    _RouteGuide_SavePosts_Handler,
		},
		{
			MethodName: "FollowUsers",
			Handler:    _RouteGuide_FollowUsers_Handler,
		},
		{
			MethodName: "UnfollowUsers",
			Handler:    _RouteGuide_UnfollowUsers_Handler,
		},
		{
			MethodName: "LoadUsers",
			Handler:    _RouteGuide_LoadUsers_Handler,
		},
		{
			MethodName: "LoadPosts",
			Handler:    _RouteGuide_LoadPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "models.proto",
}
