// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/pb/server.proto

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

type CreateServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=ServerName,proto3" json:"ServerName,omitempty"`
}

func (x *CreateServerRequest) Reset() {
	*x = CreateServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerRequest) ProtoMessage() {}

func (x *CreateServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerRequest.ProtoReflect.Descriptor instead.
func (*CreateServerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_server_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServerRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

type CreateServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerID   string `protobuf:"bytes,1,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=ServerName,proto3" json:"ServerName,omitempty"`
}

func (x *CreateServerResponse) Reset() {
	*x = CreateServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerResponse) ProtoMessage() {}

func (x *CreateServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerResponse.ProtoReflect.Descriptor instead.
func (*CreateServerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServerResponse) GetServerID() string {
	if x != nil {
		return x.ServerID
	}
	return ""
}

func (x *CreateServerResponse) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

var File_pkg_pb_server_proto protoreflect.FileDescriptor

var file_pkg_pb_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x32, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_server_proto_rawDescOnce sync.Once
	file_pkg_pb_server_proto_rawDescData = file_pkg_pb_server_proto_rawDesc
)

func file_pkg_pb_server_proto_rawDescGZIP() []byte {
	file_pkg_pb_server_proto_rawDescOnce.Do(func() {
		file_pkg_pb_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_server_proto_rawDescData)
	})
	return file_pkg_pb_server_proto_rawDescData
}

var file_pkg_pb_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_pb_server_proto_goTypes = []interface{}{
	(*CreateServerRequest)(nil),  // 0: CreateServerRequest
	(*CreateServerResponse)(nil), // 1: CreateServerResponse
}
var file_pkg_pb_server_proto_depIdxs = []int32{
	0, // 0: Server.CreateServer:input_type -> CreateServerRequest
	1, // 1: Server.CreateServer:output_type -> CreateServerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_server_proto_init() }
func file_pkg_pb_server_proto_init() {
	if File_pkg_pb_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerRequest); i {
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
		file_pkg_pb_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerResponse); i {
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
			RawDescriptor: file_pkg_pb_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_server_proto_goTypes,
		DependencyIndexes: file_pkg_pb_server_proto_depIdxs,
		MessageInfos:      file_pkg_pb_server_proto_msgTypes,
	}.Build()
	File_pkg_pb_server_proto = out.File
	file_pkg_pb_server_proto_rawDesc = nil
	file_pkg_pb_server_proto_goTypes = nil
	file_pkg_pb_server_proto_depIdxs = nil
}
