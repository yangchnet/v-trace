// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: trans/v1/errors.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Error int32

const (
	// 未找到交易记录
	Error_RECORD_NOT_FOUND Error = 0
	// 记录重复
	Error_DUPLICATE_ERR Error = 1
	// 已存在的交易
	Error_TRANS_ALREADY_EXIST Error = 2
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0: "RECORD_NOT_FOUND",
		1: "DUPLICATE_ERR",
		2: "TRANS_ALREADY_EXIST",
	}
	Error_value = map[string]int32{
		"RECORD_NOT_FOUND":    0,
		"DUPLICATE_ERR":       1,
		"TRANS_ALREADY_EXIST": 2,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_trans_v1_errors_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_trans_v1_errors_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_trans_v1_errors_proto_rawDescGZIP(), []int{0}
}

var File_trans_v1_errors_proto protoreflect.FileDescriptor

var file_trans_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x61, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a,
	0x0a, 0x10, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x44, 0x55,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x01, 0x1a, 0x04, 0xa8,
	0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x5f, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x69, 0x69, 0x70, 0x2d, 0x69, 0x63, 0x70, 0x2f,
	0x76, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trans_v1_errors_proto_rawDescOnce sync.Once
	file_trans_v1_errors_proto_rawDescData = file_trans_v1_errors_proto_rawDesc
)

func file_trans_v1_errors_proto_rawDescGZIP() []byte {
	file_trans_v1_errors_proto_rawDescOnce.Do(func() {
		file_trans_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_trans_v1_errors_proto_rawDescData)
	})
	return file_trans_v1_errors_proto_rawDescData
}

var file_trans_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trans_v1_errors_proto_goTypes = []interface{}{
	(Error)(0), // 0: trans.v1.Error
}
var file_trans_v1_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_trans_v1_errors_proto_init() }
func file_trans_v1_errors_proto_init() {
	if File_trans_v1_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trans_v1_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trans_v1_errors_proto_goTypes,
		DependencyIndexes: file_trans_v1_errors_proto_depIdxs,
		EnumInfos:         file_trans_v1_errors_proto_enumTypes,
	}.Build()
	File_trans_v1_errors_proto = out.File
	file_trans_v1_errors_proto_rawDesc = nil
	file_trans_v1_errors_proto_goTypes = nil
	file_trans_v1_errors_proto_depIdxs = nil
}
