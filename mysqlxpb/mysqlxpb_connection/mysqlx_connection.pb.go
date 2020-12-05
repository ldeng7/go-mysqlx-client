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
// source: mysqlx_connection.proto

//*
//@namespace Mysqlx::Connection

package mysqlxpb_connection

import (
	proto "github.com/golang/protobuf/proto"
	mysqlxpb "github.com/ldeng7/go-mysqlx-client/mysqlxpb"
	mysqlxpb_datatypes "github.com/ldeng7/go-mysqlx-client/mysqlxpb/mysqlxpb_datatypes"
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

//*
//Capability
//
//A tuple of a ``name`` and a @ref Mysqlx::Datatypes::Any
type Capability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string                 `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value *mysqlxpb_datatypes.Any `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
}

func (x *Capability) Reset() {
	*x = Capability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capability) ProtoMessage() {}

func (x *Capability) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capability.ProtoReflect.Descriptor instead.
func (*Capability) Descriptor() ([]byte, []int) {
	return file_mysqlx_connection_proto_rawDescGZIP(), []int{0}
}

func (x *Capability) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Capability) GetValue() *mysqlxpb_datatypes.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

//*
//Capabilities
//
//list of Capability
type Capabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capabilities []*Capability `protobuf:"bytes,1,rep,name=capabilities" json:"capabilities,omitempty"`
}

func (x *Capabilities) Reset() {
	*x = Capabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_connection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capabilities) ProtoMessage() {}

func (x *Capabilities) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_connection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capabilities.ProtoReflect.Descriptor instead.
func (*Capabilities) Descriptor() ([]byte, []int) {
	return file_mysqlx_connection_proto_rawDescGZIP(), []int{1}
}

func (x *Capabilities) GetCapabilities() []*Capability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

//*
//Get supported connection capabilities and their current state.
//
//@returns @ref Mysqlx::Connection::Capabilities or @ref Mysqlx::Error
type CapabilitiesGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CapabilitiesGet) Reset() {
	*x = CapabilitiesGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_connection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilitiesGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilitiesGet) ProtoMessage() {}

func (x *CapabilitiesGet) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_connection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilitiesGet.ProtoReflect.Descriptor instead.
func (*CapabilitiesGet) Descriptor() ([]byte, []int) {
	return file_mysqlx_connection_proto_rawDescGZIP(), []int{2}
}

//*
//Set connection capabilities atomically.
//Only provided values are changed, other values are left
//unchanged. If any of the changes fails, all changes are
//discarded.
//
//@pre active sessions == 0
//
//@returns @ref Mysqlx::Ok  or @ref Mysqlx::Error
type CapabilitiesSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capabilities *Capabilities `protobuf:"bytes,1,req,name=capabilities" json:"capabilities,omitempty"`
}

func (x *CapabilitiesSet) Reset() {
	*x = CapabilitiesSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_connection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilitiesSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilitiesSet) ProtoMessage() {}

func (x *CapabilitiesSet) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_connection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilitiesSet.ProtoReflect.Descriptor instead.
func (*CapabilitiesSet) Descriptor() ([]byte, []int) {
	return file_mysqlx_connection_proto_rawDescGZIP(), []int{3}
}

func (x *CapabilitiesSet) GetCapabilities() *Capabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

//*
//Announce to the server that the client wants to close the connection.
//
//It discards any session state of the server.
//
//@returns @ref Mysqlx::Ok
type Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Close) Reset() {
	*x = Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_connection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Close) ProtoMessage() {}

func (x *Close) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_connection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Close.ProtoReflect.Descriptor instead.
func (*Close) Descriptor() ([]byte, []int) {
	return file_mysqlx_connection_proto_rawDescGZIP(), []int{4}
}

type Compression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UncompressedSize *uint64                       `protobuf:"varint,1,opt,name=uncompressed_size,json=uncompressedSize" json:"uncompressed_size,omitempty"`
	ServerMessages   *mysqlxpb.ServerMessages_Type `protobuf:"varint,2,opt,name=server_messages,json=serverMessages,enum=Mysqlx.ServerMessages_Type" json:"server_messages,omitempty"`
	ClientMessages   *mysqlxpb.ClientMessages_Type `protobuf:"varint,3,opt,name=client_messages,json=clientMessages,enum=Mysqlx.ClientMessages_Type" json:"client_messages,omitempty"`
	Payload          []byte                        `protobuf:"bytes,4,req,name=payload" json:"payload,omitempty"`
}

func (x *Compression) Reset() {
	*x = Compression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysqlx_connection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Compression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Compression) ProtoMessage() {}

func (x *Compression) ProtoReflect() protoreflect.Message {
	mi := &file_mysqlx_connection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Compression.ProtoReflect.Descriptor instead.
func (*Compression) Descriptor() ([]byte, []int) {
	return file_mysqlx_connection_proto_rawDescGZIP(), []int{5}
}

func (x *Compression) GetUncompressedSize() uint64 {
	if x != nil && x.UncompressedSize != nil {
		return *x.UncompressedSize
	}
	return 0
}

func (x *Compression) GetServerMessages() mysqlxpb.ServerMessages_Type {
	if x != nil && x.ServerMessages != nil {
		return *x.ServerMessages
	}
	return mysqlxpb.ServerMessages_OK
}

func (x *Compression) GetClientMessages() mysqlxpb.ClientMessages_Type {
	if x != nil && x.ClientMessages != nil {
		return *x.ClientMessages
	}
	return mysqlxpb.ClientMessages_CON_CAPABILITIES_GET
}

func (x *Compression) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_mysqlx_connection_proto protoreflect.FileDescriptor

var file_mysqlx_connection_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x4d, 0x79, 0x73, 0x71, 0x6c,
	0x78, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x57, 0x0a, 0x0c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x41, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x3a, 0x04, 0x90, 0xea, 0x30, 0x02, 0x22, 0x17, 0x0a, 0x0f, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x3a, 0x04, 0x88,
	0xea, 0x30, 0x01, 0x22, 0x5c, 0x0a, 0x0f, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x53, 0x65, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x4d,
	0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x04, 0x88, 0xea, 0x30,
	0x02, 0x22, 0x0d, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x3a, 0x04, 0x88, 0xea, 0x30, 0x03,
	0x22, 0xea, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x2b, 0x0a, 0x11, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x75, 0x6e, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x44, 0x0a,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x4d,
	0x79, 0x73, 0x71, 0x6c, 0x78, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x3a, 0x08, 0x90, 0xea, 0x30, 0x13, 0x88, 0xea, 0x30, 0x2e, 0x42, 0x6e, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x63, 0x6a, 0x2e, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x64, 0x65, 0x6e, 0x67, 0x37, 0x2f, 0x67, 0x6f, 0x2d, 0x6d,
	0x79, 0x73, 0x71, 0x6c, 0x78, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x79, 0x73,
	0x71, 0x6c, 0x78, 0x70, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78, 0x70, 0x62, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x78,
	0x70, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
}

var (
	file_mysqlx_connection_proto_rawDescOnce sync.Once
	file_mysqlx_connection_proto_rawDescData = file_mysqlx_connection_proto_rawDesc
)

func file_mysqlx_connection_proto_rawDescGZIP() []byte {
	file_mysqlx_connection_proto_rawDescOnce.Do(func() {
		file_mysqlx_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_mysqlx_connection_proto_rawDescData)
	})
	return file_mysqlx_connection_proto_rawDescData
}

var file_mysqlx_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mysqlx_connection_proto_goTypes = []interface{}{
	(*Capability)(nil),                // 0: Mysqlx.Connection.Capability
	(*Capabilities)(nil),              // 1: Mysqlx.Connection.Capabilities
	(*CapabilitiesGet)(nil),           // 2: Mysqlx.Connection.CapabilitiesGet
	(*CapabilitiesSet)(nil),           // 3: Mysqlx.Connection.CapabilitiesSet
	(*Close)(nil),                     // 4: Mysqlx.Connection.Close
	(*Compression)(nil),               // 5: Mysqlx.Connection.Compression
	(*mysqlxpb_datatypes.Any)(nil),    // 6: Mysqlx.Datatypes.Any
	(mysqlxpb.ServerMessages_Type)(0), // 7: Mysqlx.ServerMessages.Type
	(mysqlxpb.ClientMessages_Type)(0), // 8: Mysqlx.ClientMessages.Type
}
var file_mysqlx_connection_proto_depIdxs = []int32{
	6, // 0: Mysqlx.Connection.Capability.value:type_name -> Mysqlx.Datatypes.Any
	0, // 1: Mysqlx.Connection.Capabilities.capabilities:type_name -> Mysqlx.Connection.Capability
	1, // 2: Mysqlx.Connection.CapabilitiesSet.capabilities:type_name -> Mysqlx.Connection.Capabilities
	7, // 3: Mysqlx.Connection.Compression.server_messages:type_name -> Mysqlx.ServerMessages.Type
	8, // 4: Mysqlx.Connection.Compression.client_messages:type_name -> Mysqlx.ClientMessages.Type
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mysqlx_connection_proto_init() }
func file_mysqlx_connection_proto_init() {
	if File_mysqlx_connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mysqlx_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capability); i {
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
		file_mysqlx_connection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capabilities); i {
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
		file_mysqlx_connection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilitiesGet); i {
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
		file_mysqlx_connection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilitiesSet); i {
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
		file_mysqlx_connection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Close); i {
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
		file_mysqlx_connection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Compression); i {
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
			RawDescriptor: file_mysqlx_connection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mysqlx_connection_proto_goTypes,
		DependencyIndexes: file_mysqlx_connection_proto_depIdxs,
		MessageInfos:      file_mysqlx_connection_proto_msgTypes,
	}.Build()
	File_mysqlx_connection_proto = out.File
	file_mysqlx_connection_proto_rawDesc = nil
	file_mysqlx_connection_proto_goTypes = nil
	file_mysqlx_connection_proto_depIdxs = nil
}