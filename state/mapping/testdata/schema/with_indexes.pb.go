// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: mapping/testdata/schema/with_indexes.proto

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

// EntityWithIndexes
type EntityWithIndexes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// one external id
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// required multiple external ids (minimum 1)
	RequiredExternalIds []string `protobuf:"bytes,3,rep,name=required_external_ids,json=requiredExternalIds,proto3" json:"required_external_ids,omitempty"`
	// optional multiple external ids (minimum 0)
	OptionalExternalIds []string `protobuf:"bytes,4,rep,name=optional_external_ids,json=optionalExternalIds,proto3" json:"optional_external_ids,omitempty"`
	Value               int32    `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EntityWithIndexes) Reset() {
	*x = EntityWithIndexes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_testdata_schema_with_indexes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityWithIndexes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityWithIndexes) ProtoMessage() {}

func (x *EntityWithIndexes) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_testdata_schema_with_indexes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityWithIndexes.ProtoReflect.Descriptor instead.
func (*EntityWithIndexes) Descriptor() ([]byte, []int) {
	return file_mapping_testdata_schema_with_indexes_proto_rawDescGZIP(), []int{0}
}

func (x *EntityWithIndexes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EntityWithIndexes) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *EntityWithIndexes) GetRequiredExternalIds() []string {
	if x != nil {
		return x.RequiredExternalIds
	}
	return nil
}

func (x *EntityWithIndexes) GetOptionalExternalIds() []string {
	if x != nil {
		return x.OptionalExternalIds
	}
	return nil
}

func (x *EntityWithIndexes) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// EntityWithIndexesList
type EntityWithIndexesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*EntityWithIndexes `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *EntityWithIndexesList) Reset() {
	*x = EntityWithIndexesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_testdata_schema_with_indexes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityWithIndexesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityWithIndexesList) ProtoMessage() {}

func (x *EntityWithIndexesList) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_testdata_schema_with_indexes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityWithIndexesList.ProtoReflect.Descriptor instead.
func (*EntityWithIndexesList) Descriptor() ([]byte, []int) {
	return file_mapping_testdata_schema_with_indexes_proto_rawDescGZIP(), []int{1}
}

func (x *EntityWithIndexesList) GetItems() []*EntityWithIndexes {
	if x != nil {
		return x.Items
	}
	return nil
}

// CreateEntityWithIndexes
type CreateEntityWithIndexes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// one external id
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// required multiple external ids (minimum 1)
	RequiredExternalIds []string `protobuf:"bytes,3,rep,name=required_external_ids,json=requiredExternalIds,proto3" json:"required_external_ids,omitempty"`
	// optional multiple external ids (minimum 0)
	OptionalExternalIds []string `protobuf:"bytes,4,rep,name=optional_external_ids,json=optionalExternalIds,proto3" json:"optional_external_ids,omitempty"`
	Value               int32    `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CreateEntityWithIndexes) Reset() {
	*x = CreateEntityWithIndexes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_testdata_schema_with_indexes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntityWithIndexes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntityWithIndexes) ProtoMessage() {}

func (x *CreateEntityWithIndexes) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_testdata_schema_with_indexes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntityWithIndexes.ProtoReflect.Descriptor instead.
func (*CreateEntityWithIndexes) Descriptor() ([]byte, []int) {
	return file_mapping_testdata_schema_with_indexes_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEntityWithIndexes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateEntityWithIndexes) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *CreateEntityWithIndexes) GetRequiredExternalIds() []string {
	if x != nil {
		return x.RequiredExternalIds
	}
	return nil
}

func (x *CreateEntityWithIndexes) GetOptionalExternalIds() []string {
	if x != nil {
		return x.OptionalExternalIds
	}
	return nil
}

func (x *CreateEntityWithIndexes) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// UpdateEntityEntityWithIndexes
type UpdateEntityWithIndexes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// one external id
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// required multiple external ids (minimum 1)
	RequiredExternalIds []string `protobuf:"bytes,3,rep,name=required_external_ids,json=requiredExternalIds,proto3" json:"required_external_ids,omitempty"`
	// optional multiple external ids (minimum 0)
	OptionalExternalIds []string `protobuf:"bytes,4,rep,name=optional_external_ids,json=optionalExternalIds,proto3" json:"optional_external_ids,omitempty"`
	Value               int32    `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateEntityWithIndexes) Reset() {
	*x = UpdateEntityWithIndexes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapping_testdata_schema_with_indexes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityWithIndexes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityWithIndexes) ProtoMessage() {}

func (x *UpdateEntityWithIndexes) ProtoReflect() protoreflect.Message {
	mi := &file_mapping_testdata_schema_with_indexes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityWithIndexes.ProtoReflect.Descriptor instead.
func (*UpdateEntityWithIndexes) Descriptor() ([]byte, []int) {
	return file_mapping_testdata_schema_with_indexes_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEntityWithIndexes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEntityWithIndexes) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *UpdateEntityWithIndexes) GetRequiredExternalIds() []string {
	if x != nil {
		return x.RequiredExternalIds
	}
	return nil
}

func (x *UpdateEntityWithIndexes) GetOptionalExternalIds() []string {
	if x != nil {
		return x.OptionalExternalIds
	}
	return nil
}

func (x *UpdateEntityWithIndexes) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_mapping_testdata_schema_with_indexes_proto protoreflect.FileDescriptor

var file_mapping_testdata_schema_with_indexes_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x57,
	0x69, 0x74, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x73, 0x12,
	0x32, 0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x48, 0x0a, 0x15, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x13, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x13, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc8,
	0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x57,
	0x69, 0x74, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x73, 0x12,
	0x32, 0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61,
	0x62, 0x2f, 0x63, 0x63, 0x6b, 0x69, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mapping_testdata_schema_with_indexes_proto_rawDescOnce sync.Once
	file_mapping_testdata_schema_with_indexes_proto_rawDescData = file_mapping_testdata_schema_with_indexes_proto_rawDesc
)

func file_mapping_testdata_schema_with_indexes_proto_rawDescGZIP() []byte {
	file_mapping_testdata_schema_with_indexes_proto_rawDescOnce.Do(func() {
		file_mapping_testdata_schema_with_indexes_proto_rawDescData = protoimpl.X.CompressGZIP(file_mapping_testdata_schema_with_indexes_proto_rawDescData)
	})
	return file_mapping_testdata_schema_with_indexes_proto_rawDescData
}

var file_mapping_testdata_schema_with_indexes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mapping_testdata_schema_with_indexes_proto_goTypes = []interface{}{
	(*EntityWithIndexes)(nil),       // 0: schema.EntityWithIndexes
	(*EntityWithIndexesList)(nil),   // 1: schema.EntityWithIndexesList
	(*CreateEntityWithIndexes)(nil), // 2: schema.CreateEntityWithIndexes
	(*UpdateEntityWithIndexes)(nil), // 3: schema.UpdateEntityWithIndexes
}
var file_mapping_testdata_schema_with_indexes_proto_depIdxs = []int32{
	0, // 0: schema.EntityWithIndexesList.items:type_name -> schema.EntityWithIndexes
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mapping_testdata_schema_with_indexes_proto_init() }
func file_mapping_testdata_schema_with_indexes_proto_init() {
	if File_mapping_testdata_schema_with_indexes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mapping_testdata_schema_with_indexes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityWithIndexes); i {
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
		file_mapping_testdata_schema_with_indexes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityWithIndexesList); i {
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
		file_mapping_testdata_schema_with_indexes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntityWithIndexes); i {
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
		file_mapping_testdata_schema_with_indexes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityWithIndexes); i {
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
			RawDescriptor: file_mapping_testdata_schema_with_indexes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mapping_testdata_schema_with_indexes_proto_goTypes,
		DependencyIndexes: file_mapping_testdata_schema_with_indexes_proto_depIdxs,
		MessageInfos:      file_mapping_testdata_schema_with_indexes_proto_msgTypes,
	}.Build()
	File_mapping_testdata_schema_with_indexes_proto = out.File
	file_mapping_testdata_schema_with_indexes_proto_rawDesc = nil
	file_mapping_testdata_schema_with_indexes_proto_goTypes = nil
	file_mapping_testdata_schema_with_indexes_proto_depIdxs = nil
}
