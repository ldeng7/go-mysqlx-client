//
// Copyright (c) 2015, 2020, Oracle and/or its affiliates. All rights reserved.
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License, version 2.0,
// as published by the Free Software Foundation.
//
// This program is also distributed with certain software (including
// but not limited to OpenSSL) that is licensed under separate terms,
// as designated in a particular file or component or in included license
// documentation.  The authors of MySQL hereby grant you an additional
// permission to link the program and your derivative works with the
// separately licensed software that they have included with MySQL.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License, version 2.0, for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.14.0
// source: mysqlx_datatypes.proto

// ifdef PROTOBUF_LITE: option optimize_for = LITE_RUNTIME;

//*
//@namespace Mysqlx::Datatypes
//@brief Data types

package mysqlxpb_datatypes

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Scalar_Type int32

const (
	Scalar_V_SINT   Scalar_Type = 1
	Scalar_V_UINT   Scalar_Type = 2
	Scalar_V_NULL   Scalar_Type = 3
	Scalar_V_OCTETS Scalar_Type = 4
	Scalar_V_DOUBLE Scalar_Type = 5
	Scalar_V_FLOAT  Scalar_Type = 6
	Scalar_V_BOOL   Scalar_Type = 7
	Scalar_V_STRING Scalar_Type = 8
)

// Enum value maps for Scalar_Type.
var (
	Scalar_Type_name = map[int32]string{
		1: "V_SINT",
		2: "V_UINT",
		3: "V_NULL",
		4: "V_OCTETS",
		5: "V_DOUBLE",
		6: "V_FLOAT",
		7: "V_BOOL",
		8: "V_STRING",
	}
	Scalar_Type_value = map[string]int32{
		"V_SINT":   1,
		"V_UINT":   2,
		"V_NULL":   3,
		"V_OCTETS": 4,
		"V_DOUBLE": 5,
		"V_FLOAT":  6,
		"V_BOOL":   7,
		"V_STRING": 8,
	}
)

func (x Scalar_Type) Enum() *Scalar_Type {
	p := new(Scalar_Type)
	*p = x
	return p
}

func (x Scalar_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scalar_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_mysqlx_datatypes_proto_enumTypes[0].Descriptor()
}

func (Scalar_Type) Type() protoreflect.EnumType {
	return &file_mysqlx_datatypes_proto_enumTypes[0]
}

func (x Scalar_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Scalar_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Scalar_Type(num)
	return nil
}

// Deprecated: Use Scalar_Type.Descriptor instead.
func (Scalar_Type) EnumDescriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{0, 0}
}

type Any_Type int32

const (
	Any_SCALAR Any_Type = 1
	Any_OBJECT Any_Type = 2
	Any_ARRAY  Any_Type = 3
)

// Enum value maps for Any_Type.
var (
	Any_Type_name = map[int32]string{
		1: "SCALAR",
		2: "OBJECT",
		3: "ARRAY",
	}
	Any_Type_value = map[string]int32{
		"SCALAR": 1,
		"OBJECT": 2,
		"ARRAY":  3,
	}
)

func (x Any_Type) Enum() *Any_Type {
	p := new(Any_Type)
	*p = x
	return p
}

func (x Any_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Any_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_mysqlx_datatypes_proto_enumTypes[1].Descriptor()
}

func (Any_Type) Type() protoreflect.EnumType {
	return &file_mysqlx_datatypes_proto_enumTypes[1]
}

func (x Any_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Any_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Any_Type(num)
	return nil
}

// Deprecated: Use Any_Type.Descriptor instead.
func (Any_Type) EnumDescriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{3, 0}
}

// a scalar
type Scalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         *Scalar_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Datatypes.Scalar_Type" json:"type,omitempty"`
	VSignedInt   *int64       `protobuf:"zigzag64,2,opt,name=v_signed_int,json=vSignedInt" json:"v_signed_int,omitempty"`
	VUnsignedInt *uint64      `protobuf:"varint,3,opt,name=v_unsigned_int,json=vUnsignedInt" json:"v_unsigned_int,omitempty"`
	// 4 is unused, was Null which doesn't have a storage anymore
	VOctets *Scalar_Octets `protobuf:"bytes,5,opt,name=v_octets,json=vOctets" json:"v_octets,omitempty"`
	VDouble *float64       `protobuf:"fixed64,6,opt,name=v_double,json=vDouble" json:"v_double,omitempty"`
	VFloat  *float32       `protobuf:"fixed32,7,opt,name=v_float,json=vFloat" json:"v_float,omitempty"`
	VBool   *bool          `protobuf:"varint,8,opt,name=v_bool,json=vBool" json:"v_bool,omitempty"`
	VString *Scalar_String `protobuf:"bytes,9,opt,name=v_string,json=vString" json:"v_string,omitempty"`
}

func (x *Scalar) Reset() {
	*x = Scalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_datatypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalar) ProtoMessage() {}

func (x *Scalar) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_datatypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalar.ProtoReflect.Descriptor instead.
func (*Scalar) Descriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{0}
}

func (x *Scalar) GetType() Scalar_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Scalar_V_SINT
}

func (x *Scalar) GetVSignedInt() int64 {
	if x != nil && x.VSignedInt != nil {
		return *x.VSignedInt
	}
	return 0
}

func (x *Scalar) GetVUnsignedInt() uint64 {
	if x != nil && x.VUnsignedInt != nil {
		return *x.VUnsignedInt
	}
	return 0
}

func (x *Scalar) GetVOctets() *Scalar_Octets {
	if x != nil {
		return x.VOctets
	}
	return nil
}

func (x *Scalar) GetVDouble() float64 {
	if x != nil && x.VDouble != nil {
		return *x.VDouble
	}
	return 0
}

func (x *Scalar) GetVFloat() float32 {
	if x != nil && x.VFloat != nil {
		return *x.VFloat
	}
	return 0
}

func (x *Scalar) GetVBool() bool {
	if x != nil && x.VBool != nil {
		return *x.VBool
	}
	return false
}

func (x *Scalar) GetVString() *Scalar_String {
	if x != nil {
		return x.VString
	}
	return nil
}

//*
//An object
type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fld []*Object_ObjectField `protobuf:"bytes,1,rep,name=fld" json:"fld,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_datatypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_datatypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{1}
}

func (x *Object) GetFld() []*Object_ObjectField {
	if x != nil {
		return x.Fld
	}
	return nil
}

//*
//An Array
type Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*Any `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (x *Array) Reset() {
	*x = Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_datatypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Array) ProtoMessage() {}

func (x *Array) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_datatypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Array.ProtoReflect.Descriptor instead.
func (*Array) Descriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{2}
}

func (x *Array) GetValue() []*Any {
	if x != nil {
		return x.Value
	}
	return nil
}

//*
//A helper to allow all field types
type Any struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   *Any_Type `protobuf:"varint,1,req,name=type,enum=Mysqlx.Datatypes.Any_Type" json:"type,omitempty"`
	Scalar *Scalar   `protobuf:"bytes,2,opt,name=scalar" json:"scalar,omitempty"`
	Obj    *Object   `protobuf:"bytes,3,opt,name=obj" json:"obj,omitempty"`
	Array  *Array    `protobuf:"bytes,4,opt,name=array" json:"array,omitempty"`
}

func (x *Any) Reset() {
	*x = Any{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_datatypes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Any) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Any) ProtoMessage() {}

func (x *Any) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_datatypes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Any.ProtoReflect.Descriptor instead.
func (*Any) Descriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{3}
}

func (x *Any) GetType() Any_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Any_SCALAR
}

func (x *Any) GetScalar() *Scalar {
	if x != nil {
		return x.Scalar
	}
	return nil
}

func (x *Any) GetObj() *Object {
	if x != nil {
		return x.Obj
	}
	return nil
}

func (x *Any) GetArray() *Array {
	if x != nil {
		return x.Array
	}
	return nil
}

//* a string with a charset/collation
type Scalar_String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     []byte  `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	Collation *uint64 `protobuf:"varint,2,opt,name=collation" json:"collation,omitempty"`
}

func (x *Scalar_String) Reset() {
	*x = Scalar_String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_datatypes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scalar_String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalar_String) ProtoMessage() {}

func (x *Scalar_String) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_datatypes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalar_String.ProtoReflect.Descriptor instead.
func (*Scalar_String) Descriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Scalar_String) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Scalar_String) GetCollation() uint64 {
	if x != nil && x.Collation != nil {
		return *x.Collation
	}
	return 0
}

//* an opaque octet sequence, with an optional content_type
//See @ref Mysqlx::Resultset::ContentType_BYTES for list of known values.
type Scalar_Octets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       []byte  `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	ContentType *uint32 `protobuf:"varint,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
}

func (x *Scalar_Octets) Reset() {
	*x = Scalar_Octets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_datatypes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scalar_Octets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalar_Octets) ProtoMessage() {}

func (x *Scalar_Octets) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_datatypes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalar_Octets.ProtoReflect.Descriptor instead.
func (*Scalar_Octets) Descriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Scalar_Octets) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Scalar_Octets) GetContentType() uint32 {
	if x != nil && x.ContentType != nil {
		return *x.ContentType
	}
	return 0
}

type Object_ObjectField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value *Any    `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
}

func (x *Object_ObjectField) Reset() {
	*x = Object_ObjectField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_datatypes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object_ObjectField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object_ObjectField) ProtoMessage() {}

func (x *Object_ObjectField) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_datatypes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object_ObjectField.ProtoReflect.Descriptor instead.
func (*Object_ObjectField) Descriptor() ([]byte, []int) {
	return file_mysqlx_datatypes_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Object_ObjectField) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *Object_ObjectField) GetValue() *Any {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_mysqlx_datatypes_proto protoreflect.FileDescriptor

var file_mysqlx_datatypes_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xb6, 0x04, 0x0a, 0x06, 0x53,
	0x63, 0x61, 0x6c, 0x61, 0x72, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x76, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0a,
	0x76, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x76, 0x5f,
	0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x76, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x08, 0x76, 0x5f, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x2e, 0x4f, 0x63, 0x74,
	0x65, 0x74, 0x73, 0x52, 0x07, 0x76, 0x4f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x76, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x76, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x5f, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x76, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x76, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x76, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x3a, 0x0a, 0x08, 0x76, 0x5f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x4d, 0x79, 0x73, 0x71,
	0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x76, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x1a, 0x3c, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x41, 0x0a, 0x06, 0x4f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x6d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06,
	0x56, 0x5f, 0x53, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x5f, 0x55, 0x49,
	0x4e, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x56, 0x5f, 0x4f, 0x43, 0x54, 0x45, 0x54, 0x53, 0x10, 0x04, 0x12, 0x0c,
	0x0a, 0x08, 0x56, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07,
	0x56, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x5f, 0x42,
	0x4f, 0x4f, 0x4c, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x08, 0x22, 0x8e, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x36,
	0x0a, 0x03, 0x66, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x4d, 0x79,
	0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x03, 0x66, 0x6c, 0x64, 0x1a, 0x4c, 0x0a, 0x0b, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x34, 0x0a, 0x05, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x2b, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d,
	0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xed, 0x01, 0x0a, 0x03, 0x41,
	0x6e, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x06, 0x73, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x12, 0x2a, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x03, 0x6f, 0x62, 0x6a,
	0x12, 0x2d, 0x0a, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22,
	0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x41, 0x4c, 0x41,
	0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x03, 0x42, 0x6c, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x63, 0x6a, 0x2e, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x64, 0x65, 0x6e, 0x67, 0x37, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x79, 0x73, 0x71,
	0x6c, 0x78, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78,
	0x70, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x70, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x70, 0x62, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73,
}

var (
	file_mysqlx_datatypes_proto_rawDescOnce sync.Once
	file_mysqlx_datatypes_proto_rawDescData = file_mysqlx_datatypes_proto_rawDesc
)

func file_mysqlx_datatypes_proto_rawDescGZIP() []byte {
	file_mysqlx_datatypes_proto_rawDescOnce.Do(func() {
		file_mysqlx_datatypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_mysqlx_datatypes_proto_rawDescData)
	})
	return file_mysqlx_datatypes_proto_rawDescData
}

var file_mysqlx_datatypes_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mysqlx_datatypes_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mysqlx_datatypes_proto_goTypes = []interface{}{
	(Scalar_Type)(0),           // 0: Mysqlx.Datatypes.Scalar.Type
	(Any_Type)(0),              // 1: Mysqlx.Datatypes.Any.Type
	(*Scalar)(nil),             // 2: Mysqlx.Datatypes.Scalar
	(*Object)(nil),             // 3: Mysqlx.Datatypes.Object
	(*Array)(nil),              // 4: Mysqlx.Datatypes.Array
	(*Any)(nil),                // 5: Mysqlx.Datatypes.Any
	(*Scalar_String)(nil),      // 6: Mysqlx.Datatypes.Scalar.String
	(*Scalar_Octets)(nil),      // 7: Mysqlx.Datatypes.Scalar.Octets
	(*Object_ObjectField)(nil), // 8: Mysqlx.Datatypes.Object.ObjectField
}
var file_mysqlx_datatypes_proto_depIdxs = []int32{
	0,  // 0: Mysqlx.Datatypes.Scalar.type:type_name -> Mysqlx.Datatypes.Scalar.Type
	7,  // 1: Mysqlx.Datatypes.Scalar.v_octets:type_name -> Mysqlx.Datatypes.Scalar.Octets
	6,  // 2: Mysqlx.Datatypes.Scalar.v_string:type_name -> Mysqlx.Datatypes.Scalar.String
	8,  // 3: Mysqlx.Datatypes.Object.fld:type_name -> Mysqlx.Datatypes.Object.ObjectField
	5,  // 4: Mysqlx.Datatypes.Array.value:type_name -> Mysqlx.Datatypes.Any
	1,  // 5: Mysqlx.Datatypes.Any.type:type_name -> Mysqlx.Datatypes.Any.Type
	2,  // 6: Mysqlx.Datatypes.Any.scalar:type_name -> Mysqlx.Datatypes.Scalar
	3,  // 7: Mysqlx.Datatypes.Any.obj:type_name -> Mysqlx.Datatypes.Object
	4,  // 8: Mysqlx.Datatypes.Any.array:type_name -> Mysqlx.Datatypes.Array
	5,  // 9: Mysqlx.Datatypes.Object.ObjectField.value:type_name -> Mysqlx.Datatypes.Any
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mysqlx_datatypes_proto_init() }
func file_mysqlx_datatypes_proto_init() {
	if File_mysqlx_datatypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mysqlx_datatypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scalar); i {
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
		file_mysqlx_datatypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_mysqlx_datatypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Array); i {
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
		file_mysqlx_datatypes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Any); i {
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
		file_mysqlx_datatypes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scalar_String); i {
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
		file_mysqlx_datatypes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scalar_Octets); i {
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
		file_mysqlx_datatypes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object_ObjectField); i {
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
			RawDescriptor: file_mysqlx_datatypes_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mysqlx_datatypes_proto_goTypes,
		DependencyIndexes: file_mysqlx_datatypes_proto_depIdxs,
		EnumInfos:         file_mysqlx_datatypes_proto_enumTypes,
		MessageInfos:      file_mysqlx_datatypes_proto_msgTypes,
	}.Build()
	File_mysqlx_datatypes_proto = out.File
	file_mysqlx_datatypes_proto_rawDesc = nil
	file_mysqlx_datatypes_proto_goTypes = nil
	file_mysqlx_datatypes_proto_depIdxs = nil
}
