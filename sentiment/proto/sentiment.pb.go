// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sentiment.proto

package sentiment

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

type Request struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab79e813d74ef963, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Request) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type Response struct {
	Score                float64  `protobuf:"fixed64,1,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab79e813d74ef963, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "sentiment.Request")
	proto.RegisterType((*Response)(nil), "sentiment.Response")
}

func init() { proto.RegisterFile("proto/sentiment.proto", fileDescriptor_ab79e813d74ef963) }

var fileDescriptor_ab79e813d74ef963 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xcd, 0x2b, 0xc9, 0xcc, 0x4d, 0xcd, 0x2b, 0xd1, 0x03, 0xf3, 0x85, 0x38,
	0xe1, 0x02, 0x4a, 0x86, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c,
	0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0x36, 0x48, 0x2c,
	0x27, 0x31, 0x2f, 0x5d, 0x82, 0x09, 0x22, 0x06, 0x62, 0x2b, 0x29, 0x70, 0x71, 0x04, 0xa5, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0x16, 0x27, 0xe7, 0x17, 0xa5, 0x82, 0x35,
	0x31, 0x06, 0x41, 0x38, 0x46, 0x8e, 0x5c, 0x9c, 0xc1, 0x30, 0x1b, 0x84, 0x4c, 0xb8, 0xd8, 0x1d,
	0xf3, 0x12, 0x73, 0x2a, 0xab, 0x52, 0x85, 0x84, 0xf4, 0x10, 0x2e, 0x81, 0xda, 0x2a, 0x25, 0x8c,
	0x22, 0x06, 0x31, 0x56, 0x89, 0xc1, 0x49, 0x30, 0x8a, 0x1f, 0xec, 0x56, 0x6b, 0xb8, 0x6c, 0x12,
	0x1b, 0x58, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x55, 0xd3, 0x83, 0xd5, 0x00, 0x00,
	0x00,
}