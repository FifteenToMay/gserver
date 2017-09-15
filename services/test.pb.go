// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	HeartbeatReq
	HeartbeatAck
	ErrorAck
	TouristsLoginReq
	TouristsLoginAck
	CreateRoleReq
	CreateRoleAck
*/
package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 心跳请求0
type HeartbeatReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeartbeatReq) Reset()                    { *m = HeartbeatReq{} }
func (m *HeartbeatReq) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatReq) ProtoMessage()               {}
func (*HeartbeatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 心跳回复1
type HeartbeatAck struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeartbeatAck) Reset()                    { *m = HeartbeatAck{} }
func (m *HeartbeatAck) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatAck) ProtoMessage()               {}
func (*HeartbeatAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 通用的错误回复2
type ErrorAck struct {
	Ackid            *int32 `protobuf:"varint,1,req,name=ackid" json:"ackid,omitempty"`
	Errid            *int32 `protobuf:"varint,2,req,name=errid" json:"errid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ErrorAck) Reset()                    { *m = ErrorAck{} }
func (m *ErrorAck) String() string            { return proto.CompactTextString(m) }
func (*ErrorAck) ProtoMessage()               {}
func (*ErrorAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ErrorAck) GetAckid() int32 {
	if m != nil && m.Ackid != nil {
		return *m.Ackid
	}
	return 0
}

func (m *ErrorAck) GetErrid() int32 {
	if m != nil && m.Errid != nil {
		return *m.Errid
	}
	return 0
}

// 请求游客登录1002
type TouristsLoginReq struct {
	Touristscot      *string `protobuf:"bytes,1,req,name=touristscot" json:"touristscot,omitempty"`
	Touristspwd      *string `protobuf:"bytes,2,req,name=touristspwd" json:"touristspwd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TouristsLoginReq) Reset()                    { *m = TouristsLoginReq{} }
func (m *TouristsLoginReq) String() string            { return proto.CompactTextString(m) }
func (*TouristsLoginReq) ProtoMessage()               {}
func (*TouristsLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TouristsLoginReq) GetTouristscot() string {
	if m != nil && m.Touristscot != nil {
		return *m.Touristscot
	}
	return ""
}

func (m *TouristsLoginReq) GetTouristspwd() string {
	if m != nil && m.Touristspwd != nil {
		return *m.Touristspwd
	}
	return ""
}

// 游客登录返回1003
type TouristsLoginAck struct {
	Str              *string `protobuf:"bytes,1,req,name=str" json:"str,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TouristsLoginAck) Reset()                    { *m = TouristsLoginAck{} }
func (m *TouristsLoginAck) String() string            { return proto.CompactTextString(m) }
func (*TouristsLoginAck) ProtoMessage()               {}
func (*TouristsLoginAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TouristsLoginAck) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

// 创建角色请求1004
type CreateRoleReq struct {
	Icon *string `protobuf:"bytes,1,req,name=icon" json:"icon,omitempty"`
	// 血脉id
	BornId           *uint32 `protobuf:"varint,2,req,name=bornId" json:"bornId,omitempty"`
	Name             *string `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateRoleReq) Reset()                    { *m = CreateRoleReq{} }
func (m *CreateRoleReq) String() string            { return proto.CompactTextString(m) }
func (*CreateRoleReq) ProtoMessage()               {}
func (*CreateRoleReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateRoleReq) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *CreateRoleReq) GetBornId() uint32 {
	if m != nil && m.BornId != nil {
		return *m.BornId
	}
	return 0
}

func (m *CreateRoleReq) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// 创建角色返回1005
type CreateRoleAck struct {
	Str              *string `protobuf:"bytes,1,req,name=str" json:"str,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateRoleAck) Reset()                    { *m = CreateRoleAck{} }
func (m *CreateRoleAck) String() string            { return proto.CompactTextString(m) }
func (*CreateRoleAck) ProtoMessage()               {}
func (*CreateRoleAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateRoleAck) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

func init() {
	proto.RegisterType((*HeartbeatReq)(nil), "services.heartbeat_req")
	proto.RegisterType((*HeartbeatAck)(nil), "services.heartbeat_ack")
	proto.RegisterType((*ErrorAck)(nil), "services.error_ack")
	proto.RegisterType((*TouristsLoginReq)(nil), "services.tourists_login_req")
	proto.RegisterType((*TouristsLoginAck)(nil), "services.tourists_login_ack")
	proto.RegisterType((*CreateRoleReq)(nil), "services.create_role_req")
	proto.RegisterType((*CreateRoleAck)(nil), "services.create_role_ack")
}

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0xf1, 0x5e, 0xaf, 0xd8, 0x91, 0x52, 0x09, 0x22, 0x5d, 0x96, 0x08, 0xe2, 0xca, 0xad,
	0xcf, 0xe0, 0xd2, 0xae, 0xdc, 0x95, 0x34, 0x0e, 0x1a, 0xac, 0x99, 0x3a, 0x19, 0xf5, 0xf5, 0xcd,
	0x4f, 0x85, 0x16, 0xef, 0xaa, 0x73, 0x3e, 0x0e, 0xdf, 0x69, 0x00, 0x04, 0x83, 0xdc, 0xcf, 0x4c,
	0x42, 0xea, 0x3c, 0x20, 0x7f, 0x3b, 0x8b, 0x41, 0x37, 0x50, 0xbf, 0xa1, 0x61, 0x19, 0xd1, 0xc8,
	0xc0, 0xf8, 0xb9, 0x05, 0xc6, 0xbe, 0xeb, 0x07, 0xa8, 0x90, 0x99, 0x38, 0x05, 0x75, 0x05, 0x87,
	0xf8, 0x71, 0x2f, 0xed, 0x49, 0xb7, 0xbb, 0x3b, 0xf4, 0x25, 0x24, 0x1a, 0x2b, 0x91, 0xee, 0x0a,
	0xcd, 0x41, 0x3f, 0x83, 0x12, 0xfa, 0x62, 0x17, 0x24, 0x0c, 0x13, 0xbd, 0x3a, 0x9f, 0xfc, 0xaa,
	0x83, 0x8b, 0x3f, 0x6a, 0x49, 0xb2, 0xa7, 0xea, 0xd7, 0x68, 0xdd, 0x98, 0x7f, 0x8a, 0x73, 0xd5,
	0x88, 0x48, 0xdf, 0xfe, 0x33, 0xa7, 0x7f, 0xbb, 0x84, 0x7d, 0x10, 0x5e, 0x8c, 0xe9, 0xd4, 0x4f,
	0xd0, 0x58, 0x8e, 0xef, 0xc0, 0x81, 0x69, 0xc2, 0x3c, 0xaf, 0xe0, 0xd4, 0x59, 0xf2, 0x4b, 0x2b,
	0xdf, 0xea, 0x1a, 0xce, 0x46, 0x62, 0xff, 0x58, 0xb6, 0xea, 0x7e, 0x49, 0xa9, 0xeb, 0xcd, 0x07,
	0xb6, 0xfb, 0xd2, 0x4d, 0xb7, 0xbe, 0xd9, 0x2a, 0x8f, 0xee, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0xf9, 0xce, 0x22, 0x6b, 0x01, 0x00, 0x00,
}
