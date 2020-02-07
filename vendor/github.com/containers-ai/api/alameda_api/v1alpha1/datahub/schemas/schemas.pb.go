// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/schemas/schemas.proto

package schemas

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

type Schema struct {
	SchemaMeta           *SchemaMeta    `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	Measurements         []*Measurement `protobuf:"bytes,2,rep,name=measurements,proto3" json:"measurements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_85ffbc323974f9ff, []int{0}
}

func (m *Schema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schema.Unmarshal(m, b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return xxx_messageInfo_Schema.Size(m)
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetSchemaMeta() *SchemaMeta {
	if m != nil {
		return m.SchemaMeta
	}
	return nil
}

func (m *Schema) GetMeasurements() []*Measurement {
	if m != nil {
		return m.Measurements
	}
	return nil
}

func init() {
	proto.RegisterType((*Schema)(nil), "containersai.alameda.v1alpha1.datahub.schemas.Schema")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/schemas/schemas.proto", fileDescriptor_85ffbc323974f9ff)
}

var fileDescriptor_85ffbc323974f9ff = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x4f, 0x87, 0x30,
	0x10, 0xc5, 0x83, 0x26, 0xff, 0xa1, 0x38, 0x31, 0x11, 0x27, 0xe2, 0xc4, 0x42, 0x2b, 0x38, 0xe9,
	0xa6, 0xce, 0x2c, 0xb8, 0x31, 0x48, 0x0e, 0xb8, 0x48, 0x13, 0x4a, 0x1b, 0x7a, 0x98, 0xf8, 0xfd,
	0xfc, 0x60, 0x06, 0x28, 0x92, 0xff, 0x06, 0xd3, 0x5d, 0x2e, 0xef, 0xfd, 0xde, 0xe5, 0xb1, 0x0c,
	0x7a, 0x50, 0xd8, 0x42, 0x05, 0x46, 0x8a, 0xef, 0x14, 0x7a, 0xd3, 0x41, 0x2a, 0x5a, 0x20, 0xe8,
	0xa6, 0x5a, 0xd8, 0xa6, 0x43, 0x05, 0x76, 0x9b, 0xdc, 0x8c, 0x9a, 0x74, 0x90, 0x34, 0x7a, 0x20,
	0x90, 0x03, 0x8e, 0x16, 0x24, 0x77, 0x00, 0xbe, 0x99, 0xb9, 0x33, 0x73, 0x67, 0xba, 0x7f, 0x3c,
	0x14, 0x41, 0x3f, 0x06, 0x5d, 0xc0, 0xc3, 0xaf, 0xc7, 0x2e, 0x1f, 0xcb, 0x3d, 0x28, 0x99, 0xbf,
	0x2a, 0x2a, 0x85, 0x04, 0xa1, 0x17, 0x79, 0xb1, 0x9f, 0x3d, 0xf3, 0x53, 0x1f, 0xf0, 0x95, 0x95,
	0x23, 0x41, 0xc1, 0xec, 0xff, 0x1e, 0x7c, 0xb2, 0x3b, 0x85, 0x60, 0xa7, 0x11, 0x15, 0x0e, 0x64,
	0xc3, 0x9b, 0xe8, 0x36, 0xf6, 0xb3, 0x97, 0x93, 0xf0, 0x7c, 0x47, 0x14, 0x57, 0xbc, 0xb7, 0xf7,
	0xf2, 0xf5, 0x4b, 0xd2, 0xac, 0x6d, 0xb4, 0x12, 0x3b, 0x35, 0x01, 0x29, 0xe6, 0x2e, 0x8e, 0xf4,
	0x52, 0x5f, 0x96, 0x4a, 0x9e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x2b, 0x72, 0x89, 0xa9,
	0x01, 0x00, 0x00,
}