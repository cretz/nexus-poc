// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: nexus/frontend/v1/config.proto

package frontendpb

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nexus_frontend_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_nexus_frontend_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_nexus_frontend_v1_config_proto_rawDescGZIP(), []int{0}
}

var File_nexus_frontend_v1_config_proto protoreflect.FileDescriptor

var file_nexus_frontend_v1_config_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x22, 0x08, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nexus_frontend_v1_config_proto_rawDescOnce sync.Once
	file_nexus_frontend_v1_config_proto_rawDescData = file_nexus_frontend_v1_config_proto_rawDesc
)

func file_nexus_frontend_v1_config_proto_rawDescGZIP() []byte {
	file_nexus_frontend_v1_config_proto_rawDescOnce.Do(func() {
		file_nexus_frontend_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_nexus_frontend_v1_config_proto_rawDescData)
	})
	return file_nexus_frontend_v1_config_proto_rawDescData
}

var file_nexus_frontend_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nexus_frontend_v1_config_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: nexus.frontend.v1.Config
}
var file_nexus_frontend_v1_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nexus_frontend_v1_config_proto_init() }
func file_nexus_frontend_v1_config_proto_init() {
	if File_nexus_frontend_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nexus_frontend_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_nexus_frontend_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nexus_frontend_v1_config_proto_goTypes,
		DependencyIndexes: file_nexus_frontend_v1_config_proto_depIdxs,
		MessageInfos:      file_nexus_frontend_v1_config_proto_msgTypes,
	}.Build()
	File_nexus_frontend_v1_config_proto = out.File
	file_nexus_frontend_v1_config_proto_rawDesc = nil
	file_nexus_frontend_v1_config_proto_goTypes = nil
	file_nexus_frontend_v1_config_proto_depIdxs = nil
}