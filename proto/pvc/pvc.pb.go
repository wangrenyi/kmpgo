// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/wangrenyi/kmpgo/proto/pvc/pvc.proto

/*
Package pvc is a generated protocol buffer package.

It is generated from these files:
	github.com/wangrenyi/kmpgo/proto/pvc/pvc.proto

It has these top-level messages:
*/
package pvc

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

// Client API for PVCInterface service

type PVCInterfaceClient interface {
	PVCCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVCUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVCUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVCDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVCDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVCGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVCList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVCWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PVCPatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
}

type pVCInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewPVCInterfaceClient(cc *grpc.ClientConn) PVCInterfaceClient {
	return &pVCInterfaceClient{cc}
}

func (c *pVCInterfaceClient) PVCCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVCInterfaceClient) PVCUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVCInterfaceClient) PVCUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCUpdateStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVCInterfaceClient) PVCDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVCInterfaceClient) PVCDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCDeleteCollection", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVCInterfaceClient) PVCGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVCInterfaceClient) PVCList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVCInterfaceClient) PVCWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCWatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pVCInterfaceClient) PVCPatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pvc.PVCInterface/PVCPatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PVCInterface service

type PVCInterfaceServer interface {
	PVCCreate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVCUpdate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVCUpdateStatus(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVCDelete(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVCDeleteCollection(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVCGet(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVCList(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVCWatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PVCPatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
}

func RegisterPVCInterfaceServer(s *grpc.Server, srv PVCInterfaceServer) {
	s.RegisterService(&_PVCInterface_serviceDesc, srv)
}

func _PVCInterface_PVCCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCCreate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVCInterface_PVCUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCUpdate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVCInterface_PVCUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCUpdateStatus(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVCInterface_PVCDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCDelete(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVCInterface_PVCDeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCDeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCDeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCDeleteCollection(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVCInterface_PVCGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCGet(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVCInterface_PVCList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCList(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVCInterface_PVCWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCWatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PVCInterface_PVCPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PVCInterfaceServer).PVCPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pvc.PVCInterface/PVCPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PVCInterfaceServer).PVCPatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PVCInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pvc.PVCInterface",
	HandlerType: (*PVCInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PVCCreate",
			Handler:    _PVCInterface_PVCCreate_Handler,
		},
		{
			MethodName: "PVCUpdate",
			Handler:    _PVCInterface_PVCUpdate_Handler,
		},
		{
			MethodName: "PVCUpdateStatus",
			Handler:    _PVCInterface_PVCUpdateStatus_Handler,
		},
		{
			MethodName: "PVCDelete",
			Handler:    _PVCInterface_PVCDelete_Handler,
		},
		{
			MethodName: "PVCDeleteCollection",
			Handler:    _PVCInterface_PVCDeleteCollection_Handler,
		},
		{
			MethodName: "PVCGet",
			Handler:    _PVCInterface_PVCGet_Handler,
		},
		{
			MethodName: "PVCList",
			Handler:    _PVCInterface_PVCList_Handler,
		},
		{
			MethodName: "PVCWatch",
			Handler:    _PVCInterface_PVCWatch_Handler,
		},
		{
			MethodName: "PVCPatch",
			Handler:    _PVCInterface_PVCPatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/wangrenyi/kmpgo/proto/pvc/pvc.proto",
}

func init() { proto.RegisterFile("github.com/wangrenyi/kmpgo/proto/pvc/pvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd2, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x06, 0x60, 0x41, 0x5c, 0x75, 0x10, 0x84, 0x7a, 0x5b, 0xf4, 0xe2, 0x03, 0xb4, 0xa0, 0xbe,
	0x41, 0x84, 0xb5, 0xe2, 0x21, 0xb8, 0x1a, 0xcf, 0xd9, 0x71, 0xec, 0x86, 0x6d, 0x33, 0x31, 0x99,
	0x56, 0x7c, 0x3e, 0x5f, 0x4c, 0x5c, 0xcb, 0xde, 0xc4, 0x43, 0x73, 0x08, 0x21, 0x64, 0xfe, 0x8f,
	0xff, 0x30, 0x50, 0x36, 0x4e, 0xd6, 0xfd, 0xaa, 0x44, 0xee, 0xaa, 0x0f, 0xeb, 0x9b, 0x48, 0xfe,
	0xd3, 0x55, 0x9b, 0x2e, 0x34, 0x5c, 0x85, 0xc8, 0xc2, 0x55, 0x18, 0xf0, 0xe7, 0x94, 0xdb, 0x57,
	0xb1, 0x1f, 0x06, 0x9c, 0xdf, 0xfc, 0x1b, 0x42, 0xee, 0x3a, 0xf6, 0xe3, 0xf5, 0x1b, 0xbd, 0xfa,
	0x3a, 0x80, 0x13, 0x6d, 0x54, 0xed, 0x85, 0xe2, 0x9b, 0x45, 0x2a, 0xee, 0xe1, 0x58, 0x1b, 0xa5,
	0x22, 0x59, 0xa1, 0xe2, 0xbc, 0x1c, 0x87, 0x6b, 0x3f, 0xf0, 0x86, 0x96, 0x14, 0x07, 0x87, 0xf4,
	0x48, 0xef, 0x3d, 0x25, 0x99, 0x5f, 0xfc, 0xf1, 0x9b, 0x02, 0xfb, 0x44, 0x97, 0x7b, 0xa3, 0xf5,
	0x1c, 0x5e, 0x33, 0x58, 0x1a, 0x4e, 0x77, 0xd6, 0x52, 0xac, 0xf4, 0x29, 0x4f, 0xbb, 0x5b, 0x6a,
	0x69, 0x7a, 0xbb, 0x27, 0x38, 0xdb, 0x59, 0x8a, 0xdb, 0x96, 0x50, 0x1c, 0xfb, 0xa9, 0xea, 0x02,
	0x66, 0xda, 0xa8, 0x05, 0xc9, 0x54, 0xe8, 0x0e, 0x0e, 0xb5, 0x51, 0x0f, 0x2e, 0x4d, 0x96, 0x6a,
	0x38, 0xd2, 0x46, 0xbd, 0x58, 0xc1, 0x75, 0x1e, 0x4a, 0x67, 0xa0, 0x56, 0xb3, 0xed, 0x32, 0x5f,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x17, 0xac, 0x61, 0x90, 0x39, 0x03, 0x00, 0x00,
}
