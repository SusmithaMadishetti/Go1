// Code generated by protoc-gen-go. DO NOT EDIT.
// source: squareroot.proto

package squareroot

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

type SquarerootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquarerootRequest) Reset()         { *m = SquarerootRequest{} }
func (m *SquarerootRequest) String() string { return proto.CompactTextString(m) }
func (*SquarerootRequest) ProtoMessage()    {}
func (*SquarerootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51098fddf27cb29, []int{0}
}

func (m *SquarerootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquarerootRequest.Unmarshal(m, b)
}
func (m *SquarerootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquarerootRequest.Marshal(b, m, deterministic)
}
func (m *SquarerootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquarerootRequest.Merge(m, src)
}
func (m *SquarerootRequest) XXX_Size() int {
	return xxx_messageInfo_SquarerootRequest.Size(m)
}
func (m *SquarerootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquarerootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquarerootRequest proto.InternalMessageInfo

func (m *SquarerootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquarerootResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquarerootResponse) Reset()         { *m = SquarerootResponse{} }
func (m *SquarerootResponse) String() string { return proto.CompactTextString(m) }
func (*SquarerootResponse) ProtoMessage()    {}
func (*SquarerootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51098fddf27cb29, []int{1}
}

func (m *SquarerootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquarerootResponse.Unmarshal(m, b)
}
func (m *SquarerootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquarerootResponse.Marshal(b, m, deterministic)
}
func (m *SquarerootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquarerootResponse.Merge(m, src)
}
func (m *SquarerootResponse) XXX_Size() int {
	return xxx_messageInfo_SquarerootResponse.Size(m)
}
func (m *SquarerootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquarerootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquarerootResponse proto.InternalMessageInfo

func (m *SquarerootResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SquarerootRequest)(nil), "squarerootpb.squarerootRequest")
	proto.RegisterType((*SquarerootResponse)(nil), "squarerootpb.squarerootResponse")
}

func init() { proto.RegisterFile("squareroot.proto", fileDescriptor_c51098fddf27cb29) }

var fileDescriptor_c51098fddf27cb29 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0x2c, 0x4d,
	0x2c, 0x4a, 0x2d, 0xca, 0xcf, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x88,
	0x14, 0x24, 0x29, 0x69, 0x73, 0x09, 0x22, 0xf8, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x62, 0x5c, 0x6c, 0x79, 0xa5, 0xb9, 0x49, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41,
	0x50, 0x9e, 0x92, 0x0e, 0x97, 0x10, 0xb2, 0xe2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0xea,
	0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xb0, 0x6a, 0xc6, 0x20, 0x28, 0xcf, 0x28, 0x8d, 0x4b, 0x30,
	0x18, 0xae, 0x3a, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x90, 0x8b, 0x0b, 0x61, 0x84,
	0x90, 0xbc, 0x1e, 0xb2, 0x63, 0xf4, 0x30, 0x5c, 0x22, 0xa5, 0x80, 0x5b, 0x01, 0xc4, 0x76, 0x25,
	0x06, 0x27, 0x9e, 0x28, 0x24, 0x23, 0x93, 0xd8, 0xc0, 0xbe, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x0c, 0xe5, 0xb1, 0x63, 0xf9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SquarerootServiceClient is the client API for SquarerootService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquarerootServiceClient interface {
	//error handling
	//this RPC will throw an exception if the sent number is negative
	//The error being sent is of type INVALID_ARGUMENT
	Squareroot(ctx context.Context, in *SquarerootRequest, opts ...grpc.CallOption) (*SquarerootResponse, error)
}

type squarerootServiceClient struct {
	cc *grpc.ClientConn
}

func NewSquarerootServiceClient(cc *grpc.ClientConn) SquarerootServiceClient {
	return &squarerootServiceClient{cc}
}

func (c *squarerootServiceClient) Squareroot(ctx context.Context, in *SquarerootRequest, opts ...grpc.CallOption) (*SquarerootResponse, error) {
	out := new(SquarerootResponse)
	err := c.cc.Invoke(ctx, "/squarerootpb.SquarerootService/squareroot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquarerootServiceServer is the server API for SquarerootService service.
type SquarerootServiceServer interface {
	//error handling
	//this RPC will throw an exception if the sent number is negative
	//The error being sent is of type INVALID_ARGUMENT
	Squareroot(context.Context, *SquarerootRequest) (*SquarerootResponse, error)
}

// UnimplementedSquarerootServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSquarerootServiceServer struct {
}

func (*UnimplementedSquarerootServiceServer) Squareroot(ctx context.Context, req *SquarerootRequest) (*SquarerootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Squareroot not implemented")
}

func RegisterSquarerootServiceServer(s *grpc.Server, srv SquarerootServiceServer) {
	s.RegisterService(&_SquarerootService_serviceDesc, srv)
}

func _SquarerootService_Squareroot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquarerootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquarerootServiceServer).Squareroot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squarerootpb.SquarerootService/Squareroot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquarerootServiceServer).Squareroot(ctx, req.(*SquarerootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SquarerootService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "squarerootpb.SquarerootService",
	HandlerType: (*SquarerootServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "squareroot",
			Handler:    _SquarerootService_Squareroot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squareroot.proto",
}
