// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/common/queries.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TimeRange_AggregateFunction int32

const (
	TimeRange_NONE TimeRange_AggregateFunction = 0
	TimeRange_MAX  TimeRange_AggregateFunction = 1
)

var TimeRange_AggregateFunction_name = map[int32]string{
	0: "NONE",
	1: "MAX",
}

var TimeRange_AggregateFunction_value = map[string]int32{
	"NONE": 0,
	"MAX":  1,
}

func (x TimeRange_AggregateFunction) String() string {
	return proto.EnumName(TimeRange_AggregateFunction_name, int32(x))
}

func (TimeRange_AggregateFunction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{0, 0}
}

type QueryCondition_Order int32

const (
	QueryCondition_NONE QueryCondition_Order = 0
	QueryCondition_ASC  QueryCondition_Order = 1
	QueryCondition_DESC QueryCondition_Order = 2
)

var QueryCondition_Order_name = map[int32]string{
	0: "NONE",
	1: "ASC",
	2: "DESC",
}

var QueryCondition_Order_value = map[string]int32{
	"NONE": 0,
	"ASC":  1,
	"DESC": 2,
}

func (x QueryCondition_Order) String() string {
	return proto.EnumName(QueryCondition_Order_name, int32(x))
}

func (QueryCondition_Order) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{1, 0}
}

//*
// Represents a time range definition
//
type TimeRange struct {
	StartTime            *timestamp.Timestamp        `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp        `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Step                 *duration.Duration          `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	AggregateFunction    TimeRange_AggregateFunction `protobuf:"varint,4,opt,name=aggregateFunction,proto3,enum=containersai.alameda.v1alpha1.datahub.common.TimeRange_AggregateFunction" json:"aggregateFunction,omitempty"`
	ApplyTime            *timestamp.Timestamp        `protobuf:"bytes,5,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{0}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TimeRange) GetStep() *duration.Duration {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *TimeRange) GetAggregateFunction() TimeRange_AggregateFunction {
	if m != nil {
		return m.AggregateFunction
	}
	return TimeRange_NONE
}

func (m *TimeRange) GetApplyTime() *timestamp.Timestamp {
	if m != nil {
		return m.ApplyTime
	}
	return nil
}

type QueryCondition struct {
	TimeRange            *TimeRange           `protobuf:"bytes,1,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	Order                QueryCondition_Order `protobuf:"varint,2,opt,name=order,proto3,enum=containersai.alameda.v1alpha1.datahub.common.QueryCondition_Order" json:"order,omitempty"`
	Limit                uint64               `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryCondition) Reset()         { *m = QueryCondition{} }
func (m *QueryCondition) String() string { return proto.CompactTextString(m) }
func (*QueryCondition) ProtoMessage()    {}
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_d602763cab07305c, []int{1}
}

func (m *QueryCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCondition.Unmarshal(m, b)
}
func (m *QueryCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCondition.Marshal(b, m, deterministic)
}
func (m *QueryCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCondition.Merge(m, src)
}
func (m *QueryCondition) XXX_Size() int {
	return xxx_messageInfo_QueryCondition.Size(m)
}
func (m *QueryCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCondition.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCondition proto.InternalMessageInfo

func (m *QueryCondition) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *QueryCondition) GetOrder() QueryCondition_Order {
	if m != nil {
		return m.Order
	}
	return QueryCondition_NONE
}

func (m *QueryCondition) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.TimeRange_AggregateFunction", TimeRange_AggregateFunction_name, TimeRange_AggregateFunction_value)
	proto.RegisterEnum("containersai.alameda.v1alpha1.datahub.common.QueryCondition_Order", QueryCondition_Order_name, QueryCondition_Order_value)
	proto.RegisterType((*TimeRange)(nil), "containersai.alameda.v1alpha1.datahub.common.TimeRange")
	proto.RegisterType((*QueryCondition)(nil), "containersai.alameda.v1alpha1.datahub.common.QueryCondition")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/common/queries.proto", fileDescriptor_d602763cab07305c)
}

var fileDescriptor_d602763cab07305c = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0xcd, 0x36, 0x71, 0xb7, 0xaf, 0x50, 0xb2, 0xc1, 0x43, 0xed, 0x41, 0x97, 0x20, 0xb2,
	0x07, 0x77, 0x86, 0x56, 0x44, 0xbc, 0xd9, 0x76, 0x57, 0xf0, 0xe0, 0x2e, 0x66, 0x8b, 0x14, 0x2f,
	0xe5, 0x6d, 0x33, 0xa6, 0x03, 0xc9, 0x4c, 0x9c, 0x4c, 0x94, 0x7e, 0x03, 0x3f, 0xb5, 0xc8, 0xcc,
	0xa4, 0x95, 0xb6, 0x87, 0xda, 0x3d, 0xbe, 0xc9, 0xf3, 0x7b, 0xde, 0x3f, 0x0f, 0x03, 0x7d, 0xcc,
	0xb1, 0x60, 0x29, 0xce, 0xb0, 0xe4, 0xf4, 0x67, 0x1f, 0xf3, 0x72, 0x89, 0x7d, 0x9a, 0xa2, 0xc6,
	0x65, 0x3d, 0xa7, 0x0b, 0x59, 0x14, 0x52, 0xd0, 0x1f, 0x35, 0x53, 0x9c, 0x55, 0xa4, 0x54, 0x52,
	0xcb, 0xe8, 0xf5, 0x42, 0x0a, 0x8d, 0x5c, 0x30, 0x55, 0x21, 0x27, 0x0d, 0x4f, 0xd6, 0x2c, 0x69,
	0x58, 0xe2, 0xd8, 0xde, 0xf3, 0x4c, 0xca, 0x2c, 0x67, 0xd4, 0xb2, 0xf3, 0xfa, 0x3b, 0x4d, 0x6b,
	0x85, 0x9a, 0x4b, 0xe1, 0xdc, 0x7a, 0x2f, 0x76, 0xff, 0x6b, 0x5e, 0xb0, 0x4a, 0x63, 0x51, 0x3a,
	0x41, 0xfc, 0xbb, 0x05, 0xed, 0x09, 0x2f, 0x58, 0x82, 0x22, 0x63, 0xd1, 0x7b, 0x80, 0x4a, 0xa3,
	0xd2, 0x33, 0x23, 0xeb, 0x7a, 0x17, 0xde, 0xe5, 0x93, 0x41, 0x8f, 0x38, 0x0f, 0xb2, 0xf6, 0x20,
	0x93, 0xb5, 0x47, 0xd2, 0xb6, 0x6a, 0x53, 0x47, 0x6f, 0xe1, 0x8c, 0x89, 0xd4, 0x81, 0x27, 0x07,
	0xc1, 0x53, 0x26, 0x52, 0x8b, 0x5d, 0x81, 0x5f, 0x69, 0x56, 0x76, 0x5b, 0x16, 0x79, 0xb6, 0x87,
	0x5c, 0x37, 0xfb, 0x24, 0x56, 0x16, 0xfd, 0x82, 0x73, 0xcc, 0x32, 0xc5, 0x32, 0xd4, 0xec, 0x63,
	0x2d, 0x16, 0xe6, 0x57, 0xd7, 0xbf, 0xf0, 0x2e, 0x3b, 0x83, 0x4f, 0xe4, 0x98, 0xcb, 0x91, 0xcd,
	0xd2, 0x64, 0xb8, 0x6b, 0x98, 0xec, 0xf7, 0x30, 0x97, 0xc1, 0xb2, 0xcc, 0x57, 0x6e, 0xc1, 0xe0,
	0xf0, 0x65, 0xac, 0xda, 0xd4, 0xf1, 0x2b, 0x38, 0xdf, 0x6b, 0x11, 0x9d, 0x81, 0x7f, 0x7b, 0x77,
	0x7b, 0x13, 0x3e, 0x8a, 0x4e, 0xa1, 0xf5, 0x79, 0x38, 0x0d, 0xbd, 0xf8, 0x8f, 0x07, 0x9d, 0x2f,
	0x35, 0x53, 0xab, 0xb1, 0x14, 0x29, 0xb7, 0xaa, 0xaf, 0x00, 0xa6, 0xdf, 0x4c, 0x99, 0x41, 0x9b,
	0x3c, 0xde, 0x3d, 0x70, 0xcf, 0xa4, 0xad, 0x37, 0x39, 0x4f, 0x21, 0x90, 0x2a, 0x65, 0xca, 0x26,
	0xd5, 0x19, 0x8c, 0x8e, 0xb3, 0xdc, 0x1e, 0x92, 0xdc, 0x19, 0xa7, 0xc4, 0x19, 0x46, 0x4f, 0x21,
	0xc8, 0x79, 0xc1, 0xb5, 0x0d, 0xd4, 0x4f, 0x5c, 0x11, 0xbf, 0x84, 0xc0, 0xaa, 0xb6, 0xd7, 0x1e,
	0xde, 0x8f, 0x43, 0xcf, 0x7c, 0xba, 0xbe, 0xb9, 0x1f, 0x87, 0x27, 0xa3, 0xd1, 0xb7, 0x0f, 0x19,
	0xd7, 0x4d, 0x23, 0xfa, 0x6f, 0xa4, 0x2b, 0xe4, 0xd4, 0x3c, 0xa0, 0xff, 0x78, 0x4c, 0xf3, 0xc7,
	0x36, 0x8b, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x56, 0xed, 0xc7, 0x7a, 0x03, 0x00,
	0x00,
}