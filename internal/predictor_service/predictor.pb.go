// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: predictor.proto

package predictor_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Greeting struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Greeting      string                 `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Greeting) Reset() {
	*x = Greeting{}
	mi := &file_predictor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Greeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeting) ProtoMessage() {}

func (x *Greeting) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeting.ProtoReflect.Descriptor instead.
func (*Greeting) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{0}
}

func (x *Greeting) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (x *Greeting) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Greeting      *Greeting              `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	From          string                 `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_predictor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

func (x *HelloRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Greeting      string                 `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_predictor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{2}
}

func (x *HelloResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

type PredictImageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bucket        string                 `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PredictImageRequest) Reset() {
	*x = PredictImageRequest{}
	mi := &file_predictor_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PredictImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictImageRequest) ProtoMessage() {}

func (x *PredictImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictImageRequest.ProtoReflect.Descriptor instead.
func (*PredictImageRequest) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{3}
}

func (x *PredictImageRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *PredictImageRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PredictImageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClassName     string                 `protobuf:"bytes,1,opt,name=className,proto3" json:"className,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PredictImageResponse) Reset() {
	*x = PredictImageResponse{}
	mi := &file_predictor_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PredictImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictImageResponse) ProtoMessage() {}

func (x *PredictImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_predictor_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictImageResponse.ProtoReflect.Descriptor instead.
func (*PredictImageResponse) Descriptor() ([]byte, []int) {
	return file_predictor_proto_rawDescGZIP(), []int{4}
}

func (x *PredictImageResponse) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

var File_predictor_proto protoreflect.FileDescriptor

const file_predictor_proto_rawDesc = "" +
	"\n" +
	"\x0fpredictor.proto\":\n" +
	"\bGreeting\x12\x1a\n" +
	"\bgreeting\x18\x01 \x01(\tR\bgreeting\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"I\n" +
	"\fHelloRequest\x12%\n" +
	"\bgreeting\x18\x01 \x01(\v2\t.GreetingR\bgreeting\x12\x12\n" +
	"\x04from\x18\x02 \x01(\tR\x04from\"+\n" +
	"\rHelloResponse\x12\x1a\n" +
	"\bgreeting\x18\x01 \x01(\tR\bgreeting\"?\n" +
	"\x13PredictImageRequest\x12\x16\n" +
	"\x06bucket\x18\x01 \x01(\tR\x06bucket\x12\x10\n" +
	"\x03key\x18\x02 \x01(\tR\x03key\"4\n" +
	"\x14PredictImageResponse\x12\x1c\n" +
	"\tclassName\x18\x01 \x01(\tR\tclassName2p\n" +
	"\tPredictor\x12&\n" +
	"\x05Hello\x12\r.HelloRequest\x1a\x0e.HelloResponse\x12;\n" +
	"\fPredictImage\x12\x14.PredictImageRequest\x1a\x15.PredictImageResponseB-Z+www.github.com/Maevlava/Matatani/Backend/mlb\x06proto3"

var (
	file_predictor_proto_rawDescOnce sync.Once
	file_predictor_proto_rawDescData []byte
)

func file_predictor_proto_rawDescGZIP() []byte {
	file_predictor_proto_rawDescOnce.Do(func() {
		file_predictor_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_predictor_proto_rawDesc), len(file_predictor_proto_rawDesc)))
	})
	return file_predictor_proto_rawDescData
}

var file_predictor_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_predictor_proto_goTypes = []any{
	(*Greeting)(nil),             // 0: Greeting
	(*HelloRequest)(nil),         // 1: HelloRequest
	(*HelloResponse)(nil),        // 2: HelloResponse
	(*PredictImageRequest)(nil),  // 3: PredictImageRequest
	(*PredictImageResponse)(nil), // 4: PredictImageResponse
}
var file_predictor_proto_depIdxs = []int32{
	0, // 0: HelloRequest.greeting:type_name -> Greeting
	1, // 1: Predictor.Hello:input_type -> HelloRequest
	3, // 2: Predictor.PredictImage:input_type -> PredictImageRequest
	2, // 3: Predictor.Hello:output_type -> HelloResponse
	4, // 4: Predictor.PredictImage:output_type -> PredictImageResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_predictor_proto_init() }
func file_predictor_proto_init() {
	if File_predictor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_predictor_proto_rawDesc), len(file_predictor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_predictor_proto_goTypes,
		DependencyIndexes: file_predictor_proto_depIdxs,
		MessageInfos:      file_predictor_proto_msgTypes,
	}.Build()
	File_predictor_proto = out.File
	file_predictor_proto_goTypes = nil
	file_predictor_proto_depIdxs = nil
}
