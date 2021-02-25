// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jobrequest.proto

package jobrequest

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

type JobRequest struct {
	Command              string   `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobRequest) Reset()         { *m = JobRequest{} }
func (m *JobRequest) String() string { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()    {}
func (*JobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_jobrequest_a0e6f9d23c3914ef, []int{0}
}
func (m *JobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobRequest.Unmarshal(m, b)
}
func (m *JobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobRequest.Marshal(b, m, deterministic)
}
func (dst *JobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRequest.Merge(dst, src)
}
func (m *JobRequest) XXX_Size() int {
	return xxx_messageInfo_JobRequest.Size(m)
}
func (m *JobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobRequest proto.InternalMessageInfo

func (m *JobRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *JobRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*JobRequest)(nil), "JobRequest")
}

func init() { proto.RegisterFile("jobrequest.proto", fileDescriptor_jobrequest_a0e6f9d23c3914ef) }

var fileDescriptor_jobrequest_a0e6f9d23c3914ef = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xca, 0x4f, 0x2a,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe2, 0xe2,
	0xf2, 0xca, 0x4f, 0x0a, 0x82, 0x88, 0x09, 0x49, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0xe6, 0x26, 0xe6,
	0xa5, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89,
	0x25, 0x89, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0x76, 0x12, 0x1b, 0xd8, 0x08, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xe5, 0xde, 0x4a, 0x56, 0x00, 0x00, 0x00,
}