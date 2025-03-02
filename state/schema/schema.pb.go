// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: schema/schema.proto

package schema

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// KeyRefId  id part of key reference
type KeyRefId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// entity type
	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	// idx name from entity type
	Idx string `protobuf:"bytes,2,opt,name=idx,proto3" json:"idx,omitempty"`
	// referred key
	RefKey []string `protobuf:"bytes,3,rep,name=ref_key,json=refKey,proto3" json:"ref_key,omitempty"`
}

func (x *KeyRefId) Reset() {
	*x = KeyRefId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRefId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRefId) ProtoMessage() {}

func (x *KeyRefId) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRefId.ProtoReflect.Descriptor instead.
func (*KeyRefId) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{0}
}

func (x *KeyRefId) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *KeyRefId) GetIdx() string {
	if x != nil {
		return x.Idx
	}
	return ""
}

func (x *KeyRefId) GetRefKey() []string {
	if x != nil {
		return x.RefKey
	}
	return nil
}

type KeyRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// entity type
	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	// idx name from entity type
	Idx string `protobuf:"bytes,2,opt,name=idx,proto3" json:"idx,omitempty"`
	// referred key
	RefKey []string `protobuf:"bytes,3,rep,name=ref_key,json=refKey,proto3" json:"ref_key,omitempty"`
	// primary key instance linked to
	PKey []string `protobuf:"bytes,4,rep,name=p_key,json=pKey,proto3" json:"p_key,omitempty"`
}

func (x *KeyRef) Reset() {
	*x = KeyRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRef) ProtoMessage() {}

func (x *KeyRef) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRef.ProtoReflect.Descriptor instead.
func (*KeyRef) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{1}
}

func (x *KeyRef) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *KeyRef) GetIdx() string {
	if x != nil {
		return x.Idx
	}
	return ""
}

func (x *KeyRef) GetRefKey() []string {
	if x != nil {
		return x.RefKey
	}
	return nil
}

func (x *KeyRef) GetPKey() []string {
	if x != nil {
		return x.PKey
	}
	return nil
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*anypb.Any `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_schema_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_schema_schema_proto_rawDescGZIP(), []int{2}
}

func (x *List) GetItems() []*anypb.Any {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_schema_schema_proto protoreflect.FileDescriptor

var file_schema_schema_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d,
	0x0a, 0x08, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x66, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x66, 0x4b, 0x65, 0x79, 0x22, 0x60, 0x0a,
	0x06, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64,
	0x78, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x66, 0x4b, 0x65, 0x79, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x4b, 0x65, 0x79, 0x22,
	0x32, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2f, 0x63, 0x63, 0x6b, 0x69,
	0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_schema_proto_rawDescOnce sync.Once
	file_schema_schema_proto_rawDescData = file_schema_schema_proto_rawDesc
)

func file_schema_schema_proto_rawDescGZIP() []byte {
	file_schema_schema_proto_rawDescOnce.Do(func() {
		file_schema_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_schema_proto_rawDescData)
	})
	return file_schema_schema_proto_rawDescData
}

var file_schema_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schema_schema_proto_goTypes = []interface{}{
	(*KeyRefId)(nil),  // 0: state.schema.KeyRefId
	(*KeyRef)(nil),    // 1: state.schema.KeyRef
	(*List)(nil),      // 2: state.schema.List
	(*anypb.Any)(nil), // 3: google.protobuf.Any
}
var file_schema_schema_proto_depIdxs = []int32{
	3, // 0: state.schema.List.items:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_schema_schema_proto_init() }
func file_schema_schema_proto_init() {
	if File_schema_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRefId); i {
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
		file_schema_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRef); i {
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
		file_schema_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
			RawDescriptor: file_schema_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_schema_proto_goTypes,
		DependencyIndexes: file_schema_schema_proto_depIdxs,
		MessageInfos:      file_schema_schema_proto_msgTypes,
	}.Build()
	File_schema_schema_proto = out.File
	file_schema_schema_proto_rawDesc = nil
	file_schema_schema_proto_goTypes = nil
	file_schema_schema_proto_depIdxs = nil
}
