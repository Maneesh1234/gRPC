// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: greet_s7_avg/greetpb/greet.proto

package greetpb

import (
	context "context"
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

type PrimeNumberDecomposition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PrimeNumberDecomposition) Reset() {
	*x = PrimeNumberDecomposition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecomposition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecomposition) ProtoMessage() {}

func (x *PrimeNumberDecomposition) ProtoReflect() protoreflect.Message {
	mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecomposition.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecomposition) Descriptor() ([]byte, []int) {
	return file_greet_s7_avg_greetpb_greet_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeNumberDecomposition) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type PrimeNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeNumberDecomposition *PrimeNumberDecomposition `protobuf:"bytes,1,opt,name=primeNumberDecomposition,proto3" json:"primeNumberDecomposition,omitempty"`
}

func (x *PrimeNumberRequest) Reset() {
	*x = PrimeNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberRequest) ProtoMessage() {}

func (x *PrimeNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return file_greet_s7_avg_greetpb_greet_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeNumberRequest) GetPrimeNumberDecomposition() *PrimeNumberDecomposition {
	if x != nil {
		return x.PrimeNumberDecomposition
	}
	return nil
}

type PrimeNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeNumberResponse) Reset() {
	*x = PrimeNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberResponse) ProtoMessage() {}

func (x *PrimeNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return file_greet_s7_avg_greetpb_greet_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeNumberResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type ComputeAverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ComputeAverageRequest) Reset() {
	*x = ComputeAverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageRequest) ProtoMessage() {}

func (x *ComputeAverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageRequest.ProtoReflect.Descriptor instead.
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return file_greet_s7_avg_greetpb_greet_proto_rawDescGZIP(), []int{3}
}

func (x *ComputeAverageRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ComputeAverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avg int64 `protobuf:"varint,1,opt,name=avg,proto3" json:"avg,omitempty"`
}

func (x *ComputeAverageResponse) Reset() {
	*x = ComputeAverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageResponse) ProtoMessage() {}

func (x *ComputeAverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_s7_avg_greetpb_greet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageResponse.ProtoReflect.Descriptor instead.
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return file_greet_s7_avg_greetpb_greet_proto_rawDescGZIP(), []int{4}
}

func (x *ComputeAverageResponse) GetAvg() int64 {
	if x != nil {
		return x.Avg
	}
	return 0
}

var File_greet_s7_avg_greetpb_greet_proto protoreflect.FileDescriptor

var file_greet_s7_avg_greetpb_greet_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x37, 0x5f, 0x61, 0x76, 0x67, 0x2f, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x37, 0x5f, 0x61, 0x76, 0x67,
	0x22, 0x2c, 0x0a, 0x18, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44,
	0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x78,
	0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x62, 0x0a, 0x18, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73,
	0x37, 0x5f, 0x61, 0x76, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x6d,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x29, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x22, 0x2a, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x76, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x76, 0x67, 0x32, 0xcf,
	0x01, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x6e,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73,
	0x37, 0x5f, 0x61, 0x76, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x5f, 0x73, 0x37, 0x5f, 0x61, 0x76, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x5f, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x37, 0x5f, 0x61, 0x76, 0x67,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73,
	0x37, 0x5f, 0x61, 0x76, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x42, 0x16, 0x5a, 0x14, 0x67, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x37, 0x5f, 0x61, 0x76, 0x67,
	0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_s7_avg_greetpb_greet_proto_rawDescOnce sync.Once
	file_greet_s7_avg_greetpb_greet_proto_rawDescData = file_greet_s7_avg_greetpb_greet_proto_rawDesc
)

func file_greet_s7_avg_greetpb_greet_proto_rawDescGZIP() []byte {
	file_greet_s7_avg_greetpb_greet_proto_rawDescOnce.Do(func() {
		file_greet_s7_avg_greetpb_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_s7_avg_greetpb_greet_proto_rawDescData)
	})
	return file_greet_s7_avg_greetpb_greet_proto_rawDescData
}

var file_greet_s7_avg_greetpb_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_greet_s7_avg_greetpb_greet_proto_goTypes = []interface{}{
	(*PrimeNumberDecomposition)(nil), // 0: greet_s7_avg.PrimeNumberDecomposition
	(*PrimeNumberRequest)(nil),       // 1: greet_s7_avg.PrimeNumberRequest
	(*PrimeNumberResponse)(nil),      // 2: greet_s7_avg.PrimeNumberResponse
	(*ComputeAverageRequest)(nil),    // 3: greet_s7_avg.ComputeAverageRequest
	(*ComputeAverageResponse)(nil),   // 4: greet_s7_avg.ComputeAverageResponse
}
var file_greet_s7_avg_greetpb_greet_proto_depIdxs = []int32{
	0, // 0: greet_s7_avg.PrimeNumberRequest.primeNumberDecomposition:type_name -> greet_s7_avg.PrimeNumberDecomposition
	1, // 1: greet_s7_avg.CalculatorService.PrimeManyTimes:input_type -> greet_s7_avg.PrimeNumberRequest
	3, // 2: greet_s7_avg.CalculatorService.ComputeAverage:input_type -> greet_s7_avg.ComputeAverageRequest
	2, // 3: greet_s7_avg.CalculatorService.PrimeManyTimes:output_type -> greet_s7_avg.PrimeNumberResponse
	4, // 4: greet_s7_avg.CalculatorService.ComputeAverage:output_type -> greet_s7_avg.ComputeAverageResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_greet_s7_avg_greetpb_greet_proto_init() }
func file_greet_s7_avg_greetpb_greet_proto_init() {
	if File_greet_s7_avg_greetpb_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_s7_avg_greetpb_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecomposition); i {
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
		file_greet_s7_avg_greetpb_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberRequest); i {
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
		file_greet_s7_avg_greetpb_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberResponse); i {
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
		file_greet_s7_avg_greetpb_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageRequest); i {
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
		file_greet_s7_avg_greetpb_greet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageResponse); i {
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
			RawDescriptor: file_greet_s7_avg_greetpb_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greet_s7_avg_greetpb_greet_proto_goTypes,
		DependencyIndexes: file_greet_s7_avg_greetpb_greet_proto_depIdxs,
		MessageInfos:      file_greet_s7_avg_greetpb_greet_proto_msgTypes,
	}.Build()
	File_greet_s7_avg_greetpb_greet_proto = out.File
	file_greet_s7_avg_greetpb_greet_proto_rawDesc = nil
	file_greet_s7_avg_greetpb_greet_proto_goTypes = nil
	file_greet_s7_avg_greetpb_greet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Server Streaming
	PrimeManyTimes(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeManyTimesClient, error)
	// Client Streaming
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) PrimeManyTimes(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/greet_s7_avg.CalculatorService/PrimeManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeManyTimesClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeManyTimesClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeManyTimesClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/greet_s7_avg.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Server Streaming
	PrimeManyTimes(*PrimeNumberRequest, CalculatorService_PrimeManyTimesServer) error
	// Client Streaming
	ComputeAverage(CalculatorService_ComputeAverageServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) PrimeManyTimes(*PrimeNumberRequest, CalculatorService_PrimeManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeManyTimes not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_PrimeManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeManyTimes(m, &calculatorServicePrimeManyTimesServer{stream})
}

type CalculatorService_PrimeManyTimesServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeManyTimesServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeManyTimesServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet_s7_avg.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeManyTimes",
			Handler:       _CalculatorService_PrimeManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet_s7_avg/greetpb/greet.proto",
}