// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reviews.proto

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

type QueryReviewsRequest struct {
	ProductID            uint64   `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryReviewsRequest) Reset()         { *m = QueryReviewsRequest{} }
func (m *QueryReviewsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryReviewsRequest) ProtoMessage()    {}
func (*QueryReviewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2a6bc9211d2c89, []int{0}
}

func (m *QueryReviewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReviewsRequest.Unmarshal(m, b)
}
func (m *QueryReviewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReviewsRequest.Marshal(b, m, deterministic)
}
func (m *QueryReviewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReviewsRequest.Merge(m, src)
}
func (m *QueryReviewsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryReviewsRequest.Size(m)
}
func (m *QueryReviewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReviewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReviewsRequest proto.InternalMessageInfo

func (m *QueryReviewsRequest) GetProductID() uint64 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

type QueryReviewsResponse struct {
	Reviews              []*Review `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryReviewsResponse) Reset()         { *m = QueryReviewsResponse{} }
func (m *QueryReviewsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryReviewsResponse) ProtoMessage()    {}
func (*QueryReviewsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2a6bc9211d2c89, []int{1}
}

func (m *QueryReviewsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryReviewsResponse.Unmarshal(m, b)
}
func (m *QueryReviewsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryReviewsResponse.Marshal(b, m, deterministic)
}
func (m *QueryReviewsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReviewsResponse.Merge(m, src)
}
func (m *QueryReviewsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryReviewsResponse.Size(m)
}
func (m *QueryReviewsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReviewsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReviewsResponse proto.InternalMessageInfo

func (m *QueryReviewsResponse) GetReviews() []*Review {
	if m != nil {
		return m.Reviews
	}
	return nil
}

type Review struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductID            uint64               `protobuf:"varint,2,opt,name=productID,proto3" json:"productID,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	CreatedTime          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Review) Reset()         { *m = Review{} }
func (m *Review) String() string { return proto.CompactTextString(m) }
func (*Review) ProtoMessage()    {}
func (*Review) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2a6bc9211d2c89, []int{2}
}

func (m *Review) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Review.Unmarshal(m, b)
}
func (m *Review) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Review.Marshal(b, m, deterministic)
}
func (m *Review) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Review.Merge(m, src)
}
func (m *Review) XXX_Size() int {
	return xxx_messageInfo_Review.Size(m)
}
func (m *Review) XXX_DiscardUnknown() {
	xxx_messageInfo_Review.DiscardUnknown(m)
}

var xxx_messageInfo_Review proto.InternalMessageInfo

func (m *Review) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Review) GetProductID() uint64 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

func (m *Review) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Review) GetCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryReviewsRequest)(nil), "QueryReviewsRequest")
	proto.RegisterType((*QueryReviewsResponse)(nil), "QueryReviewsResponse")
	proto.RegisterType((*Review)(nil), "Review")
}

func init() { proto.RegisterFile("reviews.proto", fileDescriptor_5b2a6bc9211d2c89) }

var fileDescriptor_5b2a6bc9211d2c89 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0xf7, 0x23, 0xec, 0x14, 0x3d, 0xc4, 0x15, 0x42, 0x11, 0xac, 0x3d, 0xf5, 0x94,
	0x85, 0xae, 0x17, 0x41, 0x10, 0xc4, 0x8b, 0x47, 0xc3, 0x9e, 0xbc, 0xd5, 0xed, 0x58, 0x02, 0xd6,
	0xd4, 0x7c, 0x28, 0xfe, 0x09, 0x7f, 0xb3, 0xd0, 0x34, 0xe8, 0x2e, 0x3d, 0x85, 0x79, 0x79, 0x27,
	0x0f, 0xcf, 0xc0, 0x89, 0xc1, 0x4f, 0x85, 0x5f, 0x56, 0xf4, 0x46, 0x3b, 0x9d, 0x5d, 0xb6, 0x5a,
	0xb7, 0x6f, 0xb8, 0x19, 0xa6, 0x17, 0xff, 0xba, 0x71, 0xaa, 0x43, 0xeb, 0xea, 0xae, 0x0f, 0x85,
	0x62, 0x0b, 0x67, 0x4f, 0x1e, 0xcd, 0xb7, 0x0c, 0x6b, 0x12, 0x3f, 0x3c, 0x5a, 0xc7, 0x2e, 0x60,
	0xd5, 0x1b, 0xdd, 0xf8, 0xbd, 0x7b, 0x7c, 0xe0, 0x24, 0x27, 0xe5, 0x5c, 0xfe, 0x05, 0xc5, 0x0d,
	0xac, 0x0f, 0x97, 0x6c, 0xaf, 0xdf, 0x2d, 0xb2, 0x2b, 0xa0, 0x23, 0x9e, 0x93, 0x7c, 0x56, 0xa6,
	0x15, 0x15, 0xa1, 0x22, 0x63, 0x5e, 0xfc, 0x10, 0x58, 0x86, 0x8c, 0x9d, 0x42, 0xa2, 0x9a, 0xf1,
	0xf3, 0x44, 0x35, 0x87, 0xcc, 0xe4, 0x88, 0xc9, 0x38, 0xd0, 0x0e, 0xad, 0xad, 0x5b, 0xe4, 0xb3,
	0x9c, 0x94, 0x2b, 0x19, 0x47, 0x76, 0x0b, 0xe9, 0xde, 0x60, 0xed, 0xb0, 0xd9, 0xa9, 0x0e, 0xf9,
	0x3c, 0x27, 0x65, 0x5a, 0x65, 0x22, 0x98, 0x8b, 0x68, 0x2e, 0x76, 0xd1, 0x5c, 0xfe, 0xaf, 0x57,
	0x77, 0x40, 0x47, 0x0d, 0x76, 0x0d, 0x8b, 0x41, 0x8b, 0xad, 0xc5, 0xc4, 0x4d, 0xb2, 0x73, 0x31,
	0x25, 0x7d, 0x4f, 0x9f, 0x17, 0x81, 0xb1, 0x1c, 0x9e, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5a, 0x3e, 0x19, 0x05, 0x83, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReviewsClient is the client API for Reviews service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewsClient interface {
	Query(ctx context.Context, in *QueryReviewsRequest, opts ...grpc.CallOption) (*QueryReviewsResponse, error)
}

type reviewsClient struct {
	cc *grpc.ClientConn
}

func NewReviewsClient(cc *grpc.ClientConn) ReviewsClient {
	return &reviewsClient{cc}
}

func (c *reviewsClient) Query(ctx context.Context, in *QueryReviewsRequest, opts ...grpc.CallOption) (*QueryReviewsResponse, error) {
	out := new(QueryReviewsResponse)
	err := c.cc.Invoke(ctx, "/Reviews/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewsServer is the server API for Reviews service.
type ReviewsServer interface {
	Query(context.Context, *QueryReviewsRequest) (*QueryReviewsResponse, error)
}

// UnimplementedReviewsServer can be embedded to have forward compatible implementations.
type UnimplementedReviewsServer struct {
}

func (*UnimplementedReviewsServer) Query(ctx context.Context, req *QueryReviewsRequest) (*QueryReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterReviewsServer(s *grpc.Server, srv ReviewsServer) {
	s.RegisterService(&_Reviews_serviceDesc, srv)
}

func _Reviews_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Reviews/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).Query(ctx, req.(*QueryReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reviews_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Reviews",
	HandlerType: (*ReviewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Reviews_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reviews.proto",
}
