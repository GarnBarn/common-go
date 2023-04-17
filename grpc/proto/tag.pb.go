// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: grpc/proto/tag.proto

package common_go

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

type TagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId             int32 `protobuf:"varint,1,opt,name=tagId,proto3" json:"tagId,omitempty"`
	ConsealPrivateKey bool  `protobuf:"varint,2,opt,name=consealPrivateKey,proto3" json:"consealPrivateKey,omitempty"`
}

func (x *TagRequest) Reset() {
	*x = TagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagRequest) ProtoMessage() {}

func (x *TagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagRequest.ProtoReflect.Descriptor instead.
func (*TagRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_tag_proto_rawDescGZIP(), []int{0}
}

func (x *TagRequest) GetTagId() int32 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *TagRequest) GetConsealPrivateKey() bool {
	if x != nil {
		return x.ConsealPrivateKey
	}
	return false
}

type TagExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExists bool `protobuf:"varint,1,opt,name=isExists,proto3" json:"isExists,omitempty"`
}

func (x *TagExistsResponse) Reset() {
	*x = TagExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagExistsResponse) ProtoMessage() {}

func (x *TagExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagExistsResponse.ProtoReflect.Descriptor instead.
func (*TagExistsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_tag_proto_rawDescGZIP(), []int{1}
}

func (x *TagExistsResponse) GetIsExists() bool {
	if x != nil {
		return x.IsExists
	}
	return false
}

type TagPublic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author        string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Color         string   `protobuf:"bytes,4,opt,name=Color,proto3" json:"Color,omitempty"`
	ReminderTime  []int32  `protobuf:"varint,5,rep,packed,name=reminderTime,proto3" json:"reminderTime,omitempty"`
	Subscriber    []string `protobuf:"bytes,6,rep,name=Subscriber,proto3" json:"Subscriber,omitempty"`
	SecretKeyTotp string   `protobuf:"bytes,7,opt,name=SecretKeyTotp,proto3" json:"SecretKeyTotp,omitempty"`
}

func (x *TagPublic) Reset() {
	*x = TagPublic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagPublic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagPublic) ProtoMessage() {}

func (x *TagPublic) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagPublic.ProtoReflect.Descriptor instead.
func (*TagPublic) Descriptor() ([]byte, []int) {
	return file_grpc_proto_tag_proto_rawDescGZIP(), []int{2}
}

func (x *TagPublic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TagPublic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagPublic) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *TagPublic) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TagPublic) GetReminderTime() []int32 {
	if x != nil {
		return x.ReminderTime
	}
	return nil
}

func (x *TagPublic) GetSubscriber() []string {
	if x != nil {
		return x.Subscriber
	}
	return nil
}

func (x *TagPublic) GetSecretKeyTotp() string {
	if x != nil {
		return x.SecretKeyTotp
	}
	return ""
}

var File_grpc_proto_tag_proto protoreflect.FileDescriptor

var file_grpc_proto_tag_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x61, 0x67, 0x22, 0x50, 0x0a, 0x0a, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12,
	0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x2f, 0x0a,
	0x11, 0x54, 0x61, 0x67, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0xc7,
	0x01, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x54,
	0x6f, 0x74, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x54, 0x6f, 0x74, 0x70, 0x32, 0x6c, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12,
	0x38, 0x0a, 0x0b, 0x49, 0x73, 0x54, 0x61, 0x67, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x0f,
	0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x67, 0x12, 0x0f, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x72, 0x6e, 0x42, 0x61, 0x72, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_tag_proto_rawDescOnce sync.Once
	file_grpc_proto_tag_proto_rawDescData = file_grpc_proto_tag_proto_rawDesc
)

func file_grpc_proto_tag_proto_rawDescGZIP() []byte {
	file_grpc_proto_tag_proto_rawDescOnce.Do(func() {
		file_grpc_proto_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_tag_proto_rawDescData)
	})
	return file_grpc_proto_tag_proto_rawDescData
}

var file_grpc_proto_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_proto_tag_proto_goTypes = []interface{}{
	(*TagRequest)(nil),        // 0: tag.TagRequest
	(*TagExistsResponse)(nil), // 1: tag.TagExistsResponse
	(*TagPublic)(nil),         // 2: tag.TagPublic
}
var file_grpc_proto_tag_proto_depIdxs = []int32{
	0, // 0: tag.Tag.IsTagExists:input_type -> tag.TagRequest
	0, // 1: tag.Tag.GetTag:input_type -> tag.TagRequest
	1, // 2: tag.Tag.IsTagExists:output_type -> tag.TagExistsResponse
	2, // 3: tag.Tag.GetTag:output_type -> tag.TagPublic
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_tag_proto_init() }
func file_grpc_proto_tag_proto_init() {
	if File_grpc_proto_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagRequest); i {
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
		file_grpc_proto_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagExistsResponse); i {
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
		file_grpc_proto_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagPublic); i {
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
			RawDescriptor: file_grpc_proto_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_tag_proto_goTypes,
		DependencyIndexes: file_grpc_proto_tag_proto_depIdxs,
		MessageInfos:      file_grpc_proto_tag_proto_msgTypes,
	}.Build()
	File_grpc_proto_tag_proto = out.File
	file_grpc_proto_tag_proto_rawDesc = nil
	file_grpc_proto_tag_proto_goTypes = nil
	file_grpc_proto_tag_proto_depIdxs = nil
}
