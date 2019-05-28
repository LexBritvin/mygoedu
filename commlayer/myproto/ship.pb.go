// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ship.proto

package myproto

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

type Ship struct {
	Shipname             string             `protobuf:"bytes,1,opt,name=shipname,proto3" json:"shipname,omitempty"`
	CaptainName          string             `protobuf:"bytes,2,opt,name=CaptainName,proto3" json:"CaptainName,omitempty"`
	Crew                 []*Ship_CrewMember `protobuf:"bytes,3,rep,name=crew,proto3" json:"crew,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Ship) Reset()         { *m = Ship{} }
func (m *Ship) String() string { return proto.CompactTextString(m) }
func (*Ship) ProtoMessage()    {}
func (*Ship) Descriptor() ([]byte, []int) {
	return fileDescriptor_315cb18542033454, []int{0}
}

func (m *Ship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ship.Unmarshal(m, b)
}
func (m *Ship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ship.Marshal(b, m, deterministic)
}
func (m *Ship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ship.Merge(m, src)
}
func (m *Ship) XXX_Size() int {
	return xxx_messageInfo_Ship.Size(m)
}
func (m *Ship) XXX_DiscardUnknown() {
	xxx_messageInfo_Ship.DiscardUnknown(m)
}

var xxx_messageInfo_Ship proto.InternalMessageInfo

func (m *Ship) GetShipname() string {
	if m != nil {
		return m.Shipname
	}
	return ""
}

func (m *Ship) GetCaptainName() string {
	if m != nil {
		return m.CaptainName
	}
	return ""
}

func (m *Ship) GetCrew() []*Ship_CrewMember {
	if m != nil {
		return m.Crew
	}
	return nil
}

type Ship_CrewMember struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SecClearance         int32    `protobuf:"varint,3,opt,name=secClearance,proto3" json:"secClearance,omitempty"`
	Position             string   `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ship_CrewMember) Reset()         { *m = Ship_CrewMember{} }
func (m *Ship_CrewMember) String() string { return proto.CompactTextString(m) }
func (*Ship_CrewMember) ProtoMessage()    {}
func (*Ship_CrewMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_315cb18542033454, []int{0, 0}
}

func (m *Ship_CrewMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ship_CrewMember.Unmarshal(m, b)
}
func (m *Ship_CrewMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ship_CrewMember.Marshal(b, m, deterministic)
}
func (m *Ship_CrewMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ship_CrewMember.Merge(m, src)
}
func (m *Ship_CrewMember) XXX_Size() int {
	return xxx_messageInfo_Ship_CrewMember.Size(m)
}
func (m *Ship_CrewMember) XXX_DiscardUnknown() {
	xxx_messageInfo_Ship_CrewMember.DiscardUnknown(m)
}

var xxx_messageInfo_Ship_CrewMember proto.InternalMessageInfo

func (m *Ship_CrewMember) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Ship_CrewMember) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ship_CrewMember) GetSecClearance() int32 {
	if m != nil {
		return m.SecClearance
	}
	return 0
}

func (m *Ship_CrewMember) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func init() {
	proto.RegisterType((*Ship)(nil), "myproto.Ship")
	proto.RegisterType((*Ship_CrewMember)(nil), "myproto.Ship.CrewMember")
}

func init() { proto.RegisterFile("ship.proto", fileDescriptor_315cb18542033454) }

var fileDescriptor_315cb18542033454 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x4f, 0xcd, 0x8a, 0x83, 0x30,
	0x10, 0x46, 0xcd, 0x2e, 0xbb, 0xe3, 0xb2, 0x87, 0x39, 0x05, 0x7b, 0x11, 0x4f, 0x9e, 0x52, 0x68,
	0x1f, 0xc1, 0x73, 0x7b, 0xb0, 0x4f, 0x10, 0x35, 0xe0, 0x40, 0x4d, 0x42, 0x14, 0x4a, 0x9f, 0xb8,
	0xaf, 0x51, 0x32, 0xa5, 0xb5, 0xbd, 0x7d, 0x7f, 0xcc, 0x37, 0x1f, 0xc0, 0x3c, 0x92, 0x57, 0x3e,
	0xb8, 0xc5, 0x21, 0x8c, 0xd7, 0x21, 0x68, 0xc6, 0xd5, 0x2d, 0x01, 0x71, 0x1a, 0xc9, 0x63, 0x01,
	0x3f, 0x31, 0x62, 0xf5, 0x64, 0x64, 0x52, 0x26, 0xf5, 0x6f, 0xfb, 0xe2, 0x58, 0x42, 0xde, 0x68,
	0xbf, 0x68, 0xb2, 0xc7, 0x68, 0xa7, 0x6c, 0xbf, 0x4b, 0xb8, 0x05, 0xd1, 0x04, 0x73, 0x91, 0x59,
	0x99, 0xd5, 0xf9, 0x6e, 0xa3, 0xd6, 0x06, 0x15, 0xaf, 0xab, 0x68, 0x1e, 0xcc, 0xd4, 0x99, 0xd0,
	0x72, 0xb0, 0xf0, 0x00, 0xab, 0x86, 0xff, 0x90, 0xd2, 0xc0, 0xb5, 0x5f, 0x6d, 0x4a, 0x03, 0x22,
	0x08, 0xbb, 0x36, 0x31, 0xc6, 0x0a, 0xfe, 0x66, 0xd3, 0x37, 0x67, 0xa3, 0x83, 0xb6, 0xbd, 0x91,
	0x19, 0xa7, 0x3f, 0xb4, 0x38, 0xc2, 0xbb, 0x99, 0x16, 0x72, 0x56, 0x8a, 0xc7, 0x88, 0x27, 0xef,
	0xbe, 0xf9, 0x9d, 0xfd, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xe8, 0x35, 0x2a, 0x0a, 0x01, 0x00,
	0x00,
}
