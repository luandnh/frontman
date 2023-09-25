// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: auth.proto

package pb

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

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId         string   `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	BusinessUnitId   string   `protobuf:"bytes,2,opt,name=business_unit_id,json=businessUnitId,proto3" json:"business_unit_id,omitempty"`
	UserId           string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username         string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	DatabaseHost     string   `protobuf:"bytes,5,opt,name=database_host,json=databaseHost,proto3" json:"database_host,omitempty"`
	DatabasePort     uint32   `protobuf:"varint,6,opt,name=database_port,json=databasePort,proto3" json:"database_port,omitempty"`
	DatabaseName     string   `protobuf:"bytes,7,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	DatabaseUser     string   `protobuf:"bytes,8,opt,name=database_user,json=databaseUser,proto3" json:"database_user,omitempty"`
	DatabasePassword string   `protobuf:"bytes,9,opt,name=database_password,json=databasePassword,proto3" json:"database_password,omitempty"`
	Services         []string `protobuf:"bytes,10,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *Auth) GetBusinessUnitId() string {
	if x != nil {
		return x.BusinessUnitId
	}
	return ""
}

func (x *Auth) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Auth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Auth) GetDatabaseHost() string {
	if x != nil {
		return x.DatabaseHost
	}
	return ""
}

func (x *Auth) GetDatabasePort() uint32 {
	if x != nil {
		return x.DatabasePort
	}
	return 0
}

func (x *Auth) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *Auth) GetDatabaseUser() string {
	if x != nil {
		return x.DatabaseUser
	}
	return ""
}

func (x *Auth) GetDatabasePassword() string {
	if x != nil {
		return x.DatabasePassword
	}
	return ""
}

func (x *Auth) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Auth `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetData() *Auth {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0xdf, 0x02, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x22, 0x2c, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x6c, 0x34, 0x76, 0x6e, 0x2f, 0x70, 0x69, 0x74, 0x65, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_proto_goTypes = []interface{}{
	(*Auth)(nil),         // 0: pb.Auth
	(*AuthResponse)(nil), // 1: pb.AuthResponse
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: pb.AuthResponse.data:type_name -> pb.Auth
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
