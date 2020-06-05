// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: caye.proto

package caye

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

type Status int32

const (
	Status_RUNNING Status = 0
	Status_DONE    Status = 1
	Status_FAILED  Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "RUNNING",
		1: "DONE",
		2: "FAILED",
	}
	Status_value = map[string]int32{
		"RUNNING": 0,
		"DONE":    1,
		"FAILED":  2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_caye_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_caye_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{0}
}

type BuildStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=caye.Status" json:"status,omitempty"`
}

func (x *BuildStep) Reset() {
	*x = BuildStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildStep) ProtoMessage() {}

func (x *BuildStep) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildStep.ProtoReflect.Descriptor instead.
func (*BuildStep) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{0}
}

func (x *BuildStep) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_RUNNING
}

type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status                `protobuf:"varint,1,opt,name=status,proto3,enum=caye.Status" json:"status,omitempty"`
	Steps  map[string]*BuildStep `protobuf:"bytes,2,rep,name=steps,proto3" json:"steps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{1}
}

func (x *Build) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_RUNNING
}

func (x *Build) GetSteps() map[string]*BuildStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

type StartBuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartBuildRequest) Reset() {
	*x = StartBuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBuildRequest) ProtoMessage() {}

func (x *StartBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBuildRequest.ProtoReflect.Descriptor instead.
func (*StartBuildRequest) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{2}
}

type StartBuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartBuildResponse) Reset() {
	*x = StartBuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBuildResponse) ProtoMessage() {}

func (x *StartBuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBuildResponse.ProtoReflect.Descriptor instead.
func (*StartBuildResponse) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{3}
}

func (x *StartBuildResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EndBuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error bool   `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EndBuildRequest) Reset() {
	*x = EndBuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndBuildRequest) ProtoMessage() {}

func (x *EndBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndBuildRequest.ProtoReflect.Descriptor instead.
func (*EndBuildRequest) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{4}
}

func (x *EndBuildRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EndBuildRequest) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type EndBuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndBuildResponse) Reset() {
	*x = EndBuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndBuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndBuildResponse) ProtoMessage() {}

func (x *EndBuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndBuildResponse.ProtoReflect.Descriptor instead.
func (*EndBuildResponse) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{5}
}

type StartRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *StartRunRequest) Reset() {
	*x = StartRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRunRequest) ProtoMessage() {}

func (x *StartRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRunRequest.ProtoReflect.Descriptor instead.
func (*StartRunRequest) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{6}
}

func (x *StartRunRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type StartRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartRunResponse) Reset() {
	*x = StartRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRunResponse) ProtoMessage() {}

func (x *StartRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRunResponse.ProtoReflect.Descriptor instead.
func (*StartRunResponse) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{7}
}

func (x *StartRunResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EndRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildId string `protobuf:"bytes,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	RunId   string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	Error   bool   `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *EndRunRequest) Reset() {
	*x = EndRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndRunRequest) ProtoMessage() {}

func (x *EndRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndRunRequest.ProtoReflect.Descriptor instead.
func (*EndRunRequest) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{8}
}

func (x *EndRunRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *EndRunRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *EndRunRequest) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type EndRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndRunResponse) Reset() {
	*x = EndRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndRunResponse) ProtoMessage() {}

func (x *EndRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndRunResponse.ProtoReflect.Descriptor instead.
func (*EndRunResponse) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{9}
}

type ListBuildsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBuildsRequest) Reset() {
	*x = ListBuildsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildsRequest) ProtoMessage() {}

func (x *ListBuildsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildsRequest.ProtoReflect.Descriptor instead.
func (*ListBuildsRequest) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{10}
}

type ListBuildsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Builds []*Build `protobuf:"bytes,1,rep,name=builds,proto3" json:"builds,omitempty"`
}

func (x *ListBuildsResponse) Reset() {
	*x = ListBuildsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caye_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildsResponse) ProtoMessage() {}

func (x *ListBuildsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_caye_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildsResponse.ProtoReflect.Descriptor instead.
func (*ListBuildsResponse) Descriptor() ([]byte, []int) {
	return file_caye_proto_rawDescGZIP(), []int{11}
}

func (x *ListBuildsResponse) GetBuilds() []*Build {
	if x != nil {
		return x.Builds
	}
	return nil
}

var File_caye_proto protoreflect.FileDescriptor

var file_caye_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x79, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x65, 0x70, 0x12,
	0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12,
	0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x1a, 0x49, 0x0a, 0x0a, 0x53, 0x74, 0x65, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x13,
	0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x45, 0x6e, 0x64,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x12, 0x0a, 0x10, 0x45, 0x6e, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x52,
	0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x10, 0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x06, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x06, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x2a, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x32, 0xaf, 0x02, 0x0a, 0x06, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x79, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08,
	0x45, 0x6e, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e,
	0x45, 0x6e, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x45, 0x6e, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x75, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x79,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x45, 0x6e, 0x64, 0x52, 0x75, 0x6e, 0x12, 0x13, 0x2e, 0x63,
	0x61, 0x79, 0x65, 0x2e, 0x45, 0x6e, 0x64, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x45, 0x6e, 0x64, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x17, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x79, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x69, 0x63, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x63, 0x61, 0x79,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_caye_proto_rawDescOnce sync.Once
	file_caye_proto_rawDescData = file_caye_proto_rawDesc
)

func file_caye_proto_rawDescGZIP() []byte {
	file_caye_proto_rawDescOnce.Do(func() {
		file_caye_proto_rawDescData = protoimpl.X.CompressGZIP(file_caye_proto_rawDescData)
	})
	return file_caye_proto_rawDescData
}

var file_caye_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_caye_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_caye_proto_goTypes = []interface{}{
	(Status)(0),                // 0: caye.Status
	(*BuildStep)(nil),          // 1: caye.BuildStep
	(*Build)(nil),              // 2: caye.Build
	(*StartBuildRequest)(nil),  // 3: caye.StartBuildRequest
	(*StartBuildResponse)(nil), // 4: caye.StartBuildResponse
	(*EndBuildRequest)(nil),    // 5: caye.EndBuildRequest
	(*EndBuildResponse)(nil),   // 6: caye.EndBuildResponse
	(*StartRunRequest)(nil),    // 7: caye.StartRunRequest
	(*StartRunResponse)(nil),   // 8: caye.StartRunResponse
	(*EndRunRequest)(nil),      // 9: caye.EndRunRequest
	(*EndRunResponse)(nil),     // 10: caye.EndRunResponse
	(*ListBuildsRequest)(nil),  // 11: caye.ListBuildsRequest
	(*ListBuildsResponse)(nil), // 12: caye.ListBuildsResponse
	nil,                        // 13: caye.Build.StepsEntry
}
var file_caye_proto_depIdxs = []int32{
	0,  // 0: caye.BuildStep.status:type_name -> caye.Status
	0,  // 1: caye.Build.status:type_name -> caye.Status
	13, // 2: caye.Build.steps:type_name -> caye.Build.StepsEntry
	2,  // 3: caye.ListBuildsResponse.builds:type_name -> caye.Build
	1,  // 4: caye.Build.StepsEntry.value:type_name -> caye.BuildStep
	3,  // 5: caye.Builds.StartBuild:input_type -> caye.StartBuildRequest
	5,  // 6: caye.Builds.EndBuild:input_type -> caye.EndBuildRequest
	7,  // 7: caye.Builds.StartRun:input_type -> caye.StartRunRequest
	9,  // 8: caye.Builds.EndRun:input_type -> caye.EndRunRequest
	11, // 9: caye.Builds.List:input_type -> caye.ListBuildsRequest
	4,  // 10: caye.Builds.StartBuild:output_type -> caye.StartBuildResponse
	6,  // 11: caye.Builds.EndBuild:output_type -> caye.EndBuildResponse
	8,  // 12: caye.Builds.StartRun:output_type -> caye.StartRunResponse
	10, // 13: caye.Builds.EndRun:output_type -> caye.EndRunResponse
	12, // 14: caye.Builds.List:output_type -> caye.ListBuildsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_caye_proto_init() }
func file_caye_proto_init() {
	if File_caye_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_caye_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildStep); i {
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
		file_caye_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_caye_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBuildRequest); i {
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
		file_caye_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBuildResponse); i {
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
		file_caye_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndBuildRequest); i {
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
		file_caye_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndBuildResponse); i {
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
		file_caye_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRunRequest); i {
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
		file_caye_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRunResponse); i {
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
		file_caye_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndRunRequest); i {
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
		file_caye_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndRunResponse); i {
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
		file_caye_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildsRequest); i {
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
		file_caye_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildsResponse); i {
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
			RawDescriptor: file_caye_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_caye_proto_goTypes,
		DependencyIndexes: file_caye_proto_depIdxs,
		EnumInfos:         file_caye_proto_enumTypes,
		MessageInfos:      file_caye_proto_msgTypes,
	}.Build()
	File_caye_proto = out.File
	file_caye_proto_rawDesc = nil
	file_caye_proto_goTypes = nil
	file_caye_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BuildsClient is the client API for Builds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildsClient interface {
	StartBuild(ctx context.Context, in *StartBuildRequest, opts ...grpc.CallOption) (*StartBuildResponse, error)
	EndBuild(ctx context.Context, in *EndBuildRequest, opts ...grpc.CallOption) (*EndBuildResponse, error)
	StartRun(ctx context.Context, in *StartRunRequest, opts ...grpc.CallOption) (*StartRunResponse, error)
	EndRun(ctx context.Context, in *EndRunRequest, opts ...grpc.CallOption) (*EndRunResponse, error)
	List(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsResponse, error)
}

type buildsClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildsClient(cc grpc.ClientConnInterface) BuildsClient {
	return &buildsClient{cc}
}

func (c *buildsClient) StartBuild(ctx context.Context, in *StartBuildRequest, opts ...grpc.CallOption) (*StartBuildResponse, error) {
	out := new(StartBuildResponse)
	err := c.cc.Invoke(ctx, "/caye.Builds/StartBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) EndBuild(ctx context.Context, in *EndBuildRequest, opts ...grpc.CallOption) (*EndBuildResponse, error) {
	out := new(EndBuildResponse)
	err := c.cc.Invoke(ctx, "/caye.Builds/EndBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) StartRun(ctx context.Context, in *StartRunRequest, opts ...grpc.CallOption) (*StartRunResponse, error) {
	out := new(StartRunResponse)
	err := c.cc.Invoke(ctx, "/caye.Builds/StartRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) EndRun(ctx context.Context, in *EndRunRequest, opts ...grpc.CallOption) (*EndRunResponse, error) {
	out := new(EndRunResponse)
	err := c.cc.Invoke(ctx, "/caye.Builds/EndRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) List(ctx context.Context, in *ListBuildsRequest, opts ...grpc.CallOption) (*ListBuildsResponse, error) {
	out := new(ListBuildsResponse)
	err := c.cc.Invoke(ctx, "/caye.Builds/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildsServer is the server API for Builds service.
type BuildsServer interface {
	StartBuild(context.Context, *StartBuildRequest) (*StartBuildResponse, error)
	EndBuild(context.Context, *EndBuildRequest) (*EndBuildResponse, error)
	StartRun(context.Context, *StartRunRequest) (*StartRunResponse, error)
	EndRun(context.Context, *EndRunRequest) (*EndRunResponse, error)
	List(context.Context, *ListBuildsRequest) (*ListBuildsResponse, error)
}

// UnimplementedBuildsServer can be embedded to have forward compatible implementations.
type UnimplementedBuildsServer struct {
}

func (*UnimplementedBuildsServer) StartBuild(context.Context, *StartBuildRequest) (*StartBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBuild not implemented")
}
func (*UnimplementedBuildsServer) EndBuild(context.Context, *EndBuildRequest) (*EndBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndBuild not implemented")
}
func (*UnimplementedBuildsServer) StartRun(context.Context, *StartRunRequest) (*StartRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRun not implemented")
}
func (*UnimplementedBuildsServer) EndRun(context.Context, *EndRunRequest) (*EndRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndRun not implemented")
}
func (*UnimplementedBuildsServer) List(context.Context, *ListBuildsRequest) (*ListBuildsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterBuildsServer(s *grpc.Server, srv BuildsServer) {
	s.RegisterService(&_Builds_serviceDesc, srv)
}

func _Builds_StartBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).StartBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caye.Builds/StartBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).StartBuild(ctx, req.(*StartBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_EndBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).EndBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caye.Builds/EndBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).EndBuild(ctx, req.(*EndBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_StartRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).StartRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caye.Builds/StartRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).StartRun(ctx, req.(*StartRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_EndRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).EndRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caye.Builds/EndRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).EndRun(ctx, req.(*EndRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caye.Builds/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).List(ctx, req.(*ListBuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "caye.Builds",
	HandlerType: (*BuildsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartBuild",
			Handler:    _Builds_StartBuild_Handler,
		},
		{
			MethodName: "EndBuild",
			Handler:    _Builds_EndBuild_Handler,
		},
		{
			MethodName: "StartRun",
			Handler:    _Builds_StartRun_Handler,
		},
		{
			MethodName: "EndRun",
			Handler:    _Builds_EndRun_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Builds_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "caye.proto",
}