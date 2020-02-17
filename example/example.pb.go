// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/example.proto

package example

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/wores/protoc-gen-pipeline/pipeline"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Example struct {
	Text                 string                `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	WrapText             *wrappers.StringValue `protobuf:"bytes,2,opt,name=wrapText,proto3" json:"wrapText,omitempty"`
	Texts                []string              `protobuf:"bytes,3,rep,name=texts,proto3" json:"texts,omitempty"`
	Inner                *Example_Inner        `protobuf:"bytes,4,opt,name=inner,proto3" json:"inner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Example) Reset()         { *m = Example{} }
func (m *Example) String() string { return proto.CompactTextString(m) }
func (*Example) ProtoMessage()    {}
func (*Example) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c78cffa5d645ba4, []int{0}
}
func (m *Example) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Example.Unmarshal(m, b)
}
func (m *Example) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Example.Marshal(b, m, deterministic)
}
func (m *Example) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example.Merge(m, src)
}
func (m *Example) XXX_Size() int {
	return xxx_messageInfo_Example.Size(m)
}
func (m *Example) XXX_DiscardUnknown() {
	xxx_messageInfo_Example.DiscardUnknown(m)
}

var xxx_messageInfo_Example proto.InternalMessageInfo

func (m *Example) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Example) GetWrapText() *wrappers.StringValue {
	if m != nil {
		return m.WrapText
	}
	return nil
}

func (m *Example) GetTexts() []string {
	if m != nil {
		return m.Texts
	}
	return nil
}

func (m *Example) GetInner() *Example_Inner {
	if m != nil {
		return m.Inner
	}
	return nil
}

type Example_Inner struct {
	Text                 string                `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	WrapText             *wrappers.StringValue `protobuf:"bytes,2,opt,name=wrapText,proto3" json:"wrapText,omitempty"`
	Texts                []string              `protobuf:"bytes,3,rep,name=texts,proto3" json:"texts,omitempty"`
	Inner                *Example_Inner        `protobuf:"bytes,4,opt,name=inner,proto3" json:"inner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Example_Inner) Reset()         { *m = Example_Inner{} }
func (m *Example_Inner) String() string { return proto.CompactTextString(m) }
func (*Example_Inner) ProtoMessage()    {}
func (*Example_Inner) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c78cffa5d645ba4, []int{0, 0}
}
func (m *Example_Inner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Example_Inner.Unmarshal(m, b)
}
func (m *Example_Inner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Example_Inner.Marshal(b, m, deterministic)
}
func (m *Example_Inner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Example_Inner.Merge(m, src)
}
func (m *Example_Inner) XXX_Size() int {
	return xxx_messageInfo_Example_Inner.Size(m)
}
func (m *Example_Inner) XXX_DiscardUnknown() {
	xxx_messageInfo_Example_Inner.DiscardUnknown(m)
}

var xxx_messageInfo_Example_Inner proto.InternalMessageInfo

func (m *Example_Inner) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Example_Inner) GetWrapText() *wrappers.StringValue {
	if m != nil {
		return m.WrapText
	}
	return nil
}

func (m *Example_Inner) GetTexts() []string {
	if m != nil {
		return m.Texts
	}
	return nil
}

func (m *Example_Inner) GetInner() *Example_Inner {
	if m != nil {
		return m.Inner
	}
	return nil
}

func init() {
	proto.RegisterType((*Example)(nil), "example.Example")
	proto.RegisterType((*Example_Inner)(nil), "example.Example.Inner")
}

func init() { proto.RegisterFile("example/example.proto", fileDescriptor_1c78cffa5d645ba4) }

var fileDescriptor_1c78cffa5d645ba4 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x94, 0x5c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x38, 0xa9, 0x34, 0x4d, 0xbf, 0xbc,
	0x28, 0xb1, 0xa0, 0x20, 0xb5, 0xa8, 0x18, 0xa2, 0x50, 0x4a, 0xbc, 0x20, 0xb3, 0x20, 0x35, 0x27,
	0x33, 0x2f, 0x55, 0x1f, 0xc6, 0x80, 0x48, 0x28, 0xfd, 0x66, 0xe2, 0x62, 0x77, 0x85, 0x18, 0x22,
	0x24, 0xc3, 0xc5, 0x52, 0x92, 0x5a, 0x51, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xb1,
	0x6b, 0x36, 0x2f, 0x8b, 0x10, 0x93, 0x00, 0x63, 0x10, 0x58, 0x54, 0xc8, 0x89, 0x8b, 0x03, 0x64,
	0x68, 0x08, 0x48, 0x05, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x8c, 0x1e, 0xc4, 0x56, 0x3d, 0x98,
	0xad, 0x7a, 0xc1, 0x25, 0x45, 0x99, 0x79, 0xe9, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x48, 0xfa, 0xe1,
	0xfa, 0x84, 0x94, 0xb8, 0x58, 0x41, 0x66, 0x15, 0x4b, 0x30, 0x2b, 0x30, 0x6b, 0x70, 0x3a, 0xf1,
	0xec, 0x9a, 0xcd, 0xcb, 0x21, 0xc5, 0xc6, 0x05, 0x51, 0x06, 0x91, 0x12, 0xd2, 0xe1, 0x62, 0xcd,
	0xcc, 0xcb, 0x4b, 0x2d, 0x92, 0x60, 0x01, 0x5b, 0x22, 0xa6, 0x07, 0xf3, 0x32, 0xd4, 0x99, 0x7a,
	0x9e, 0x20, 0xd9, 0x20, 0x88, 0x22, 0xa9, 0xdd, 0x8c, 0x5c, 0xac, 0x60, 0x81, 0xa1, 0xe8, 0x7a,
	0x27, 0xbd, 0x28, 0x9d, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf2,
	0xfc, 0xa2, 0xd4, 0x62, 0x48, 0x14, 0x26, 0xeb, 0xa6, 0xa7, 0xe6, 0xe9, 0xc2, 0x23, 0x0d, 0x6a,
	0x48, 0x12, 0x1b, 0x58, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x93, 0x91, 0x60, 0xe9, 0x0f,
	0x02, 0x00, 0x00,
}
