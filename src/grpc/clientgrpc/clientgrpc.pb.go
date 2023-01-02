// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: clientgrpc.proto

package clientgrpc

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotpToken string `protobuf:"bytes,1,opt,name=totpToken,proto3" json:"totpToken,omitempty"`
	TotpUri   string `protobuf:"bytes,2,opt,name=totpUri,proto3" json:"totpUri,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetTotpToken() string {
	if x != nil {
		return x.TotpToken
	}
	return ""
}

func (x *LoginResponse) GetTotpUri() string {
	if x != nil {
		return x.TotpUri
	}
	return ""
}

type LoginConfirmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotpToken string `protobuf:"bytes,1,opt,name=totpToken,proto3" json:"totpToken,omitempty"`
	TotpCode  string `protobuf:"bytes,2,opt,name=totpCode,proto3" json:"totpCode,omitempty"`
}

func (x *LoginConfirmRequest) Reset() {
	*x = LoginConfirmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginConfirmRequest) ProtoMessage() {}

func (x *LoginConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginConfirmRequest.ProtoReflect.Descriptor instead.
func (*LoginConfirmRequest) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{2}
}

func (x *LoginConfirmRequest) GetTotpToken() string {
	if x != nil {
		return x.TotpToken
	}
	return ""
}

func (x *LoginConfirmRequest) GetTotpCode() string {
	if x != nil {
		return x.TotpCode
	}
	return ""
}

type LoginConfirmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionToken string `protobuf:"bytes,1,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *LoginConfirmResponse) Reset() {
	*x = LoginConfirmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginConfirmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginConfirmResponse) ProtoMessage() {}

func (x *LoginConfirmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginConfirmResponse.ProtoReflect.Descriptor instead.
func (*LoginConfirmResponse) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{3}
}

func (x *LoginConfirmResponse) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type GetUserDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionToken string `protobuf:"bytes,1,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *GetUserDetailsRequest) Reset() {
	*x = GetUserDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailsRequest) ProtoMessage() {}

func (x *GetUserDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetUserDetailsRequest) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserDetailsRequest) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type GetUserDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type           int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Roles          []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	TotpConfigured bool     `protobuf:"varint,4,opt,name=totpConfigured,proto3" json:"totpConfigured,omitempty"`
	TotpUri        string   `protobuf:"bytes,5,opt,name=totpUri,proto3" json:"totpUri,omitempty"`
}

func (x *GetUserDetailsResponse) Reset() {
	*x = GetUserDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailsResponse) ProtoMessage() {}

func (x *GetUserDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetUserDetailsResponse) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserDetailsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUserDetailsResponse) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GetUserDetailsResponse) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *GetUserDetailsResponse) GetTotpConfigured() bool {
	if x != nil {
		return x.TotpConfigured
	}
	return false
}

func (x *GetUserDetailsResponse) GetTotpUri() string {
	if x != nil {
		return x.TotpUri
	}
	return ""
}

type AuthenticateMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineToken string `protobuf:"bytes,1,opt,name=machineToken,proto3" json:"machineToken,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Ip           string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *AuthenticateMachineRequest) Reset() {
	*x = AuthenticateMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateMachineRequest) ProtoMessage() {}

func (x *AuthenticateMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateMachineRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateMachineRequest) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{6}
}

func (x *AuthenticateMachineRequest) GetMachineToken() string {
	if x != nil {
		return x.MachineToken
	}
	return ""
}

func (x *AuthenticateMachineRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthenticateMachineRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type AuthenticateMachineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type  int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Roles []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *AuthenticateMachineResponse) Reset() {
	*x = AuthenticateMachineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateMachineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateMachineResponse) ProtoMessage() {}

func (x *AuthenticateMachineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateMachineResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateMachineResponse) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{7}
}

func (x *AuthenticateMachineResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthenticateMachineResponse) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AuthenticateMachineResponse) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionToken string `protobuf:"bytes,1,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{8}
}

func (x *LogoutRequest) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SessionId string `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientgrpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clientgrpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_clientgrpc_proto_rawDescGZIP(), []int{9}
}

func (x *LogoutResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LogoutResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

var File_clientgrpc_proto protoreflect.FileDescriptor

var file_clientgrpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x22, 0x40,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x47, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x6f, 0x74, 0x70, 0x55, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x6f, 0x74, 0x70, 0x55, 0x72, 0x69, 0x22, 0x4f, 0x0a, 0x13, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x6f, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x6f, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x70,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x74, 0x6f, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x74, 0x70, 0x55, 0x72, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x6f, 0x74, 0x70, 0x55, 0x72, 0x69, 0x22, 0x66, 0x0a, 0x1a, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x22, 0x57, 0x0a, 0x1b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x3e, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x32, 0xa2, 0x03, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x51, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x21, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a,
	0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clientgrpc_proto_rawDescOnce sync.Once
	file_clientgrpc_proto_rawDescData = file_clientgrpc_proto_rawDesc
)

func file_clientgrpc_proto_rawDescGZIP() []byte {
	file_clientgrpc_proto_rawDescOnce.Do(func() {
		file_clientgrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_clientgrpc_proto_rawDescData)
	})
	return file_clientgrpc_proto_rawDescData
}

var file_clientgrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_clientgrpc_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),                // 0: clientgrpc.LoginRequest
	(*LoginResponse)(nil),               // 1: clientgrpc.LoginResponse
	(*LoginConfirmRequest)(nil),         // 2: clientgrpc.LoginConfirmRequest
	(*LoginConfirmResponse)(nil),        // 3: clientgrpc.LoginConfirmResponse
	(*GetUserDetailsRequest)(nil),       // 4: clientgrpc.GetUserDetailsRequest
	(*GetUserDetailsResponse)(nil),      // 5: clientgrpc.GetUserDetailsResponse
	(*AuthenticateMachineRequest)(nil),  // 6: clientgrpc.AuthenticateMachineRequest
	(*AuthenticateMachineResponse)(nil), // 7: clientgrpc.AuthenticateMachineResponse
	(*LogoutRequest)(nil),               // 8: clientgrpc.LogoutRequest
	(*LogoutResponse)(nil),              // 9: clientgrpc.LogoutResponse
}
var file_clientgrpc_proto_depIdxs = []int32{
	0, // 0: clientgrpc.ClientService.Login:input_type -> clientgrpc.LoginRequest
	2, // 1: clientgrpc.ClientService.LoginConfirm:input_type -> clientgrpc.LoginConfirmRequest
	4, // 2: clientgrpc.ClientService.GetUserDetails:input_type -> clientgrpc.GetUserDetailsRequest
	8, // 3: clientgrpc.ClientService.Logout:input_type -> clientgrpc.LogoutRequest
	6, // 4: clientgrpc.ClientService.AuthenticateMachine:input_type -> clientgrpc.AuthenticateMachineRequest
	1, // 5: clientgrpc.ClientService.Login:output_type -> clientgrpc.LoginResponse
	3, // 6: clientgrpc.ClientService.LoginConfirm:output_type -> clientgrpc.LoginConfirmResponse
	5, // 7: clientgrpc.ClientService.GetUserDetails:output_type -> clientgrpc.GetUserDetailsResponse
	9, // 8: clientgrpc.ClientService.Logout:output_type -> clientgrpc.LogoutResponse
	7, // 9: clientgrpc.ClientService.AuthenticateMachine:output_type -> clientgrpc.AuthenticateMachineResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clientgrpc_proto_init() }
func file_clientgrpc_proto_init() {
	if File_clientgrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clientgrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_clientgrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_clientgrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginConfirmRequest); i {
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
		file_clientgrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginConfirmResponse); i {
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
		file_clientgrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailsRequest); i {
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
		file_clientgrpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailsResponse); i {
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
		file_clientgrpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateMachineRequest); i {
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
		file_clientgrpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateMachineResponse); i {
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
		file_clientgrpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_clientgrpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResponse); i {
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
			RawDescriptor: file_clientgrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clientgrpc_proto_goTypes,
		DependencyIndexes: file_clientgrpc_proto_depIdxs,
		MessageInfos:      file_clientgrpc_proto_msgTypes,
	}.Build()
	File_clientgrpc_proto = out.File
	file_clientgrpc_proto_rawDesc = nil
	file_clientgrpc_proto_goTypes = nil
	file_clientgrpc_proto_depIdxs = nil
}
