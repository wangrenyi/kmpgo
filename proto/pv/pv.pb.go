// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/wangrenyi/kmpgo/proto/pv/pv.proto

/*
Package pv is a generated protocol buffer package.

It is generated from these files:
	github.com/wangrenyi/kmpgo/proto/pv/pv.proto

It has these top-level messages:
*/
package pv

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

// Client API for PVInterface service

type PVInterfaceClient interface {
	PVCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVPatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
}

type pVInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewPVInterfaceClient(cc *grpc.ClientConn) PVInterfaceClient {
	return &pVInterfaceClient{cc}
}

func (c *pVInterfaceClient) PVCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVInterfaceClient) PVUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVInterfaceClient) PVUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVUpdateStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVInterfaceClient) PVDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVInterfaceClient) PVDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVDeleteCollection", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVInterfaceClient) PVGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVInterfaceClient) PVList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVInterfaceClient) PVWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVWatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVInterfaceClient) PVPatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pv.PVInterface/PVPatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PVInterface service

type PVInterfaceServer interface {
	PVCreate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVUpdate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVUpdateStatus(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVDelete(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVDeleteCollection(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVGet(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVList(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVWatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVPatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
}

func RegisterPVInterfaceServer(s *grpc.Server, srv PVInterfaceServer) {
	s.RegisterService(&_PVInterface_serviceDesc, srv)
}

func _PVInterface_PVCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVCreate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVInterface_PVUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVUpdate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVInterface_PVUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVUpdateStatus(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVInterface_PVDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVDelete(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVInterface_PVDeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVDeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVDeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVDeleteCollection(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVInterface_PVGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVGet(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVInterface_PVList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVList(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVInterface_PVWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVWatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVInterface_PVPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVInterfaceServer).PVPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PVInterface/PVPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVInterfaceServer).PVPatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PVInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.PVInterface",
	HandlerType: (*PVInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PVCreate",
			Handler:    _PVInterface_PVCreate_Handler,
		},
		{
			MethodName: "PVUpdate",
			Handler:    _PVInterface_PVUpdate_Handler,
		},
		{
			MethodName: "PVUpdateStatus",
			Handler:    _PVInterface_PVUpdateStatus_Handler,
		},
		{
			MethodName: "PVDelete",
			Handler:    _PVInterface_PVDelete_Handler,
		},
		{
			MethodName: "PVDeleteCollection",
			Handler:    _PVInterface_PVDeleteCollection_Handler,
		},
		{
			MethodName: "PVGet",
			Handler:    _PVInterface_PVGet_Handler,
		},
		{
			MethodName: "PVList",
			Handler:    _PVInterface_PVList_Handler,
		},
		{
			MethodName: "PVWatch",
			Handler:    _PVInterface_PVWatch_Handler,
		},
		{
			MethodName: "PVPatch",
			Handler:    _PVInterface_PVPatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/wangrenyi/kmpgo/proto/pv/pv.proto",
}

func init() { proto.RegisterFile("github.com/wangrenyi/kmpgo/proto/pv/pv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd2, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0x11, 0x5a, 0x65, 0x04, 0x0f, 0x7b, 0x2c, 0x7a, 0xf1, 0x2c, 0x1b, 0x50, 0xdf,
	0xa0, 0x62, 0x5d, 0x10, 0x0c, 0x2e, 0xc6, 0x73, 0x1a, 0xc7, 0x6d, 0xe8, 0x6e, 0x26, 0x26, 0xb3,
	0x11, 0x1f, 0xcd, 0xb7, 0x13, 0xda, 0xee, 0x51, 0x3c, 0x24, 0xa7, 0x04, 0x66, 0xfe, 0x8f, 0xff,
	0x30, 0x70, 0xdd, 0x59, 0xde, 0x8c, 0xeb, 0xda, 0xd0, 0x20, 0xbe, 0xb4, 0xeb, 0x02, 0xba, 0x6f,
	0x2b, 0xb6, 0x83, 0xef, 0x48, 0xf8, 0x40, 0x4c, 0xc2, 0x27, 0xe1, 0x53, 0xbd, 0xfb, 0x57, 0xc7,
	0x3e, 0x2d, 0xee, 0xfe, 0x4d, 0x18, 0x1a, 0x06, 0x72, 0x87, 0x67, 0x9f, 0xbc, 0xf9, 0x99, 0xc1,
	0x99, 0x54, 0x8d, 0x63, 0x0c, 0x1f, 0xda, 0x60, 0xd5, 0xc0, 0xa9, 0x54, 0xcb, 0x80, 0x9a, 0xb1,
	0xba, 0xa8, 0x0f, 0xab, 0x8d, 0x4b, 0xb4, 0xc5, 0x16, 0x43, 0xb2, 0x06, 0x5f, 0xf0, 0x73, 0xc4,
	0xc8, 0x8b, 0xcb, 0x3f, 0xa6, 0xd1, 0x93, 0x8b, 0x78, 0x75, 0xb4, 0xa7, 0x5e, 0xfd, 0x7b, 0x01,
	0xea, 0x19, 0xce, 0x27, 0xaa, 0x65, 0xcd, 0x63, 0x2c, 0xd2, 0xed, 0x1e, 0x7b, 0xcc, 0xef, 0xd6,
	0x42, 0x35, 0x51, 0x4b, 0xea, 0x7b, 0x34, 0x6c, 0xc9, 0xe5, 0xa2, 0x0f, 0x30, 0x93, 0x6a, 0x85,
	0x9c, 0xeb, 0xac, 0x60, 0x2e, 0xd5, 0x93, 0x8d, 0xd9, 0xd0, 0x23, 0x9c, 0x48, 0xf5, 0xa6, 0xd9,
	0x6c, 0x8a, 0x48, 0xb2, 0x80, 0xb4, 0x9e, 0xef, 0x4e, 0xf8, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x4d, 0xdc, 0xe9, 0x2c, 0x03, 0x00, 0x00,
}
