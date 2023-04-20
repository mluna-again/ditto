// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: cookies/cookies.proto

package cookies

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

// CookieTravelerClient is the client API for CookieTraveler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CookieTravelerClient interface {
	SendCookie(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Cookie, error)
}

type cookieTravelerClient struct {
	cc grpc.ClientConnInterface
}

func NewCookieTravelerClient(cc grpc.ClientConnInterface) CookieTravelerClient {
	return &cookieTravelerClient{cc}
}

func (c *cookieTravelerClient) SendCookie(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Cookie, error) {
	out := new(Cookie)
	err := c.cc.Invoke(ctx, "/cookies.CookieTraveler/SendCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CookieTravelerServer is the server API for CookieTraveler service.
// All implementations must embed UnimplementedCookieTravelerServer
// for forward compatibility
type CookieTravelerServer interface {
	SendCookie(context.Context, *Cookie) (*Cookie, error)
	mustEmbedUnimplementedCookieTravelerServer()
}

// UnimplementedCookieTravelerServer must be embedded to have forward compatible implementations.
type UnimplementedCookieTravelerServer struct {
}

func (UnimplementedCookieTravelerServer) SendCookie(context.Context, *Cookie) (*Cookie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCookie not implemented")
}
func (UnimplementedCookieTravelerServer) mustEmbedUnimplementedCookieTravelerServer() {}

// UnsafeCookieTravelerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CookieTravelerServer will
// result in compilation errors.
type UnsafeCookieTravelerServer interface {
	mustEmbedUnimplementedCookieTravelerServer()
}

func RegisterCookieTravelerServer(s grpc.ServiceRegistrar, srv CookieTravelerServer) {
	s.RegisterService(&CookieTraveler_ServiceDesc, srv)
}

func _CookieTraveler_SendCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CookieTravelerServer).SendCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cookies.CookieTraveler/SendCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CookieTravelerServer).SendCookie(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

// CookieTraveler_ServiceDesc is the grpc.ServiceDesc for CookieTraveler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CookieTraveler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cookies.CookieTraveler",
	HandlerType: (*CookieTravelerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCookie",
			Handler:    _CookieTraveler_SendCookie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cookies/cookies.proto",
}