// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/user-response/delivery/grpc/response_grpc/response.proto

package response_grpc

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

type Response struct {
	ID                   int64                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FreelancerId         int64                `protobuf:"varint,2,opt,name=FreelancerId,proto3" json:"FreelancerId,omitempty"`
	JobId                int64                `protobuf:"varint,3,opt,name=JobId,proto3" json:"JobId,omitempty"`
	Files                string               `protobuf:"bytes,4,opt,name=Files,proto3" json:"Files,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,5,opt,name=Date,proto3" json:"Date,omitempty"`
	StatusManager        string               `protobuf:"bytes,6,opt,name=StatusManager,proto3" json:"StatusManager,omitempty"`
	StatusFreelancer     string               `protobuf:"bytes,7,opt,name=StatusFreelancer,proto3" json:"StatusFreelancer,omitempty"`
	PaymentAmount        float32              `protobuf:"fixed32,8,opt,name=PaymentAmount,proto3" json:"PaymentAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_43dede0a4d2699ce, []int{0}
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

func (m *Response) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Response) GetFreelancerId() int64 {
	if m != nil {
		return m.FreelancerId
	}
	return 0
}

func (m *Response) GetJobId() int64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *Response) GetFiles() string {
	if m != nil {
		return m.Files
	}
	return ""
}

func (m *Response) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Response) GetStatusManager() string {
	if m != nil {
		return m.StatusManager
	}
	return ""
}

func (m *Response) GetStatusFreelancer() string {
	if m != nil {
		return m.StatusFreelancer
	}
	return ""
}

func (m *Response) GetPaymentAmount() float32 {
	if m != nil {
		return m.PaymentAmount
	}
	return 0
}

type ResponseID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseID) Reset()         { *m = ResponseID{} }
func (m *ResponseID) String() string { return proto.CompactTextString(m) }
func (*ResponseID) ProtoMessage()    {}
func (*ResponseID) Descriptor() ([]byte, []int) {
	return fileDescriptor_43dede0a4d2699ce, []int{1}
}

func (m *ResponseID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseID.Unmarshal(m, b)
}
func (m *ResponseID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseID.Marshal(b, m, deterministic)
}
func (m *ResponseID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseID.Merge(m, src)
}
func (m *ResponseID) XXX_Size() int {
	return xxx_messageInfo_ResponseID.Size(m)
}
func (m *ResponseID) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseID.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseID proto.InternalMessageInfo

func (m *ResponseID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*Response)(nil), "response_grpc.Response")
	proto.RegisterType((*ResponseID)(nil), "response_grpc.ResponseID")
}

func init() {
	proto.RegisterFile("internal/app/user-response/delivery/grpc/response_grpc/response.proto", fileDescriptor_43dede0a4d2699ce)
}

var fileDescriptor_43dede0a4d2699ce = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0x69, 0xf7, 0xe7, 0xdd, 0x7b, 0x74, 0x2a, 0x41, 0x30, 0x0e, 0xc1, 0x32, 0xbc, 0x28,
	0x82, 0x09, 0xcc, 0x3b, 0xef, 0x84, 0x3a, 0xac, 0x30, 0x90, 0xea, 0xbd, 0x64, 0xeb, 0xb1, 0x14,
	0xda, 0xa4, 0x24, 0xa9, 0xb0, 0xaf, 0xe1, 0x27, 0x96, 0xa5, 0x96, 0x11, 0xc5, 0xcb, 0xf3, 0xfb,
	0x9d, 0x1c, 0x9e, 0x3c, 0xf0, 0x50, 0x4a, 0x8b, 0x5a, 0x8a, 0x8a, 0x8b, 0xa6, 0xe1, 0xad, 0x41,
	0x7d, 0xa3, 0xd1, 0x34, 0x4a, 0x1a, 0xe4, 0x39, 0x56, 0xe5, 0x07, 0xea, 0x2d, 0x2f, 0x74, 0xb3,
	0xe1, 0x3d, 0x7e, 0xf3, 0x26, 0xd6, 0x68, 0x65, 0x15, 0x99, 0x7a, 0x76, 0x76, 0x59, 0x28, 0x55,
	0x54, 0xc8, 0x9d, 0x5c, 0xb7, 0xef, 0xdc, 0x96, 0x35, 0x1a, 0x2b, 0xea, 0xa6, 0xdb, 0x9f, 0x7f,
	0x86, 0x30, 0xc9, 0xbe, 0x9f, 0x90, 0x23, 0x08, 0xd3, 0x84, 0x06, 0x51, 0x10, 0x0f, 0xb2, 0x30,
	0x4d, 0xc8, 0x1c, 0x0e, 0x97, 0x1a, 0xb1, 0x12, 0x72, 0x83, 0x3a, 0xcd, 0x69, 0xe8, 0x8c, 0xc7,
	0xc8, 0x29, 0x8c, 0x9e, 0xd4, 0x3a, 0xcd, 0xe9, 0xc0, 0xc9, 0x6e, 0xd8, 0xd1, 0x65, 0x59, 0xa1,
	0xa1, 0xc3, 0x28, 0x88, 0xff, 0x67, 0xdd, 0x40, 0x18, 0x0c, 0x13, 0x61, 0x91, 0x8e, 0xa2, 0x20,
	0x3e, 0x58, 0xcc, 0x58, 0x17, 0x8e, 0xf5, 0xe1, 0xd8, 0x6b, 0x1f, 0x2e, 0x73, 0x7b, 0xe4, 0x0a,
	0xa6, 0x2f, 0x56, 0xd8, 0xd6, 0xac, 0x84, 0x14, 0x05, 0x6a, 0x3a, 0x76, 0xd7, 0x7c, 0x48, 0xae,
	0xe1, 0xa4, 0x03, 0xfb, 0x5c, 0xf4, 0x9f, 0x5b, 0xfc, 0xc5, 0x77, 0x17, 0x9f, 0xc5, 0xb6, 0x46,
	0x69, 0xef, 0x6b, 0xd5, 0x4a, 0x4b, 0x27, 0x51, 0x10, 0x87, 0x99, 0x0f, 0xe7, 0x17, 0x00, 0x7d,
	0x27, 0x69, 0xf2, 0xb3, 0x95, 0xc5, 0x0a, 0x8e, 0x7b, 0xfb, 0x28, 0x64, 0x5e, 0xa1, 0x26, 0x77,
	0x30, 0x5c, 0x96, 0x32, 0x27, 0xe7, 0xcc, 0xab, 0x9f, 0xed, 0xaf, 0xcc, 0xce, 0xfe, 0x50, 0xeb,
	0xb1, 0xfb, 0xfe, 0xed, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x81, 0x08, 0x82, 0x01, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResponseHandlerClient is the client API for ResponseHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResponseHandlerClient interface {
	Find(ctx context.Context, in *ResponseID, opts ...grpc.CallOption) (*Response, error)
}

type responseHandlerClient struct {
	cc *grpc.ClientConn
}

func NewResponseHandlerClient(cc *grpc.ClientConn) ResponseHandlerClient {
	return &responseHandlerClient{cc}
}

func (c *responseHandlerClient) Find(ctx context.Context, in *ResponseID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/response_grpc.ResponseHandler/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResponseHandlerServer is the server API for ResponseHandler service.
type ResponseHandlerServer interface {
	Find(context.Context, *ResponseID) (*Response, error)
}

// UnimplementedResponseHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedResponseHandlerServer struct {
}

func (*UnimplementedResponseHandlerServer) Find(ctx context.Context, req *ResponseID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterResponseHandlerServer(s *grpc.Server, srv ResponseHandlerServer) {
	s.RegisterService(&_ResponseHandler_serviceDesc, srv)
}

func _ResponseHandler_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResponseID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResponseHandlerServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/response_grpc.ResponseHandler/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResponseHandlerServer).Find(ctx, req.(*ResponseID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResponseHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "response_grpc.ResponseHandler",
	HandlerType: (*ResponseHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _ResponseHandler_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/user-response/delivery/grpc/response_grpc/response.proto",
}
