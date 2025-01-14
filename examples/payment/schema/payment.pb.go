// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: payment/schema/payment.proto

package schema

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_schema_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_payment_schema_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_payment_schema_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Payment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Payment) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PaymentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Payment `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PaymentList) Reset() {
	*x = PaymentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_schema_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentList) ProtoMessage() {}

func (x *PaymentList) ProtoReflect() protoreflect.Message {
	mi := &file_payment_schema_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentList.ProtoReflect.Descriptor instead.
func (*PaymentList) Descriptor() ([]byte, []int) {
	return file_payment_schema_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentList) GetItems() []*Payment {
	if x != nil {
		return x.Items
	}
	return nil
}

type PaymentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PaymentEvent) Reset() {
	*x = PaymentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_schema_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentEvent) ProtoMessage() {}

func (x *PaymentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_payment_schema_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentEvent.ProtoReflect.Descriptor instead.
func (*PaymentEvent) Descriptor() ([]byte, []int) {
	return file_payment_schema_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PaymentEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PaymentEvent) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_payment_schema_payment_proto protoreflect.FileDescriptor

var file_payment_schema_payment_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x45, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a,
	0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x4a, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x37,
	0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2f, 0x63, 0x63, 0x6b, 0x69, 0x74, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_schema_payment_proto_rawDescOnce sync.Once
	file_payment_schema_payment_proto_rawDescData = file_payment_schema_payment_proto_rawDesc
)

func file_payment_schema_payment_proto_rawDescGZIP() []byte {
	file_payment_schema_payment_proto_rawDescOnce.Do(func() {
		file_payment_schema_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_schema_payment_proto_rawDescData)
	})
	return file_payment_schema_payment_proto_rawDescData
}

var file_payment_schema_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payment_schema_payment_proto_goTypes = []interface{}{
	(*Payment)(nil),      // 0: schema.Payment
	(*PaymentList)(nil),  // 1: schema.PaymentList
	(*PaymentEvent)(nil), // 2: schema.PaymentEvent
}
var file_payment_schema_payment_proto_depIdxs = []int32{
	0, // 0: schema.PaymentList.items:type_name -> schema.Payment
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_schema_payment_proto_init() }
func file_payment_schema_payment_proto_init() {
	if File_payment_schema_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_schema_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_payment_schema_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentList); i {
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
		file_payment_schema_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentEvent); i {
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
			RawDescriptor: file_payment_schema_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payment_schema_payment_proto_goTypes,
		DependencyIndexes: file_payment_schema_payment_proto_depIdxs,
		MessageInfos:      file_payment_schema_payment_proto_msgTypes,
	}.Build()
	File_payment_schema_payment_proto = out.File
	file_payment_schema_payment_proto_rawDesc = nil
	file_payment_schema_payment_proto_goTypes = nil
	file_payment_schema_payment_proto_depIdxs = nil
}
