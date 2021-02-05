// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metric.proto

package events

import (
	fmt "fmt"
	metrics "github.com/batchcorp/schemas/build/go/events/metrics"
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

type Metric_Type int32

const (
	Metric_UNSET   Metric_Type = 0
	Metric_COUNTER Metric_Type = 1
	Metric_ERROR   Metric_Type = 2
)

var Metric_Type_name = map[int32]string{
	0: "UNSET",
	1: "COUNTER",
	2: "ERROR",
}

var Metric_Type_value = map[string]int32{
	"UNSET":   0,
	"COUNTER": 1,
	"ERROR":   2,
}

func (x Metric_Type) String() string {
	return proto.EnumName(Metric_Type_name, int32(x))
}

func (Metric_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_da41641f55bff5df, []int{0, 0}
}

// This message is emitted by mlib and consumed by the metrics service.
// We use this envelope type so that the metrics service can use a single
// consume func to handle multiple message types.
type Metric struct {
	Type Metric_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.Metric_Type" json:"type,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*Metric_Counter
	//	*Metric_Error
	Event                isMetric_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_da41641f55bff5df, []int{0}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetType() Metric_Type {
	if m != nil {
		return m.Type
	}
	return Metric_UNSET
}

type isMetric_Event interface {
	isMetric_Event()
}

type Metric_Counter struct {
	Counter *metrics.Counter `protobuf:"bytes,100,opt,name=counter,proto3,oneof"`
}

type Metric_Error struct {
	Error *metrics.Error `protobuf:"bytes,101,opt,name=error,proto3,oneof"`
}

func (*Metric_Counter) isMetric_Event() {}

func (*Metric_Error) isMetric_Event() {}

func (m *Metric) GetEvent() isMetric_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Metric) GetCounter() *metrics.Counter {
	if x, ok := m.GetEvent().(*Metric_Counter); ok {
		return x.Counter
	}
	return nil
}

func (m *Metric) GetError() *metrics.Error {
	if x, ok := m.GetEvent().(*Metric_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Metric) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Metric_Counter)(nil),
		(*Metric_Error)(nil),
	}
}

func init() {
	proto.RegisterEnum("events.Metric_Type", Metric_Type_name, Metric_Type_value)
	proto.RegisterType((*Metric)(nil), "events.Metric")
}

func init() { proto.RegisterFile("metric.proto", fileDescriptor_da41641f55bff5df) }

var fileDescriptor_da41641f55bff5df = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x6a, 0xc2, 0x40,
	0x10, 0x86, 0xb3, 0xc5, 0x24, 0x74, 0x2c, 0x12, 0x56, 0x0a, 0xc1, 0x93, 0x78, 0x68, 0x2d, 0xc8,
	0x2c, 0xd8, 0x37, 0x50, 0x02, 0x5e, 0xaa, 0xb0, 0x8d, 0x97, 0xde, 0xcc, 0x3a, 0x98, 0x40, 0xe3,
	0x86, 0xcd, 0xa6, 0xe0, 0xb3, 0xf5, 0xe5, 0x4a, 0x76, 0x63, 0xaf, 0xdf, 0xff, 0xf1, 0xcf, 0xfc,
	0xf0, 0x54, 0x93, 0x35, 0x95, 0xc2, 0xc6, 0x68, 0xab, 0x79, 0x44, 0x3f, 0x74, 0xb5, 0xed, 0xec,
	0xd9, 0xd3, 0x56, 0x28, 0xdd, 0x5d, 0x2d, 0x19, 0x1f, 0xcf, 0xa6, 0x77, 0x4c, 0xc6, 0xe8, 0x01,
	0x2e, 0x7e, 0x19, 0x44, 0x1f, 0x8e, 0xf3, 0x57, 0x18, 0xd9, 0x5b, 0x43, 0x29, 0x9b, 0xb3, 0xe5,
	0x64, 0x3d, 0x45, 0xdf, 0x86, 0x3e, 0xc5, 0xfc, 0xd6, 0x90, 0x74, 0x02, 0x5f, 0x41, 0x3c, 0x34,
	0xa7, 0xe7, 0x39, 0x5b, 0x8e, 0xd7, 0x09, 0x0e, 0xd5, 0xb8, 0xf5, 0x7c, 0x17, 0xc8, 0xbb, 0xc2,
	0x5f, 0x20, 0x74, 0x07, 0x53, 0x72, 0xee, 0xe4, 0xdf, 0xcd, 0x7a, 0xba, 0x0b, 0xa4, 0x8f, 0x17,
	0x6f, 0x30, 0xea, 0x6f, 0xf0, 0x47, 0x08, 0x8f, 0xfb, 0xcf, 0x2c, 0x4f, 0x02, 0x3e, 0x86, 0x78,
	0x7b, 0x38, 0xee, 0xf3, 0x4c, 0x26, 0xac, 0xe7, 0x99, 0x94, 0x07, 0x99, 0x3c, 0x6c, 0x62, 0x08,
	0xdd, 0x73, 0x1b, 0xfc, 0x5a, 0x5d, 0x2a, 0x5b, 0x76, 0x05, 0x2a, 0x5d, 0x8b, 0xe2, 0x64, 0x55,
	0xa9, 0xb4, 0x69, 0x44, 0xab, 0x4a, 0xaa, 0x4f, 0xad, 0x28, 0xba, 0xea, 0xfb, 0x2c, 0x2e, 0x5a,
	0xf8, 0x2d, 0x45, 0xe4, 0x46, 0xbf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x21, 0xfb, 0x63, 0xea,
	0x38, 0x01, 0x00, 0x00,
}