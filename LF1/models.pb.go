// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: models.proto

package routeguide

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

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username       string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password       string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	FollowingNames []string `protobuf:"bytes,3,rep,name=FollowingNames,proto3" json:"FollowingNames,omitempty"`
	FollowersNames []string `protobuf:"bytes,4,rep,name=FollowersNames,proto3" json:"FollowersNames,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *Users) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Users) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Users) GetFollowingNames() []string {
	if x != nil {
		return x.FollowingNames
	}
	return nil
}

func (x *Users) GetFollowersNames() []string {
	if x != nil {
		return x.FollowersNames
	}
	return nil
}

type Posts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Timestamp string   `protobuf:"bytes,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Posts     []string `protobuf:"bytes,3,rep,name=Posts,proto3" json:"Posts,omitempty"`
}

func (x *Posts) Reset() {
	*x = Posts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Posts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Posts) ProtoMessage() {}

func (x *Posts) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Posts.ProtoReflect.Descriptor instead.
func (*Posts) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *Posts) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Posts) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Posts) GetPosts() []string {
	if x != nil {
		return x.Posts
	}
	return nil
}

type FollowingsPosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowingsPosts []*Posts `protobuf:"bytes,1,rep,name=FollowingsPosts,proto3" json:"FollowingsPosts,omitempty"`
}

func (x *FollowingsPosts) Reset() {
	*x = FollowingsPosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingsPosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingsPosts) ProtoMessage() {}

func (x *FollowingsPosts) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingsPosts.ProtoReflect.Descriptor instead.
func (*FollowingsPosts) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{2}
}

func (x *FollowingsPosts) GetFollowingsPosts() []*Posts {
	if x != nil {
		return x.FollowingsPosts
	}
	return nil
}

type Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Search) Reset() {
	*x = Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Search) ProtoMessage() {}

func (x *Search) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Search.ProtoReflect.Descriptor instead.
func (*Search) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{3}
}

func (x *Search) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Followings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usernames []string `protobuf:"bytes,1,rep,name=usernames,proto3" json:"usernames,omitempty"`
}

func (x *Followings) Reset() {
	*x = Followings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Followings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Followings) ProtoMessage() {}

func (x *Followings) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Followings.ProtoReflect.Descriptor instead.
func (*Followings) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{4}
}

func (x *Followings) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

type SaveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Saved bool `protobuf:"varint,1,opt,name=saved,proto3" json:"saved,omitempty"`
}

func (x *SaveReply) Reset() {
	*x = SaveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveReply) ProtoMessage() {}

func (x *SaveReply) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveReply.ProtoReflect.Descriptor instead.
func (*SaveReply) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{5}
}

func (x *SaveReply) GetSaved() bool {
	if x != nil {
		return x.Saved
	}
	return false
}

type FollowReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followed bool `protobuf:"varint,1,opt,name=followed,proto3" json:"followed,omitempty"`
}

func (x *FollowReply) Reset() {
	*x = FollowReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowReply) ProtoMessage() {}

func (x *FollowReply) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowReply.ProtoReflect.Descriptor instead.
func (*FollowReply) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{6}
}

func (x *FollowReply) GetFollowed() bool {
	if x != nil {
		return x.Followed
	}
	return false
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x05, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x05,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x4e, 0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x0a, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x22, 0x29, 0x0a, 0x0b, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x32, 0xf6, 0x02, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47,
	0x75, 0x69, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64,
	0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x09, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x15, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x17, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x17, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x4c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x12, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x4c, 0x6f,
	0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x1a,
	0x1b, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x00, 0x42, 0x68,
	0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x42, 0x0f, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_models_proto_goTypes = []interface{}{
	(*Users)(nil),           // 0: routeguide.Users
	(*Posts)(nil),           // 1: routeguide.Posts
	(*FollowingsPosts)(nil), // 2: routeguide.FollowingsPosts
	(*Search)(nil),          // 3: routeguide.Search
	(*Followings)(nil),      // 4: routeguide.Followings
	(*SaveReply)(nil),       // 5: routeguide.SaveReply
	(*FollowReply)(nil),     // 6: routeguide.FollowReply
}
var file_models_proto_depIdxs = []int32{
	1, // 0: routeguide.FollowingsPosts.FollowingsPosts:type_name -> routeguide.Posts
	0, // 1: routeguide.RouteGuide.SaveUsers:input_type -> routeguide.Users
	1, // 2: routeguide.RouteGuide.SavePosts:input_type -> routeguide.Posts
	3, // 3: routeguide.RouteGuide.FollowUsers:input_type -> routeguide.Search
	3, // 4: routeguide.RouteGuide.UnfollowUsers:input_type -> routeguide.Search
	3, // 5: routeguide.RouteGuide.LoadUsers:input_type -> routeguide.Search
	4, // 6: routeguide.RouteGuide.LoadPosts:input_type -> routeguide.Followings
	5, // 7: routeguide.RouteGuide.SaveUsers:output_type -> routeguide.SaveReply
	5, // 8: routeguide.RouteGuide.SavePosts:output_type -> routeguide.SaveReply
	6, // 9: routeguide.RouteGuide.FollowUsers:output_type -> routeguide.FollowReply
	6, // 10: routeguide.RouteGuide.UnfollowUsers:output_type -> routeguide.FollowReply
	0, // 11: routeguide.RouteGuide.LoadUsers:output_type -> routeguide.Users
	2, // 12: routeguide.RouteGuide.LoadPosts:output_type -> routeguide.FollowingsPosts
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Posts); i {
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
		file_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingsPosts); i {
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
		file_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Search); i {
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
		file_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Followings); i {
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
		file_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveReply); i {
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
		file_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowReply); i {
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
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}
