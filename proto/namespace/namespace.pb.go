// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/wangrenyi/kmpgo/proto/namespace/namespace.proto

/*
Package namespace is a generated protocol buffer package.

It is generated from these files:
	github.com/wangrenyi/kmpgo/proto/namespace/namespace.proto

It has these top-level messages:
*/
package namespace

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

// Client API for NamespaceInterface service

type NamespaceInterfaceClient interface {
	NamespaceCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NamespaceUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NamespaceUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NamespaceDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NamespaceGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NamespaceList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NamespaceWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NamespacePatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	// Namespace expansion
	NamespaceFinalize(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
}

type namespaceInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceInterfaceClient(cc *grpc.ClientConn) NamespaceInterfaceClient {
	return &namespaceInterfaceClient{cc}
}

func (c *namespaceInterfaceClient) NamespaceCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespaceCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceInterfaceClient) NamespaceUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespaceUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceInterfaceClient) NamespaceUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespaceUpdateStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceInterfaceClient) NamespaceDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespaceDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceInterfaceClient) NamespaceGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespaceGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceInterfaceClient) NamespaceList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespaceList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceInterfaceClient) NamespaceWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespaceWatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceInterfaceClient) NamespacePatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespacePatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceInterfaceClient) NamespaceFinalize(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/namespace.NamespaceInterface/NamespaceFinalize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NamespaceInterface service

type NamespaceInterfaceServer interface {
	NamespaceCreate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NamespaceUpdate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NamespaceUpdateStatus(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NamespaceDelete(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NamespaceGet(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NamespaceList(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NamespaceWatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NamespacePatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	// Namespace expansion
	NamespaceFinalize(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
}

func RegisterNamespaceInterfaceServer(s *grpc.Server, srv NamespaceInterfaceServer) {
	s.RegisterService(&_NamespaceInterface_serviceDesc, srv)
}

func _NamespaceInterface_NamespaceCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespaceCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespaceCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespaceCreate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceInterface_NamespaceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespaceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespaceUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespaceUpdate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceInterface_NamespaceUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespaceUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespaceUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespaceUpdateStatus(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceInterface_NamespaceDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespaceDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespaceDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespaceDelete(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceInterface_NamespaceGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespaceGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespaceGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespaceGet(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceInterface_NamespaceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespaceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespaceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespaceList(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceInterface_NamespaceWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespaceWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespaceWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespaceWatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceInterface_NamespacePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespacePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespacePatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespacePatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceInterface_NamespaceFinalize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceInterfaceServer).NamespaceFinalize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespace.NamespaceInterface/NamespaceFinalize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceInterfaceServer).NamespaceFinalize(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NamespaceInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "namespace.NamespaceInterface",
	HandlerType: (*NamespaceInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NamespaceCreate",
			Handler:    _NamespaceInterface_NamespaceCreate_Handler,
		},
		{
			MethodName: "NamespaceUpdate",
			Handler:    _NamespaceInterface_NamespaceUpdate_Handler,
		},
		{
			MethodName: "NamespaceUpdateStatus",
			Handler:    _NamespaceInterface_NamespaceUpdateStatus_Handler,
		},
		{
			MethodName: "NamespaceDelete",
			Handler:    _NamespaceInterface_NamespaceDelete_Handler,
		},
		{
			MethodName: "NamespaceGet",
			Handler:    _NamespaceInterface_NamespaceGet_Handler,
		},
		{
			MethodName: "NamespaceList",
			Handler:    _NamespaceInterface_NamespaceList_Handler,
		},
		{
			MethodName: "NamespaceWatch",
			Handler:    _NamespaceInterface_NamespaceWatch_Handler,
		},
		{
			MethodName: "NamespacePatch",
			Handler:    _NamespaceInterface_NamespacePatch_Handler,
		},
		{
			MethodName: "NamespaceFinalize",
			Handler:    _NamespaceInterface_NamespaceFinalize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/wangrenyi/kmpgo/proto/namespace/namespace.proto",
}

func init() {
	proto.RegisterFile("github.com/wangrenyi/kmpgo/proto/namespace/namespace.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd3, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0x05, 0x70, 0x2f, 0x0a, 0x06, 0xff, 0xe0, 0x82, 0x97, 0xa2, 0x17, 0x3f, 0xc0, 0x2e, 0xa8,
	0x27, 0xaf, 0x8a, 0x52, 0xd0, 0x5a, 0x5a, 0xd4, 0xf3, 0x34, 0x3e, 0xb7, 0xa1, 0xcd, 0x4c, 0x4c,
	0x66, 0x2b, 0xfa, 0xc5, 0xbd, 0x0a, 0xb5, 0x46, 0x14, 0xc4, 0xc3, 0xee, 0x29, 0x61, 0x1e, 0xf9,
	0xf1, 0x20, 0x8c, 0x39, 0xab, 0x9d, 0x4e, 0x9b, 0x49, 0x69, 0xc5, 0x57, 0x2f, 0xc4, 0x75, 0x04,
	0xbf, 0xba, 0x6a, 0xe6, 0x43, 0x2d, 0x55, 0x88, 0xa2, 0x52, 0x31, 0x79, 0xa4, 0x40, 0x16, 0xdf,
	0xb7, 0x72, 0x99, 0x14, 0x9b, 0x79, 0xd0, 0x3b, 0xfd, 0x97, 0xb1, 0xe2, 0xbd, 0xf0, 0xea, 0xf8,
	0x04, 0x8e, 0xdf, 0xd7, 0x4d, 0x31, 0xf8, 0x32, 0xfa, 0xac, 0x88, 0x4f, 0x64, 0x51, 0x0c, 0xcd,
	0x6e, 0x9e, 0x9e, 0x47, 0x90, 0xa2, 0x38, 0x28, 0x57, 0x0f, 0xfb, 0xbc, 0x90, 0x19, 0xc6, 0x88,
	0x0b, 0x67, 0x31, 0xc2, 0x73, 0x83, 0xa4, 0xbd, 0xc3, 0x3f, 0xd2, 0x14, 0x84, 0x13, 0x8e, 0xd6,
	0x7e, 0x88, 0x77, 0xe1, 0xb1, 0x03, 0xf1, 0xde, 0xec, 0xff, 0x12, 0xc7, 0x4a, 0xda, 0xa4, 0x2e,
	0x9b, 0x5e, 0x60, 0x8e, 0xf6, 0x4d, 0x6f, 0xcc, 0x56, 0x16, 0xaf, 0xa0, 0x6d, 0xb9, 0x81, 0xd9,
	0xce, 0xdc, 0xb5, 0x4b, 0xad, 0xbd, 0x5b, 0xb3, 0x93, 0xbd, 0x07, 0x52, 0x3b, 0xed, 0x12, 0x1c,
	0x76, 0x01, 0x8e, 0xcc, 0x5e, 0x06, 0x2f, 0x1d, 0xd3, 0xdc, 0xbd, 0xb5, 0xfd, 0x94, 0xc9, 0xc6,
	0x72, 0x01, 0x4e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x53, 0xa7, 0x32, 0x3c, 0x7f, 0x03, 0x00,
	0x00,
}