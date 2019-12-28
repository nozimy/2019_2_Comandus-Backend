// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/chat_app/chat/delivery/grpc/chat_grpc/chat.proto

package chat_grpc

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

type Nothing struct {
	Dummy                bool     `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbf4e45455396621, []int{0}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func (m *Nothing) GetDummy() bool {
	if m != nil {
		return m.Dummy
	}
	return false
}

type Chat struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	User                 int64    `protobuf:"varint,2,opt,name=User,proto3" json:"User,omitempty"`
	Support              int64    `protobuf:"varint,3,opt,name=Support,proto3" json:"Support,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Proposal             int64    `protobuf:"varint,5,opt,name=Proposal,proto3" json:"Proposal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chat) Reset()         { *m = Chat{} }
func (m *Chat) String() string { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()    {}
func (*Chat) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbf4e45455396621, []int{1}
}

func (m *Chat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chat.Unmarshal(m, b)
}
func (m *Chat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chat.Marshal(b, m, deterministic)
}
func (m *Chat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chat.Merge(m, src)
}
func (m *Chat) XXX_Size() int {
	return xxx_messageInfo_Chat.Size(m)
}
func (m *Chat) XXX_DiscardUnknown() {
	xxx_messageInfo_Chat.DiscardUnknown(m)
}

var xxx_messageInfo_Chat proto.InternalMessageInfo

func (m *Chat) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Chat) GetUser() int64 {
	if m != nil {
		return m.User
	}
	return 0
}

func (m *Chat) GetSupport() int64 {
	if m != nil {
		return m.Support
	}
	return 0
}

func (m *Chat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chat) GetProposal() int64 {
	if m != nil {
		return m.Proposal
	}
	return 0
}

func init() {
	proto.RegisterType((*Nothing)(nil), "chat_grpc.Nothing")
	proto.RegisterType((*Chat)(nil), "chat_grpc.Chat")
}

func init() {
	proto.RegisterFile("internal/chat_app/chat/delivery/grpc/chat_grpc/chat.proto", fileDescriptor_fbf4e45455396621)
}

var fileDescriptor_fbf4e45455396621 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xbf, 0x6b, 0xc3, 0x40,
	0x0c, 0x85, 0xb1, 0xe3, 0xfc, 0x52, 0xa1, 0x05, 0xd1, 0xe1, 0xc8, 0xd2, 0x90, 0x29, 0x74, 0xb0,
	0xa1, 0x9d, 0x32, 0x27, 0x43, 0xb3, 0x84, 0xe2, 0xd2, 0xb9, 0xa8, 0xb1, 0x88, 0x0d, 0xb6, 0xef,
	0x50, 0x94, 0x42, 0xfe, 0xfb, 0x62, 0x85, 0xde, 0xd0, 0xed, 0x7b, 0xdf, 0x7b, 0x70, 0x27, 0xd8,
	0x34, 0xbd, 0xb2, 0xf4, 0xd4, 0x16, 0xc7, 0x9a, 0xf4, 0x8b, 0x42, 0x30, 0x28, 0x2a, 0x6e, 0x9b,
	0x1f, 0x96, 0x6b, 0x71, 0x92, 0x70, 0xbc, 0x75, 0x91, 0xf2, 0x20, 0x5e, 0x3d, 0xce, 0xa3, 0x5d,
	0x3d, 0xc1, 0xf4, 0xe0, 0xb5, 0x6e, 0xfa, 0x13, 0x3e, 0xc2, 0xb8, 0xba, 0x74, 0xdd, 0xd5, 0x25,
	0xcb, 0x64, 0x3d, 0x2b, 0x6f, 0x61, 0xa5, 0x90, 0x6d, 0x6b, 0x52, 0xbc, 0x87, 0x74, 0xbf, 0xb3,
	0x6a, 0x54, 0xa6, 0xfb, 0x1d, 0x22, 0x64, 0x9f, 0x67, 0x16, 0x97, 0x9a, 0x31, 0x46, 0x07, 0xd3,
	0x8f, 0x4b, 0x08, 0x5e, 0xd4, 0x8d, 0x4c, 0xff, 0xc5, 0x61, 0x7d, 0xa0, 0x8e, 0x5d, 0xb6, 0x4c,
	0xd6, 0xf3, 0xd2, 0x18, 0x17, 0x30, 0x7b, 0x17, 0x1f, 0xfc, 0x99, 0x5a, 0x37, 0xb6, 0x79, 0xcc,
	0x2f, 0x1b, 0xb8, 0x1b, 0x5e, 0x7d, 0xa3, 0xbe, 0x6a, 0x59, 0xf0, 0x19, 0x26, 0x5b, 0x61, 0x52,
	0xc6, 0x87, 0x3c, 0xfe, 0x3d, 0x1f, 0x16, 0x8b, 0xff, 0xe2, 0x7b, 0x62, 0x37, 0xbe, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x2f, 0xf7, 0x09, 0x88, 0x20, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatHandlerClient is the client API for ChatHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatHandlerClient interface {
	Create(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Chat, error)
}

type chatHandlerClient struct {
	cc *grpc.ClientConn
}

func NewChatHandlerClient(cc *grpc.ClientConn) ChatHandlerClient {
	return &chatHandlerClient{cc}
}

func (c *chatHandlerClient) Create(ctx context.Context, in *Chat, opts ...grpc.CallOption) (*Chat, error) {
	out := new(Chat)
	err := c.cc.Invoke(ctx, "/chat_grpc.ChatHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatHandlerServer is the server API for ChatHandler service.
type ChatHandlerServer interface {
	Create(context.Context, *Chat) (*Chat, error)
}

// UnimplementedChatHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedChatHandlerServer struct {
}

func (*UnimplementedChatHandlerServer) Create(ctx context.Context, req *Chat) (*Chat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterChatHandlerServer(s *grpc.Server, srv ChatHandlerServer) {
	s.RegisterService(&_ChatHandler_serviceDesc, srv)
}

func _ChatHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_grpc.ChatHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatHandlerServer).Create(ctx, req.(*Chat))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat_grpc.ChatHandler",
	HandlerType: (*ChatHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChatHandler_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/chat_app/chat/delivery/grpc/chat_grpc/chat.proto",
}