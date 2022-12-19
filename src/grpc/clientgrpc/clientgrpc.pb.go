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
	Totp      string `protobuf:"bytes,2,opt,name=totp,proto3" json:"totp,omitempty"`
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

func (x *LoginConfirmRequest) GetTotp() string {
	if x != nil {
		return x.Totp
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
	0x52, 0x07, 0x74, 0x6f, 0x74, 0x70, 0x55, 0x72, 0x69, 0x22, 0x47, 0x0a, 0x13, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f,
	0x74, 0x70, 0x22, 0x3a, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa0,
	0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x1f,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_clientgrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_clientgrpc_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),         // 0: clientgrpc.LoginRequest
	(*LoginResponse)(nil),        // 1: clientgrpc.LoginResponse
	(*LoginConfirmRequest)(nil),  // 2: clientgrpc.LoginConfirmRequest
	(*LoginConfirmResponse)(nil), // 3: clientgrpc.LoginConfirmResponse
}
var file_clientgrpc_proto_depIdxs = []int32{
	0, // 0: clientgrpc.ClientService.Login:input_type -> clientgrpc.LoginRequest
	2, // 1: clientgrpc.ClientService.LoginConfirm:input_type -> clientgrpc.LoginConfirmRequest
	1, // 2: clientgrpc.ClientService.Login:output_type -> clientgrpc.LoginResponse
	3, // 3: clientgrpc.ClientService.LoginConfirm:output_type -> clientgrpc.LoginConfirmResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_clientgrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
