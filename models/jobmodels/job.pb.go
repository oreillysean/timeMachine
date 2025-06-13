// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: models/jobmodels/job.proto

package jobmodels

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

// Used to create jobs
type JobCreationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TriggerTime int64  `protobuf:"varint,2,opt,name=TriggerTime,proto3" json:"TriggerTime,omitempty"`
	Meta        []byte `protobuf:"bytes,3,opt,name=Meta,proto3" json:"Meta,omitempty"`
	Route       string `protobuf:"bytes,4,opt,name=Route,proto3" json:"Route,omitempty"`
	Collection  string `protobuf:"bytes,5,opt,name=Collection,proto3" json:"Collection,omitempty"`
}

func (x *JobCreationDetails) Reset() {
	*x = JobCreationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_jobmodels_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobCreationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCreationDetails) ProtoMessage() {}

func (x *JobCreationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_models_jobmodels_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobCreationDetails.ProtoReflect.Descriptor instead.
func (*JobCreationDetails) Descriptor() ([]byte, []int) {
	return file_models_jobmodels_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobCreationDetails) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *JobCreationDetails) GetTriggerTime() int64 {
	if x != nil {
		return x.TriggerTime
	}
	return 0
}

func (x *JobCreationDetails) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *JobCreationDetails) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *JobCreationDetails) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

// Used to fetch and delete job
type JobFetchDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=Collection,proto3" json:"Collection,omitempty"`
}

func (x *JobFetchDetails) Reset() {
	*x = JobFetchDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_jobmodels_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobFetchDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobFetchDetails) ProtoMessage() {}

func (x *JobFetchDetails) ProtoReflect() protoreflect.Message {
	mi := &file_models_jobmodels_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobFetchDetails.ProtoReflect.Descriptor instead.
func (*JobFetchDetails) Descriptor() ([]byte, []int) {
	return file_models_jobmodels_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobFetchDetails) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *JobFetchDetails) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

// Empty message because grpc doesnt allow methods without return
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_jobmodels_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_models_jobmodels_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_models_jobmodels_job_proto_rawDescGZIP(), []int{2}
}

var File_models_jobmodels_job_proto protoreflect.FileDescriptor

var file_models_jobmodels_job_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6a, 0x6f,
	0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x4a, 0x6f, 0x62, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20,
	0x0a, 0x0b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x0f, 0x4a, 0x6f,
	0x62, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x61, 0x72, 0x74, 0x68, 0x69, 0x6b, 0x72, 0x61, 0x6f, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b, 0x6a, 0x6f, 0x62,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_jobmodels_job_proto_rawDescOnce sync.Once
	file_models_jobmodels_job_proto_rawDescData = file_models_jobmodels_job_proto_rawDesc
)

func file_models_jobmodels_job_proto_rawDescGZIP() []byte {
	file_models_jobmodels_job_proto_rawDescOnce.Do(func() {
		file_models_jobmodels_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_jobmodels_job_proto_rawDescData)
	})
	return file_models_jobmodels_job_proto_rawDescData
}

var file_models_jobmodels_job_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_models_jobmodels_job_proto_goTypes = []interface{}{
	(*JobCreationDetails)(nil), // 0: jobmodels.JobCreationDetails
	(*JobFetchDetails)(nil),    // 1: jobmodels.JobFetchDetails
	(*Empty)(nil),              // 2: jobmodels.Empty
}
var file_models_jobmodels_job_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_jobmodels_job_proto_init() }
func file_models_jobmodels_job_proto_init() {
	if File_models_jobmodels_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_jobmodels_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobCreationDetails); i {
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
		file_models_jobmodels_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobFetchDetails); i {
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
		file_models_jobmodels_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_models_jobmodels_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_jobmodels_job_proto_goTypes,
		DependencyIndexes: file_models_jobmodels_job_proto_depIdxs,
		MessageInfos:      file_models_jobmodels_job_proto_msgTypes,
	}.Build()
	File_models_jobmodels_job_proto = out.File
	file_models_jobmodels_job_proto_rawDesc = nil
	file_models_jobmodels_job_proto_goTypes = nil
	file_models_jobmodels_job_proto_depIdxs = nil
}
