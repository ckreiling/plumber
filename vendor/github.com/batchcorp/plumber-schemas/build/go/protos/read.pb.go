// Code generated by protoc-gen-go. DO NOT EDIT.
// source: read.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
	opts "github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	records "github.com/batchcorp/plumber-schemas/build/go/protos/records"
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

type CreateReadRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth      `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	Read                 *opts.ReadOptions `protobuf:"bytes,1,opt,name=read,proto3" json:"read,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateReadRequest) Reset()         { *m = CreateReadRequest{} }
func (m *CreateReadRequest) String() string { return proto.CompactTextString(m) }
func (*CreateReadRequest) ProtoMessage()    {}
func (*CreateReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{0}
}

func (m *CreateReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReadRequest.Unmarshal(m, b)
}
func (m *CreateReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReadRequest.Marshal(b, m, deterministic)
}
func (m *CreateReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReadRequest.Merge(m, src)
}
func (m *CreateReadRequest) XXX_Size() int {
	return xxx_messageInfo_CreateReadRequest.Size(m)
}
func (m *CreateReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReadRequest proto.InternalMessageInfo

func (m *CreateReadRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *CreateReadRequest) GetRead() *opts.ReadOptions {
	if m != nil {
		return m.Read
	}
	return nil
}

type CreateReadResponse struct {
	Status *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	// Assigned and returned by plumber-server to identify a successful read request
	ReadId               string   `protobuf:"bytes,1,opt,name=read_id,json=readId,proto3" json:"read_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReadResponse) Reset()         { *m = CreateReadResponse{} }
func (m *CreateReadResponse) String() string { return proto.CompactTextString(m) }
func (*CreateReadResponse) ProtoMessage()    {}
func (*CreateReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{1}
}

func (m *CreateReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReadResponse.Unmarshal(m, b)
}
func (m *CreateReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReadResponse.Marshal(b, m, deterministic)
}
func (m *CreateReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReadResponse.Merge(m, src)
}
func (m *CreateReadResponse) XXX_Size() int {
	return xxx_messageInfo_CreateReadResponse.Size(m)
}
func (m *CreateReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReadResponse proto.InternalMessageInfo

func (m *CreateReadResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateReadResponse) GetReadId() string {
	if m != nil {
		return m.ReadId
	}
	return ""
}

type StopReadRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ReadId               string       `protobuf:"bytes,1,opt,name=read_id,json=readId,proto3" json:"read_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StopReadRequest) Reset()         { *m = StopReadRequest{} }
func (m *StopReadRequest) String() string { return proto.CompactTextString(m) }
func (*StopReadRequest) ProtoMessage()    {}
func (*StopReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{2}
}

func (m *StopReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopReadRequest.Unmarshal(m, b)
}
func (m *StopReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopReadRequest.Marshal(b, m, deterministic)
}
func (m *StopReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopReadRequest.Merge(m, src)
}
func (m *StopReadRequest) XXX_Size() int {
	return xxx_messageInfo_StopReadRequest.Size(m)
}
func (m *StopReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopReadRequest proto.InternalMessageInfo

func (m *StopReadRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *StopReadRequest) GetReadId() string {
	if m != nil {
		return m.ReadId
	}
	return ""
}

type StopReadResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StopReadResponse) Reset()         { *m = StopReadResponse{} }
func (m *StopReadResponse) String() string { return proto.CompactTextString(m) }
func (*StopReadResponse) ProtoMessage()    {}
func (*StopReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{3}
}

func (m *StopReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopReadResponse.Unmarshal(m, b)
}
func (m *StopReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopReadResponse.Marshal(b, m, deterministic)
}
func (m *StopReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopReadResponse.Merge(m, src)
}
func (m *StopReadResponse) XXX_Size() int {
	return xxx_messageInfo_StopReadResponse.Size(m)
}
func (m *StopReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopReadResponse proto.InternalMessageInfo

func (m *StopReadResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ResumeReadRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ReadId               string       `protobuf:"bytes,1,opt,name=read_id,json=readId,proto3" json:"read_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResumeReadRequest) Reset()         { *m = ResumeReadRequest{} }
func (m *ResumeReadRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeReadRequest) ProtoMessage()    {}
func (*ResumeReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{4}
}

func (m *ResumeReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeReadRequest.Unmarshal(m, b)
}
func (m *ResumeReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeReadRequest.Marshal(b, m, deterministic)
}
func (m *ResumeReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeReadRequest.Merge(m, src)
}
func (m *ResumeReadRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeReadRequest.Size(m)
}
func (m *ResumeReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeReadRequest proto.InternalMessageInfo

func (m *ResumeReadRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *ResumeReadRequest) GetReadId() string {
	if m != nil {
		return m.ReadId
	}
	return ""
}

type ResumeReadResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResumeReadResponse) Reset()         { *m = ResumeReadResponse{} }
func (m *ResumeReadResponse) String() string { return proto.CompactTextString(m) }
func (*ResumeReadResponse) ProtoMessage()    {}
func (*ResumeReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{5}
}

func (m *ResumeReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeReadResponse.Unmarshal(m, b)
}
func (m *ResumeReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeReadResponse.Marshal(b, m, deterministic)
}
func (m *ResumeReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeReadResponse.Merge(m, src)
}
func (m *ResumeReadResponse) XXX_Size() int {
	return xxx_messageInfo_ResumeReadResponse.Size(m)
}
func (m *ResumeReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeReadResponse proto.InternalMessageInfo

func (m *ResumeReadResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type DeleteReadRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ReadId               string       `protobuf:"bytes,1,opt,name=read_id,json=readId,proto3" json:"read_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteReadRequest) Reset()         { *m = DeleteReadRequest{} }
func (m *DeleteReadRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteReadRequest) ProtoMessage()    {}
func (*DeleteReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{6}
}

func (m *DeleteReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReadRequest.Unmarshal(m, b)
}
func (m *DeleteReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReadRequest.Marshal(b, m, deterministic)
}
func (m *DeleteReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReadRequest.Merge(m, src)
}
func (m *DeleteReadRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteReadRequest.Size(m)
}
func (m *DeleteReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReadRequest proto.InternalMessageInfo

func (m *DeleteReadRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *DeleteReadRequest) GetReadId() string {
	if m != nil {
		return m.ReadId
	}
	return ""
}

type DeleteReadResponse struct {
	Status               *common.Status `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteReadResponse) Reset()         { *m = DeleteReadResponse{} }
func (m *DeleteReadResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteReadResponse) ProtoMessage()    {}
func (*DeleteReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{7}
}

func (m *DeleteReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReadResponse.Unmarshal(m, b)
}
func (m *DeleteReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReadResponse.Marshal(b, m, deterministic)
}
func (m *DeleteReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReadResponse.Merge(m, src)
}
func (m *DeleteReadResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteReadResponse.Size(m)
}
func (m *DeleteReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReadResponse proto.InternalMessageInfo

func (m *DeleteReadResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type StartReadRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	ReadId               string       `protobuf:"bytes,1,opt,name=read_id,json=readId,proto3" json:"read_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StartReadRequest) Reset()         { *m = StartReadRequest{} }
func (m *StartReadRequest) String() string { return proto.CompactTextString(m) }
func (*StartReadRequest) ProtoMessage()    {}
func (*StartReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{8}
}

func (m *StartReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartReadRequest.Unmarshal(m, b)
}
func (m *StartReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartReadRequest.Marshal(b, m, deterministic)
}
func (m *StartReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartReadRequest.Merge(m, src)
}
func (m *StartReadRequest) XXX_Size() int {
	return xxx_messageInfo_StartReadRequest.Size(m)
}
func (m *StartReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartReadRequest proto.InternalMessageInfo

func (m *StartReadRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *StartReadRequest) GetReadId() string {
	if m != nil {
		return m.ReadId
	}
	return ""
}

type StartReadResponse struct {
	Records              []*records.ReadRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	Status               *common.Status        `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StartReadResponse) Reset()         { *m = StartReadResponse{} }
func (m *StartReadResponse) String() string { return proto.CompactTextString(m) }
func (*StartReadResponse) ProtoMessage()    {}
func (*StartReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{9}
}

func (m *StartReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartReadResponse.Unmarshal(m, b)
}
func (m *StartReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartReadResponse.Marshal(b, m, deterministic)
}
func (m *StartReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartReadResponse.Merge(m, src)
}
func (m *StartReadResponse) XXX_Size() int {
	return xxx_messageInfo_StartReadResponse.Size(m)
}
func (m *StartReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartReadResponse proto.InternalMessageInfo

func (m *StartReadResponse) GetRecords() []*records.ReadRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *StartReadResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetAllReadsRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAllReadsRequest) Reset()         { *m = GetAllReadsRequest{} }
func (m *GetAllReadsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllReadsRequest) ProtoMessage()    {}
func (*GetAllReadsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{10}
}

func (m *GetAllReadsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllReadsRequest.Unmarshal(m, b)
}
func (m *GetAllReadsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllReadsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllReadsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllReadsRequest.Merge(m, src)
}
func (m *GetAllReadsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllReadsRequest.Size(m)
}
func (m *GetAllReadsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllReadsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllReadsRequest proto.InternalMessageInfo

func (m *GetAllReadsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type GetAllReadsResponse struct {
	Read                 []*opts.ReadOptions `protobuf:"bytes,1,rep,name=read,proto3" json:"read,omitempty"`
	Status               *common.Status      `protobuf:"bytes,1000,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetAllReadsResponse) Reset()         { *m = GetAllReadsResponse{} }
func (m *GetAllReadsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllReadsResponse) ProtoMessage()    {}
func (*GetAllReadsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b10ec61df6818dd, []int{11}
}

func (m *GetAllReadsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllReadsResponse.Unmarshal(m, b)
}
func (m *GetAllReadsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllReadsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllReadsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllReadsResponse.Merge(m, src)
}
func (m *GetAllReadsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllReadsResponse.Size(m)
}
func (m *GetAllReadsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllReadsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllReadsResponse proto.InternalMessageInfo

func (m *GetAllReadsResponse) GetRead() []*opts.ReadOptions {
	if m != nil {
		return m.Read
	}
	return nil
}

func (m *GetAllReadsResponse) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateReadRequest)(nil), "protos.CreateReadRequest")
	proto.RegisterType((*CreateReadResponse)(nil), "protos.CreateReadResponse")
	proto.RegisterType((*StopReadRequest)(nil), "protos.StopReadRequest")
	proto.RegisterType((*StopReadResponse)(nil), "protos.StopReadResponse")
	proto.RegisterType((*ResumeReadRequest)(nil), "protos.ResumeReadRequest")
	proto.RegisterType((*ResumeReadResponse)(nil), "protos.ResumeReadResponse")
	proto.RegisterType((*DeleteReadRequest)(nil), "protos.DeleteReadRequest")
	proto.RegisterType((*DeleteReadResponse)(nil), "protos.DeleteReadResponse")
	proto.RegisterType((*StartReadRequest)(nil), "protos.StartReadRequest")
	proto.RegisterType((*StartReadResponse)(nil), "protos.StartReadResponse")
	proto.RegisterType((*GetAllReadsRequest)(nil), "protos.GetAllReadsRequest")
	proto.RegisterType((*GetAllReadsResponse)(nil), "protos.GetAllReadsResponse")
}

func init() { proto.RegisterFile("read.proto", fileDescriptor_7b10ec61df6818dd) }

var fileDescriptor_7b10ec61df6818dd = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x86, 0x09, 0x09, 0x09, 0xab, 0x1c, 0xb2, 0x56, 0x58, 0xd6, 0xe4, 0x14, 0x7c, 0xca, 0x61,
	0xd7, 0x82, 0x6c, 0xd8, 0xe3, 0x42, 0xb2, 0x81, 0xd2, 0x4b, 0x0b, 0x4a, 0xdb, 0x43, 0xa1, 0x14,
	0xd9, 0x1e, 0x62, 0x53, 0x3b, 0x72, 0x35, 0xd2, 0xa1, 0x4f, 0xd1, 0xd7, 0xec, 0x63, 0x14, 0x59,
	0x0e, 0x71, 0x0b, 0xa5, 0xc4, 0xa4, 0x27, 0xdb, 0xa3, 0xdf, 0xdf, 0x7c, 0x63, 0x18, 0x13, 0xa2,
	0x40, 0x24, 0x61, 0xa9, 0xa4, 0x96, 0xb4, 0x5f, 0x5d, 0x70, 0xe2, 0xc5, 0xb2, 0x28, 0xe4, 0x8e,
	0x09, 0xa3, 0x53, 0x77, 0x34, 0x19, 0xd7, 0x25, 0xd4, 0x42, 0x1b, 0xac, 0x8b, 0x54, 0x41, 0x2c,
	0x55, 0x82, 0x2c, 0x12, 0x08, 0x75, 0x6d, 0x24, 0x4b, 0x8d, 0xec, 0x00, 0x0d, 0x1e, 0x88, 0xf7,
	0x5f, 0x81, 0xd0, 0xc0, 0x41, 0x24, 0x1c, 0x1e, 0x0d, 0xa0, 0xa6, 0x33, 0xd2, 0xb3, 0x70, 0xff,
	0xf9, 0x62, 0xda, 0x99, 0x0d, 0xe7, 0x63, 0x97, 0xc5, 0xd0, 0x75, 0x09, 0x97, 0x46, 0xa7, 0xbc,
	0x4a, 0xd0, 0x5f, 0xa4, 0x67, 0x61, 0x7e, 0xa7, 0x0a, 0xfa, 0xfb, 0xa0, 0xed, 0x12, 0x5a, 0xe2,
	0x65, 0xa9, 0x33, 0xb9, 0x43, 0x5e, 0xa5, 0x82, 0x3b, 0x42, 0x9b, 0xcd, 0xb0, 0x94, 0x3b, 0x04,
	0x1a, 0x92, 0xbe, 0xf3, 0xf6, 0x5f, 0x06, 0x15, 0xe6, 0xc7, 0xbb, 0x7e, 0x9b, 0xea, 0x94, 0xd7,
	0x29, 0xfa, 0x93, 0x0c, 0x2c, 0xed, 0x3e, 0x73, 0x6d, 0xbf, 0xf1, 0xbe, 0x7d, 0x3c, 0x4f, 0x82,
	0x2b, 0x32, 0xda, 0x68, 0x59, 0xb6, 0x9b, 0xe4, 0x43, 0xea, 0x8a, 0x7c, 0x3f, 0x50, 0xdb, 0x29,
	0x07, 0x37, 0xc4, 0xe3, 0x80, 0xa6, 0x80, 0x13, 0xbb, 0xad, 0x09, 0x6d, 0x72, 0xdb, 0xdb, 0xad,
	0x21, 0x07, 0xfd, 0x05, 0x76, 0x4d, 0x6e, 0x4b, 0xbb, 0x6b, 0xfb, 0xfd, 0x85, 0xd2, 0x27, 0x96,
	0x7b, 0x22, 0x5e, 0x03, 0x5b, 0xbb, 0x2d, 0x6c, 0xba, 0x5a, 0x1a, 0xbf, 0x33, 0xed, 0xce, 0x86,
	0xf3, 0xc9, 0x9e, 0x5c, 0x97, 0x43, 0x17, 0xb7, 0xf7, 0x7c, 0x1f, 0x3d, 0x7a, 0xa2, 0x7f, 0x84,
	0x9e, 0x81, 0x5e, 0xe6, 0xb9, 0x85, 0xe1, 0xd1, 0x33, 0x05, 0x48, 0xc6, 0x6f, 0xde, 0xaf, 0xe5,
	0x0f, 0xbb, 0xd8, 0xfd, 0x7c, 0x17, 0x8f, 0x95, 0x5e, 0xfd, 0xbd, 0x5d, 0x6c, 0x33, 0x9d, 0x9a,
	0xc8, 0x9e, 0xb3, 0x48, 0xe8, 0x38, 0x8d, 0xa5, 0x2a, 0x59, 0x99, 0x9b, 0x22, 0x02, 0xf5, 0x1b,
	0xe3, 0x14, 0x0a, 0x81, 0x2c, 0x32, 0x59, 0x9e, 0xb0, 0xad, 0x64, 0x8e, 0x16, 0xb9, 0xbf, 0xd6,
	0x9f, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x6e, 0xd4, 0x7b, 0xca, 0x04, 0x00, 0x00,
}
