// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ID struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_e121d07081ab1e22, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

type DateTime struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateTime) Reset()         { *m = DateTime{} }
func (m *DateTime) String() string { return proto.CompactTextString(m) }
func (*DateTime) ProtoMessage()    {}
func (*DateTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_e121d07081ab1e22, []int{1}
}
func (m *DateTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateTime.Unmarshal(m, b)
}
func (m *DateTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateTime.Marshal(b, m, deterministic)
}
func (dst *DateTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateTime.Merge(dst, src)
}
func (m *DateTime) XXX_Size() int {
	return xxx_messageInfo_DateTime.Size(m)
}
func (m *DateTime) XXX_DiscardUnknown() {
	xxx_messageInfo_DateTime.DiscardUnknown(m)
}

var xxx_messageInfo_DateTime proto.InternalMessageInfo

type Checksum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checksum) Reset()         { *m = Checksum{} }
func (m *Checksum) String() string { return proto.CompactTextString(m) }
func (*Checksum) ProtoMessage()    {}
func (*Checksum) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_e121d07081ab1e22, []int{2}
}
func (m *Checksum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checksum.Unmarshal(m, b)
}
func (m *Checksum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checksum.Marshal(b, m, deterministic)
}
func (dst *Checksum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checksum.Merge(dst, src)
}
func (m *Checksum) XXX_Size() int {
	return xxx_messageInfo_Checksum.Size(m)
}
func (m *Checksum) XXX_DiscardUnknown() {
	xxx_messageInfo_Checksum.DiscardUnknown(m)
}

var xxx_messageInfo_Checksum proto.InternalMessageInfo

type UID struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UID) Reset()         { *m = UID{} }
func (m *UID) String() string { return proto.CompactTextString(m) }
func (*UID) ProtoMessage()    {}
func (*UID) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_e121d07081ab1e22, []int{3}
}
func (m *UID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UID.Unmarshal(m, b)
}
func (m *UID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UID.Marshal(b, m, deterministic)
}
func (dst *UID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UID.Merge(dst, src)
}
func (m *UID) XXX_Size() int {
	return xxx_messageInfo_UID.Size(m)
}
func (m *UID) XXX_DiscardUnknown() {
	xxx_messageInfo_UID.DiscardUnknown(m)
}

var xxx_messageInfo_UID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ID)(nil), "proto.ID")
	proto.RegisterType((*DateTime)(nil), "proto.DateTime")
	proto.RegisterType((*Checksum)(nil), "proto.Checksum")
	proto.RegisterType((*UID)(nil), "proto.UID")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_types_e121d07081ab1e22) }

var fileDescriptor_types_e121d07081ab1e22 = []byte{
	// 80 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x2c, 0x5c, 0x4c, 0x9e,
	0x2e, 0x4a, 0x5c, 0x5c, 0x1c, 0x2e, 0x89, 0x25, 0xa9, 0x21, 0x99, 0xb9, 0xa9, 0x20, 0xb6, 0x73,
	0x46, 0x6a, 0x72, 0x76, 0x71, 0x69, 0xae, 0x12, 0x2b, 0x17, 0x73, 0xa8, 0xa7, 0x4b, 0x12, 0x1b,
	0x58, 0xad, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x3c, 0xfe, 0xb2, 0x41, 0x00, 0x00, 0x00,
}
