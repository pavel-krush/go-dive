// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: dive/v1/dive.proto

package diveContracts

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Activity_Type int32

const (
	Activity_TYPE_UNKNOWN     Activity_Type = 0
	Activity_TYPE_BACKGROUND  Activity_Type = 1
	Activity_TYPE_GRPC_UNARY  Activity_Type = 2
	Activity_TYPE_GRPC_STREAM Activity_Type = 3
)

// Enum value maps for Activity_Type.
var (
	Activity_Type_name = map[int32]string{
		0: "TYPE_UNKNOWN",
		1: "TYPE_BACKGROUND",
		2: "TYPE_GRPC_UNARY",
		3: "TYPE_GRPC_STREAM",
	}
	Activity_Type_value = map[string]int32{
		"TYPE_UNKNOWN":     0,
		"TYPE_BACKGROUND":  1,
		"TYPE_GRPC_UNARY":  2,
		"TYPE_GRPC_STREAM": 3,
	}
)

func (x Activity_Type) Enum() *Activity_Type {
	p := new(Activity_Type)
	*p = x
	return p
}

func (x Activity_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activity_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_dive_v1_dive_proto_enumTypes[0].Descriptor()
}

func (Activity_Type) Type() protoreflect.EnumType {
	return &file_dive_v1_dive_proto_enumTypes[0]
}

func (x Activity_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activity_Type.Descriptor instead.
func (Activity_Type) EnumDescriptor() ([]byte, []int) {
	return file_dive_v1_dive_proto_rawDescGZIP(), []int{1, 0}
}

type DiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiveRequest) Reset() {
	*x = DiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dive_v1_dive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiveRequest) ProtoMessage() {}

func (x *DiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dive_v1_dive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiveRequest.ProtoReflect.Descriptor instead.
func (*DiveRequest) Descriptor() ([]byte, []int) {
	return file_dive_v1_dive_proto_rawDescGZIP(), []int{0}
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        Activity_Type          `protobuf:"varint,1,opt,name=type,proto3,enum=dive.v1.Activity_Type" json:"type,omitempty"`
	Id          string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	StartedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	Duration    *durationpb.Duration   `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Request     *anypb.Any             `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	Props       *structpb.Struct       `protobuf:"bytes,6,opt,name=props,proto3" json:"props,omitempty"`
	GrpcService string                 `protobuf:"bytes,7,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	GrpcMethod  string                 `protobuf:"bytes,8,opt,name=grpc_method,json=grpcMethod,proto3" json:"grpc_method,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dive_v1_dive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_dive_v1_dive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_dive_v1_dive_proto_rawDescGZIP(), []int{1}
}

func (x *Activity) GetType() Activity_Type {
	if x != nil {
		return x.Type
	}
	return Activity_TYPE_UNKNOWN
}

func (x *Activity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Activity) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Activity) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Activity) GetRequest() *anypb.Any {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Activity) GetProps() *structpb.Struct {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *Activity) GetGrpcService() string {
	if x != nil {
		return x.GrpcService
	}
	return ""
}

func (x *Activity) GetGrpcMethod() string {
	if x != nil {
		return x.GrpcMethod
	}
	return ""
}

type DiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activities []*Activity `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
}

func (x *DiveResponse) Reset() {
	*x = DiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dive_v1_dive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiveResponse) ProtoMessage() {}

func (x *DiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dive_v1_dive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiveResponse.ProtoReflect.Descriptor instead.
func (*DiveResponse) Descriptor() ([]byte, []int) {
	return file_dive_v1_dive_proto_rawDescGZIP(), []int{2}
}

func (x *DiveResponse) GetActivities() []*Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

var File_dive_v1_dive_proto protoreflect.FileDescriptor

var file_dive_v1_dive_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x76, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x44, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xb5, 0x03, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x64, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x22, 0x58, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x50, 0x43, 0x5f, 0x55,
	0x4e, 0x41, 0x52, 0x59, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47,
	0x52, 0x50, 0x43, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x10, 0x03, 0x22, 0x41, 0x0a, 0x0c,
	0x44, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x64, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32,
	0x54, 0x0a, 0x0b, 0x44, 0x69, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45,
	0x0a, 0x04, 0x44, 0x69, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x64, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64,
	0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x64, 0x69,
	0x76, 0x65, 0x2f, 0x76, 0x31, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x76, 0x65, 0x6c, 0x2d, 0x6b, 0x72, 0x75, 0x73, 0x68, 0x2f,
	0x64, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dive_v1_dive_proto_rawDescOnce sync.Once
	file_dive_v1_dive_proto_rawDescData = file_dive_v1_dive_proto_rawDesc
)

func file_dive_v1_dive_proto_rawDescGZIP() []byte {
	file_dive_v1_dive_proto_rawDescOnce.Do(func() {
		file_dive_v1_dive_proto_rawDescData = protoimpl.X.CompressGZIP(file_dive_v1_dive_proto_rawDescData)
	})
	return file_dive_v1_dive_proto_rawDescData
}

var file_dive_v1_dive_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dive_v1_dive_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dive_v1_dive_proto_goTypes = []interface{}{
	(Activity_Type)(0),            // 0: dive.v1.Activity.Type
	(*DiveRequest)(nil),           // 1: dive.v1.DiveRequest
	(*Activity)(nil),              // 2: dive.v1.Activity
	(*DiveResponse)(nil),          // 3: dive.v1.DiveResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 5: google.protobuf.Duration
	(*anypb.Any)(nil),             // 6: google.protobuf.Any
	(*structpb.Struct)(nil),       // 7: google.protobuf.Struct
}
var file_dive_v1_dive_proto_depIdxs = []int32{
	0, // 0: dive.v1.Activity.type:type_name -> dive.v1.Activity.Type
	4, // 1: dive.v1.Activity.started_at:type_name -> google.protobuf.Timestamp
	5, // 2: dive.v1.Activity.duration:type_name -> google.protobuf.Duration
	6, // 3: dive.v1.Activity.request:type_name -> google.protobuf.Any
	7, // 4: dive.v1.Activity.props:type_name -> google.protobuf.Struct
	2, // 5: dive.v1.DiveResponse.activities:type_name -> dive.v1.Activity
	1, // 6: dive.v1.DiveService.Dive:input_type -> dive.v1.DiveRequest
	3, // 7: dive.v1.DiveService.Dive:output_type -> dive.v1.DiveResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_dive_v1_dive_proto_init() }
func file_dive_v1_dive_proto_init() {
	if File_dive_v1_dive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dive_v1_dive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiveRequest); i {
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
		file_dive_v1_dive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_dive_v1_dive_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiveResponse); i {
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
			RawDescriptor: file_dive_v1_dive_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dive_v1_dive_proto_goTypes,
		DependencyIndexes: file_dive_v1_dive_proto_depIdxs,
		EnumInfos:         file_dive_v1_dive_proto_enumTypes,
		MessageInfos:      file_dive_v1_dive_proto_msgTypes,
	}.Build()
	File_dive_v1_dive_proto = out.File
	file_dive_v1_dive_proto_rawDesc = nil
	file_dive_v1_dive_proto_goTypes = nil
	file_dive_v1_dive_proto_depIdxs = nil
}