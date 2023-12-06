// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: kwil/tx/v1/call.proto

package txpb

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

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body *CallRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// auth_type is the type of authenticator that will be used to derive
	// identifier from the sender.
	AuthType string `protobuf:"bytes,2,opt,name=auth_type,proto3" json:"auth_type,omitempty"`
	Sender   []byte `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_tx_v1_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_tx_v1_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_kwil_tx_v1_call_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetBody() *CallRequest_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *CallRequest) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *CallRequest) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_tx_v1_call_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_tx_v1_call_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_kwil_tx_v1_call_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

type CallRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CallRequest_Body) Reset() {
	*x = CallRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_tx_v1_call_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest_Body) ProtoMessage() {}

func (x *CallRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_tx_v1_call_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest_Body.ProtoReflect.Descriptor instead.
func (*CallRequest_Body) Descriptor() ([]byte, []int) {
	return file_kwil_tx_v1_call_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CallRequest_Body) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_kwil_tx_v1_call_proto protoreflect.FileDescriptor

var file_kwil_tx_v1_call_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x74, 0x78, 0x22, 0x8f, 0x01, 0x0a, 0x0b,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x78, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x20, 0x0a, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x26, 0x0a,
	0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x77, 0x69, 0x6c, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x6b, 0x77, 0x69,
	0x6c, 0x2d, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x78, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kwil_tx_v1_call_proto_rawDescOnce sync.Once
	file_kwil_tx_v1_call_proto_rawDescData = file_kwil_tx_v1_call_proto_rawDesc
)

func file_kwil_tx_v1_call_proto_rawDescGZIP() []byte {
	file_kwil_tx_v1_call_proto_rawDescOnce.Do(func() {
		file_kwil_tx_v1_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_kwil_tx_v1_call_proto_rawDescData)
	})
	return file_kwil_tx_v1_call_proto_rawDescData
}

var file_kwil_tx_v1_call_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kwil_tx_v1_call_proto_goTypes = []interface{}{
	(*CallRequest)(nil),      // 0: tx.CallRequest
	(*CallResponse)(nil),     // 1: tx.CallResponse
	(*CallRequest_Body)(nil), // 2: tx.CallRequest.Body
}
var file_kwil_tx_v1_call_proto_depIdxs = []int32{
	2, // 0: tx.CallRequest.body:type_name -> tx.CallRequest.Body
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kwil_tx_v1_call_proto_init() }
func file_kwil_tx_v1_call_proto_init() {
	if File_kwil_tx_v1_call_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kwil_tx_v1_call_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_kwil_tx_v1_call_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_kwil_tx_v1_call_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest_Body); i {
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
			RawDescriptor: file_kwil_tx_v1_call_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kwil_tx_v1_call_proto_goTypes,
		DependencyIndexes: file_kwil_tx_v1_call_proto_depIdxs,
		MessageInfos:      file_kwil_tx_v1_call_proto_msgTypes,
	}.Build()
	File_kwil_tx_v1_call_proto = out.File
	file_kwil_tx_v1_call_proto_rawDesc = nil
	file_kwil_tx_v1_call_proto_goTypes = nil
	file_kwil_tx_v1_call_proto_depIdxs = nil
}