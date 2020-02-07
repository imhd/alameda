// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/common/types.proto

package common

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

type ResourceBoundary int32

const (
	ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED ResourceBoundary = 0
	ResourceBoundary_RESOURCE_RAW                ResourceBoundary = 1
	ResourceBoundary_RESOURCE_UPPER_BOUND        ResourceBoundary = 2
	ResourceBoundary_RESOURCE_LOWER_BOUND        ResourceBoundary = 3
)

var ResourceBoundary_name = map[int32]string{
	0: "RESOURCE_BOUNDARY_UNDEFINED",
	1: "RESOURCE_RAW",
	2: "RESOURCE_UPPER_BOUND",
	3: "RESOURCE_LOWER_BOUND",
}

var ResourceBoundary_value = map[string]int32{
	"RESOURCE_BOUNDARY_UNDEFINED": 0,
	"RESOURCE_RAW":                1,
	"RESOURCE_UPPER_BOUND":        2,
	"RESOURCE_LOWER_BOUND":        3,
}

func (x ResourceBoundary) String() string {
	return proto.EnumName(ResourceBoundary_name, int32(x))
}

func (ResourceBoundary) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ef50cddc61837d8f, []int{0}
}

type ResourceQuota int32

const (
	ResourceQuota_RESOURCE_QUOTA_UNDEFINED ResourceQuota = 0
	ResourceQuota_RESOURCE_LIMIT           ResourceQuota = 1
	ResourceQuota_RESOURCE_REQUEST         ResourceQuota = 2
	ResourceQuota_RESOURCE_INITIAL_LIMIT   ResourceQuota = 3
	ResourceQuota_RESOURCE_INITIAL_REQUEST ResourceQuota = 4
)

var ResourceQuota_name = map[int32]string{
	0: "RESOURCE_QUOTA_UNDEFINED",
	1: "RESOURCE_LIMIT",
	2: "RESOURCE_REQUEST",
	3: "RESOURCE_INITIAL_LIMIT",
	4: "RESOURCE_INITIAL_REQUEST",
}

var ResourceQuota_value = map[string]int32{
	"RESOURCE_QUOTA_UNDEFINED": 0,
	"RESOURCE_LIMIT":           1,
	"RESOURCE_REQUEST":         2,
	"RESOURCE_INITIAL_LIMIT":   3,
	"RESOURCE_INITIAL_REQUEST": 4,
}

func (x ResourceQuota) String() string {
	return proto.EnumName(ResourceQuota_name, int32(x))
}

func (ResourceQuota) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ef50cddc61837d8f, []int{1}
}

type DataType int32

const (
	DataType_DATATYPE_UNDEFINED DataType = 0
	DataType_DATATYPE_BOOL      DataType = 1
	DataType_DATATYPE_INT       DataType = 2
	DataType_DATATYPE_INT8      DataType = 3
	DataType_DATATYPE_INT16     DataType = 4
	DataType_DATATYPE_INT32     DataType = 5
	DataType_DATATYPE_INT64     DataType = 6
	DataType_DATATYPE_UINT      DataType = 7
	DataType_DATATYPE_UINT8     DataType = 8
	DataType_DATATYPE_UINT16    DataType = 9
	DataType_DATATYPE_UINT32    DataType = 10
	DataType_DATATYPE_UTIN64    DataType = 11
	DataType_DATATYPE_FLOAT32   DataType = 12
	DataType_DATATYPE_FLOAT64   DataType = 13
	DataType_DATATYPE_STRING    DataType = 14
)

var DataType_name = map[int32]string{
	0:  "DATATYPE_UNDEFINED",
	1:  "DATATYPE_BOOL",
	2:  "DATATYPE_INT",
	3:  "DATATYPE_INT8",
	4:  "DATATYPE_INT16",
	5:  "DATATYPE_INT32",
	6:  "DATATYPE_INT64",
	7:  "DATATYPE_UINT",
	8:  "DATATYPE_UINT8",
	9:  "DATATYPE_UINT16",
	10: "DATATYPE_UINT32",
	11: "DATATYPE_UTIN64",
	12: "DATATYPE_FLOAT32",
	13: "DATATYPE_FLOAT64",
	14: "DATATYPE_STRING",
}

var DataType_value = map[string]int32{
	"DATATYPE_UNDEFINED": 0,
	"DATATYPE_BOOL":      1,
	"DATATYPE_INT":       2,
	"DATATYPE_INT8":      3,
	"DATATYPE_INT16":     4,
	"DATATYPE_INT32":     5,
	"DATATYPE_INT64":     6,
	"DATATYPE_UINT":      7,
	"DATATYPE_UINT8":     8,
	"DATATYPE_UINT16":    9,
	"DATATYPE_UINT32":    10,
	"DATATYPE_UTIN64":    11,
	"DATATYPE_FLOAT32":   12,
	"DATATYPE_FLOAT64":   13,
	"DATATYPE_STRING":    14,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ef50cddc61837d8f, []int{2}
}

type ColumnType int32

const (
	ColumnType_COLUMNTYPE_UDEFINED ColumnType = 0
	ColumnType_COLUMNTYPE_TAG      ColumnType = 1
	ColumnType_COLUMNTYPE_FIELD    ColumnType = 2
)

var ColumnType_name = map[int32]string{
	0: "COLUMNTYPE_UDEFINED",
	1: "COLUMNTYPE_TAG",
	2: "COLUMNTYPE_FIELD",
}

var ColumnType_value = map[string]int32{
	"COLUMNTYPE_UDEFINED": 0,
	"COLUMNTYPE_TAG":      1,
	"COLUMNTYPE_FIELD":    2,
}

func (x ColumnType) String() string {
	return proto.EnumName(ColumnType_name, int32(x))
}

func (ColumnType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ef50cddc61837d8f, []int{3}
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.ResourceBoundary", ResourceBoundary_name, ResourceBoundary_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.ResourceQuota", ResourceQuota_name, ResourceQuota_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.DataType", DataType_name, DataType_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.ColumnType", ColumnType_name, ColumnType_value)
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/common/types.proto", fileDescriptor_ef50cddc61837d8f)
}

var fileDescriptor_ef50cddc61837d8f = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x0b, 0xa4, 0x69, 0x3a, 0x0d, 0xe9, 0x74, 0x13, 0xa5, 0x51, 0x5b, 0xa9, 0x77, 0xd4,
	0xda, 0x22, 0x20, 0xab, 0xc7, 0xda, 0xd8, 0x44, 0x2b, 0x39, 0xbb, 0x60, 0x76, 0x15, 0xa5, 0x17,
	0xb4, 0x01, 0xab, 0xb1, 0x04, 0x5e, 0x0b, 0xec, 0x4a, 0x3c, 0x46, 0x1f, 0xa8, 0xef, 0x56, 0x99,
	0x3f, 0x86, 0x4d, 0x2e, 0xb9, 0xfe, 0xe6, 0xfb, 0xe6, 0x1b, 0x8f, 0x77, 0xc0, 0x56, 0x33, 0x35,
	0x8f, 0xa7, 0x6a, 0xac, 0xb2, 0xc4, 0xfe, 0xd3, 0x56, 0xb3, 0xec, 0x51, 0xb5, 0xed, 0xa9, 0xca,
	0xd5, 0x63, 0xf1, 0x60, 0x4f, 0xf4, 0x7c, 0xae, 0x53, 0x3b, 0x5f, 0x65, 0xf1, 0xd2, 0xca, 0x16,
	0x3a, 0xd7, 0xe4, 0xdb, 0x44, 0xa7, 0xb9, 0x4a, 0xd2, 0x78, 0xb1, 0x54, 0x89, 0xb5, 0x75, 0x5b,
	0x3b, 0xa7, 0xb5, 0x75, 0x5a, 0x1b, 0x67, 0x6b, 0x05, 0x18, 0xc5, 0x4b, 0x5d, 0x2c, 0x26, 0xb1,
	0xa7, 0x8b, 0x74, 0xaa, 0x16, 0x2b, 0xf2, 0x15, 0x3e, 0x47, 0xc1, 0x88, 0xcb, 0xa8, 0x17, 0x8c,
	0x3d, 0x2e, 0x99, 0xef, 0x46, 0xf7, 0x63, 0xc9, 0xfc, 0xa0, 0x4f, 0x59, 0xe0, 0xe3, 0x2b, 0x82,
	0x70, 0x5a, 0x09, 0x22, 0xf7, 0x0e, 0x6b, 0xe4, 0x0a, 0x2e, 0x2a, 0x22, 0x07, 0x83, 0x20, 0xda,
	0x18, 0xb1, 0x6e, 0x54, 0x42, 0x7e, 0x57, 0x55, 0x1a, 0xad, 0xbf, 0x35, 0x68, 0xee, 0xb2, 0x87,
	0x85, 0xce, 0x15, 0xf9, 0x02, 0x57, 0x95, 0x76, 0x28, 0xb9, 0x70, 0x8d, 0x54, 0x02, 0x67, 0xfb,
	0x4e, 0xf4, 0x96, 0x0a, 0xac, 0x91, 0x0b, 0xc0, 0xfd, 0x24, 0xc1, 0x50, 0x06, 0x23, 0x81, 0x75,
	0xf2, 0x09, 0x2e, 0x2b, 0x4a, 0x19, 0x15, 0xd4, 0x0d, 0xb7, 0x8e, 0x86, 0x91, 0xb1, 0xab, 0xed,
	0x9c, 0x47, 0xad, 0x7f, 0x75, 0x38, 0xf1, 0x55, 0xae, 0xc4, 0x2a, 0x8b, 0xc9, 0x25, 0x10, 0xdf,
	0x15, 0xae, 0xb8, 0x1f, 0x04, 0xc6, 0x20, 0x1f, 0xa0, 0x59, 0x71, 0x8f, 0xf3, 0x10, 0x6b, 0xe5,
	0x46, 0x2a, 0x44, 0x59, 0x39, 0xc3, 0xa1, 0x88, 0x32, 0xf1, 0x03, 0x1b, 0xe5, 0x07, 0x1c, 0xa2,
	0xb6, 0x83, 0x47, 0x4f, 0x59, 0xe7, 0x1a, 0x5f, 0x3f, 0x65, 0x4e, 0x17, 0x8f, 0x8d, 0x76, 0xb2,
	0x4c, 0x78, 0x63, 0xc8, 0xe4, 0x3a, 0xe2, 0x84, 0x9c, 0xc3, 0x7b, 0x83, 0xb5, 0x1d, 0x7c, 0xfb,
	0x0c, 0x76, 0xae, 0x11, 0x4c, 0x28, 0x28, 0x73, 0xba, 0xf8, 0xae, 0x5c, 0x67, 0x05, 0xfb, 0x21,
	0x77, 0x4b, 0xe9, 0xe9, 0x73, 0xea, 0x74, 0xb1, 0x69, 0x34, 0x18, 0x89, 0x88, 0xb2, 0x1b, 0x3c,
	0x6b, 0x71, 0x80, 0x9e, 0x9e, 0x15, 0xf3, 0x74, 0xbd, 0xc0, 0x8f, 0x70, 0xde, 0xe3, 0xa1, 0xbc,
	0x65, 0x9b, 0x14, 0xe3, 0x57, 0x1e, 0x14, 0x84, 0x7b, 0xb3, 0xf9, 0x95, 0x07, 0xac, 0x4f, 0x83,
	0xd0, 0xc7, 0xba, 0xe7, 0xfd, 0xfa, 0xf9, 0x3b, 0xc9, 0xb7, 0x0f, 0xd6, 0xde, 0x3f, 0xed, 0xef,
	0x2a, 0xb1, 0xcb, 0x8b, 0x78, 0xc1, 0x75, 0x3c, 0x1c, 0xaf, 0x0f, 0xa3, 0xf3, 0x3f, 0x00, 0x00,
	0xff, 0xff, 0xac, 0x63, 0x99, 0x0f, 0x4b, 0x03, 0x00, 0x00,
}
