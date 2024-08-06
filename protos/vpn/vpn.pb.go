// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: vpn.proto

package vpn

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

type AuthType int32

const (
	AuthType_NONE AuthType = 0
	AuthType_XOR  AuthType = 1
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "NONE",
		1: "XOR",
	}
	AuthType_value = map[string]int32{
		"NONE": 0,
		"XOR":  1,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_vpn_proto_enumTypes[0].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_vpn_proto_enumTypes[0]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{0}
}

type StatusCode int32

const (
	StatusCode_OK                    StatusCode = 0
	StatusCode_UNKNOWN_ERROR         StatusCode = 1
	StatusCode_INVALID_AUTHORIZATION StatusCode = 2
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "OK",
		1: "UNKNOWN_ERROR",
		2: "INVALID_AUTHORIZATION",
	}
	StatusCode_value = map[string]int32{
		"OK":                    0,
		"UNKNOWN_ERROR":         1,
		"INVALID_AUTHORIZATION": 2,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_vpn_proto_enumTypes[1].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_vpn_proto_enumTypes[1]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{1}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType AuthType `protobuf:"varint,1,opt,name=auth_type,json=authType,proto3,enum=AuthType" json:"auth_type,omitempty"`
	Data     string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetAuthType() AuthType {
	if x != nil {
		return x.AuthType
	}
	return AuthType_NONE
}

func (x *ConnectRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TunnelId string     `protobuf:"bytes,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	Data     string     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Code     StatusCode `protobuf:"varint,3,opt,name=code,proto3,enum=StatusCode" json:"code,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectResponse) GetTunnelId() string {
	if x != nil {
		return x.TunnelId
	}
	return ""
}

func (x *ConnectResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *ConnectResponse) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_OK
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{2}
}

func (x *DisconnectRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=StatusCode" json:"code,omitempty"`
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{3}
}

func (x *DisconnectResponse) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_OK
}

type RouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TunnelId string   `protobuf:"bytes,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	Route    []string `protobuf:"bytes,2,rep,name=route,proto3" json:"route,omitempty"`
}

func (x *RouteRequest) Reset() {
	*x = RouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRequest) ProtoMessage() {}

func (x *RouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRequest.ProtoReflect.Descriptor instead.
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{4}
}

func (x *RouteRequest) GetTunnelId() string {
	if x != nil {
		return x.TunnelId
	}
	return ""
}

func (x *RouteRequest) GetRoute() []string {
	if x != nil {
		return x.Route
	}
	return nil
}

type RouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=StatusCode" json:"code,omitempty"`
}

func (x *RouteResponse) Reset() {
	*x = RouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteResponse) ProtoMessage() {}

func (x *RouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteResponse.ProtoReflect.Descriptor instead.
func (*RouteResponse) Descriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{5}
}

func (x *RouteResponse) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_OK
}

type DataPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TunnelId string `protobuf:"bytes,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataPacket) Reset() {
	*x = DataPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vpn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPacket) ProtoMessage() {}

func (x *DataPacket) ProtoReflect() protoreflect.Message {
	mi := &file_vpn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPacket.ProtoReflect.Descriptor instead.
func (*DataPacket) Descriptor() ([]byte, []int) {
	return file_vpn_proto_rawDescGZIP(), []int{6}
}

func (x *DataPacket) GetTunnelId() string {
	if x != nil {
		return x.TunnelId
	}
	return ""
}

func (x *DataPacket) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_vpn_proto protoreflect.FileDescriptor

var file_vpn_proto_rawDesc = []byte{
	0x0a, 0x09, 0x76, 0x70, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x09, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x63, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x27,
	0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x41,
	0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x22, 0x30, 0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x3d, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x2a, 0x1d, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x4f, 0x52, 0x10,
	0x01, 0x2a, 0x42, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x02, 0x32, 0xfa, 0x01, 0x0a, 0x03, 0x56, 0x50, 0x4e, 0x12, 0x2c, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x0f, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0b, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x1a, 0x0b, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6b, 0x75, 0x6d, 0x69, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x76, 0x70, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vpn_proto_rawDescOnce sync.Once
	file_vpn_proto_rawDescData = file_vpn_proto_rawDesc
)

func file_vpn_proto_rawDescGZIP() []byte {
	file_vpn_proto_rawDescOnce.Do(func() {
		file_vpn_proto_rawDescData = protoimpl.X.CompressGZIP(file_vpn_proto_rawDescData)
	})
	return file_vpn_proto_rawDescData
}

var file_vpn_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_vpn_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_vpn_proto_goTypes = []any{
	(AuthType)(0),              // 0: AuthType
	(StatusCode)(0),            // 1: StatusCode
	(*ConnectRequest)(nil),     // 2: ConnectRequest
	(*ConnectResponse)(nil),    // 3: ConnectResponse
	(*DisconnectRequest)(nil),  // 4: DisconnectRequest
	(*DisconnectResponse)(nil), // 5: DisconnectResponse
	(*RouteRequest)(nil),       // 6: RouteRequest
	(*RouteResponse)(nil),      // 7: RouteResponse
	(*DataPacket)(nil),         // 8: DataPacket
}
var file_vpn_proto_depIdxs = []int32{
	0, // 0: ConnectRequest.auth_type:type_name -> AuthType
	1, // 1: ConnectResponse.code:type_name -> StatusCode
	1, // 2: DisconnectResponse.code:type_name -> StatusCode
	1, // 3: RouteResponse.code:type_name -> StatusCode
	2, // 4: VPN.Connect:input_type -> ConnectRequest
	4, // 5: VPN.Disconnect:input_type -> DisconnectRequest
	6, // 6: VPN.RegisterRoute:input_type -> RouteRequest
	6, // 7: VPN.UnregisterRoute:input_type -> RouteRequest
	8, // 8: VPN.TransferData:input_type -> DataPacket
	3, // 9: VPN.Connect:output_type -> ConnectResponse
	5, // 10: VPN.Disconnect:output_type -> DisconnectResponse
	7, // 11: VPN.RegisterRoute:output_type -> RouteResponse
	7, // 12: VPN.UnregisterRoute:output_type -> RouteResponse
	8, // 13: VPN.TransferData:output_type -> DataPacket
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vpn_proto_init() }
func file_vpn_proto_init() {
	if File_vpn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vpn_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectRequest); i {
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
		file_vpn_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectResponse); i {
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
		file_vpn_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DisconnectRequest); i {
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
		file_vpn_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DisconnectResponse); i {
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
		file_vpn_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RouteRequest); i {
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
		file_vpn_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RouteResponse); i {
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
		file_vpn_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DataPacket); i {
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
			RawDescriptor: file_vpn_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vpn_proto_goTypes,
		DependencyIndexes: file_vpn_proto_depIdxs,
		EnumInfos:         file_vpn_proto_enumTypes,
		MessageInfos:      file_vpn_proto_msgTypes,
	}.Build()
	File_vpn_proto = out.File
	file_vpn_proto_rawDesc = nil
	file_vpn_proto_goTypes = nil
	file_vpn_proto_depIdxs = nil
}