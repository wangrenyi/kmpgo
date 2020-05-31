// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/wangrenyi/kmpgo/proto/endpoint/endpoint.proto

/*
Package endpoint is a generated protocol buffer package.

It is generated from these files:
	github.com/wangrenyi/kmpgo/proto/endpoint/endpoint.proto

It has these top-level messages:
*/
package endpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/wangrenyi/kmpgo/proto/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EndpointInterface service

type EndpointInterfaceClient interface {
	EndpointCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	EndpointUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	EndpointDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	EndpointDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	EndpointGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	EndpointList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	EndpointWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	EndpointPatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
}

type endpointInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointInterfaceClient(cc *grpc.ClientConn) EndpointInterfaceClient {
	return &endpointInterfaceClient{cc}
}

func (c *endpointInterfaceClient) EndpointCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/endpoint.EndpointInterface/EndpointCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointInterfaceClient) EndpointUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/endpoint.EndpointInterface/EndpointUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointInterfaceClient) EndpointDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/endpoint.EndpointInterface/EndpointDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointInterfaceClient) EndpointDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/endpoint.EndpointInterface/EndpointDeleteCollection", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointInterfaceClient) EndpointGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/endpoint.EndpointInterface/EndpointGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointInterfaceClient) EndpointList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/endpoint.EndpointInterface/EndpointList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointInterfaceClient) EndpointWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/endpoint.EndpointInterface/EndpointWatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointInterfaceClient) EndpointPatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/endpoint.EndpointInterface/EndpointPatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EndpointInterface service

type EndpointInterfaceServer interface {
	EndpointCreate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	EndpointUpdate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	EndpointDelete(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	EndpointDeleteCollection(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	EndpointGet(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	EndpointList(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	EndpointWatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	EndpointPatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
}

func RegisterEndpointInterfaceServer(s *grpc.Server, srv EndpointInterfaceServer) {
	s.RegisterService(&_EndpointInterface_serviceDesc, srv)
}

func _EndpointInterface_EndpointCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointInterfaceServer).EndpointCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.EndpointInterface/EndpointCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointInterfaceServer).EndpointCreate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointInterface_EndpointUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointInterfaceServer).EndpointUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.EndpointInterface/EndpointUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointInterfaceServer).EndpointUpdate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointInterface_EndpointDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointInterfaceServer).EndpointDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.EndpointInterface/EndpointDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointInterfaceServer).EndpointDelete(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointInterface_EndpointDeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointInterfaceServer).EndpointDeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.EndpointInterface/EndpointDeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointInterfaceServer).EndpointDeleteCollection(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointInterface_EndpointGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointInterfaceServer).EndpointGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.EndpointInterface/EndpointGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointInterfaceServer).EndpointGet(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointInterface_EndpointList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointInterfaceServer).EndpointList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.EndpointInterface/EndpointList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointInterfaceServer).EndpointList(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointInterface_EndpointWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointInterfaceServer).EndpointWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.EndpointInterface/EndpointWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointInterfaceServer).EndpointWatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndpointInterface_EndpointPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointInterfaceServer).EndpointPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoint.EndpointInterface/EndpointPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointInterfaceServer).EndpointPatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "endpoint.EndpointInterface",
	HandlerType: (*EndpointInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EndpointCreate",
			Handler:    _EndpointInterface_EndpointCreate_Handler,
		},
		{
			MethodName: "EndpointUpdate",
			Handler:    _EndpointInterface_EndpointUpdate_Handler,
		},
		{
			MethodName: "EndpointDelete",
			Handler:    _EndpointInterface_EndpointDelete_Handler,
		},
		{
			MethodName: "EndpointDeleteCollection",
			Handler:    _EndpointInterface_EndpointDeleteCollection_Handler,
		},
		{
			MethodName: "EndpointGet",
			Handler:    _EndpointInterface_EndpointGet_Handler,
		},
		{
			MethodName: "EndpointList",
			Handler:    _EndpointInterface_EndpointList_Handler,
		},
		{
			MethodName: "EndpointWatch",
			Handler:    _EndpointInterface_EndpointWatch_Handler,
		},
		{
			MethodName: "EndpointPatch",
			Handler:    _EndpointInterface_EndpointPatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/wangrenyi/kmpgo/proto/endpoint/endpoint.proto",
}

func init() {
	proto.RegisterFile("github.com/wangrenyi/kmpgo/proto/endpoint/endpoint.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd2, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x70, 0x2f, 0x8a, 0xc4, 0x3f, 0xe0, 0x9e, 0xa4, 0xe8, 0xc5, 0x07, 0xd8, 0x80, 0x7a,
	0xf0, 0x5e, 0x45, 0x0a, 0xf5, 0x0f, 0x8a, 0x88, 0xc7, 0x34, 0xfd, 0xdc, 0x86, 0x6e, 0x66, 0x62,
	0x32, 0xad, 0xf8, 0x32, 0x3e, 0xab, 0x50, 0x37, 0xa2, 0x07, 0xe9, 0x21, 0x7b, 0x4a, 0xe6, 0x1b,
	0xbe, 0xdf, 0x69, 0xd4, 0x45, 0xe3, 0x64, 0xb6, 0x98, 0xd4, 0x96, 0xbd, 0x7e, 0x37, 0xd4, 0x44,
	0xd0, 0x87, 0xd3, 0x73, 0x1f, 0x1a, 0xd6, 0x21, 0xb2, 0xb0, 0x06, 0x4d, 0x03, 0x3b, 0x92, 0x9f,
	0x4f, 0xbd, 0xca, 0xab, 0xed, 0x3c, 0x0f, 0xce, 0xd7, 0x1a, 0x96, 0xbd, 0x67, 0xea, 0x9e, 0xef,
	0xfe, 0xe9, 0xe7, 0xa6, 0x3a, 0xb8, 0xea, 0x88, 0x11, 0x09, 0xe2, 0xab, 0xb1, 0xa8, 0xee, 0xd4,
	0x7e, 0x0e, 0x87, 0x11, 0x46, 0x50, 0x1d, 0xd5, 0x5d, 0x6d, 0x44, 0x4b, 0x9e, 0xe3, 0x11, 0x71,
	0xe9, 0x2c, 0x1e, 0xf0, 0xb6, 0x40, 0x92, 0xc1, 0xf1, 0x3f, 0xdb, 0x14, 0x98, 0x12, 0x4e, 0x36,
	0x7e, 0x83, 0x4f, 0x61, 0xda, 0x2f, 0x78, 0x89, 0x16, 0xe5, 0xe0, 0x8b, 0x3a, 0xfc, 0x0b, 0x0e,
	0xb9, 0x6d, 0x61, 0xc5, 0x31, 0x95, 0xd2, 0x63, 0xb5, 0x93, 0xe9, 0x6b, 0x48, 0xa9, 0x76, 0xa3,
	0x76, 0xb3, 0x36, 0x76, 0xa9, 0x98, 0xbb, 0x55, 0x7b, 0x99, 0x7b, 0x36, 0x62, 0x67, 0x3d, 0x7a,
	0xf7, 0x3d, 0x78, 0x93, 0xad, 0xd5, 0x9d, 0x9e, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xbb,
	0x21, 0xde, 0x23, 0x03, 0x00, 0x00,
}