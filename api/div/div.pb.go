// Code generated by protoc-gen-go. DO NOT EDIT.
// source: div/div/div.proto

package div

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

type Error_Code int32

const (
	Error_UNAUTHORIZED          Error_Code = 0
	Error_BAD_REQUEST           Error_Code = 1
	Error_INTERNAL_SERVER_ERROR Error_Code = 2
)

var Error_Code_name = map[int32]string{
	0: "UNAUTHORIZED",
	1: "BAD_REQUEST",
	2: "INTERNAL_SERVER_ERROR",
}

var Error_Code_value = map[string]int32{
	"UNAUTHORIZED":          0,
	"BAD_REQUEST":           1,
	"INTERNAL_SERVER_ERROR": 2,
}

func (x Error_Code) String() string {
	return proto.EnumName(Error_Code_name, int32(x))
}

func (Error_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b3c29ef71785dab, []int{2, 0}
}

type Request struct {
	Term1                int64    `protobuf:"varint,1,opt,name=term1,proto3" json:"term1,omitempty"`
	Term2                int64    `protobuf:"varint,2,opt,name=term2,proto3" json:"term2,omitempty"`
	IsAuthorized         bool     `protobuf:"varint,3,opt,name=isAuthorized,proto3" json:"isAuthorized,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b3c29ef71785dab, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTerm1() int64 {
	if m != nil {
		return m.Term1
	}
	return 0
}

func (m *Request) GetTerm2() int64 {
	if m != nil {
		return m.Term2
	}
	return 0
}

func (m *Request) GetIsAuthorized() bool {
	if m != nil {
		return m.IsAuthorized
	}
	return false
}

type Response struct {
	Div                  int64    `protobuf:"varint,1,opt,name=div,proto3" json:"div,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b3c29ef71785dab, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetDiv() int64 {
	if m != nil {
		return m.Div
	}
	return 0
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=div.Error_Code" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b3c29ef71785dab, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() Error_Code {
	if m != nil {
		return m.Code
	}
	return Error_UNAUTHORIZED
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("div.Error_Code", Error_Code_name, Error_Code_value)
	proto.RegisterType((*Request)(nil), "div.Request")
	proto.RegisterType((*Response)(nil), "div.Response")
	proto.RegisterType((*Error)(nil), "div.Error")
}

func init() { proto.RegisterFile("div/div/div.proto", fileDescriptor_7b3c29ef71785dab) }

var fileDescriptor_7b3c29ef71785dab = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x9b, 0xa6, 0xfd, 0xb7, 0xff, 0x69, 0xb4, 0x71, 0x50, 0x88, 0x9e, 0xc2, 0xf6, 0x52,
	0x3c, 0x44, 0x8c, 0x77, 0x21, 0x9a, 0x05, 0x0b, 0x92, 0xe2, 0x34, 0x11, 0xf4, 0x12, 0xd4, 0x2c,
	0x9a, 0x43, 0xdd, 0xba, 0x9b, 0xe6, 0xe0, 0x67, 0xf0, 0x43, 0x4b, 0x36, 0x29, 0xc5, 0xc3, 0xc2,
	0xbe, 0xdf, 0xcc, 0xbe, 0x9d, 0x79, 0x70, 0x54, 0x94, 0xf5, 0x45, 0x77, 0x82, 0x8d, 0x92, 0x95,
	0x44, 0xbb, 0x28, 0x6b, 0xf6, 0x04, 0x23, 0x12, 0x5f, 0x5b, 0xa1, 0x2b, 0x3c, 0x86, 0x61, 0x25,
	0xd4, 0xfa, 0xd2, 0xb3, 0x7c, 0x6b, 0x6e, 0x53, 0x2b, 0x76, 0x34, 0xf4, 0xfa, 0x7b, 0x1a, 0x22,
	0x03, 0xa7, 0xd4, 0xd1, 0xb6, 0xfa, 0x90, 0xaa, 0xfc, 0x16, 0x85, 0x67, 0xfb, 0xd6, 0x7c, 0x4c,
	0x7f, 0x18, 0xbb, 0x86, 0x31, 0x09, 0xbd, 0x91, 0x9f, 0x5a, 0xa0, 0x0b, 0xcd, 0x6f, 0x9d, 0x73,
	0x73, 0x45, 0x1f, 0x86, 0x42, 0x29, 0xa9, 0x8c, 0xef, 0x24, 0x84, 0xa0, 0x19, 0x8c, 0x37, 0x84,
	0xda, 0x02, 0xfb, 0xb1, 0x60, 0x68, 0x00, 0xce, 0x60, 0xf0, 0x26, 0x0b, 0x61, 0x9e, 0x1f, 0x86,
	0xd3, 0x7d, 0x6b, 0x70, 0x2b, 0x0b, 0x41, 0xa6, 0x88, 0x1e, 0x8c, 0xd6, 0x42, 0xeb, 0x97, 0x77,
	0x61, 0x2c, 0xff, 0xd3, 0x4e, 0xb2, 0x18, 0x06, 0x4d, 0x1f, 0xba, 0xe0, 0x64, 0x49, 0x94, 0xa5,
	0x77, 0x4b, 0x5a, 0x3c, 0xf3, 0xd8, 0xed, 0xe1, 0x14, 0x26, 0x37, 0x51, 0x9c, 0x13, 0x7f, 0xc8,
	0xf8, 0x2a, 0x75, 0x2d, 0x3c, 0x85, 0x93, 0x45, 0x92, 0x72, 0x4a, 0xa2, 0xfb, 0x7c, 0xc5, 0xe9,
	0x91, 0x53, 0xce, 0x89, 0x96, 0xe4, 0xf6, 0xc3, 0x73, 0xb0, 0xe3, 0xb2, 0xc6, 0x19, 0xf4, 0x63,
	0x89, 0x8e, 0x99, 0xa1, 0x4b, 0xee, 0xec, 0xa0, 0x53, 0xed, 0xb2, 0xac, 0xf7, 0xfa, 0xcf, 0x24,
	0x7c, 0xf5, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xf7, 0xcd, 0x77, 0x76, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DivClient is the client API for Div service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DivClient interface {
	Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type divClient struct {
	cc *grpc.ClientConn
}

func NewDivClient(cc *grpc.ClientConn) DivClient {
	return &divClient{cc}
}

func (c *divClient) Do(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/div.Div/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DivServer is the server API for Div service.
type DivServer interface {
	Do(context.Context, *Request) (*Response, error)
}

// UnimplementedDivServer can be embedded to have forward compatible implementations.
type UnimplementedDivServer struct {
}

func (*UnimplementedDivServer) Do(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}

func RegisterDivServer(s *grpc.Server, srv DivServer) {
	s.RegisterService(&_Div_serviceDesc, srv)
}

func _Div_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/div.Div/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivServer).Do(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Div_serviceDesc = grpc.ServiceDesc{
	ServiceName: "div.Div",
	HandlerType: (*DivServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Div_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "div/div/div.proto",
}