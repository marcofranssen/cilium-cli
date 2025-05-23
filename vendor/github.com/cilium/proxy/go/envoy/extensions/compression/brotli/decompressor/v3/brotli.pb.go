// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.26.1
// source: envoy/extensions/compression/brotli/decompressor/v3/brotli.proto

package decompressorv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type Brotli struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, disables "canny" ring buffer allocation strategy.
	// Ring buffer is allocated according to window size, despite the real size of the content.
	DisableRingBufferReallocation bool `protobuf:"varint,1,opt,name=disable_ring_buffer_reallocation,json=disableRingBufferReallocation,proto3" json:"disable_ring_buffer_reallocation,omitempty"`
	// Value for decompressor's next output buffer. If not set, defaults to 4096.
	ChunkSize *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
}

func (x *Brotli) Reset() {
	*x = Brotli{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brotli) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brotli) ProtoMessage() {}

func (x *Brotli) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brotli.ProtoReflect.Descriptor instead.
func (*Brotli) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDescGZIP(), []int{0}
}

func (x *Brotli) GetDisableRingBufferReallocation() bool {
	if x != nil {
		return x.DisableRingBufferReallocation
	}
	return false
}

func (x *Brotli) GetChunkSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.ChunkSize
	}
	return nil
}

var File_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto protoreflect.FileDescriptor

var file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x62,
	0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x62, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9c, 0x01, 0x0a, 0x06, 0x42, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x12, 0x47, 0x0a, 0x20, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x69, 0x6e,
	0x67, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x2a, 0x07, 0x18, 0x80, 0x80, 0x04,
	0x28, 0x80, 0x20, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x42, 0xc5,
	0x01, 0x0a, 0x41, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x72,
	0x6f, 0x74, 0x6c, 0x69, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x42, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x6c, 0x69, 0x2f,
	0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x3b,
	0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDescOnce sync.Once
	file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDescData = file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDesc
)

func file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDescGZIP() []byte {
	file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDescData)
	})
	return file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDescData
}

var file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_goTypes = []interface{}{
	(*Brotli)(nil),                 // 0: envoy.extensions.compression.brotli.decompressor.v3.Brotli
	(*wrapperspb.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
}
var file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.compression.brotli.decompressor.v3.Brotli.chunk_size:type_name -> google.protobuf.UInt32Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_init() }
func file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_init() {
	if File_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brotli); i {
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
			RawDescriptor: file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_msgTypes,
	}.Build()
	File_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto = out.File
	file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_rawDesc = nil
	file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_goTypes = nil
	file_envoy_extensions_compression_brotli_decompressor_v3_brotli_proto_depIdxs = nil
}
