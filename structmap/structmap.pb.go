// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: structmap/structmap.proto

package structmap

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_structmap_structmap_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50000,
		Name:          "structmap.enable",
		Tag:           "varint,50000,opt,name=enable",
		Filename:      "structmap/structmap.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional bool enable = 50000;
	E_Enable = &file_structmap_structmap_proto_extTypes[0]
)

var File_structmap_structmap_proto protoreflect.FileDescriptor

var file_structmap_structmap_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6d, 0x61, 0x70, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x6d, 0x61, 0x70, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x39, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x72, 0x65, 0x76, 0x65, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x6d, 0x61, 0x70, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6d, 0x61, 0x70, 0x3b, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x6d, 0x61, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_structmap_structmap_proto_goTypes = []interface{}{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
}
var file_structmap_structmap_proto_depIdxs = []int32{
	0, // 0: structmap.enable:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_structmap_structmap_proto_init() }
func file_structmap_structmap_proto_init() {
	if File_structmap_structmap_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_structmap_structmap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_structmap_structmap_proto_goTypes,
		DependencyIndexes: file_structmap_structmap_proto_depIdxs,
		ExtensionInfos:    file_structmap_structmap_proto_extTypes,
	}.Build()
	File_structmap_structmap_proto = out.File
	file_structmap_structmap_proto_rawDesc = nil
	file_structmap_structmap_proto_goTypes = nil
	file_structmap_structmap_proto_depIdxs = nil
}
