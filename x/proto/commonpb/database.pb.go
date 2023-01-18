// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: kwil/common/database.proto

package commonpb

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

type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner      string      `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Tables     []*Table    `protobuf:"bytes,3,rep,name=tables,proto3" json:"tables,omitempty"`
	SqlQueries []*SQLQuery `protobuf:"bytes,4,rep,name=sql_queries,json=sqlQueries,proto3" json:"sql_queries,omitempty"`
	Roles      []*Role     `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
	Indexes    []*Index    `protobuf:"bytes,6,rep,name=indexes,proto3" json:"indexes,omitempty"`
}

func (x *Database) Reset() {
	*x = Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{0}
}

func (x *Database) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Database) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Database) GetTables() []*Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

func (x *Database) GetSqlQueries() []*SQLQuery {
	if x != nil {
		return x.SqlQueries
	}
	return nil
}

func (x *Database) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Database) GetIndexes() []*Index {
	if x != nil {
		return x.Indexes
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Columns []*Column `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{1}
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       DataType     `protobuf:"varint,2,opt,name=type,proto3,enum=common.DataType" json:"type,omitempty"`
	Attributes []*Attribute `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{2}
}

func (x *Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Column) GetType() DataType {
	if x != nil {
		return x.Type
	}
	return DataType_INVALID_TYPE
}

func (x *Column) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  AttributeType `protobuf:"varint,1,opt,name=type,proto3,enum=common.AttributeType" json:"type,omitempty"`
	Value []byte        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{3}
}

func (x *Attribute) GetType() AttributeType {
	if x != nil {
		return x.Type
	}
	return AttributeType_INVALID_ATTRIBUTE
}

func (x *Attribute) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SQLQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       QueryType      `protobuf:"varint,2,opt,name=type,proto3,enum=common.QueryType" json:"type,omitempty"`
	Table      string         `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	Parameters []*Parameter   `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty"`
	Where      []*WhereClause `protobuf:"bytes,5,rep,name=where,proto3" json:"where,omitempty"`
}

func (x *SQLQuery) Reset() {
	*x = SQLQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLQuery) ProtoMessage() {}

func (x *SQLQuery) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLQuery.ProtoReflect.Descriptor instead.
func (*SQLQuery) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{4}
}

func (x *SQLQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SQLQuery) GetType() QueryType {
	if x != nil {
		return x.Type
	}
	return QueryType_QUERY_INVALID
}

func (x *SQLQuery) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *SQLQuery) GetParameters() []*Parameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *SQLQuery) GetWhere() []*WhereClause {
	if x != nil {
		return x.Where
	}
	return nil
}

type Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Column   string       `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
	Static   bool         `protobuf:"varint,3,opt,name=static,proto3" json:"static,omitempty"`
	Value    string       `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Modifier ModifierType `protobuf:"varint,5,opt,name=modifier,proto3,enum=common.ModifierType" json:"modifier,omitempty"`
}

func (x *Parameter) Reset() {
	*x = Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parameter) ProtoMessage() {}

func (x *Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parameter.ProtoReflect.Descriptor instead.
func (*Parameter) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{5}
}

func (x *Parameter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Parameter) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *Parameter) GetStatic() bool {
	if x != nil {
		return x.Static
	}
	return false
}

func (x *Parameter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Parameter) GetModifier() ModifierType {
	if x != nil {
		return x.Modifier
	}
	return ModifierType_NO_MODIFIER
}

type WhereClause struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Column   string             `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
	Static   bool               `protobuf:"varint,3,opt,name=static,proto3" json:"static,omitempty"`
	Value    string             `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Modifier ModifierType       `protobuf:"varint,5,opt,name=modifier,proto3,enum=common.ModifierType" json:"modifier,omitempty"`
	Operator ComparisonOperator `protobuf:"varint,6,opt,name=operator,proto3,enum=common.ComparisonOperator" json:"operator,omitempty"`
}

func (x *WhereClause) Reset() {
	*x = WhereClause{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhereClause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhereClause) ProtoMessage() {}

func (x *WhereClause) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhereClause.ProtoReflect.Descriptor instead.
func (*WhereClause) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{6}
}

func (x *WhereClause) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WhereClause) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *WhereClause) GetStatic() bool {
	if x != nil {
		return x.Static
	}
	return false
}

func (x *WhereClause) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *WhereClause) GetModifier() ModifierType {
	if x != nil {
		return x.Modifier
	}
	return ModifierType_NO_MODIFIER
}

func (x *WhereClause) GetOperator() ComparisonOperator {
	if x != nil {
		return x.Operator
	}
	return ComparisonOperator_OPERATOR_INVALID
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Default     bool     `protobuf:"varint,2,opt,name=default,proto3" json:"default,omitempty"`
	Permissions []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{7}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *Role) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Table   string    `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Columns []string  `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	Using   IndexType `protobuf:"varint,4,opt,name=using,proto3,enum=common.IndexType" json:"using,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kwil_common_database_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_kwil_common_database_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_kwil_common_database_proto_rawDescGZIP(), []int{8}
}

func (x *Index) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Index) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Index) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Index) GetUsing() IndexType {
	if x != nil {
		return x.Using
	}
	return IndexType_INVALID_INDEX
}

var File_kwil_common_database_proto protoreflect.FileDescriptor

var file_kwil_common_database_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x17, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x73,
	0x71, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x51, 0x4c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x0a, 0x73, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22,
	0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x05, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x22, 0x75, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x09, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x08, 0x53, 0x51, 0x4c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x57, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x52, 0x05, 0x77, 0x68,
	0x65, 0x72, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xd1, 0x01,
	0x0a, 0x0b, 0x57, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0x56, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x74, 0x0a, 0x05, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x42,
	0x17, 0x5a, 0x15, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kwil_common_database_proto_rawDescOnce sync.Once
	file_kwil_common_database_proto_rawDescData = file_kwil_common_database_proto_rawDesc
)

func file_kwil_common_database_proto_rawDescGZIP() []byte {
	file_kwil_common_database_proto_rawDescOnce.Do(func() {
		file_kwil_common_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_kwil_common_database_proto_rawDescData)
	})
	return file_kwil_common_database_proto_rawDescData
}

var file_kwil_common_database_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_kwil_common_database_proto_goTypes = []interface{}{
	(*Database)(nil),        // 0: common.Database
	(*Table)(nil),           // 1: common.Table
	(*Column)(nil),          // 2: common.Column
	(*Attribute)(nil),       // 3: common.Attribute
	(*SQLQuery)(nil),        // 4: common.SQLQuery
	(*Parameter)(nil),       // 5: common.Parameter
	(*WhereClause)(nil),     // 6: common.WhereClause
	(*Role)(nil),            // 7: common.Role
	(*Index)(nil),           // 8: common.Index
	(DataType)(0),           // 9: common.DataType
	(AttributeType)(0),      // 10: common.AttributeType
	(QueryType)(0),          // 11: common.QueryType
	(ModifierType)(0),       // 12: common.ModifierType
	(ComparisonOperator)(0), // 13: common.ComparisonOperator
	(IndexType)(0),          // 14: common.IndexType
}
var file_kwil_common_database_proto_depIdxs = []int32{
	1,  // 0: common.Database.tables:type_name -> common.Table
	4,  // 1: common.Database.sql_queries:type_name -> common.SQLQuery
	7,  // 2: common.Database.roles:type_name -> common.Role
	8,  // 3: common.Database.indexes:type_name -> common.Index
	2,  // 4: common.Table.columns:type_name -> common.Column
	9,  // 5: common.Column.type:type_name -> common.DataType
	3,  // 6: common.Column.attributes:type_name -> common.Attribute
	10, // 7: common.Attribute.type:type_name -> common.AttributeType
	11, // 8: common.SQLQuery.type:type_name -> common.QueryType
	5,  // 9: common.SQLQuery.parameters:type_name -> common.Parameter
	6,  // 10: common.SQLQuery.where:type_name -> common.WhereClause
	12, // 11: common.Parameter.modifier:type_name -> common.ModifierType
	12, // 12: common.WhereClause.modifier:type_name -> common.ModifierType
	13, // 13: common.WhereClause.operator:type_name -> common.ComparisonOperator
	14, // 14: common.Index.using:type_name -> common.IndexType
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_kwil_common_database_proto_init() }
func file_kwil_common_database_proto_init() {
	if File_kwil_common_database_proto != nil {
		return
	}
	file_kwil_common_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kwil_common_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Database); i {
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
		file_kwil_common_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_kwil_common_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
		file_kwil_common_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_kwil_common_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLQuery); i {
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
		file_kwil_common_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parameter); i {
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
		file_kwil_common_database_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhereClause); i {
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
		file_kwil_common_database_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_kwil_common_database_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
			RawDescriptor: file_kwil_common_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kwil_common_database_proto_goTypes,
		DependencyIndexes: file_kwil_common_database_proto_depIdxs,
		MessageInfos:      file_kwil_common_database_proto_msgTypes,
	}.Build()
	File_kwil_common_database_proto = out.File
	file_kwil_common_database_proto_rawDesc = nil
	file_kwil_common_database_proto_goTypes = nil
	file_kwil_common_database_proto_depIdxs = nil
}
