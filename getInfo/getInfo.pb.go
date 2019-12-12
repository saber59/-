// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getInfo.proto

package getInfo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Rec struct {
	R                    string   `protobuf:"bytes,1,opt,name=R,proto3" json:"R,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rec) Reset()         { *m = Rec{} }
func (m *Rec) String() string { return proto.CompactTextString(m) }
func (*Rec) ProtoMessage()    {}
func (*Rec) Descriptor() ([]byte, []int) {
	return fileDescriptor_13db4f652dbf8fe9, []int{0}
}

func (m *Rec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rec.Unmarshal(m, b)
}
func (m *Rec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rec.Marshal(b, m, deterministic)
}
func (m *Rec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rec.Merge(m, src)
}
func (m *Rec) XXX_Size() int {
	return xxx_messageInfo_Rec.Size(m)
}
func (m *Rec) XXX_DiscardUnknown() {
	xxx_messageInfo_Rec.DiscardUnknown(m)
}

var xxx_messageInfo_Rec proto.InternalMessageInfo

func (m *Rec) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

type Send struct {
	StuName              string   `protobuf:"bytes,1,opt,name=StuName,proto3" json:"StuName,omitempty"`
	DormNo               string   `protobuf:"bytes,2,opt,name=DormNo,proto3" json:"DormNo,omitempty"`
	BuildNum             string   `protobuf:"bytes,3,opt,name=BuildNum,proto3" json:"BuildNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Send) Reset()         { *m = Send{} }
func (m *Send) String() string { return proto.CompactTextString(m) }
func (*Send) ProtoMessage()    {}
func (*Send) Descriptor() ([]byte, []int) {
	return fileDescriptor_13db4f652dbf8fe9, []int{1}
}

func (m *Send) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Send.Unmarshal(m, b)
}
func (m *Send) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Send.Marshal(b, m, deterministic)
}
func (m *Send) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Send.Merge(m, src)
}
func (m *Send) XXX_Size() int {
	return xxx_messageInfo_Send.Size(m)
}
func (m *Send) XXX_DiscardUnknown() {
	xxx_messageInfo_Send.DiscardUnknown(m)
}

var xxx_messageInfo_Send proto.InternalMessageInfo

func (m *Send) GetStuName() string {
	if m != nil {
		return m.StuName
	}
	return ""
}

func (m *Send) GetDormNo() string {
	if m != nil {
		return m.DormNo
	}
	return ""
}

func (m *Send) GetBuildNum() string {
	if m != nil {
		return m.BuildNum
	}
	return ""
}

func init() {
	proto.RegisterType((*Rec)(nil), "getInfo.Rec")
	proto.RegisterType((*Send)(nil), "getInfo.Send")
}

func init() { proto.RegisterFile("getInfo.proto", fileDescriptor_13db4f652dbf8fe9) }

var fileDescriptor_13db4f652dbf8fe9 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0x2d, 0xf1,
	0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x84, 0xb9,
	0x98, 0x83, 0x52, 0x93, 0x85, 0x78, 0xb8, 0x18, 0x83, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x18, 0x83, 0x94, 0x42, 0xb8, 0x58, 0x82, 0x53, 0xf3, 0x52, 0x84, 0x24, 0xb8, 0xd8, 0x83, 0x4b,
	0x4a, 0xfd, 0x12, 0x73, 0x53, 0xa1, 0x72, 0x30, 0xae, 0x90, 0x18, 0x17, 0x9b, 0x4b, 0x7e, 0x51,
	0xae, 0x5f, 0xbe, 0x04, 0x13, 0x58, 0x02, 0xca, 0x13, 0x92, 0xe2, 0xe2, 0x70, 0x2a, 0xcd, 0xcc,
	0x49, 0xf1, 0x2b, 0xcd, 0x95, 0x60, 0x06, 0xcb, 0xc0, 0xf9, 0x46, 0x06, 0x5c, 0x30, 0x5b, 0x85,
	0x54, 0xb9, 0x58, 0xc0, 0x34, 0x8f, 0x1e, 0xcc, 0x59, 0x41, 0xa9, 0xc9, 0x52, 0xbc, 0x70, 0x1e,
	0xc8, 0x76, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x63, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11,
	0x79, 0x22, 0xec, 0xbd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GetInfoClient is the client API for GetInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetInfoClient interface {
	Info(ctx context.Context, in *Rec, opts ...grpc.CallOption) (*Send, error)
}

type getInfoClient struct {
	cc *grpc.ClientConn
}

func NewGetInfoClient(cc *grpc.ClientConn) GetInfoClient {
	return &getInfoClient{cc}
}

func (c *getInfoClient) Info(ctx context.Context, in *Rec, opts ...grpc.CallOption) (*Send, error) {
	out := new(Send)
	err := c.cc.Invoke(ctx, "/getInfo.getInfo/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetInfoServer is the server API for GetInfo service.
type GetInfoServer interface {
	Info(context.Context, *Rec) (*Send, error)
}

// UnimplementedGetInfoServer can be embedded to have forward compatible implementations.
type UnimplementedGetInfoServer struct {
}

func (*UnimplementedGetInfoServer) Info(ctx context.Context, req *Rec) (*Send, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterGetInfoServer(s *grpc.Server, srv GetInfoServer) {
	s.RegisterService(&_GetInfo_serviceDesc, srv)
}

func _GetInfo_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetInfoServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/getInfo.getInfo/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetInfoServer).Info(ctx, req.(*Rec))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "getInfo.getInfo",
	HandlerType: (*GetInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _GetInfo_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getInfo.proto",
}
