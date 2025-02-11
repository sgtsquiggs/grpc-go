// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payloads.proto

package grpc_testing

import (
	fmt "fmt"
	proto "github.com/sgtsquiggs/protobuf/proto"
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

type ByteBufferParams struct {
	ReqSize              int32    `protobuf:"varint,1,opt,name=req_size,json=reqSize,proto3" json:"req_size,omitempty"`
	RespSize             int32    `protobuf:"varint,2,opt,name=resp_size,json=respSize,proto3" json:"resp_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByteBufferParams) Reset()         { *m = ByteBufferParams{} }
func (m *ByteBufferParams) String() string { return proto.CompactTextString(m) }
func (*ByteBufferParams) ProtoMessage()    {}
func (*ByteBufferParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a075f58f70088c8, []int{0}
}

func (m *ByteBufferParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByteBufferParams.Unmarshal(m, b)
}
func (m *ByteBufferParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByteBufferParams.Marshal(b, m, deterministic)
}
func (m *ByteBufferParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByteBufferParams.Merge(m, src)
}
func (m *ByteBufferParams) XXX_Size() int {
	return xxx_messageInfo_ByteBufferParams.Size(m)
}
func (m *ByteBufferParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ByteBufferParams.DiscardUnknown(m)
}

var xxx_messageInfo_ByteBufferParams proto.InternalMessageInfo

func (m *ByteBufferParams) GetReqSize() int32 {
	if m != nil {
		return m.ReqSize
	}
	return 0
}

func (m *ByteBufferParams) GetRespSize() int32 {
	if m != nil {
		return m.RespSize
	}
	return 0
}

type SimpleProtoParams struct {
	ReqSize              int32    `protobuf:"varint,1,opt,name=req_size,json=reqSize,proto3" json:"req_size,omitempty"`
	RespSize             int32    `protobuf:"varint,2,opt,name=resp_size,json=respSize,proto3" json:"resp_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleProtoParams) Reset()         { *m = SimpleProtoParams{} }
func (m *SimpleProtoParams) String() string { return proto.CompactTextString(m) }
func (*SimpleProtoParams) ProtoMessage()    {}
func (*SimpleProtoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a075f58f70088c8, []int{1}
}

func (m *SimpleProtoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleProtoParams.Unmarshal(m, b)
}
func (m *SimpleProtoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleProtoParams.Marshal(b, m, deterministic)
}
func (m *SimpleProtoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleProtoParams.Merge(m, src)
}
func (m *SimpleProtoParams) XXX_Size() int {
	return xxx_messageInfo_SimpleProtoParams.Size(m)
}
func (m *SimpleProtoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleProtoParams.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleProtoParams proto.InternalMessageInfo

func (m *SimpleProtoParams) GetReqSize() int32 {
	if m != nil {
		return m.ReqSize
	}
	return 0
}

func (m *SimpleProtoParams) GetRespSize() int32 {
	if m != nil {
		return m.RespSize
	}
	return 0
}

type ComplexProtoParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComplexProtoParams) Reset()         { *m = ComplexProtoParams{} }
func (m *ComplexProtoParams) String() string { return proto.CompactTextString(m) }
func (*ComplexProtoParams) ProtoMessage()    {}
func (*ComplexProtoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a075f58f70088c8, []int{2}
}

func (m *ComplexProtoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexProtoParams.Unmarshal(m, b)
}
func (m *ComplexProtoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexProtoParams.Marshal(b, m, deterministic)
}
func (m *ComplexProtoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexProtoParams.Merge(m, src)
}
func (m *ComplexProtoParams) XXX_Size() int {
	return xxx_messageInfo_ComplexProtoParams.Size(m)
}
func (m *ComplexProtoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexProtoParams.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexProtoParams proto.InternalMessageInfo

type PayloadConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*PayloadConfig_BytebufParams
	//	*PayloadConfig_SimpleParams
	//	*PayloadConfig_ComplexParams
	Payload              isPayloadConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PayloadConfig) Reset()         { *m = PayloadConfig{} }
func (m *PayloadConfig) String() string { return proto.CompactTextString(m) }
func (*PayloadConfig) ProtoMessage()    {}
func (*PayloadConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a075f58f70088c8, []int{3}
}

func (m *PayloadConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadConfig.Unmarshal(m, b)
}
func (m *PayloadConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadConfig.Marshal(b, m, deterministic)
}
func (m *PayloadConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadConfig.Merge(m, src)
}
func (m *PayloadConfig) XXX_Size() int {
	return xxx_messageInfo_PayloadConfig.Size(m)
}
func (m *PayloadConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadConfig proto.InternalMessageInfo

type isPayloadConfig_Payload interface {
	isPayloadConfig_Payload()
}

type PayloadConfig_BytebufParams struct {
	BytebufParams *ByteBufferParams `protobuf:"bytes,1,opt,name=bytebuf_params,json=bytebufParams,proto3,oneof"`
}

type PayloadConfig_SimpleParams struct {
	SimpleParams *SimpleProtoParams `protobuf:"bytes,2,opt,name=simple_params,json=simpleParams,proto3,oneof"`
}

type PayloadConfig_ComplexParams struct {
	ComplexParams *ComplexProtoParams `protobuf:"bytes,3,opt,name=complex_params,json=complexParams,proto3,oneof"`
}

func (*PayloadConfig_BytebufParams) isPayloadConfig_Payload() {}

func (*PayloadConfig_SimpleParams) isPayloadConfig_Payload() {}

func (*PayloadConfig_ComplexParams) isPayloadConfig_Payload() {}

func (m *PayloadConfig) GetPayload() isPayloadConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PayloadConfig) GetBytebufParams() *ByteBufferParams {
	if x, ok := m.GetPayload().(*PayloadConfig_BytebufParams); ok {
		return x.BytebufParams
	}
	return nil
}

func (m *PayloadConfig) GetSimpleParams() *SimpleProtoParams {
	if x, ok := m.GetPayload().(*PayloadConfig_SimpleParams); ok {
		return x.SimpleParams
	}
	return nil
}

func (m *PayloadConfig) GetComplexParams() *ComplexProtoParams {
	if x, ok := m.GetPayload().(*PayloadConfig_ComplexParams); ok {
		return x.ComplexParams
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PayloadConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PayloadConfig_BytebufParams)(nil),
		(*PayloadConfig_SimpleParams)(nil),
		(*PayloadConfig_ComplexParams)(nil),
	}
}

func init() {
	proto.RegisterType((*ByteBufferParams)(nil), "grpc.testing.ByteBufferParams")
	proto.RegisterType((*SimpleProtoParams)(nil), "grpc.testing.SimpleProtoParams")
	proto.RegisterType((*ComplexProtoParams)(nil), "grpc.testing.ComplexProtoParams")
	proto.RegisterType((*PayloadConfig)(nil), "grpc.testing.PayloadConfig")
}

func init() { proto.RegisterFile("payloads.proto", fileDescriptor_3a075f58f70088c8) }

var fileDescriptor_3a075f58f70088c8 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0xac, 0xcc,
	0xc9, 0x4f, 0x4c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x57, 0xf2, 0xe2, 0x12, 0x70, 0xaa, 0x2c, 0x49,
	0x75, 0x2a, 0x4d, 0x4b, 0x4b, 0x2d, 0x0a, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0x16, 0x92, 0xe4, 0xe2,
	0x28, 0x4a, 0x2d, 0x8c, 0x2f, 0xce, 0xac, 0x4a, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62,
	0x2f, 0x4a, 0x2d, 0x0c, 0xce, 0xac, 0x4a, 0x15, 0x92, 0xe6, 0xe2, 0x2c, 0x4a, 0x2d, 0x2e, 0x80,
	0xc8, 0x31, 0x81, 0xe5, 0x38, 0x40, 0x02, 0x20, 0x49, 0x25, 0x6f, 0x2e, 0xc1, 0xe0, 0xcc, 0xdc,
	0x82, 0x9c, 0xd4, 0x00, 0x90, 0x45, 0x14, 0x1a, 0x26, 0xc2, 0x25, 0xe4, 0x9c, 0x0f, 0x32, 0xac,
	0x02, 0xc9, 0x34, 0xa5, 0x6f, 0x8c, 0x5c, 0xbc, 0x01, 0x10, 0xff, 0x38, 0xe7, 0xe7, 0xa5, 0x65,
	0xa6, 0x0b, 0xb9, 0x73, 0xf1, 0x25, 0x55, 0x96, 0xa4, 0x26, 0x95, 0xa6, 0xc5, 0x17, 0x80, 0xd5,
	0x80, 0x6d, 0xe1, 0x36, 0x92, 0xd3, 0x43, 0xf6, 0xa7, 0x1e, 0xba, 0x27, 0x3d, 0x18, 0x82, 0x78,
	0xa1, 0xfa, 0xa0, 0x0e, 0x75, 0xe3, 0xe2, 0x2d, 0x06, 0xbb, 0x1e, 0x66, 0x0e, 0x13, 0xd8, 0x1c,
	0x79, 0x54, 0x73, 0x30, 0x3c, 0xe8, 0xc1, 0x10, 0xc4, 0x03, 0xd1, 0x07, 0x35, 0xc7, 0x93, 0x8b,
	0x2f, 0x19, 0xe2, 0x70, 0x98, 0x41, 0xcc, 0x60, 0x83, 0x14, 0x50, 0x0d, 0xc2, 0xf4, 0x1c, 0xc8,
	0x49, 0x50, 0x9d, 0x10, 0x01, 0x27, 0x4e, 0x2e, 0x76, 0x68, 0xe4, 0x25, 0xb1, 0x81, 0x23, 0xcf,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x8c, 0x18, 0x4e, 0xce, 0x01, 0x00, 0x00,
}
