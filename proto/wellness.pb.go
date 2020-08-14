// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: wellness.proto

package wellness

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListLessonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DayOfWeek    *wrappers.Int32Value  `protobuf:"bytes,1,opt,name=dayOfWeek,proto3" json:"dayOfWeek,omitempty"`
	BeginHours   *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=beginHours,proto3" json:"beginHours,omitempty"`
	BeginMinutes *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=beginMinutes,proto3" json:"beginMinutes,omitempty"`
	TypeSearch   *wrappers.StringValue `protobuf:"bytes,4,opt,name=typeSearch,proto3" json:"typeSearch,omitempty"`
}

func (x *ListLessonsRequest) Reset() {
	*x = ListLessonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wellness_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLessonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLessonsRequest) ProtoMessage() {}

func (x *ListLessonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wellness_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLessonsRequest.ProtoReflect.Descriptor instead.
func (*ListLessonsRequest) Descriptor() ([]byte, []int) {
	return file_wellness_proto_rawDescGZIP(), []int{0}
}

func (x *ListLessonsRequest) GetDayOfWeek() *wrappers.Int32Value {
	if x != nil {
		return x.DayOfWeek
	}
	return nil
}

func (x *ListLessonsRequest) GetBeginHours() *wrappers.Int32Value {
	if x != nil {
		return x.BeginHours
	}
	return nil
}

func (x *ListLessonsRequest) GetBeginMinutes() *wrappers.Int32Value {
	if x != nil {
		return x.BeginMinutes
	}
	return nil
}

func (x *ListLessonsRequest) GetTypeSearch() *wrappers.StringValue {
	if x != nil {
		return x.TypeSearch
	}
	return nil
}

type ListRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRoomsRequest) Reset() {
	*x = ListRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wellness_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsRequest) ProtoMessage() {}

func (x *ListRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wellness_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsRequest.ProtoReflect.Descriptor instead.
func (*ListRoomsRequest) Descriptor() ([]byte, []int) {
	return file_wellness_proto_rawDescGZIP(), []int{1}
}

type GetRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int32 `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
}

func (x *GetRoomRequest) Reset() {
	*x = GetRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wellness_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomRequest) ProtoMessage() {}

func (x *GetRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wellness_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomRequest.ProtoReflect.Descriptor instead.
func (*GetRoomRequest) Descriptor() ([]byte, []int) {
	return file_wellness_proto_rawDescGZIP(), []int{2}
}

func (x *GetRoomRequest) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type RoomsArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *RoomsArray) Reset() {
	*x = RoomsArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wellness_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomsArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomsArray) ProtoMessage() {}

func (x *RoomsArray) ProtoReflect() protoreflect.Message {
	mi := &file_wellness_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomsArray.ProtoReflect.Descriptor instead.
func (*RoomsArray) Descriptor() ([]byte, []int) {
	return file_wellness_proto_rawDescGZIP(), []int{3}
}

func (x *RoomsArray) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wellness_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_wellness_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_wellness_proto_rawDescGZIP(), []int{4}
}

func (x *Room) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LessonsArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lessons []*Lesson `protobuf:"bytes,1,rep,name=lessons,proto3" json:"lessons,omitempty"`
}

func (x *LessonsArray) Reset() {
	*x = LessonsArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wellness_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LessonsArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LessonsArray) ProtoMessage() {}

func (x *LessonsArray) ProtoReflect() protoreflect.Message {
	mi := &file_wellness_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LessonsArray.ProtoReflect.Descriptor instead.
func (*LessonsArray) Descriptor() ([]byte, []int) {
	return file_wellness_proto_rawDescGZIP(), []int{5}
}

func (x *LessonsArray) GetLessons() []*Lesson {
	if x != nil {
		return x.Lessons
	}
	return nil
}

type Lesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DayOfWeek    int32  `protobuf:"varint,3,opt,name=dayOfWeek,proto3" json:"dayOfWeek,omitempty"`
	BeginHours   int32  `protobuf:"varint,4,opt,name=beginHours,proto3" json:"beginHours,omitempty"`
	BeginMinutes int32  `protobuf:"varint,5,opt,name=beginMinutes,proto3" json:"beginMinutes,omitempty"`
	EndHours     int32  `protobuf:"varint,6,opt,name=endHours,proto3" json:"endHours,omitempty"`
	EndMinutes   int32  `protobuf:"varint,7,opt,name=endMinutes,proto3" json:"endMinutes,omitempty"`
	Duration     int32  `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	NbPlace      int32  `protobuf:"varint,9,opt,name=NbPlace,proto3" json:"NbPlace,omitempty"`
}

func (x *Lesson) Reset() {
	*x = Lesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wellness_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lesson) ProtoMessage() {}

func (x *Lesson) ProtoReflect() protoreflect.Message {
	mi := &file_wellness_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lesson.ProtoReflect.Descriptor instead.
func (*Lesson) Descriptor() ([]byte, []int) {
	return file_wellness_proto_rawDescGZIP(), []int{6}
}

func (x *Lesson) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lesson) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lesson) GetDayOfWeek() int32 {
	if x != nil {
		return x.DayOfWeek
	}
	return 0
}

func (x *Lesson) GetBeginHours() int32 {
	if x != nil {
		return x.BeginHours
	}
	return 0
}

func (x *Lesson) GetBeginMinutes() int32 {
	if x != nil {
		return x.BeginMinutes
	}
	return 0
}

func (x *Lesson) GetEndHours() int32 {
	if x != nil {
		return x.EndHours
	}
	return 0
}

func (x *Lesson) GetEndMinutes() int32 {
	if x != nil {
		return x.EndMinutes
	}
	return 0
}

func (x *Lesson) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Lesson) GetNbPlace() int32 {
	if x != nil {
		return x.NbPlace
	}
	return 0
}

var File_wellness_proto protoreflect.FileDescriptor

var file_wellness_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8e, 0x05, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x8e, 0x01, 0x0a, 0x09, 0x64, 0x61, 0x79,
	0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x53, 0x92, 0x41, 0x47, 0x32, 0x45,
	0x50, 0x65, 0x72, 0x6d, 0x65, 0x74, 0x20, 0x64, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x72, 0x65,
	0x72, 0x20, 0x73, 0x75, 0x72, 0x20, 0x6c, 0x65, 0x20, 0x6a, 0x6f, 0x75, 0x72, 0x20, 0x64, 0x65,
	0x20, 0x6c, 0x61, 0x20, 0x73, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x65, 0x20, 0x28, 0x30, 0x20, 0x64,
	0x69, 0x6d, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x20, 0xc3, 0xa0, 0x20, 0x36, 0x20, 0x53, 0x61, 0x6d,
	0x65, 0x64, 0x69, 0x29, 0x2e, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x06, 0x28, 0x00, 0x52, 0x09,
	0x64, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x87, 0x01, 0x0a, 0x0a, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x4a, 0x92, 0x41, 0x3e,
	0x32, 0x3c, 0x50, 0x65, 0x72, 0x6d, 0x65, 0x74, 0x20, 0x64, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x74,
	0x72, 0x65, 0x72, 0x20, 0x73, 0x75, 0x72, 0x20, 0x6c, 0x27, 0x68, 0x65, 0x75, 0x72, 0x65, 0x20,
	0x64, 0x65, 0x20, 0x64, 0xc3, 0xa9, 0x62, 0x75, 0x74, 0x20, 0x64, 0x75, 0x20, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x27, 0x20, 0x28, 0x30, 0x20, 0xc3, 0xa0, 0x20, 0x32, 0x33, 0x29, 0x2e, 0xfa, 0x42,
	0x06, 0x1a, 0x04, 0x18, 0x17, 0x28, 0x00, 0x52, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x0c, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x4e, 0x92, 0x41, 0x42, 0x32, 0x40, 0x50, 0x65,
	0x72, 0x6d, 0x65, 0x74, 0x20, 0x64, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x72, 0x65, 0x72, 0x20,
	0x73, 0x75, 0x72, 0x20, 0x6c, 0x65, 0x73, 0x20, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x20,
	0x64, 0x65, 0x20, 0x64, 0xc3, 0xa9, 0x62, 0x75, 0x74, 0x20, 0x64, 0x75, 0x20, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x27, 0x20, 0x28, 0x30, 0x20, 0xc3, 0xa0, 0x20, 0x35, 0x39, 0x29, 0x2e, 0xfa, 0x42,
	0x06, 0x1a, 0x04, 0x18, 0x3b, 0x28, 0x00, 0x52, 0x0c, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x4d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0xca, 0x01, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x8b, 0x01, 0x92, 0x41, 0x6d, 0x32, 0x6b,
	0x50, 0x65, 0x72, 0x6d, 0x65, 0x74, 0x20, 0x64, 0x65, 0x20, 0x64, 0xc3, 0xa9, 0x66, 0x69, 0x6e,
	0x69, 0x72, 0x20, 0x6c, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x64, 0x65, 0x20, 0x66, 0x69,
	0x6c, 0x74, 0x72, 0x61, 0x67, 0x65, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x27, 0x68, 0x65, 0x75, 0x72,
	0x65, 0x20, 0x65, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x6d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x20, 0x28, 0x6c, 0x74, 0x2c, 0x20, 0x67, 0x74, 0x2c, 0x20, 0x67, 0x74, 0x65, 0x2c, 0x20,
	0x6c, 0x74, 0x65, 0x2c, 0x20, 0x65, 0x71, 0x29, 0x2e, 0x20, 0x50, 0x61, 0x72, 0x20, 0x64, 0xc3,
	0xa9, 0x66, 0x61, 0x75, 0x74, 0x20, 0x65, 0x71, 0x75, 0x61, 0x6c, 0xfa, 0x42, 0x18, 0x72, 0x16,
	0x52, 0x02, 0x6c, 0x74, 0x52, 0x02, 0x67, 0x74, 0x52, 0x03, 0x67, 0x74, 0x65, 0x52, 0x03, 0x6c,
	0x74, 0x65, 0x52, 0x02, 0x65, 0x71, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x22, 0x4b, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x3d,
	0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x17, 0x92,
	0x41, 0x14, 0x32, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x6f, 0x66,
	0x20, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x2a, 0x0a,
	0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x0c, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x45, 0x0a, 0x07, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x77, 0x65, 0x6c,
	0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x42, 0x19, 0x92, 0x41,
	0x16, 0x32, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x22, 0xb0, 0x04, 0x0a, 0x06, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1a, 0x92, 0x41, 0x17, 0x32, 0x15, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x61, 0x6e, 0x74, 0x20, 0x64, 0x75, 0x20, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x2e, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x92, 0x41, 0x0f, 0x32, 0x0d, 0x4e, 0x6f, 0x6d, 0x20,
	0x64, 0x75, 0x20, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x5d, 0x0a, 0x09, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x3f, 0x92, 0x41, 0x3c, 0x32, 0x3a, 0x4a, 0x6f, 0x75, 0x72, 0x20, 0x64, 0x65,
	0x20, 0x6c, 0x61, 0x20, 0x73, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x65, 0x20, 0x64, 0x75, 0x20, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x20, 0x28, 0x30, 0x20, 0x70, 0x6f, 0x75, 0x72, 0x20, 0x64, 0x69, 0x6d,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x20, 0xc3, 0xa0, 0x20, 0x36, 0x20, 0x53, 0x61, 0x6d, 0x65, 0x64,
	0x69, 0x29, 0x2e, 0x52, 0x09, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x3e,
	0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x1e, 0x92, 0x41, 0x1b, 0x32, 0x19, 0x48, 0x65, 0x75, 0x72, 0x65, 0x20, 0x64,
	0x65, 0x20, 0x64, 0xc3, 0xa9, 0x62, 0x75, 0x74, 0x20, 0x64, 0x75, 0x20, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x2e, 0x52, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x43,
	0x0a, 0x0c, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x1f, 0x92, 0x41, 0x1c, 0x32, 0x1a, 0x4d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x20, 0x64, 0x65, 0x20, 0x64, 0xc3, 0xa9, 0x62, 0x75, 0x74, 0x20, 0x64, 0x75, 0x20, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x2e, 0x52, 0x0c, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1b, 0x92, 0x41, 0x18, 0x32, 0x16, 0x48, 0x65, 0x75, 0x72,
	0x65, 0x20, 0x64, 0x65, 0x20, 0x66, 0x69, 0x6e, 0x20, 0x64, 0x75, 0x20, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x2e, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x0a,
	0x65, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x1c, 0x92, 0x41, 0x19, 0x32, 0x17, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x20, 0x64, 0x65,
	0x20, 0x66, 0x69, 0x6e, 0x20, 0x64, 0x75, 0x20, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x52, 0x0a,
	0x65, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x21, 0x92, 0x41,
	0x1e, 0x32, 0x1c, 0x44, 0x75, 0x72, 0xc3, 0xa9, 0x65, 0x20, 0x64, 0x75, 0x20, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x20, 0x28, 0x65, 0x6e, 0x20, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x29, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x4e, 0x62, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1e, 0x92, 0x41, 0x1b, 0x32,
	0x19, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x20, 0x64, 0x65, 0x20, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x20, 0x64, 0x75, 0x20, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x52, 0x07, 0x4e, 0x62, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x32, 0xe8, 0x04, 0x0a, 0x0f, 0x57, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x1a, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x92, 0x41,
	0x32, 0x0a, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x1a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x12, 0xaf, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x4f, 0x66, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x18, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x69, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x7d, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x92, 0x41, 0x40, 0x0a, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x20, 0x6f, 0x66,
	0x20, 0x61, 0x20, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c,
	0x6c, 0x20, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20,
	0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x12, 0x87, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x18, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x77, 0x65,
	0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x22, 0x52, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2f,
	0x7b, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x7d, 0x92, 0x41, 0x34, 0x0a, 0x05, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x12, 0x0c, 0x47, 0x65, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x6f, 0x6f, 0x6d,
	0x1a, 0x1d, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x72, 0x6f, 0x6f, 0x6d, 0x20, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x50, 0x46, 0x53, 0x2e, 0x12,
	0x8c, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x12,
	0x1c, 0x2e, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x92, 0x41,
	0x2d, 0x0a, 0x07, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x6f, 0x66, 0x20, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x1a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x92,
	0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65,
	0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2d, 0x77, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x92, 0x41, 0x54, 0x12,
	0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x02, 0x01, 0x72, 0x47, 0x0a, 0x11, 0x57, 0x65,
	0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x32, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x2d, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x77, 0x65, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73,
	0x2d, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wellness_proto_rawDescOnce sync.Once
	file_wellness_proto_rawDescData = file_wellness_proto_rawDesc
)

func file_wellness_proto_rawDescGZIP() []byte {
	file_wellness_proto_rawDescOnce.Do(func() {
		file_wellness_proto_rawDescData = protoimpl.X.CompressGZIP(file_wellness_proto_rawDescData)
	})
	return file_wellness_proto_rawDescData
}

var file_wellness_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_wellness_proto_goTypes = []interface{}{
	(*ListLessonsRequest)(nil),   // 0: wellness.ListLessonsRequest
	(*ListRoomsRequest)(nil),     // 1: wellness.ListRoomsRequest
	(*GetRoomRequest)(nil),       // 2: wellness.GetRoomRequest
	(*RoomsArray)(nil),           // 3: wellness.RoomsArray
	(*Room)(nil),                 // 4: wellness.Room
	(*LessonsArray)(nil),         // 5: wellness.LessonsArray
	(*Lesson)(nil),               // 6: wellness.Lesson
	(*wrappers.Int32Value)(nil),  // 7: google.protobuf.Int32Value
	(*wrappers.StringValue)(nil), // 8: google.protobuf.StringValue
}
var file_wellness_proto_depIdxs = []int32{
	7,  // 0: wellness.ListLessonsRequest.dayOfWeek:type_name -> google.protobuf.Int32Value
	7,  // 1: wellness.ListLessonsRequest.beginHours:type_name -> google.protobuf.Int32Value
	7,  // 2: wellness.ListLessonsRequest.beginMinutes:type_name -> google.protobuf.Int32Value
	8,  // 3: wellness.ListLessonsRequest.typeSearch:type_name -> google.protobuf.StringValue
	4,  // 4: wellness.RoomsArray.rooms:type_name -> wellness.Room
	6,  // 5: wellness.LessonsArray.lessons:type_name -> wellness.Lesson
	1,  // 6: wellness.WellnessService.ListRooms:input_type -> wellness.ListRoomsRequest
	2,  // 7: wellness.WellnessService.GetLessonsOfRoom:input_type -> wellness.GetRoomRequest
	2,  // 8: wellness.WellnessService.GetRoom:input_type -> wellness.GetRoomRequest
	0,  // 9: wellness.WellnessService.ListLessons:input_type -> wellness.ListLessonsRequest
	3,  // 10: wellness.WellnessService.ListRooms:output_type -> wellness.RoomsArray
	5,  // 11: wellness.WellnessService.GetLessonsOfRoom:output_type -> wellness.LessonsArray
	4,  // 12: wellness.WellnessService.GetRoom:output_type -> wellness.Room
	5,  // 13: wellness.WellnessService.ListLessons:output_type -> wellness.LessonsArray
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_wellness_proto_init() }
func file_wellness_proto_init() {
	if File_wellness_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wellness_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLessonsRequest); i {
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
		file_wellness_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsRequest); i {
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
		file_wellness_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomRequest); i {
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
		file_wellness_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomsArray); i {
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
		file_wellness_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_wellness_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LessonsArray); i {
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
		file_wellness_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lesson); i {
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
			RawDescriptor: file_wellness_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wellness_proto_goTypes,
		DependencyIndexes: file_wellness_proto_depIdxs,
		MessageInfos:      file_wellness_proto_msgTypes,
	}.Build()
	File_wellness_proto = out.File
	file_wellness_proto_rawDesc = nil
	file_wellness_proto_goTypes = nil
	file_wellness_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WellnessServiceClient is the client API for WellnessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WellnessServiceClient interface {
	ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*RoomsArray, error)
	GetLessonsOfRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*LessonsArray, error)
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*Room, error)
	ListLessons(ctx context.Context, in *ListLessonsRequest, opts ...grpc.CallOption) (*LessonsArray, error)
}

type wellnessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWellnessServiceClient(cc grpc.ClientConnInterface) WellnessServiceClient {
	return &wellnessServiceClient{cc}
}

func (c *wellnessServiceClient) ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*RoomsArray, error) {
	out := new(RoomsArray)
	err := c.cc.Invoke(ctx, "/wellness.WellnessService/ListRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wellnessServiceClient) GetLessonsOfRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*LessonsArray, error) {
	out := new(LessonsArray)
	err := c.cc.Invoke(ctx, "/wellness.WellnessService/GetLessonsOfRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wellnessServiceClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/wellness.WellnessService/GetRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wellnessServiceClient) ListLessons(ctx context.Context, in *ListLessonsRequest, opts ...grpc.CallOption) (*LessonsArray, error) {
	out := new(LessonsArray)
	err := c.cc.Invoke(ctx, "/wellness.WellnessService/ListLessons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WellnessServiceServer is the server API for WellnessService service.
type WellnessServiceServer interface {
	ListRooms(context.Context, *ListRoomsRequest) (*RoomsArray, error)
	GetLessonsOfRoom(context.Context, *GetRoomRequest) (*LessonsArray, error)
	GetRoom(context.Context, *GetRoomRequest) (*Room, error)
	ListLessons(context.Context, *ListLessonsRequest) (*LessonsArray, error)
}

// UnimplementedWellnessServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWellnessServiceServer struct {
}

func (*UnimplementedWellnessServiceServer) ListRooms(context.Context, *ListRoomsRequest) (*RoomsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRooms not implemented")
}
func (*UnimplementedWellnessServiceServer) GetLessonsOfRoom(context.Context, *GetRoomRequest) (*LessonsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLessonsOfRoom not implemented")
}
func (*UnimplementedWellnessServiceServer) GetRoom(context.Context, *GetRoomRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (*UnimplementedWellnessServiceServer) ListLessons(context.Context, *ListLessonsRequest) (*LessonsArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLessons not implemented")
}

func RegisterWellnessServiceServer(s *grpc.Server, srv WellnessServiceServer) {
	s.RegisterService(&_WellnessService_serviceDesc, srv)
}

func _WellnessService_ListRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WellnessServiceServer).ListRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wellness.WellnessService/ListRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WellnessServiceServer).ListRooms(ctx, req.(*ListRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WellnessService_GetLessonsOfRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WellnessServiceServer).GetLessonsOfRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wellness.WellnessService/GetLessonsOfRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WellnessServiceServer).GetLessonsOfRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WellnessService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WellnessServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wellness.WellnessService/GetRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WellnessServiceServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WellnessService_ListLessons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLessonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WellnessServiceServer).ListLessons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wellness.WellnessService/ListLessons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WellnessServiceServer).ListLessons(ctx, req.(*ListLessonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WellnessService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wellness.WellnessService",
	HandlerType: (*WellnessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRooms",
			Handler:    _WellnessService_ListRooms_Handler,
		},
		{
			MethodName: "GetLessonsOfRoom",
			Handler:    _WellnessService_GetLessonsOfRoom_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _WellnessService_GetRoom_Handler,
		},
		{
			MethodName: "ListLessons",
			Handler:    _WellnessService_ListLessons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wellness.proto",
}
