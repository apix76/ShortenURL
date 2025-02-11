// Code generated by protoc-gen-go-grpcServer. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpcServer v1.5.1
// - protoc             v3.12.4
// source: Proto/Request.proto

package Proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcServer package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ShortenURL_GetShortenURL_FullMethodName = "/ShortenURL/GetShortenURL"
	ShortenURL_GetAllURL_FullMethodName     = "/ShortenURL/GetAllURL"
)

// ShortenURLClient is the client API for ShortenURL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenURLClient interface {
	GetShortenURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*ShortURL, error)
	GetAllURL(ctx context.Context, in *ShortURL, opts ...grpc.CallOption) (*URL, error)
}

type shortenURLClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenURLClient(cc grpc.ClientConnInterface) ShortenURLClient {
	return &shortenURLClient{cc}
}

func (c *shortenURLClient) GetShortenURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*ShortURL, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShortURL)
	err := c.cc.Invoke(ctx, ShortenURL_GetShortenURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenURLClient) GetAllURL(ctx context.Context, in *ShortURL, opts ...grpc.CallOption) (*URL, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(URL)
	err := c.cc.Invoke(ctx, ShortenURL_GetAllURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenURLServer is the server API for ShortenURL service.
// All implementations must embed UnimplementedShortenURLServer
// for forward compatibility.
type ShortenURLServer interface {
	GetShortenURL(context.Context, *URL) (*ShortURL, error)
	GetAllURL(context.Context, *ShortURL) (*URL, error)
	mustEmbedUnimplementedShortenURLServer()
}

// UnimplementedShortenURLServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShortenURLServer struct{}

func (UnimplementedShortenURLServer) GetShortenURL(context.Context, *URL) (*ShortURL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortenURL not implemented")
}
func (UnimplementedShortenURLServer) GetAllURL(context.Context, *ShortURL) (*URL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllURL not implemented")
}
func (UnimplementedShortenURLServer) mustEmbedUnimplementedShortenURLServer() {}
func (UnimplementedShortenURLServer) testEmbeddedByValue()                    {}

// UnsafeShortenURLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenURLServer will
// result in compilation errors.
type UnsafeShortenURLServer interface {
	mustEmbedUnimplementedShortenURLServer()
}

func RegisterShortenURLServer(s grpc.ServiceRegistrar, srv ShortenURLServer) {
	// If the following call pancis, it indicates UnimplementedShortenURLServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShortenURL_ServiceDesc, srv)
}

func _ShortenURL_GetShortenURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenURLServer).GetShortenURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenURL_GetShortenURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenURLServer).GetShortenURL(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenURL_GetAllURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortURL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenURLServer).GetAllURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenURL_GetAllURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenURLServer).GetAllURL(ctx, req.(*ShortURL))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortenURL_ServiceDesc is the grpc.ServiceDesc for ShortenURL service.
// It's only intended for direct use with grpcServer.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortenURL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ShortenURL",
	HandlerType: (*ShortenURLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShortenURL",
			Handler:    _ShortenURL_GetShortenURL_Handler,
		},
		{
			MethodName: "GetAllURL",
			Handler:    _ShortenURL_GetAllURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/Request.proto",
}
