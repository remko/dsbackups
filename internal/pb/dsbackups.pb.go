// Reverse engineered with protoc --decode_raw

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: dsbackups.proto

package pb

import (
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

type OverallExportMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exports []*ExportMetadataEntry `protobuf:"bytes,1,rep,name=exports,proto3" json:"exports,omitempty"`
}

func (x *OverallExportMetadata) Reset() {
	*x = OverallExportMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsbackups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverallExportMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverallExportMetadata) ProtoMessage() {}

func (x *OverallExportMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_dsbackups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverallExportMetadata.ProtoReflect.Descriptor instead.
func (*OverallExportMetadata) Descriptor() ([]byte, []int) {
	return file_dsbackups_proto_rawDescGZIP(), []int{0}
}

func (x *OverallExportMetadata) GetExports() []*ExportMetadataEntry {
	if x != nil {
		return x.Exports
	}
	return nil
}

type ExportMetadataEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind  *ExportMetadataEntryKind `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Path  string                   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Count int64                    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Size  int64                    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ExportMetadataEntry) Reset() {
	*x = ExportMetadataEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsbackups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetadataEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetadataEntry) ProtoMessage() {}

func (x *ExportMetadataEntry) ProtoReflect() protoreflect.Message {
	mi := &file_dsbackups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetadataEntry.ProtoReflect.Descriptor instead.
func (*ExportMetadataEntry) Descriptor() ([]byte, []int) {
	return file_dsbackups_proto_rawDescGZIP(), []int{1}
}

func (x *ExportMetadataEntry) GetKind() *ExportMetadataEntryKind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *ExportMetadataEntry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ExportMetadataEntry) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ExportMetadataEntry) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ExportMetadataEntryKind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unknown1 int32  `protobuf:"varint,1,opt,name=unknown1,proto3" json:"unknown1,omitempty"`
	Kind     string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Unknown2 int32  `protobuf:"varint,3,opt,name=unknown2,proto3" json:"unknown2,omitempty"`
}

func (x *ExportMetadataEntryKind) Reset() {
	*x = ExportMetadataEntryKind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsbackups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetadataEntryKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetadataEntryKind) ProtoMessage() {}

func (x *ExportMetadataEntryKind) ProtoReflect() protoreflect.Message {
	mi := &file_dsbackups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetadataEntryKind.ProtoReflect.Descriptor instead.
func (*ExportMetadataEntryKind) Descriptor() ([]byte, []int) {
	return file_dsbackups_proto_rawDescGZIP(), []int{2}
}

func (x *ExportMetadataEntryKind) GetUnknown1() int32 {
	if x != nil {
		return x.Unknown1
	}
	return 0
}

func (x *ExportMetadataEntryKind) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ExportMetadataEntryKind) GetUnknown2() int32 {
	if x != nil {
		return x.Unknown2
	}
	return 0
}

type ExportMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *ExportMetadataOperation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Items     *ExportMetadataItems     `protobuf:"bytes,2,opt,name=items,proto3" json:"items,omitempty"`
}

func (x *ExportMetadata) Reset() {
	*x = ExportMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsbackups_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetadata) ProtoMessage() {}

func (x *ExportMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_dsbackups_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetadata.ProtoReflect.Descriptor instead.
func (*ExportMetadata) Descriptor() ([]byte, []int) {
	return file_dsbackups_proto_rawDescGZIP(), []int{3}
}

func (x *ExportMetadata) GetOperation() *ExportMetadataOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *ExportMetadata) GetItems() *ExportMetadataItems {
	if x != nil {
		return x.Items
	}
	return nil
}

type ExportMetadataOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`  // "export_entities"
	StartTime int64  `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"` // timestamp in microseconds (seems to be something else in emulator exports)
	EndTime   int64  `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`     // timestamp in microseconds (seems to be something else in emulator exports)
}

func (x *ExportMetadataOperation) Reset() {
	*x = ExportMetadataOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsbackups_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetadataOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetadataOperation) ProtoMessage() {}

func (x *ExportMetadataOperation) ProtoReflect() protoreflect.Message {
	mi := &file_dsbackups_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetadataOperation.ProtoReflect.Descriptor instead.
func (*ExportMetadataOperation) Descriptor() ([]byte, []int) {
	return file_dsbackups_proto_rawDescGZIP(), []int{4}
}

func (x *ExportMetadataOperation) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *ExportMetadataOperation) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *ExportMetadataOperation) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type ExportMetadataItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind    string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Outputs []string `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *ExportMetadataItems) Reset() {
	*x = ExportMetadataItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dsbackups_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportMetadataItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportMetadataItems) ProtoMessage() {}

func (x *ExportMetadataItems) ProtoReflect() protoreflect.Message {
	mi := &file_dsbackups_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportMetadataItems.ProtoReflect.Descriptor instead.
func (*ExportMetadataItems) Descriptor() ([]byte, []int) {
	return file_dsbackups_proto_rawDescGZIP(), []int{5}
}

func (x *ExportMetadataItems) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ExportMetadataItems) GetOutputs() []string {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_dsbackups_proto protoreflect.FileDescriptor

var file_dsbackups_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x73, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x64, 0x73, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x22, 0x51, 0x0a, 0x15,
	0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x73, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22,
	0x8b, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x36, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x73, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x73, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x65, 0x0a,
	0x17, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x32, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x73, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x73, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x6f, 0x0a, 0x17, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x43, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6d, 0x6b, 0x6f, 0x2f, 0x64, 0x73, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dsbackups_proto_rawDescOnce sync.Once
	file_dsbackups_proto_rawDescData = file_dsbackups_proto_rawDesc
)

func file_dsbackups_proto_rawDescGZIP() []byte {
	file_dsbackups_proto_rawDescOnce.Do(func() {
		file_dsbackups_proto_rawDescData = protoimpl.X.CompressGZIP(file_dsbackups_proto_rawDescData)
	})
	return file_dsbackups_proto_rawDescData
}

var file_dsbackups_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dsbackups_proto_goTypes = []interface{}{
	(*OverallExportMetadata)(nil),   // 0: dsbackups.OverallExportMetadata
	(*ExportMetadataEntry)(nil),     // 1: dsbackups.ExportMetadataEntry
	(*ExportMetadataEntryKind)(nil), // 2: dsbackups.ExportMetadataEntryKind
	(*ExportMetadata)(nil),          // 3: dsbackups.ExportMetadata
	(*ExportMetadataOperation)(nil), // 4: dsbackups.ExportMetadataOperation
	(*ExportMetadataItems)(nil),     // 5: dsbackups.ExportMetadataItems
}
var file_dsbackups_proto_depIdxs = []int32{
	1, // 0: dsbackups.OverallExportMetadata.exports:type_name -> dsbackups.ExportMetadataEntry
	2, // 1: dsbackups.ExportMetadataEntry.kind:type_name -> dsbackups.ExportMetadataEntryKind
	4, // 2: dsbackups.ExportMetadata.operation:type_name -> dsbackups.ExportMetadataOperation
	5, // 3: dsbackups.ExportMetadata.items:type_name -> dsbackups.ExportMetadataItems
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dsbackups_proto_init() }
func file_dsbackups_proto_init() {
	if File_dsbackups_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dsbackups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverallExportMetadata); i {
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
		file_dsbackups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetadataEntry); i {
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
		file_dsbackups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetadataEntryKind); i {
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
		file_dsbackups_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetadata); i {
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
		file_dsbackups_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetadataOperation); i {
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
		file_dsbackups_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportMetadataItems); i {
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
			RawDescriptor: file_dsbackups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dsbackups_proto_goTypes,
		DependencyIndexes: file_dsbackups_proto_depIdxs,
		MessageInfos:      file_dsbackups_proto_msgTypes,
	}.Build()
	File_dsbackups_proto = out.File
	file_dsbackups_proto_rawDesc = nil
	file_dsbackups_proto_goTypes = nil
	file_dsbackups_proto_depIdxs = nil
}
