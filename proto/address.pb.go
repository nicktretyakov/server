// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: address.proto

package proto

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

type AddressStatus int32

const (
	AddressStatus_ADDRESS_STATUS           AddressStatus = 0
	AddressStatus_INITIAL                  AddressStatus = 1 // Начальный
	AddressStatus_DECLINED                 AddressStatus = 2 // Отклонен
	AddressStatus_ON_REGISTRATION          AddressStatus = 3 //На регистрации
	AddressStatus_CONFIRMED                AddressStatus = 4 // Утвержден
	AddressStatus_DONE                     AddressStatus = 5 // Завершен
	AddressStatus_ON_AGREE                 AddressStatus = 6 // На согласовании
	AddressStatus_FINALIZE_ON_REGISTRATION AddressStatus = 7 // Завершение
	AddressStatus_FINALIZE_ON_AGREE        AddressStatus = 8 // Завершение согласуется
	AddressStatus_FINALIZE_DECLINED        AddressStatus = 9 // завершение
)

// Enum value maps for AddressStatus.
var (
	AddressStatus_name = map[int32]string{
		0: "ADDRESS_STATUS",
		1: "INITIAL",
		2: "DECLINED",
		3: "ON_REGISTRATION",
		4: "CONFIRMED",
		5: "DONE",
		6: "ON_AGREE",
		7: "FINALIZE_ON_REGISTRATION",
		8: "FINALIZE_ON_AGREE",
		9: "FINALIZE_DECLINED",
	}
	AddressStatus_value = map[string]int32{
		"ADDRESS_STATUS":           0,
		"INITIAL":                  1,
		"DECLINED":                 2,
		"ON_REGISTRATION":          3,
		"CONFIRMED":                4,
		"DONE":                     5,
		"ON_AGREE":                 6,
		"FINALIZE_ON_REGISTRATION": 7,
		"FINALIZE_ON_AGREE":        8,
		"FINALIZE_DECLINED":        9,
	}
)

func (x AddressStatus) Enum() *AddressStatus {
	p := new(AddressStatus)
	*p = x
	return p
}

func (x AddressStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_address_proto_enumTypes[0].Descriptor()
}

func (AddressStatus) Type() protoreflect.EnumType {
	return &file_address_proto_enumTypes[0]
}

func (x AddressStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressStatus.Descriptor instead.
func (AddressStatus) EnumDescriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

type AddressType int32

const (
	AddressType_UNKNOWN_OBJECT AddressType = 0
	AddressType_BOOKING        AddressType = 1
	AddressType_ROOM           AddressType = 2
)

// Enum value maps for AddressType.
var (
	AddressType_name = map[int32]string{
		0: "UNKNOWN_OBJECT",
		1: "BOOKING",
		2: "ROOM",
	}
	AddressType_value = map[string]int32{
		"UNKNOWN_OBJECT": 0,
		"BOOKING":        1,
		"ROOM":           2,
	}
)

func (x AddressType) Enum() *AddressType {
	p := new(AddressType)
	*p = x
	return p
}

func (x AddressType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressType) Descriptor() protoreflect.EnumDescriptor {
	return file_address_proto_enumTypes[1].Descriptor()
}

func (AddressType) Type() protoreflect.EnumType {
	return &file_address_proto_enumTypes[1]
}

func (x AddressType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressType.Descriptor instead.
func (AddressType) EnumDescriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{1}
}

type AddressState int32

const (
	AddressState_UNKNOWN_STATE AddressState = 0
	AddressState_PUBLISHED     AddressState = 1
	AddressState_ARCHIVED      AddressState = 2
)

// Enum value maps for AddressState.
var (
	AddressState_name = map[int32]string{
		0: "UNKNOWN_STATE",
		1: "PUBLISHED",
		2: "ARCHIVED",
	}
	AddressState_value = map[string]int32{
		"UNKNOWN_STATE": 0,
		"PUBLISHED":     1,
		"ARCHIVED":      2,
	}
)

func (x AddressState) Enum() *AddressState {
	p := new(AddressState)
	*p = x
	return p
}

func (x AddressState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressState) Descriptor() protoreflect.EnumDescriptor {
	return file_address_proto_enumTypes[2].Descriptor()
}

func (AddressState) Type() protoreflect.EnumType {
	return &file_address_proto_enumTypes[2]
}

func (x AddressState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressState.Descriptor instead.
func (AddressState) EnumDescriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{2}
}

type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	FileName string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Size     uint64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *Attachment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attachment) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Attachment) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Attachment) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Timeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Timeline) Reset() {
	*x = Timeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeline) ProtoMessage() {}

func (x *Timeline) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeline.ProtoReflect.Descriptor instead.
func (*Timeline) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{1}
}

func (x *Timeline) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *Timeline) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{2}
}

func (x *Link) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Link) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Link) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x5e, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x32, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x42, 0x0a, 0x04,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2a, 0xc6, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52,
	0x4d, 0x45, 0x44, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x05, 0x12,
	0x0c, 0x0a, 0x08, 0x4f, 0x4e, 0x5f, 0x41, 0x47, 0x52, 0x45, 0x45, 0x10, 0x06, 0x12, 0x1c, 0x0a,
	0x18, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x47,
	0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x46,
	0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x41, 0x47, 0x52, 0x45, 0x45,
	0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x5f, 0x44,
	0x45, 0x43, 0x4c, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x09, 0x2a, 0x38, 0x0a, 0x0b, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x4f,
	0x4d, 0x10, 0x02, 0x2a, 0x3e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53,
	0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45,
	0x44, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData = file_address_proto_rawDesc
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_proto_rawDescData)
	})
	return file_address_proto_rawDescData
}

var file_address_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_address_proto_goTypes = []interface{}{
	(AddressStatus)(0), // 0: booking.AddressStatus
	(AddressType)(0),   // 1: booking.AddressType
	(AddressState)(0),  // 2: booking.AddressState
	(*Attachment)(nil), // 3: booking.Attachment
	(*Timeline)(nil),   // 4: booking.Timeline
	(*Link)(nil),       // 5: booking.Link
}
var file_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
		file_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeline); i {
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
		file_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
			RawDescriptor: file_address_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		EnumInfos:         file_address_proto_enumTypes,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_rawDesc = nil
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}
