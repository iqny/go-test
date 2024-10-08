// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package api

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

// Timestamp from public import google/protobuf/timestamp.proto
type Timestamp = timestamp.Timestamp

type ListResponse struct {
	Order                []*Order `protobuf:"bytes,2,rep,name=Order,proto3" json:"Order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetOrder() []*Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Order struct {
	Id                   int64                `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Birthday             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

type PageRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageRequest) Reset()         { *m = PageRequest{} }
func (m *PageRequest) String() string { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()    {}
func (*PageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *PageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageRequest.Unmarshal(m, b)
}
func (m *PageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageRequest.Marshal(b, m, deterministic)
}
func (m *PageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageRequest.Merge(m, src)
}
func (m *PageRequest) XXX_Size() int {
	return xxx_messageInfo_PageRequest.Size(m)
}
func (m *PageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PageRequest proto.InternalMessageInfo

func (m *PageRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func init() {
	proto.RegisterType((*ListResponse)(nil), "main.ListResponse")
	proto.RegisterType((*Order)(nil), "main.Order")
	proto.RegisterType((*PageRequest)(nil), "main.PageRequest")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x3f, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0xc9, 0x1f, 0xaa, 0xf2, 0x52, 0x21, 0xf1, 0xa6, 0x28, 0x0b, 0x69, 0xa6, 0x2c, 0x38,
	0x22, 0x48, 0x2c, 0x88, 0x85, 0xad, 0x12, 0x82, 0xca, 0x30, 0xb1, 0x39, 0xcd, 0x23, 0x58, 0x22,
	0x75, 0xb0, 0x5d, 0x24, 0xbe, 0x3d, 0xb2, 0xdd, 0x56, 0xdd, 0xce, 0xe7, 0xf3, 0xf9, 0x77, 0x90,
	0x29, 0xdd, 0x93, 0x66, 0x93, 0x56, 0x56, 0x61, 0x3a, 0x0a, 0xb9, 0x2d, 0xae, 0x07, 0xa5, 0x86,
	0x6f, 0x6a, 0xbc, 0xd7, 0xed, 0x3e, 0x1b, 0x2b, 0x47, 0x32, 0x56, 0x8c, 0x53, 0x88, 0x55, 0xb7,
	0xb0, 0x78, 0x96, 0xc6, 0x72, 0x32, 0x93, 0xda, 0x1a, 0xc2, 0x25, 0x9c, 0xbf, 0xba, 0x96, 0x3c,
	0x2e, 0x93, 0x3a, 0x6b, 0x33, 0xe6, 0x6a, 0x98, 0xb7, 0x78, 0xb8, 0xa9, 0x36, 0xfb, 0x08, 0x5e,
	0x42, 0xbc, 0xea, 0xf3, 0xa8, 0x8c, 0xea, 0x84, 0xc7, 0xab, 0x1e, 0x11, 0xd2, 0x17, 0x31, 0x52,
	0x1e, 0x97, 0x51, 0x7d, 0xc1, 0xbd, 0xc6, 0x7b, 0x98, 0x77, 0x52, 0xdb, 0xaf, 0x5e, 0xfc, 0xe5,
	0x49, 0x19, 0xd5, 0x59, 0x5b, 0xb0, 0xc0, 0xc4, 0x0e, 0x4c, 0xec, 0xfd, 0xc0, 0xc4, 0x8f, 0xd9,
	0x6a, 0x09, 0xd9, 0x5a, 0x0c, 0xc4, 0xe9, 0x67, 0x47, 0xc6, 0xba, 0x6a, 0x77, 0xdc, 0x7f, 0xe6,
	0x75, 0xfb, 0x08, 0x0b, 0xcf, 0xf1, 0x46, 0xfa, 0x57, 0x6e, 0x08, 0x6f, 0x20, 0x75, 0x53, 0xf0,
	0x2a, 0x30, 0x9f, 0x3c, 0x2f, 0x30, 0x58, 0xa7, 0x4b, 0x9f, 0xe6, 0x1f, 0x33, 0xd6, 0x3c, 0x88,
	0x49, 0xae, 0xcf, 0xba, 0x99, 0x67, 0xb9, 0xfb, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x65, 0x24,
	0x81, 0x42, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	List(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) List(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/main.OrderService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	List(context.Context, *PageRequest) (*ListResponse, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) List(ctx context.Context, req *PageRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.OrderService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).List(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _OrderService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
