// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: kwil/txsvc/executables.proto

package txpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	commonpb "kwil/x/proto/commonpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetExecutablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Owner    string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *GetExecutablesRequest) Reset() {
	*x = GetExecutablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_txsvc_executables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExecutablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExecutablesRequest) ProtoMessage() {}

func (x *GetExecutablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_txsvc_executables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExecutablesRequest.ProtoReflect.Descriptor instead.
func (*GetExecutablesRequest) Descriptor() ([]byte, []int) {
	return file_kwil_txsvc_executables_proto_rawDescGZIP(), []int{0}
}

func (x *GetExecutablesRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *GetExecutablesRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type GetExecutablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Executables []*commonpb.Executable `protobuf:"bytes,1,rep,name=executables,proto3" json:"executables,omitempty"`
}

func (x *GetExecutablesResponse) Reset() {
	*x = GetExecutablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_txsvc_executables_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExecutablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExecutablesResponse) ProtoMessage() {}

func (x *GetExecutablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_txsvc_executables_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExecutablesResponse.ProtoReflect.Descriptor instead.
func (*GetExecutablesResponse) Descriptor() ([]byte, []int) {
	return file_kwil_txsvc_executables_proto_rawDescGZIP(), []int{1}
}

func (x *GetExecutablesResponse) GetExecutables() []*commonpb.Executable {
	if x != nil {
		return x.Executables
	}
	return nil
}

type GetExecutablesByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetExecutablesByIdRequest) Reset() {
	*x = GetExecutablesByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_txsvc_executables_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExecutablesByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExecutablesByIdRequest) ProtoMessage() {}

func (x *GetExecutablesByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_txsvc_executables_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExecutablesByIdRequest.ProtoReflect.Descriptor instead.
func (*GetExecutablesByIdRequest) Descriptor() ([]byte, []int) {
	return file_kwil_txsvc_executables_proto_rawDescGZIP(), []int{2}
}

func (x *GetExecutablesByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_kwil_txsvc_executables_proto protoreflect.FileDescriptor

var file_kwil_txsvc_executables_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x78, 0x73, 0x76, 0x63, 0x1a, 0x1c, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x4e,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x2b,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x13, 0x5a, 0x11, 0x6b,
	0x77, 0x69, 0x6c, 0x2f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x78, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kwil_txsvc_executables_proto_rawDescOnce sync.Once
	file_kwil_txsvc_executables_proto_rawDescData = file_kwil_txsvc_executables_proto_rawDesc
)

func file_kwil_txsvc_executables_proto_rawDescGZIP() []byte {
	file_kwil_txsvc_executables_proto_rawDescOnce.Do(func() {
		file_kwil_txsvc_executables_proto_rawDescData = protoimpl.X.CompressGZIP(file_kwil_txsvc_executables_proto_rawDescData)
	})
	return file_kwil_txsvc_executables_proto_rawDescData
}

var file_kwil_txsvc_executables_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kwil_txsvc_executables_proto_goTypes = []interface{}{
	(*GetExecutablesRequest)(nil),     // 0: txsvc.GetExecutablesRequest
	(*GetExecutablesResponse)(nil),    // 1: txsvc.GetExecutablesResponse
	(*GetExecutablesByIdRequest)(nil), // 2: txsvc.GetExecutablesByIdRequest
	(*commonpb.Executable)(nil),       // 3: common.Executable
}
var file_kwil_txsvc_executables_proto_depIdxs = []int32{
	3, // 0: txsvc.GetExecutablesResponse.executables:type_name -> common.Executable
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kwil_txsvc_executables_proto_init() }
func file_kwil_txsvc_executables_proto_init() {
	if File_kwil_txsvc_executables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kwil_txsvc_executables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExecutablesRequest); i {
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
		file_kwil_txsvc_executables_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExecutablesResponse); i {
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
		file_kwil_txsvc_executables_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExecutablesByIdRequest); i {
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
			RawDescriptor: file_kwil_txsvc_executables_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kwil_txsvc_executables_proto_goTypes,
		DependencyIndexes: file_kwil_txsvc_executables_proto_depIdxs,
		MessageInfos:      file_kwil_txsvc_executables_proto_msgTypes,
	}.Build()
	File_kwil_txsvc_executables_proto = out.File
	file_kwil_txsvc_executables_proto_rawDesc = nil
	file_kwil_txsvc_executables_proto_goTypes = nil
	file_kwil_txsvc_executables_proto_depIdxs = nil
}