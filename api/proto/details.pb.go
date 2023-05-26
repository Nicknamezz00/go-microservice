// Code generated by protoc-gen-go. DO NOT EDIT.
// source: details.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetDetailsRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailsRequest) Reset()         { *m = GetDetailsRequest{} }
func (m *GetDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDetailsRequest) ProtoMessage()    {}
func (*GetDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c457112169d4fb2c, []int{0}
}

func (m *GetDetailsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailsRequest.Unmarshal(m, b)
}
func (m *GetDetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailsRequest.Marshal(b, m, deterministic)
}
func (m *GetDetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailsRequest.Merge(m, src)
}
func (m *GetDetailsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDetailsRequest.Size(m)
}
func (m *GetDetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailsRequest proto.InternalMessageInfo

func (m *GetDetailsRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Detail struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                float32              `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	CreatedTime          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Detail) Reset()         { *m = Detail{} }
func (m *Detail) String() string { return proto.CompactTextString(m) }
func (*Detail) ProtoMessage()    {}
func (*Detail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c457112169d4fb2c, []int{1}
}

func (m *Detail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Detail.Unmarshal(m, b)
}
func (m *Detail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Detail.Marshal(b, m, deterministic)
}
func (m *Detail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Detail.Merge(m, src)
}
func (m *Detail) XXX_Size() int {
	return xxx_messageInfo_Detail.Size(m)
}
func (m *Detail) XXX_DiscardUnknown() {
	xxx_messageInfo_Detail.DiscardUnknown(m)
}

var xxx_messageInfo_Detail proto.InternalMessageInfo

func (m *Detail) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Detail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Detail) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Detail) GetCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDetailsRequest)(nil), "GetDetailsRequest")
	proto.RegisterType((*Detail)(nil), "Detail")
}

func init() { proto.RegisterFile("details.proto", fileDescriptor_c457112169d4fb2c) }

var fileDescriptor_c457112169d4fb2c = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x46, 0x49, 0x6e, 0xef, 0x16, 0xe7, 0x50, 0x70, 0xb0, 0x08, 0xdb, 0x18, 0xd6, 0x26, 0x8d,
	0x39, 0x38, 0x5b, 0x2b, 0x11, 0xae, 0x0f, 0x57, 0xd9, 0xc5, 0xcb, 0x78, 0x04, 0x36, 0x66, 0xdd,
	0x9d, 0xed, 0xfd, 0xe9, 0xc2, 0xc6, 0x05, 0xf1, 0xaa, 0x99, 0x6f, 0xe6, 0xcd, 0xf0, 0xe0, 0x3a,
	0x10, 0xfb, 0xd8, 0x8d, 0xb6, 0x1f, 0x32, 0xe7, 0xe6, 0xfe, 0x9c, 0xf3, 0xb9, 0xa3, 0xdd, 0x9c,
	0xde, 0xa7, 0x8f, 0x1d, 0xc7, 0x44, 0x23, 0xfb, 0xd4, 0x17, 0xa0, 0x7d, 0x80, 0xdb, 0x03, 0xf1,
	0x6b, 0x39, 0x72, 0xf4, 0x35, 0xd1, 0xc8, 0x78, 0x03, 0x32, 0x06, 0x25, 0xb4, 0x30, 0x95, 0x93,
	0x31, 0xb4, 0xdf, 0x02, 0x36, 0x05, 0xf9, 0xbf, 0x42, 0x84, 0xea, 0xd3, 0x27, 0x52, 0x52, 0x0b,
	0x73, 0xe5, 0xe6, 0x1e, 0xef, 0x60, 0xdd, 0x0f, 0xf1, 0x44, 0x6a, 0xa5, 0x85, 0x91, 0xae, 0x04,
	0x7c, 0x86, 0xed, 0x69, 0x20, 0xcf, 0x14, 0x8e, 0x31, 0x91, 0xaa, 0xb4, 0x30, 0xdb, 0x7d, 0x63,
	0x8b, 0xa0, 0x5d, 0x04, 0xed, 0x71, 0x11, 0x74, 0x7f, 0xf1, 0xfd, 0x23, 0xd4, 0xbf, 0x92, 0xd8,
	0xc2, 0xea, 0x40, 0x8c, 0x68, 0x2f, 0xc4, 0x9b, 0xda, 0x96, 0xc1, 0x4b, 0xfd, 0xb6, 0x2e, 0x1f,
	0x37, 0x73, 0x79, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xab, 0xf8, 0x3a, 0x65, 0x18, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DetailsClient is the client API for Details service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetailsClient interface {
	Get(ctx context.Context, in *GetDetailsRequest, opts ...grpc.CallOption) (*Detail, error)
}

type detailsClient struct {
	cc *grpc.ClientConn
}

func NewDetailsClient(cc *grpc.ClientConn) DetailsClient {
	return &detailsClient{cc}
}

func (c *detailsClient) Get(ctx context.Context, in *GetDetailsRequest, opts ...grpc.CallOption) (*Detail, error) {
	out := new(Detail)
	err := c.cc.Invoke(ctx, "/Details/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailsServer is the server API for Details service.
type DetailsServer interface {
	Get(context.Context, *GetDetailsRequest) (*Detail, error)
}

// UnimplementedDetailsServer can be embedded to have forward compatible implementations.
type UnimplementedDetailsServer struct {
}

func (*UnimplementedDetailsServer) Get(ctx context.Context, req *GetDetailsRequest) (*Detail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterDetailsServer(s *grpc.Server, srv DetailsServer) {
	s.RegisterService(&_Details_serviceDesc, srv)
}

func _Details_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Details/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailsServer).Get(ctx, req.(*GetDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Details_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Details",
	HandlerType: (*DetailsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Details_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "details.proto",
}
