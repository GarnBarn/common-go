// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/tag.proto

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

// TagClient is the client API for Tag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagClient interface {
	IsTagExists(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*TagExistsResponse, error)
	GetTag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*TagPublic, error)
}

type tagClient struct {
	cc grpc.ClientConnInterface
}

func NewTagClient(cc grpc.ClientConnInterface) TagClient {
	return &tagClient{cc}
}

func (c *tagClient) IsTagExists(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*TagExistsResponse, error) {
	out := new(TagExistsResponse)
	err := c.cc.Invoke(ctx, "/tag.Tag/IsTagExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagClient) GetTag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*TagPublic, error) {
	out := new(TagPublic)
	err := c.cc.Invoke(ctx, "/tag.Tag/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServer is the server API for Tag service.
// All implementations must embed UnimplementedTagServer
// for forward compatibility
type TagServer interface {
	IsTagExists(context.Context, *TagRequest) (*TagExistsResponse, error)
	GetTag(context.Context, *TagRequest) (*TagPublic, error)
	mustEmbedUnimplementedTagServer()
}

// UnimplementedTagServer must be embedded to have forward compatible implementations.
type UnimplementedTagServer struct {
}

func (UnimplementedTagServer) IsTagExists(context.Context, *TagRequest) (*TagExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsTagExists not implemented")
}
func (UnimplementedTagServer) GetTag(context.Context, *TagRequest) (*TagPublic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedTagServer) mustEmbedUnimplementedTagServer() {}

// UnsafeTagServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServer will
// result in compilation errors.
type UnsafeTagServer interface {
	mustEmbedUnimplementedTagServer()
}

func RegisterTagServer(s grpc.ServiceRegistrar, srv TagServer) {
	s.RegisterService(&Tag_ServiceDesc, srv)
}

func _Tag_IsTagExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).IsTagExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tag.Tag/IsTagExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).IsTagExists(ctx, req.(*TagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tag_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tag.Tag/GetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).GetTag(ctx, req.(*TagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tag_ServiceDesc is the grpc.ServiceDesc for Tag service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tag_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tag.Tag",
	HandlerType: (*TagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsTagExists",
			Handler:    _Tag_IsTagExists_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _Tag_GetTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tag.proto",
}