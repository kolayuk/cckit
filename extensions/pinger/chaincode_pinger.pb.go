// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: chaincode_pinger.proto

package pinger

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// stores time and certificate of ping tx creator
type PingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvokerId           string                 `protobuf:"bytes,1,opt,name=invoker_id,json=invokerId,proto3" json:"invoker_id,omitempty"`
	InvokerCert         []byte                 `protobuf:"bytes,2,opt,name=invoker_cert,json=invokerCert,proto3" json:"invoker_cert,omitempty"`
	EndorsingServerTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=endorsing_server_time,json=endorsingServerTime,proto3" json:"endorsing_server_time,omitempty"`
	TxTime              *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=tx_time,json=txTime,proto3" json:"tx_time,omitempty"`
}

func (x *PingInfo) Reset() {
	*x = PingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_pinger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingInfo) ProtoMessage() {}

func (x *PingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_pinger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingInfo.ProtoReflect.Descriptor instead.
func (*PingInfo) Descriptor() ([]byte, []int) {
	return file_chaincode_pinger_proto_rawDescGZIP(), []int{0}
}

func (x *PingInfo) GetInvokerId() string {
	if x != nil {
		return x.InvokerId
	}
	return ""
}

func (x *PingInfo) GetInvokerCert() []byte {
	if x != nil {
		return x.InvokerCert
	}
	return nil
}

func (x *PingInfo) GetEndorsingServerTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndorsingServerTime
	}
	return nil
}

func (x *PingInfo) GetTxTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TxTime
	}
	return nil
}

var File_chaincode_pinger_proto protoreflect.FileDescriptor

var file_chaincode_pinger_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x12, 0x4e, 0x0a, 0x15, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x13, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x06, 0x74, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x78, 0x0a, 0x16, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2f, 0x63, 0x63,
	0x6b, 0x69, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaincode_pinger_proto_rawDescOnce sync.Once
	file_chaincode_pinger_proto_rawDescData = file_chaincode_pinger_proto_rawDesc
)

func file_chaincode_pinger_proto_rawDescGZIP() []byte {
	file_chaincode_pinger_proto_rawDescOnce.Do(func() {
		file_chaincode_pinger_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaincode_pinger_proto_rawDescData)
	})
	return file_chaincode_pinger_proto_rawDescData
}

var file_chaincode_pinger_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chaincode_pinger_proto_goTypes = []interface{}{
	(*PingInfo)(nil),              // 0: extensions.pinger.PingInfo
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_chaincode_pinger_proto_depIdxs = []int32{
	1, // 0: extensions.pinger.PingInfo.endorsing_server_time:type_name -> google.protobuf.Timestamp
	1, // 1: extensions.pinger.PingInfo.tx_time:type_name -> google.protobuf.Timestamp
	2, // 2: extensions.pinger.ChaincodePingerService.Ping:input_type -> google.protobuf.Empty
	0, // 3: extensions.pinger.ChaincodePingerService.Ping:output_type -> extensions.pinger.PingInfo
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chaincode_pinger_proto_init() }
func file_chaincode_pinger_proto_init() {
	if File_chaincode_pinger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaincode_pinger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chaincode_pinger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chaincode_pinger_proto_goTypes,
		DependencyIndexes: file_chaincode_pinger_proto_depIdxs,
		MessageInfos:      file_chaincode_pinger_proto_msgTypes,
	}.Build()
	File_chaincode_pinger_proto = out.File
	file_chaincode_pinger_proto_rawDesc = nil
	file_chaincode_pinger_proto_goTypes = nil
	file_chaincode_pinger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChaincodePingerServiceClient is the client API for ChaincodePingerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChaincodePingerServiceClient interface {
	// ping chaincode
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingInfo, error)
}

type chaincodePingerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChaincodePingerServiceClient(cc grpc.ClientConnInterface) ChaincodePingerServiceClient {
	return &chaincodePingerServiceClient{cc}
}

func (c *chaincodePingerServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingInfo, error) {
	out := new(PingInfo)
	err := c.cc.Invoke(ctx, "/extensions.pinger.ChaincodePingerService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaincodePingerServiceServer is the server API for ChaincodePingerService service.
type ChaincodePingerServiceServer interface {
	// ping chaincode
	Ping(context.Context, *emptypb.Empty) (*PingInfo, error)
}

// UnimplementedChaincodePingerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChaincodePingerServiceServer struct {
}

func (*UnimplementedChaincodePingerServiceServer) Ping(context.Context, *emptypb.Empty) (*PingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterChaincodePingerServiceServer(s *grpc.Server, srv ChaincodePingerServiceServer) {
	s.RegisterService(&_ChaincodePingerService_serviceDesc, srv)
}

func _ChaincodePingerService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodePingerServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extensions.pinger.ChaincodePingerService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodePingerServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChaincodePingerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "extensions.pinger.ChaincodePingerService",
	HandlerType: (*ChaincodePingerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ChaincodePingerService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chaincode_pinger.proto",
}