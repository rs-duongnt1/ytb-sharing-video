// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/ytb_sharing.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password        string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword string `protobuf:"bytes,3,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
}

func (x *UserRegisterRequest) Reset() {
	*x = UserRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRequest) ProtoMessage() {}

func (x *UserRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRequest.ProtoReflect.Descriptor instead.
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRegisterRequest) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

type UserRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserRegisterResponse) Reset() {
	*x = UserRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterResponse) ProtoMessage() {}

func (x *UserRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterResponse.ProtoReflect.Descriptor instead.
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegisterResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserRegisterResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRegisterResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginRequest) Reset() {
	*x = UserLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRequest) ProtoMessage() {}

func (x *UserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRequest.ProtoReflect.Descriptor instead.
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{2}
}

func (x *UserLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserLoginResponse) Reset() {
	*x = UserLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResponse) ProtoMessage() {}

func (x *UserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResponse.ProtoReflect.Descriptor instead.
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{4}
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfoResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url         string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	VideoId     string `protobuf:"bytes,5,opt,name=videoId,proto3" json:"videoId,omitempty"`
	User        *User  `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{7}
}

func (x *Video) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Video) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *Video) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type VideoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoListRequest) Reset() {
	*x = VideoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoListRequest) ProtoMessage() {}

func (x *VideoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoListRequest.ProtoReflect.Descriptor instead.
func (*VideoListRequest) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{8}
}

type VideoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *VideoListResponse) Reset() {
	*x = VideoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoListResponse) ProtoMessage() {}

func (x *VideoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoListResponse.ProtoReflect.Descriptor instead.
func (*VideoListResponse) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{9}
}

func (x *VideoListResponse) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

type VideoShareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VideoShareRequest) Reset() {
	*x = VideoShareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoShareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoShareRequest) ProtoMessage() {}

func (x *VideoShareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoShareRequest.ProtoReflect.Descriptor instead.
func (*VideoShareRequest) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{10}
}

func (x *VideoShareRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VideoShareRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VideoShareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url         string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	VideoId     string `protobuf:"bytes,5,opt,name=videoId,proto3" json:"videoId,omitempty"`
	User        *User  `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *VideoShareResponse) Reset() {
	*x = VideoShareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ytb_sharing_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoShareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoShareResponse) ProtoMessage() {}

func (x *VideoShareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ytb_sharing_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoShareResponse.ProtoReflect.Descriptor instead.
func (*VideoShareResponse) Descriptor() ([]byte, []int) {
	return file_proto_ytb_sharing_proto_rawDescGZIP(), []int{11}
}

func (x *VideoShareResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoShareResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VideoShareResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoShareResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VideoShareResponse) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *VideoShareResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_proto_ytb_sharing_proto protoreflect.FileDescriptor

var file_proto_ytb_sharing_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x79, 0x74, 0x62, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x52, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x44, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x29, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x11, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x38, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2c, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xa2, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x12, 0x0a, 0x10,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3f, 0x0a, 0x11, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x22, 0x3b, 0x0a, 0x11, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xaf,
	0x01, 0x0a, 0x12, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x32, 0x94, 0x04, 0x0a, 0x11, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x6f, 0x72, 0x69, 0x75, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a,
	0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x59, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x79,
	0x74, 0x62, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x79, 0x74, 0x62,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x6a, 0x0a, 0x0a, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x1e, 0x2e, 0x79, 0x74, 0x62, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x79, 0x74, 0x62, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x12, 0x5e, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x79, 0x74, 0x62, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x42, 0x74, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x79,
	0x74, 0x62, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0f, 0x59, 0x74, 0x62, 0x53,
	0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x08, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x59, 0x58, 0x58, 0xaa, 0x02, 0x0a,
	0x59, 0x74, 0x62, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x0a, 0x59, 0x74, 0x62,
	0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x16, 0x59, 0x74, 0x62, 0x53, 0x68, 0x61,
	0x72, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0a, 0x59, 0x74, 0x62, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ytb_sharing_proto_rawDescOnce sync.Once
	file_proto_ytb_sharing_proto_rawDescData = file_proto_ytb_sharing_proto_rawDesc
)

func file_proto_ytb_sharing_proto_rawDescGZIP() []byte {
	file_proto_ytb_sharing_proto_rawDescOnce.Do(func() {
		file_proto_ytb_sharing_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ytb_sharing_proto_rawDescData)
	})
	return file_proto_ytb_sharing_proto_rawDescData
}

var file_proto_ytb_sharing_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_ytb_sharing_proto_goTypes = []interface{}{
	(*UserRegisterRequest)(nil),  // 0: ytb_sharing.UserRegisterRequest
	(*UserRegisterResponse)(nil), // 1: ytb_sharing.UserRegisterResponse
	(*UserLoginRequest)(nil),     // 2: ytb_sharing.UserLoginRequest
	(*UserLoginResponse)(nil),    // 3: ytb_sharing.UserLoginResponse
	(*UserInfoRequest)(nil),      // 4: ytb_sharing.UserInfoRequest
	(*UserInfoResponse)(nil),     // 5: ytb_sharing.UserInfoResponse
	(*User)(nil),                 // 6: ytb_sharing.User
	(*Video)(nil),                // 7: ytb_sharing.Video
	(*VideoListRequest)(nil),     // 8: ytb_sharing.VideoListRequest
	(*VideoListResponse)(nil),    // 9: ytb_sharing.VideoListResponse
	(*VideoShareRequest)(nil),    // 10: ytb_sharing.VideoShareRequest
	(*VideoShareResponse)(nil),   // 11: ytb_sharing.VideoShareResponse
}
var file_proto_ytb_sharing_proto_depIdxs = []int32{
	6,  // 0: ytb_sharing.Video.user:type_name -> ytb_sharing.User
	7,  // 1: ytb_sharing.VideoListResponse.videos:type_name -> ytb_sharing.Video
	6,  // 2: ytb_sharing.VideoShareResponse.user:type_name -> ytb_sharing.User
	0,  // 3: ytb_sharing.BuildoriumService.UserRegister:input_type -> ytb_sharing.UserRegisterRequest
	2,  // 4: ytb_sharing.BuildoriumService.UserLogin:input_type -> ytb_sharing.UserLoginRequest
	4,  // 5: ytb_sharing.BuildoriumService.UserInfo:input_type -> ytb_sharing.UserInfoRequest
	10, // 6: ytb_sharing.BuildoriumService.VideoShare:input_type -> ytb_sharing.VideoShareRequest
	8,  // 7: ytb_sharing.BuildoriumService.VideoList:input_type -> ytb_sharing.VideoListRequest
	1,  // 8: ytb_sharing.BuildoriumService.UserRegister:output_type -> ytb_sharing.UserRegisterResponse
	3,  // 9: ytb_sharing.BuildoriumService.UserLogin:output_type -> ytb_sharing.UserLoginResponse
	5,  // 10: ytb_sharing.BuildoriumService.UserInfo:output_type -> ytb_sharing.UserInfoResponse
	11, // 11: ytb_sharing.BuildoriumService.VideoShare:output_type -> ytb_sharing.VideoShareResponse
	9,  // 12: ytb_sharing.BuildoriumService.VideoList:output_type -> ytb_sharing.VideoListResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_ytb_sharing_proto_init() }
func file_proto_ytb_sharing_proto_init() {
	if File_proto_ytb_sharing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ytb_sharing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterRequest); i {
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
		file_proto_ytb_sharing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterResponse); i {
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
		file_proto_ytb_sharing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRequest); i {
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
		file_proto_ytb_sharing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResponse); i {
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
		file_proto_ytb_sharing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRequest); i {
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
		file_proto_ytb_sharing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
		file_proto_ytb_sharing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_ytb_sharing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_proto_ytb_sharing_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoListRequest); i {
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
		file_proto_ytb_sharing_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoListResponse); i {
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
		file_proto_ytb_sharing_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoShareRequest); i {
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
		file_proto_ytb_sharing_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoShareResponse); i {
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
			RawDescriptor: file_proto_ytb_sharing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ytb_sharing_proto_goTypes,
		DependencyIndexes: file_proto_ytb_sharing_proto_depIdxs,
		MessageInfos:      file_proto_ytb_sharing_proto_msgTypes,
	}.Build()
	File_proto_ytb_sharing_proto = out.File
	file_proto_ytb_sharing_proto_rawDesc = nil
	file_proto_ytb_sharing_proto_goTypes = nil
	file_proto_ytb_sharing_proto_depIdxs = nil
}