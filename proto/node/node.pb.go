// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/wangrenyi/kmpgo/proto/node/node.proto

/*
Package node is a generated protocol buffer package.

It is generated from these files:
	github.com/wangrenyi/kmpgo/proto/node/node.proto

It has these top-level messages:
*/
package node

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

// Client API for NodeInterface service

type NodeInterfaceClient interface {
	NodeCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NodeUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NodeUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NodeDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NodeDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NodeGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NodeList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NodeWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	NodePatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	// Node expansion
	NodePatchStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
}

type nodeInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewNodeInterfaceClient(cc *grpc.ClientConn) NodeInterfaceClient {
	return &nodeInterfaceClient{cc}
}

func (c *nodeInterfaceClient) NodeCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodeCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodeUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodeUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodeUpdateStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodeDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodeDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodeDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodeDeleteCollection", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodeGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodeGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodeList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodeList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodeWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodeWatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodePatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodePatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeInterfaceClient) NodePatchStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/node.NodeInterface/NodePatchStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeInterface service

type NodeInterfaceServer interface {
	NodeCreate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NodeUpdate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NodeUpdateStatus(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NodeDelete(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NodeDeleteCollection(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NodeGet(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NodeList(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NodeWatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	NodePatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	// Node expansion
	NodePatchStatus(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
}

func RegisterNodeInterfaceServer(s *grpc.Server, srv NodeInterfaceServer) {
	s.RegisterService(&_NodeInterface_serviceDesc, srv)
}

func _NodeInterface_NodeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodeCreate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodeUpdate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodeUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodeUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodeUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodeUpdateStatus(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodeDelete(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodeDeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodeDeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodeDeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodeDeleteCollection(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodeGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodeGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodeGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodeGet(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodeList(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodeWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodeWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodeWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodeWatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodePatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodePatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeInterface_NodePatchStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeInterfaceServer).NodePatchStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/node.NodeInterface/NodePatchStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeInterfaceServer).NodePatchStatus(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "node.NodeInterface",
	HandlerType: (*NodeInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodeCreate",
			Handler:    _NodeInterface_NodeCreate_Handler,
		},
		{
			MethodName: "NodeUpdate",
			Handler:    _NodeInterface_NodeUpdate_Handler,
		},
		{
			MethodName: "NodeUpdateStatus",
			Handler:    _NodeInterface_NodeUpdateStatus_Handler,
		},
		{
			MethodName: "NodeDelete",
			Handler:    _NodeInterface_NodeDelete_Handler,
		},
		{
			MethodName: "NodeDeleteCollection",
			Handler:    _NodeInterface_NodeDeleteCollection_Handler,
		},
		{
			MethodName: "NodeGet",
			Handler:    _NodeInterface_NodeGet_Handler,
		},
		{
			MethodName: "NodeList",
			Handler:    _NodeInterface_NodeList_Handler,
		},
		{
			MethodName: "NodeWatch",
			Handler:    _NodeInterface_NodeWatch_Handler,
		},
		{
			MethodName: "NodePatch",
			Handler:    _NodeInterface_NodePatch_Handler,
		},
		{
			MethodName: "NodePatchStatus",
			Handler:    _NodeInterface_NodePatchStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/wangrenyi/kmpgo/proto/node/node.proto",
}

func init() { proto.RegisterFile("github.com/wangrenyi/kmpgo/proto/node/node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd3, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x06, 0x70, 0x0f, 0x52, 0x75, 0x40, 0x94, 0xe0, 0xa9, 0xe8, 0xc5, 0x07, 0x48, 0x44, 0x7d,
	0x83, 0x0a, 0x1a, 0x15, 0xa9, 0x96, 0xe2, 0x79, 0xbb, 0xf9, 0x4c, 0x97, 0x26, 0x3b, 0x71, 0x77,
	0x52, 0xf1, 0x69, 0x7c, 0x55, 0xd9, 0x58, 0xcd, 0x49, 0x7a, 0xc8, 0x5e, 0xf6, 0xff, 0xfc, 0xf8,
	0x0e, 0x3b, 0x74, 0x51, 0x1a, 0x59, 0xb6, 0x8b, 0x54, 0x73, 0x9d, 0x7d, 0x28, 0x5b, 0x3a, 0xd8,
	0x4f, 0x93, 0xad, 0xea, 0xa6, 0xe4, 0xac, 0x71, 0x2c, 0x9c, 0x59, 0x2e, 0xd0, 0x0d, 0x69, 0xb7,
	0x4f, 0x76, 0xc3, 0x7a, 0x7c, 0xbd, 0xb5, 0x4e, 0x73, 0x5d, 0xb3, 0xdd, 0x4c, 0x3f, 0xb5, 0x97,
	0x5f, 0x23, 0x3a, 0x7c, 0xe2, 0x02, 0xb9, 0x15, 0xb8, 0x37, 0xa5, 0x91, 0x3c, 0x10, 0x85, 0x83,
	0x89, 0x83, 0x12, 0x24, 0xa7, 0xe9, 0xe6, 0x79, 0x6e, 0xd7, 0xbc, 0xc2, 0x0c, 0x6e, 0x6d, 0x34,
	0x5e, 0xf0, 0xde, 0xc2, 0xcb, 0xf8, 0xec, 0x9f, 0x5b, 0xdf, 0xb0, 0xf5, 0x38, 0xdf, 0xf9, 0xc5,
	0xe6, 0x4d, 0x11, 0x01, 0x7b, 0xa6, 0xe3, 0x1e, 0x9b, 0x89, 0x92, 0xd6, 0x47, 0xca, 0x77, 0x83,
	0x0a, 0xc3, 0xf3, 0xcd, 0xe9, 0xa4, 0xc7, 0x26, 0x5c, 0x55, 0xd0, 0x62, 0xd8, 0x0e, 0x65, 0xef,
	0x68, 0x2f, 0xb0, 0xb7, 0x90, 0xa1, 0x52, 0x4e, 0xfb, 0x41, 0x7a, 0x34, 0x7e, 0x30, 0x75, 0x4f,
	0x07, 0x81, 0x7a, 0x55, 0xa2, 0x97, 0x91, 0xac, 0x69, 0x0c, 0x6b, 0x4a, 0x47, 0x7f, 0x56, 0x94,
	0x2f, 0xb2, 0x18, 0x75, 0x8d, 0x72, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x27, 0xd5, 0x9d,
	0x98, 0x03, 0x00, 0x00,
}