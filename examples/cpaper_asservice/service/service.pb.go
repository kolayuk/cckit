// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import schema "github.com/s7techlab/cckit/examples/cpaper_asservice/schema"

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

// Client API for CPaper service

type CPaperClient interface {
	List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*schema.CommercialPaperList, error)
	Get(ctx context.Context, in *schema.CommercialPaperId, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	GetByExternalId(ctx context.Context, in *schema.ExternalId, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	Issue(ctx context.Context, in *schema.IssueCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	Buy(ctx context.Context, in *schema.BuyCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	Redeem(ctx context.Context, in *schema.RedeemCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
	Delete(ctx context.Context, in *schema.CommercialPaperId, opts ...grpc.CallOption) (*schema.CommercialPaper, error)
}

type cPaperClient struct {
	cc *grpc.ClientConn
}

func NewCPaperClient(cc *grpc.ClientConn) CPaperClient {
	return &cPaperClient{cc}
}

func (c *cPaperClient) List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*schema.CommercialPaperList, error) {
	out := new(schema.CommercialPaperList)
	err := grpc.Invoke(ctx, "/service.CPaper/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Get(ctx context.Context, in *schema.CommercialPaperId, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := grpc.Invoke(ctx, "/service.CPaper/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) GetByExternalId(ctx context.Context, in *schema.ExternalId, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := grpc.Invoke(ctx, "/service.CPaper/GetByExternalId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Issue(ctx context.Context, in *schema.IssueCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := grpc.Invoke(ctx, "/service.CPaper/Issue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Buy(ctx context.Context, in *schema.BuyCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := grpc.Invoke(ctx, "/service.CPaper/Buy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Redeem(ctx context.Context, in *schema.RedeemCommercialPaper, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := grpc.Invoke(ctx, "/service.CPaper/Redeem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPaperClient) Delete(ctx context.Context, in *schema.CommercialPaperId, opts ...grpc.CallOption) (*schema.CommercialPaper, error) {
	out := new(schema.CommercialPaper)
	err := grpc.Invoke(ctx, "/service.CPaper/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CPaper service

type CPaperServer interface {
	List(context.Context, *google_protobuf1.Empty) (*schema.CommercialPaperList, error)
	Get(context.Context, *schema.CommercialPaperId) (*schema.CommercialPaper, error)
	GetByExternalId(context.Context, *schema.ExternalId) (*schema.CommercialPaper, error)
	Issue(context.Context, *schema.IssueCommercialPaper) (*schema.CommercialPaper, error)
	Buy(context.Context, *schema.BuyCommercialPaper) (*schema.CommercialPaper, error)
	Redeem(context.Context, *schema.RedeemCommercialPaper) (*schema.CommercialPaper, error)
	Delete(context.Context, *schema.CommercialPaperId) (*schema.CommercialPaper, error)
}

func RegisterCPaperServer(s *grpc.Server, srv CPaperServer) {
	s.RegisterService(&_CPaper_serviceDesc, srv)
}

func _CPaper_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CPaper/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).List(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.CommercialPaperId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CPaper/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Get(ctx, req.(*schema.CommercialPaperId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_GetByExternalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.ExternalId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).GetByExternalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CPaper/GetByExternalId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).GetByExternalId(ctx, req.(*schema.ExternalId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.IssueCommercialPaper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CPaper/Issue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Issue(ctx, req.(*schema.IssueCommercialPaper))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.BuyCommercialPaper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CPaper/Buy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Buy(ctx, req.(*schema.BuyCommercialPaper))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Redeem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.RedeemCommercialPaper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Redeem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CPaper/Redeem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Redeem(ctx, req.(*schema.RedeemCommercialPaper))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPaper_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(schema.CommercialPaperId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPaperServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.CPaper/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPaperServer).Delete(ctx, req.(*schema.CommercialPaperId))
	}
	return interceptor(ctx, in, info, handler)
}

var _CPaper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.CPaper",
	HandlerType: (*CPaperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CPaper_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CPaper_Get_Handler,
		},
		{
			MethodName: "GetByExternalId",
			Handler:    _CPaper_GetByExternalId_Handler,
		},
		{
			MethodName: "Issue",
			Handler:    _CPaper_Issue_Handler,
		},
		{
			MethodName: "Buy",
			Handler:    _CPaper_Buy_Handler,
		},
		{
			MethodName: "Redeem",
			Handler:    _CPaper_Redeem_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CPaper_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xd1, 0xaa, 0xd3, 0x30,
	0x18, 0xc7, 0xc1, 0xcd, 0x0e, 0x23, 0x73, 0xf2, 0xa9, 0x13, 0xbb, 0x89, 0xd8, 0x1b, 0xc1, 0x8b,
	0x06, 0xf4, 0xc2, 0xfb, 0xcd, 0x31, 0x07, 0x03, 0x45, 0x10, 0xc4, 0x1b, 0x49, 0xd3, 0xcf, 0x2d,
	0xd8, 0x34, 0x25, 0x49, 0x65, 0x65, 0xec, 0xc6, 0x57, 0xf0, 0xd1, 0x7c, 0x85, 0xf3, 0x1e, 0xe7,
	0xd0, 0x34, 0xe1, 0xc0, 0x81, 0x72, 0x06, 0xe7, 0xaa, 0xfc, 0xfb, 0xe5, 0xff, 0xfb, 0x11, 0xbe,
	0x90, 0xb1, 0x41, 0xfd, 0x47, 0x70, 0x4c, 0x2b, 0xad, 0xac, 0x82, 0x91, 0x8f, 0xf1, 0x7c, 0xa7,
	0xd4, 0xae, 0x40, 0xca, 0x2a, 0x41, 0x59, 0x59, 0x2a, 0xcb, 0xac, 0x50, 0xa5, 0xe9, 0x8e, 0xc5,
	0x33, 0x3f, 0x75, 0x29, 0xab, 0x7f, 0x51, 0x94, 0x95, 0x6d, 0xfc, 0xf0, 0xd3, 0x4e, 0xd8, 0x7d,
	0x9d, 0xa5, 0x5c, 0x49, 0x6a, 0x3e, 0x58, 0xe4, 0xfb, 0x82, 0x65, 0x94, 0xf3, 0xdf, 0xc2, 0x52,
	0x3c, 0x30, 0x59, 0x15, 0x68, 0x28, 0xaf, 0x58, 0x85, 0xfa, 0x27, 0x33, 0x5e, 0x48, 0x0d, 0xdf,
	0xa3, 0x64, 0xfe, 0xd3, 0x91, 0xde, 0x5d, 0x0e, 0x49, 0xb4, 0xfc, 0xd2, 0x1e, 0x84, 0x2d, 0x19,
	0x6e, 0x85, 0xb1, 0x30, 0x4d, 0x3b, 0x75, 0x1a, 0xd4, 0xe9, 0xaa, 0x55, 0xc7, 0xb3, 0xd4, 0x37,
	0x97, 0x4a, 0x4a, 0xd4, 0x5c, 0xb0, 0xc2, 0x15, 0xdb, 0x52, 0x32, 0xf9, 0xfb, 0xff, 0xe2, 0xdf,
	0xbd, 0x07, 0x30, 0xf2, 0x5a, 0xc8, 0xc8, 0x60, 0x8d, 0x16, 0x5e, 0xf4, 0x94, 0x36, 0x79, 0xfc,
	0xbc, 0x67, 0x94, 0xbc, 0x71, 0xac, 0xd7, 0xf0, 0xca, 0xb3, 0xe8, 0x51, 0x18, 0x53, 0xa3, 0x3e,
	0xd1, 0x63, 0x77, 0xa5, 0xb2, 0x96, 0x19, 0xea, 0x13, 0xfc, 0x20, 0x93, 0x35, 0xda, 0x45, 0xb3,
	0x3a, 0x58, 0xd4, 0x25, 0x2b, 0x36, 0x39, 0x40, 0x80, 0x5e, 0xff, 0xeb, 0x17, 0xc5, 0x4e, 0xf4,
	0x14, 0x20, 0x88, 0xf0, 0x60, 0x45, 0x4e, 0x8f, 0x22, 0x3f, 0xc1, 0x37, 0x72, 0x7f, 0xd3, 0x5a,
	0x61, 0x1e, 0xda, 0x2e, 0xde, 0x40, 0xf4, 0xb3, 0x9f, 0x39, 0xf6, 0x24, 0x19, 0x07, 0xb6, 0xbb,
	0x03, 0x7c, 0x26, 0x83, 0x45, 0xdd, 0x40, 0x1c, 0x6a, 0x8b, 0xba, 0x39, 0x1b, 0xf9, 0xc4, 0x21,
	0xc7, 0xc9, 0xc3, 0x80, 0xcc, 0xea, 0x06, 0xbe, 0x93, 0xe8, 0x2b, 0xe6, 0x88, 0x12, 0x5e, 0x86,
	0x5e, 0x97, 0xcf, 0xc6, 0x4e, 0x1d, 0xf6, 0x71, 0xf2, 0x28, 0x60, 0x75, 0xc7, 0x43, 0x12, 0x7d,
	0xc4, 0x02, 0x2d, 0xde, 0x65, 0x89, 0x6f, 0x6f, 0x5b, 0x62, 0x16, 0xb9, 0x67, 0xf6, 0xfe, 0x2a,
	0x00, 0x00, 0xff, 0xff, 0x1d, 0xa1, 0x70, 0x55, 0x27, 0x03, 0x00, 0x00,
}