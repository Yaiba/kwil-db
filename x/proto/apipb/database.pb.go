// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: kwil/apisvc/database.proto

package apipb

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

type CreateDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Fee       string `protobuf:"bytes,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Operation int32  `protobuf:"varint,5,opt,name=operation,proto3" json:"operation,omitempty"`
	Crud      int32  `protobuf:"varint,6,opt,name=crud,proto3" json:"crud,omitempty"`
	From      string `protobuf:"bytes,7,opt,name=from,proto3" json:"from,omitempty"`
	Signature string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *CreateDatabaseRequest) Reset() {
	*x = CreateDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseRequest) ProtoMessage() {}

func (x *CreateDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseRequest.ProtoReflect.Descriptor instead.
func (*CreateDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDatabaseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateDatabaseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDatabaseRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateDatabaseRequest) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *CreateDatabaseRequest) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *CreateDatabaseRequest) GetCrud() int32 {
	if x != nil {
		return x.Crud
	}
	return 0
}

func (x *CreateDatabaseRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateDatabaseRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type CreateDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDatabaseResponse) Reset() {
	*x = CreateDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseResponse) ProtoMessage() {}

func (x *CreateDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseResponse.ProtoReflect.Descriptor instead.
func (*CreateDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{1}
}

type UpdateDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner        string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Fee          string `protobuf:"bytes,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Operation    int32  `protobuf:"varint,5,opt,name=operation,proto3" json:"operation,omitempty"`
	Crud         int32  `protobuf:"varint,6,opt,name=crud,proto3" json:"crud,omitempty"`
	Instructions string `protobuf:"bytes,7,opt,name=instructions,proto3" json:"instructions,omitempty"`
	From         string `protobuf:"bytes,8,opt,name=from,proto3" json:"from,omitempty"`
	Nonce        string `protobuf:"bytes,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Signature    string `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *UpdateDatabaseRequest) Reset() {
	*x = UpdateDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatabaseRequest) ProtoMessage() {}

func (x *UpdateDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatabaseRequest.ProtoReflect.Descriptor instead.
func (*UpdateDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDatabaseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *UpdateDatabaseRequest) GetCrud() int32 {
	if x != nil {
		return x.Crud
	}
	return 0
}

func (x *UpdateDatabaseRequest) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type UpdateDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateDatabaseResponse) Reset() {
	*x = UpdateDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatabaseResponse) ProtoMessage() {}

func (x *UpdateDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatabaseResponse.ProtoReflect.Descriptor instead.
func (*UpdateDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{3}
}

type ListDatabasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDatabasesRequest) Reset() {
	*x = ListDatabasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatabasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabasesRequest) ProtoMessage() {}

func (x *ListDatabasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabasesRequest.ProtoReflect.Descriptor instead.
func (*ListDatabasesRequest) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{4}
}

type ListDatabasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDatabasesResponse) Reset() {
	*x = ListDatabasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDatabasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabasesResponse) ProtoMessage() {}

func (x *ListDatabasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabasesResponse.ProtoReflect.Descriptor instead.
func (*ListDatabasesResponse) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{5}
}

type GetDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDatabaseRequest) Reset() {
	*x = GetDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseRequest) ProtoMessage() {}

func (x *GetDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseRequest.ProtoReflect.Descriptor instead.
func (*GetDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{6}
}

func (x *GetDatabaseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDatabaseResponse) Reset() {
	*x = GetDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseResponse) ProtoMessage() {}

func (x *GetDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseResponse.ProtoReflect.Descriptor instead.
func (*GetDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{7}
}

type DeleteDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDatabaseRequest) Reset() {
	*x = DeleteDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseRequest) ProtoMessage() {}

func (x *DeleteDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDatabaseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDatabaseResponse) Reset() {
	*x = DeleteDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_apisvc_database_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseResponse) ProtoMessage() {}

func (x *DeleteDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_apisvc_database_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseResponse.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_kwil_apisvc_database_proto_rawDescGZIP(), []int{9}
}

var File_kwil_apisvc_database_proto protoreflect.FileDescriptor

var file_kwil_apisvc_database_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x76, 0x63, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70,
	0x69, 0x73, 0x76, 0x63, 0x22, 0xc5, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x75, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x72, 0x75, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x18, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x02, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72,
	0x75, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x72, 0x75, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x78, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_kwil_apisvc_database_proto_rawDescOnce sync.Once
	file_kwil_apisvc_database_proto_rawDescData = file_kwil_apisvc_database_proto_rawDesc
)

func file_kwil_apisvc_database_proto_rawDescGZIP() []byte {
	file_kwil_apisvc_database_proto_rawDescOnce.Do(func() {
		file_kwil_apisvc_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_kwil_apisvc_database_proto_rawDescData)
	})
	return file_kwil_apisvc_database_proto_rawDescData
}

var file_kwil_apisvc_database_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_kwil_apisvc_database_proto_goTypes = []interface{}{
	(*CreateDatabaseRequest)(nil),  // 0: apisvc.CreateDatabaseRequest
	(*CreateDatabaseResponse)(nil), // 1: apisvc.CreateDatabaseResponse
	(*UpdateDatabaseRequest)(nil),  // 2: apisvc.UpdateDatabaseRequest
	(*UpdateDatabaseResponse)(nil), // 3: apisvc.UpdateDatabaseResponse
	(*ListDatabasesRequest)(nil),   // 4: apisvc.ListDatabasesRequest
	(*ListDatabasesResponse)(nil),  // 5: apisvc.ListDatabasesResponse
	(*GetDatabaseRequest)(nil),     // 6: apisvc.GetDatabaseRequest
	(*GetDatabaseResponse)(nil),    // 7: apisvc.GetDatabaseResponse
	(*DeleteDatabaseRequest)(nil),  // 8: apisvc.DeleteDatabaseRequest
	(*DeleteDatabaseResponse)(nil), // 9: apisvc.DeleteDatabaseResponse
}
var file_kwil_apisvc_database_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kwil_apisvc_database_proto_init() }
func file_kwil_apisvc_database_proto_init() {
	if File_kwil_apisvc_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kwil_apisvc_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatabaseRequest); i {
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
		file_kwil_apisvc_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatabaseResponse); i {
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
		file_kwil_apisvc_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDatabaseRequest); i {
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
		file_kwil_apisvc_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDatabaseResponse); i {
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
		file_kwil_apisvc_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatabasesRequest); i {
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
		file_kwil_apisvc_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDatabasesResponse); i {
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
		file_kwil_apisvc_database_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatabaseRequest); i {
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
		file_kwil_apisvc_database_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatabaseResponse); i {
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
		file_kwil_apisvc_database_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatabaseRequest); i {
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
		file_kwil_apisvc_database_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatabaseResponse); i {
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
			RawDescriptor: file_kwil_apisvc_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kwil_apisvc_database_proto_goTypes,
		DependencyIndexes: file_kwil_apisvc_database_proto_depIdxs,
		MessageInfos:      file_kwil_apisvc_database_proto_msgTypes,
	}.Build()
	File_kwil_apisvc_database_proto = out.File
	file_kwil_apisvc_database_proto_rawDesc = nil
	file_kwil_apisvc_database_proto_goTypes = nil
	file_kwil_apisvc_database_proto_depIdxs = nil
}
