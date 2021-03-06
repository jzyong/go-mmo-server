// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.4.0
// source: mid.proto

package message

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

type MID int32

const (
	MID_None MID = 0
	//服务器内部消息
	MID_ServerRegisterUpdateReq MID = 110001 //服务器注册更新
	MID_ServerListReq           MID = 110003 //服务器间请求游戏服务器列表
	MID_ServerListRes           MID = 110004 //服务器间响应游戏服务器列表
)

// Enum value maps for MID.
var (
	MID_name = map[int32]string{
		0:      "None",
		110001: "ServerRegisterUpdateReq",
		110003: "ServerListReq",
		110004: "ServerListRes",
	}
	MID_value = map[string]int32{
		"None":                    0,
		"ServerRegisterUpdateReq": 110001,
		"ServerListReq":           110003,
		"ServerListRes":           110004,
	}
)

func (x MID) Enum() *MID {
	p := new(MID)
	*p = x
	return p
}

func (x MID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MID) Descriptor() protoreflect.EnumDescriptor {
	return file_mid_proto_enumTypes[0].Descriptor()
}

func (MID) Type() protoreflect.EnumType {
	return &file_mid_proto_enumTypes[0]
}

func (x MID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MID.Descriptor instead.
func (MID) EnumDescriptor() ([]byte, []int) {
	return file_mid_proto_rawDescGZIP(), []int{0}
}

var File_mid_proto protoreflect.FileDescriptor

var file_mid_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x58, 0x0a, 0x03, 0x4d, 0x49, 0x44,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x17, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x10, 0xb1, 0xdb, 0x06, 0x12, 0x13, 0x0a, 0x0d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x10, 0xb3, 0xdb, 0x06, 0x12, 0x13,
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x10,
	0xb4, 0xdb, 0x06, 0x42, 0x1e, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6d, 0x6f, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5a, 0x0b, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mid_proto_rawDescOnce sync.Once
	file_mid_proto_rawDescData = file_mid_proto_rawDesc
)

func file_mid_proto_rawDescGZIP() []byte {
	file_mid_proto_rawDescOnce.Do(func() {
		file_mid_proto_rawDescData = protoimpl.X.CompressGZIP(file_mid_proto_rawDescData)
	})
	return file_mid_proto_rawDescData
}

var file_mid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mid_proto_goTypes = []interface{}{
	(MID)(0), // 0: ProtoMessage.MID
}
var file_mid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mid_proto_init() }
func file_mid_proto_init() {
	if File_mid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mid_proto_goTypes,
		DependencyIndexes: file_mid_proto_depIdxs,
		EnumInfos:         file_mid_proto_enumTypes,
	}.Build()
	File_mid_proto = out.File
	file_mid_proto_rawDesc = nil
	file_mid_proto_goTypes = nil
	file_mid_proto_depIdxs = nil
}
