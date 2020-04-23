// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: basic-test/basic.proto

package basic_test

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

var File_basic_test_basic_proto protoreflect.FileDescriptor

var file_basic_test_basic_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x32,
	0x10, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x17, 0x5a, 0x15, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3b,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_basic_test_basic_proto_goTypes = []interface{}{}
var file_basic_test_basic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_basic_test_basic_proto_init() }
func file_basic_test_basic_proto_init() {
	if File_basic_test_basic_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basic_test_basic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basic_test_basic_proto_goTypes,
		DependencyIndexes: file_basic_test_basic_proto_depIdxs,
	}.Build()
	File_basic_test_basic_proto = out.File
	file_basic_test_basic_proto_rawDesc = nil
	file_basic_test_basic_proto_goTypes = nil
	file_basic_test_basic_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterServiceClient interface {
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

// GreeterServiceServer is the server API for GreeterService service.
type GreeterServiceServer interface {
}

// UnimplementedGreeterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func RegisterGreeterServiceServer(s *grpc.Server, srv GreeterServiceServer) {
	s.RegisterService(&_GreeterService_serviceDesc, srv)
}

var _GreeterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "basic-test/basic.proto",
}
