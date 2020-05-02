// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: service.proto

package sampler

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image      []byte  `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	ImageShape []int32 `protobuf:"varint,2,rep,packed,name=image_shape,json=imageShape,proto3" json:"image_shape,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Reply) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Reply) GetImageShape() []int32 {
	if x != nil {
		return x.ImageShape
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x05,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x68, 0x61, 0x70, 0x65, 0x32, 0x4f, 0x0a, 0x0d,
	0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x72, 0x12, 0x3e, 0x0a,
	0x0b, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_proto_goTypes = []interface{}{
	(*Reply)(nil),       // 0: camera_sampler.Reply
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_service_proto_depIdxs = []int32{
	1, // 0: camera_sampler.CameraSampler.SampleImage:input_type -> google.protobuf.Empty
	0, // 1: camera_sampler.CameraSampler.SampleImage:output_type -> camera_sampler.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CameraSamplerClient is the client API for CameraSampler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CameraSamplerClient interface {
	SampleImage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Reply, error)
}

type cameraSamplerClient struct {
	cc grpc.ClientConnInterface
}

func NewCameraSamplerClient(cc grpc.ClientConnInterface) CameraSamplerClient {
	return &cameraSamplerClient{cc}
}

func (c *cameraSamplerClient) SampleImage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/camera_sampler.CameraSampler/SampleImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CameraSamplerServer is the server API for CameraSampler service.
type CameraSamplerServer interface {
	SampleImage(context.Context, *empty.Empty) (*Reply, error)
}

// UnimplementedCameraSamplerServer can be embedded to have forward compatible implementations.
type UnimplementedCameraSamplerServer struct {
}

func (*UnimplementedCameraSamplerServer) SampleImage(context.Context, *empty.Empty) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SampleImage not implemented")
}

func RegisterCameraSamplerServer(s *grpc.Server, srv CameraSamplerServer) {
	s.RegisterService(&_CameraSampler_serviceDesc, srv)
}

func _CameraSampler_SampleImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CameraSamplerServer).SampleImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/camera_sampler.CameraSampler/SampleImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CameraSamplerServer).SampleImage(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CameraSampler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "camera_sampler.CameraSampler",
	HandlerType: (*CameraSamplerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SampleImage",
			Handler:    _CameraSampler_SampleImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}