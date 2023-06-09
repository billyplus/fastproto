// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: test/oneof.proto

package test

import (
	_ "github.com/billyplus/fastproto/options"
	pb "github.com/billyplus/fastproto/test/pb"
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

// full proto
type OneOfProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	//
	// Types that are assignable to TestOneof:
	//	*OneOfProto_VInt32
	//	*OneOfProto_VInt64
	//	*OneOfProto_VUint32
	//	*OneOfProto_VUint64
	//	*OneOfProto_VString
	//	*OneOfProto_VBytes
	//	*OneOfProto_VBool
	//	*OneOfProto_SInt32
	//	*OneOfProto_SInt64
	//	*OneOfProto_Fixed32
	//	*OneOfProto_Fixed64
	//	*OneOfProto_Sfixed32
	//	*OneOfProto_Sfixed64
	//	*OneOfProto_Float32
	//	*OneOfProto_Float64
	//	*OneOfProto_MActor
	//	*OneOfProto_Outer
	TestOneof isOneOfProto_TestOneof `protobuf_oneof:"test_oneof"`
}

func (x *OneOfProto) Reset() {
	*x = OneOfProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_oneof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfProto) ProtoMessage() {}

func (x *OneOfProto) ProtoReflect() protoreflect.Message {
	mi := &file_test_oneof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfProto.ProtoReflect.Descriptor instead.
func (*OneOfProto) Descriptor() ([]byte, []int) {
	return file_test_oneof_proto_rawDescGZIP(), []int{0}
}

func (m *OneOfProto) GetTestOneof() isOneOfProto_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (x *OneOfProto) GetVInt32() int32 {
	if x, ok := x.GetTestOneof().(*OneOfProto_VInt32); ok {
		return x.VInt32
	}
	return 0
}

func (x *OneOfProto) GetVInt64() int64 {
	if x, ok := x.GetTestOneof().(*OneOfProto_VInt64); ok {
		return x.VInt64
	}
	return 0
}

func (x *OneOfProto) GetVUint32() uint32 {
	if x, ok := x.GetTestOneof().(*OneOfProto_VUint32); ok {
		return x.VUint32
	}
	return 0
}

func (x *OneOfProto) GetVUint64() uint64 {
	if x, ok := x.GetTestOneof().(*OneOfProto_VUint64); ok {
		return x.VUint64
	}
	return 0
}

func (x *OneOfProto) GetVString() string {
	if x, ok := x.GetTestOneof().(*OneOfProto_VString); ok {
		return x.VString
	}
	return ""
}

func (x *OneOfProto) GetVBytes() []byte {
	if x, ok := x.GetTestOneof().(*OneOfProto_VBytes); ok {
		return x.VBytes
	}
	return nil
}

func (x *OneOfProto) GetVBool() bool {
	if x, ok := x.GetTestOneof().(*OneOfProto_VBool); ok {
		return x.VBool
	}
	return false
}

func (x *OneOfProto) GetSInt32() int32 {
	if x, ok := x.GetTestOneof().(*OneOfProto_SInt32); ok {
		return x.SInt32
	}
	return 0
}

func (x *OneOfProto) GetSInt64() int64 {
	if x, ok := x.GetTestOneof().(*OneOfProto_SInt64); ok {
		return x.SInt64
	}
	return 0
}

func (x *OneOfProto) GetFixed32() uint32 {
	if x, ok := x.GetTestOneof().(*OneOfProto_Fixed32); ok {
		return x.Fixed32
	}
	return 0
}

func (x *OneOfProto) GetFixed64() uint64 {
	if x, ok := x.GetTestOneof().(*OneOfProto_Fixed64); ok {
		return x.Fixed64
	}
	return 0
}

func (x *OneOfProto) GetSfixed32() int32 {
	if x, ok := x.GetTestOneof().(*OneOfProto_Sfixed32); ok {
		return x.Sfixed32
	}
	return 0
}

func (x *OneOfProto) GetSfixed64() int64 {
	if x, ok := x.GetTestOneof().(*OneOfProto_Sfixed64); ok {
		return x.Sfixed64
	}
	return 0
}

func (x *OneOfProto) GetFloat32() float32 {
	if x, ok := x.GetTestOneof().(*OneOfProto_Float32); ok {
		return x.Float32
	}
	return 0
}

func (x *OneOfProto) GetFloat64() float64 {
	if x, ok := x.GetTestOneof().(*OneOfProto_Float64); ok {
		return x.Float64
	}
	return 0
}

func (x *OneOfProto) GetMActor() *OtherMessage {
	if x, ok := x.GetTestOneof().(*OneOfProto_MActor); ok {
		return x.MActor
	}
	return nil
}

func (x *OneOfProto) GetOuter() *pb.OuterMsg {
	if x, ok := x.GetTestOneof().(*OneOfProto_Outer); ok {
		return x.Outer
	}
	return nil
}

type isOneOfProto_TestOneof interface {
	isOneOfProto_TestOneof()
}

type OneOfProto_VInt32 struct {
	VInt32 int32 `protobuf:"varint,1,opt,name=v_int32,json=vInt32,proto3,oneof"`
}

type OneOfProto_VInt64 struct {
	VInt64 int64 `protobuf:"varint,2,opt,name=v_int64,json=vInt64,proto3,oneof"`
}

type OneOfProto_VUint32 struct {
	VUint32 uint32 `protobuf:"varint,3,opt,name=v_uint32,json=vUint32,proto3,oneof"`
}

type OneOfProto_VUint64 struct {
	VUint64 uint64 `protobuf:"varint,4,opt,name=v_uint64,json=vUint64,proto3,oneof"`
}

type OneOfProto_VString struct {
	VString string `protobuf:"bytes,5,opt,name=v_string,json=vString,proto3,oneof"`
}

type OneOfProto_VBytes struct {
	VBytes []byte `protobuf:"bytes,23,opt,name=v_bytes,json=vBytes,proto3,oneof"`
}

type OneOfProto_VBool struct {
	VBool bool `protobuf:"varint,6,opt,name=v_bool,json=vBool,proto3,oneof"`
}

type OneOfProto_SInt32 struct {
	//
	SInt32 int32 `protobuf:"zigzag32,7,opt,name=s_int32,json=sInt32,proto3,oneof"`
}

type OneOfProto_SInt64 struct {
	SInt64 int64 `protobuf:"zigzag64,8,opt,name=s_int64,json=sInt64,proto3,oneof"`
}

type OneOfProto_Fixed32 struct {
	Fixed32 uint32 `protobuf:"fixed32,9,opt,name=fixed32,proto3,oneof"`
}

type OneOfProto_Fixed64 struct {
	Fixed64 uint64 `protobuf:"fixed64,10,opt,name=fixed64,proto3,oneof"`
}

type OneOfProto_Sfixed32 struct {
	Sfixed32 int32 `protobuf:"fixed32,11,opt,name=sfixed32,proto3,oneof"`
}

type OneOfProto_Sfixed64 struct {
	Sfixed64 int64 `protobuf:"fixed64,12,opt,name=sfixed64,proto3,oneof"`
}

type OneOfProto_Float32 struct {
	//
	Float32 float32 `protobuf:"fixed32,20,opt,name=float32,proto3,oneof"`
}

type OneOfProto_Float64 struct {
	Float64 float64 `protobuf:"fixed64,21,opt,name=float64,proto3,oneof"`
}

type OneOfProto_MActor struct {
	MActor *OtherMessage `protobuf:"bytes,231,opt,name=m_actor,json=mActor,proto3,oneof"`
}

type OneOfProto_Outer struct {
	Outer *pb.OuterMsg `protobuf:"bytes,281,opt,name=outer,proto3,oneof"`
}

func (*OneOfProto_VInt32) isOneOfProto_TestOneof() {}

func (*OneOfProto_VInt64) isOneOfProto_TestOneof() {}

func (*OneOfProto_VUint32) isOneOfProto_TestOneof() {}

func (*OneOfProto_VUint64) isOneOfProto_TestOneof() {}

func (*OneOfProto_VString) isOneOfProto_TestOneof() {}

func (*OneOfProto_VBytes) isOneOfProto_TestOneof() {}

func (*OneOfProto_VBool) isOneOfProto_TestOneof() {}

func (*OneOfProto_SInt32) isOneOfProto_TestOneof() {}

func (*OneOfProto_SInt64) isOneOfProto_TestOneof() {}

func (*OneOfProto_Fixed32) isOneOfProto_TestOneof() {}

func (*OneOfProto_Fixed64) isOneOfProto_TestOneof() {}

func (*OneOfProto_Sfixed32) isOneOfProto_TestOneof() {}

func (*OneOfProto_Sfixed64) isOneOfProto_TestOneof() {}

func (*OneOfProto_Float32) isOneOfProto_TestOneof() {}

func (*OneOfProto_Float64) isOneOfProto_TestOneof() {}

func (*OneOfProto_MActor) isOneOfProto_TestOneof() {}

func (*OneOfProto_Outer) isOneOfProto_TestOneof() {}

var File_test_oneof_proto protoreflect.FileDescriptor

var file_test_oneof_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x94, 0x04, 0x0a, 0x0a, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x0a, 0x07, 0x76, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x06, 0x76, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x07, 0x76,
	0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06,
	0x76, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1b, 0x0a, 0x08, 0x76, 0x5f, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x07, 0x76, 0x55, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x1b, 0x0a, 0x08, 0x76, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x07, 0x76, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x1b, 0x0a, 0x08, 0x76, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x76, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a,
	0x07, 0x76, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x06, 0x76, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x06, 0x76, 0x5f, 0x62, 0x6f,
	0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x05, 0x76, 0x42, 0x6f, 0x6f,
	0x6c, 0x12, 0x19, 0x0a, 0x07, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x11, 0x48, 0x00, 0x52, 0x06, 0x73, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x07,
	0x73, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x12, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x07, 0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x06, 0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12,
	0x1c, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0f, 0x48, 0x00, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1c, 0x0a,
	0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10, 0x48,
	0x00, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x07, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x07,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x07, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x36, 0x34, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x36, 0x34, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x5f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0xe7,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x41, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x99, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x29, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x79, 0x70, 0x6c, 0x75, 0x73,
	0x2f, 0x66, 0x61, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0xb0,
	0xac, 0x1b, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_oneof_proto_rawDescOnce sync.Once
	file_test_oneof_proto_rawDescData = file_test_oneof_proto_rawDesc
)

func file_test_oneof_proto_rawDescGZIP() []byte {
	file_test_oneof_proto_rawDescOnce.Do(func() {
		file_test_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_oneof_proto_rawDescData)
	})
	return file_test_oneof_proto_rawDescData
}

var file_test_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_test_oneof_proto_goTypes = []interface{}{
	(*OneOfProto)(nil),   // 0: test.OneOfProto
	(*OtherMessage)(nil), // 1: test.OtherMessage
	(*pb.OuterMsg)(nil),  // 2: pb.OuterMsg
}
var file_test_oneof_proto_depIdxs = []int32{
	1, // 0: test.OneOfProto.m_actor:type_name -> test.OtherMessage
	2, // 1: test.OneOfProto.outer:type_name -> pb.OuterMsg
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_test_oneof_proto_init() }
func file_test_oneof_proto_init() {
	if File_test_oneof_proto != nil {
		return
	}
	file_test_msg_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_test_oneof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfProto); i {
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
	file_test_oneof_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OneOfProto_VInt32)(nil),
		(*OneOfProto_VInt64)(nil),
		(*OneOfProto_VUint32)(nil),
		(*OneOfProto_VUint64)(nil),
		(*OneOfProto_VString)(nil),
		(*OneOfProto_VBytes)(nil),
		(*OneOfProto_VBool)(nil),
		(*OneOfProto_SInt32)(nil),
		(*OneOfProto_SInt64)(nil),
		(*OneOfProto_Fixed32)(nil),
		(*OneOfProto_Fixed64)(nil),
		(*OneOfProto_Sfixed32)(nil),
		(*OneOfProto_Sfixed64)(nil),
		(*OneOfProto_Float32)(nil),
		(*OneOfProto_Float64)(nil),
		(*OneOfProto_MActor)(nil),
		(*OneOfProto_Outer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_oneof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_oneof_proto_goTypes,
		DependencyIndexes: file_test_oneof_proto_depIdxs,
		MessageInfos:      file_test_oneof_proto_msgTypes,
	}.Build()
	File_test_oneof_proto = out.File
	file_test_oneof_proto_rawDesc = nil
	file_test_oneof_proto_goTypes = nil
	file_test_oneof_proto_depIdxs = nil
}
