// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bio-routing/bio-rd/net/api/net.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IP_Version int32

const (
	IP_IPv4 IP_Version = 0
	IP_IPv6 IP_Version = 1
)

var IP_Version_name = map[int32]string{
	0: "IPv4",
	1: "IPv6",
}

var IP_Version_value = map[string]int32{
	"IPv4": 0,
	"IPv6": 1,
}

func (x IP_Version) String() string {
	return proto.EnumName(IP_Version_name, int32(x))
}

func (IP_Version) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e879b68d7a71dcc0, []int{1, 0}
}

type Prefix struct {
	Address              *IP      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Pfxlen               uint32   `protobuf:"varint,2,opt,name=pfxlen,proto3" json:"pfxlen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prefix) Reset()         { *m = Prefix{} }
func (m *Prefix) String() string { return proto.CompactTextString(m) }
func (*Prefix) ProtoMessage()    {}
func (*Prefix) Descriptor() ([]byte, []int) {
	return fileDescriptor_e879b68d7a71dcc0, []int{0}
}

func (m *Prefix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prefix.Unmarshal(m, b)
}
func (m *Prefix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prefix.Marshal(b, m, deterministic)
}
func (m *Prefix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prefix.Merge(m, src)
}
func (m *Prefix) XXX_Size() int {
	return xxx_messageInfo_Prefix.Size(m)
}
func (m *Prefix) XXX_DiscardUnknown() {
	xxx_messageInfo_Prefix.DiscardUnknown(m)
}

var xxx_messageInfo_Prefix proto.InternalMessageInfo

func (m *Prefix) GetAddress() *IP {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Prefix) GetPfxlen() uint32 {
	if m != nil {
		return m.Pfxlen
	}
	return 0
}

type IP struct {
	Higher               uint64     `protobuf:"varint,1,opt,name=higher,proto3" json:"higher,omitempty"`
	Lower                uint64     `protobuf:"varint,2,opt,name=lower,proto3" json:"lower,omitempty"`
	Version              IP_Version `protobuf:"varint,3,opt,name=version,proto3,enum=bio.net.IP_Version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IP) Reset()         { *m = IP{} }
func (m *IP) String() string { return proto.CompactTextString(m) }
func (*IP) ProtoMessage()    {}
func (*IP) Descriptor() ([]byte, []int) {
	return fileDescriptor_e879b68d7a71dcc0, []int{1}
}

func (m *IP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IP.Unmarshal(m, b)
}
func (m *IP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IP.Marshal(b, m, deterministic)
}
func (m *IP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IP.Merge(m, src)
}
func (m *IP) XXX_Size() int {
	return xxx_messageInfo_IP.Size(m)
}
func (m *IP) XXX_DiscardUnknown() {
	xxx_messageInfo_IP.DiscardUnknown(m)
}

var xxx_messageInfo_IP proto.InternalMessageInfo

func (m *IP) GetHigher() uint64 {
	if m != nil {
		return m.Higher
	}
	return 0
}

func (m *IP) GetLower() uint64 {
	if m != nil {
		return m.Lower
	}
	return 0
}

func (m *IP) GetVersion() IP_Version {
	if m != nil {
		return m.Version
	}
	return IP_IPv4
}

func init() {
	proto.RegisterEnum("bio.net.IP_Version", IP_Version_name, IP_Version_value)
	proto.RegisterType((*Prefix)(nil), "bio.net.Prefix")
	proto.RegisterType((*IP)(nil), "bio.net.IP")
}

func init() {
	proto.RegisterFile("github.com/bio-routing/bio-rd/net/api/net.proto", fileDescriptor_e879b68d7a71dcc0)
}

var fileDescriptor_e879b68d7a71dcc0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xca, 0xcc, 0xd7, 0x2d, 0xca, 0x2f, 0x2d, 0xc9,
	0xcc, 0x4b, 0x87, 0xb0, 0x53, 0xf4, 0xf3, 0x52, 0x4b, 0xf4, 0x13, 0x0b, 0x32, 0x41, 0xb4, 0x5e,
	0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x7b, 0x52, 0x66, 0xbe, 0x5e, 0x5e, 0x6a, 0x89, 0x92, 0x3b,
	0x17, 0x5b, 0x40, 0x51, 0x6a, 0x5a, 0x66, 0x85, 0x90, 0x2a, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51,
	0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xb7, 0x1e, 0x54, 0x91, 0x9e, 0x67,
	0x40, 0x10, 0x4c, 0x4e, 0x48, 0x8c, 0x8b, 0xad, 0x20, 0xad, 0x22, 0x27, 0x35, 0x4f, 0x82, 0x49,
	0x81, 0x51, 0x83, 0x37, 0x08, 0xca, 0x53, 0x6a, 0x60, 0xe4, 0x62, 0xf2, 0x0c, 0x00, 0x49, 0x67,
	0x64, 0xa6, 0x67, 0xa4, 0x16, 0x81, 0x0d, 0x61, 0x09, 0x82, 0xf2, 0x84, 0x44, 0xb8, 0x58, 0x73,
	0xf2, 0xcb, 0x53, 0x8b, 0xc0, 0xba, 0x58, 0x82, 0x20, 0x1c, 0x21, 0x5d, 0x2e, 0xf6, 0xb2, 0xd4,
	0xa2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x61, 0x24, 0x3b, 0xf5,
	0xc2, 0x20, 0x52, 0x41, 0x30, 0x35, 0x4a, 0xb2, 0x5c, 0xec, 0x50, 0x31, 0x21, 0x0e, 0x2e, 0x16,
	0xcf, 0x80, 0x32, 0x13, 0x01, 0x06, 0x28, 0xcb, 0x4c, 0x80, 0xd1, 0x49, 0x3d, 0x4a, 0x95, 0xa8,
	0x70, 0x48, 0x62, 0x03, 0x07, 0x82, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xba, 0xb6, 0x37,
	0x37, 0x01, 0x00, 0x00,
}
