// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: basic-test/serverstream-api.proto

package basic_test

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//Data Structure
type NotificationMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationMessage) Reset() {
	*x = NotificationMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_test_serverstream_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationMessage) ProtoMessage() {}

func (x *NotificationMessage) ProtoReflect() protoreflect.Message {
	mi := &file_basic_test_serverstream_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationMessage.ProtoReflect.Descriptor instead.
func (*NotificationMessage) Descriptor() ([]byte, []int) {
	return file_basic_test_serverstream_api_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *NotificationMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//Request and Response Payload
type NotificationMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requester string `protobuf:"bytes,1,opt,name=requester,proto3" json:"requester,omitempty"`
	Criteria  string `protobuf:"bytes,2,opt,name=criteria,proto3" json:"criteria,omitempty"`
}

func (x *NotificationMessageRequest) Reset() {
	*x = NotificationMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_test_serverstream_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationMessageRequest) ProtoMessage() {}

func (x *NotificationMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basic_test_serverstream_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationMessageRequest.ProtoReflect.Descriptor instead.
func (*NotificationMessageRequest) Descriptor() ([]byte, []int) {
	return file_basic_test_serverstream_api_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationMessageRequest) GetRequester() string {
	if x != nil {
		return x.Requester
	}
	return ""
}

func (x *NotificationMessageRequest) GetCriteria() string {
	if x != nil {
		return x.Criteria
	}
	return ""
}

type NotificationMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode      string               `protobuf:"bytes,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	ResponsePayload *NotificationMessage `protobuf:"bytes,2,opt,name=responsePayload,proto3" json:"responsePayload,omitempty"`
}

func (x *NotificationMessageResponse) Reset() {
	*x = NotificationMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_test_serverstream_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationMessageResponse) ProtoMessage() {}

func (x *NotificationMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basic_test_serverstream_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationMessageResponse.ProtoReflect.Descriptor instead.
func (*NotificationMessageResponse) Descriptor() ([]byte, []int) {
	return file_basic_test_serverstream_api_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationMessageResponse) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

func (x *NotificationMessageResponse) GetResponsePayload() *NotificationMessage {
	if x != nil {
		return x.ResponsePayload
	}
	return nil
}

var File_basic_test_serverstream_api_proto protoreflect.FileDescriptor

var file_basic_test_serverstream_api_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22,
	0x47, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x56, 0x0a, 0x1a, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x22, 0x88, 0x01, 0x0a, 0x1b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x49, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x88, 0x01, 0x0a, 0x1a,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x13, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x26, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x17, 0x5a, 0x15, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2d,
	0x74, 0x65, 0x73, 0x74, 0x3b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basic_test_serverstream_api_proto_rawDescOnce sync.Once
	file_basic_test_serverstream_api_proto_rawDescData = file_basic_test_serverstream_api_proto_rawDesc
)

func file_basic_test_serverstream_api_proto_rawDescGZIP() []byte {
	file_basic_test_serverstream_api_proto_rawDescOnce.Do(func() {
		file_basic_test_serverstream_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_test_serverstream_api_proto_rawDescData)
	})
	return file_basic_test_serverstream_api_proto_rawDescData
}

var file_basic_test_serverstream_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_basic_test_serverstream_api_proto_goTypes = []interface{}{
	(*NotificationMessage)(nil),         // 0: basic_test.NotificationMessage
	(*NotificationMessageRequest)(nil),  // 1: basic_test.NotificationMessageRequest
	(*NotificationMessageResponse)(nil), // 2: basic_test.NotificationMessageResponse
}
var file_basic_test_serverstream_api_proto_depIdxs = []int32{
	0, // 0: basic_test.NotificationMessageResponse.responsePayload:type_name -> basic_test.NotificationMessage
	1, // 1: basic_test.NotificationMessageService.NotificationMessage:input_type -> basic_test.NotificationMessageRequest
	2, // 2: basic_test.NotificationMessageService.NotificationMessage:output_type -> basic_test.NotificationMessageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_basic_test_serverstream_api_proto_init() }
func file_basic_test_serverstream_api_proto_init() {
	if File_basic_test_serverstream_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basic_test_serverstream_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationMessage); i {
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
		file_basic_test_serverstream_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationMessageRequest); i {
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
		file_basic_test_serverstream_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationMessageResponse); i {
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
			RawDescriptor: file_basic_test_serverstream_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basic_test_serverstream_api_proto_goTypes,
		DependencyIndexes: file_basic_test_serverstream_api_proto_depIdxs,
		MessageInfos:      file_basic_test_serverstream_api_proto_msgTypes,
	}.Build()
	File_basic_test_serverstream_api_proto = out.File
	file_basic_test_serverstream_api_proto_rawDesc = nil
	file_basic_test_serverstream_api_proto_goTypes = nil
	file_basic_test_serverstream_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationMessageServiceClient is the client API for NotificationMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationMessageServiceClient interface {
	//Server-Streaming API
	NotificationMessage(ctx context.Context, in *NotificationMessageRequest, opts ...grpc.CallOption) (NotificationMessageService_NotificationMessageClient, error)
}

type notificationMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationMessageServiceClient(cc grpc.ClientConnInterface) NotificationMessageServiceClient {
	return &notificationMessageServiceClient{cc}
}

func (c *notificationMessageServiceClient) NotificationMessage(ctx context.Context, in *NotificationMessageRequest, opts ...grpc.CallOption) (NotificationMessageService_NotificationMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotificationMessageService_serviceDesc.Streams[0], "/basic_test.NotificationMessageService/NotificationMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationMessageServiceNotificationMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationMessageService_NotificationMessageClient interface {
	Recv() (*NotificationMessageResponse, error)
	grpc.ClientStream
}

type notificationMessageServiceNotificationMessageClient struct {
	grpc.ClientStream
}

func (x *notificationMessageServiceNotificationMessageClient) Recv() (*NotificationMessageResponse, error) {
	m := new(NotificationMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationMessageServiceServer is the server API for NotificationMessageService service.
type NotificationMessageServiceServer interface {
	//Server-Streaming API
	NotificationMessage(*NotificationMessageRequest, NotificationMessageService_NotificationMessageServer) error
}

// UnimplementedNotificationMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationMessageServiceServer struct {
}

func (*UnimplementedNotificationMessageServiceServer) NotificationMessage(*NotificationMessageRequest, NotificationMessageService_NotificationMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method NotificationMessage not implemented")
}

func RegisterNotificationMessageServiceServer(s *grpc.Server, srv NotificationMessageServiceServer) {
	s.RegisterService(&_NotificationMessageService_serviceDesc, srv)
}

func _NotificationMessageService_NotificationMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationMessageServiceServer).NotificationMessage(m, &notificationMessageServiceNotificationMessageServer{stream})
}

type NotificationMessageService_NotificationMessageServer interface {
	Send(*NotificationMessageResponse) error
	grpc.ServerStream
}

type notificationMessageServiceNotificationMessageServer struct {
	grpc.ServerStream
}

func (x *notificationMessageServiceNotificationMessageServer) Send(m *NotificationMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NotificationMessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "basic_test.NotificationMessageService",
	HandlerType: (*NotificationMessageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NotificationMessage",
			Handler:       _NotificationMessageService_NotificationMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "basic-test/serverstream-api.proto",
}