// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/company/delivery/grpc/company_grpc/company.proto

package company_grpc

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

type CompanyWithUser struct {
	MyCompany            *Company `protobuf:"bytes,1,opt,name=MyCompany,proto3" json:"MyCompany,omitempty"`
	UserID               int64    `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompanyWithUser) Reset()         { *m = CompanyWithUser{} }
func (m *CompanyWithUser) String() string { return proto.CompactTextString(m) }
func (*CompanyWithUser) ProtoMessage()    {}
func (*CompanyWithUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_5755ad00c494f094, []int{0}
}

func (m *CompanyWithUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyWithUser.Unmarshal(m, b)
}
func (m *CompanyWithUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyWithUser.Marshal(b, m, deterministic)
}
func (m *CompanyWithUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyWithUser.Merge(m, src)
}
func (m *CompanyWithUser) XXX_Size() int {
	return xxx_messageInfo_CompanyWithUser.Size(m)
}
func (m *CompanyWithUser) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyWithUser.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyWithUser proto.InternalMessageInfo

func (m *CompanyWithUser) GetMyCompany() *Company {
	if m != nil {
		return m.MyCompany
	}
	return nil
}

func (m *CompanyWithUser) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

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
	return fileDescriptor_5755ad00c494f094, []int{1}
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

type Company struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CompanyName          string   `protobuf:"bytes,2,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	Site                 string   `protobuf:"bytes,3,opt,name=Site,proto3" json:"Site,omitempty"`
	TagLine              string   `protobuf:"bytes,4,opt,name=TagLine,proto3" json:"TagLine,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Country              string   `protobuf:"bytes,6,opt,name=Country,proto3" json:"Country,omitempty"`
	City                 string   `protobuf:"bytes,7,opt,name=City,proto3" json:"City,omitempty"`
	Address              string   `protobuf:"bytes,8,opt,name=Address,proto3" json:"Address,omitempty"`
	Phone                string   `protobuf:"bytes,9,opt,name=Phone,proto3" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Company) Reset()         { *m = Company{} }
func (m *Company) String() string { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()    {}
func (*Company) Descriptor() ([]byte, []int) {
	return fileDescriptor_5755ad00c494f094, []int{2}
}

func (m *Company) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company.Unmarshal(m, b)
}
func (m *Company) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company.Marshal(b, m, deterministic)
}
func (m *Company) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company.Merge(m, src)
}
func (m *Company) XXX_Size() int {
	return xxx_messageInfo_Company.Size(m)
}
func (m *Company) XXX_DiscardUnknown() {
	xxx_messageInfo_Company.DiscardUnknown(m)
}

var xxx_messageInfo_Company proto.InternalMessageInfo

func (m *Company) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Company) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *Company) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *Company) GetTagLine() string {
	if m != nil {
		return m.TagLine
	}
	return ""
}

func (m *Company) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Company) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Company) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Company) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Company) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type CompanyID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompanyID) Reset()         { *m = CompanyID{} }
func (m *CompanyID) String() string { return proto.CompactTextString(m) }
func (*CompanyID) ProtoMessage()    {}
func (*CompanyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_5755ad00c494f094, []int{3}
}

func (m *CompanyID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyID.Unmarshal(m, b)
}
func (m *CompanyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyID.Marshal(b, m, deterministic)
}
func (m *CompanyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyID.Merge(m, src)
}
func (m *CompanyID) XXX_Size() int {
	return xxx_messageInfo_CompanyID.Size(m)
}
func (m *CompanyID) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyID.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyID proto.InternalMessageInfo

func (m *CompanyID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type UserID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_5755ad00c494f094, []int{4}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*CompanyWithUser)(nil), "company_grpc.CompanyWithUser")
	proto.RegisterType((*Nothing)(nil), "company_grpc.Nothing")
	proto.RegisterType((*Company)(nil), "company_grpc.Company")
	proto.RegisterType((*CompanyID)(nil), "company_grpc.CompanyID")
	proto.RegisterType((*UserID)(nil), "company_grpc.UserID")
}

func init() {
	proto.RegisterFile("internal/app/company/delivery/grpc/company_grpc/company.proto", fileDescriptor_5755ad00c494f094)
}

var fileDescriptor_5755ad00c494f094 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5f, 0x4b, 0xfb, 0x30,
	0x14, 0xa5, 0x5b, 0xb7, 0xae, 0x77, 0xbf, 0xdf, 0x84, 0x30, 0x35, 0x28, 0xe2, 0xe8, 0xd3, 0x9e,
	0x56, 0xd8, 0xc0, 0xa7, 0xf9, 0x20, 0xad, 0x62, 0x41, 0x87, 0x54, 0xc5, 0x37, 0xa5, 0xae, 0x61,
	0x0b, 0xac, 0x69, 0x49, 0x33, 0xa1, 0xdf, 0xcf, 0x2f, 0xe3, 0xb7, 0x90, 0xfc, 0x29, 0xd6, 0x31,
	0xdf, 0x72, 0xce, 0xb9, 0xe7, 0x24, 0xb9, 0xf7, 0xc2, 0x25, 0x65, 0x82, 0x70, 0x96, 0x6c, 0xfc,
	0xa4, 0x28, 0xfc, 0x65, 0x9e, 0x15, 0x09, 0xab, 0xfc, 0x94, 0x6c, 0xe8, 0x07, 0xe1, 0x95, 0xbf,
	0xe2, 0xc5, 0xb2, 0x66, 0xdf, 0x9a, 0x60, 0x52, 0xf0, 0x5c, 0xe4, 0xe8, 0x5f, 0x53, 0xf3, 0x5e,
	0xe1, 0x20, 0xd0, 0xf8, 0x85, 0x8a, 0xf5, 0x73, 0x49, 0x38, 0x9a, 0x81, 0x7b, 0x5f, 0x19, 0x12,
	0x5b, 0x23, 0x6b, 0xdc, 0x9f, 0x1e, 0x4e, 0x9a, 0xa6, 0x89, 0x11, 0xe3, 0x9f, 0x3a, 0x74, 0x04,
	0x5d, 0x69, 0x8e, 0x42, 0xdc, 0x1a, 0x59, 0xe3, 0x76, 0x6c, 0x90, 0x77, 0x0e, 0xce, 0x22, 0x17,
	0x6b, 0xca, 0x56, 0x68, 0x08, 0x9d, 0x74, 0x9b, 0x65, 0x3a, 0xb3, 0x17, 0x6b, 0xe0, 0x7d, 0x59,
	0xe0, 0xd4, 0x21, 0x03, 0x68, 0x45, 0xa1, 0x92, 0xdb, 0x71, 0x2b, 0x0a, 0xd1, 0x08, 0xfa, 0x46,
	0x5a, 0x24, 0x19, 0x51, 0xc9, 0x6e, 0xdc, 0xa4, 0x10, 0x02, 0xfb, 0x91, 0x0a, 0x82, 0xdb, 0x4a,
	0x52, 0x67, 0x84, 0xc1, 0x79, 0x4a, 0x56, 0x77, 0x94, 0x11, 0x6c, 0x2b, 0xba, 0x86, 0x32, 0x2f,
	0x24, 0xe5, 0x92, 0xd3, 0x42, 0xd0, 0x9c, 0xe1, 0x8e, 0xce, 0x6b, 0x50, 0xd2, 0x1b, 0xe4, 0x5b,
	0x26, 0x78, 0x85, 0xbb, 0xda, 0x6b, 0xa0, 0xbc, 0x29, 0xa0, 0xa2, 0xc2, 0x8e, 0xbe, 0x49, 0x9e,
	0x65, 0xf5, 0x55, 0x9a, 0x72, 0x52, 0x96, 0xb8, 0xa7, 0xab, 0x0d, 0x94, 0x7f, 0x7d, 0x58, 0xe7,
	0x8c, 0x60, 0x57, 0xf1, 0x1a, 0x78, 0xa7, 0xe0, 0x9a, 0xc7, 0x47, 0xe1, 0xee, 0x67, 0x3d, 0x5c,
	0x77, 0x70, 0x57, 0x99, 0x7e, 0x5a, 0x30, 0x30, 0xbe, 0xdb, 0x84, 0xa5, 0x1b, 0xc2, 0xd1, 0x1c,
	0xfe, 0x07, 0x9c, 0x24, 0x82, 0xd4, 0xad, 0x1b, 0xfe, 0x9e, 0x90, 0x4e, 0x3a, 0xd9, 0x3f, 0x37,
	0x74, 0x01, 0xf6, 0x0d, 0x65, 0x29, 0x3a, 0xde, 0x2b, 0xff, 0xed, 0x9b, 0x83, 0x7d, 0x9d, 0x52,
	0x81, 0xce, 0xf6, 0xca, 0xf5, 0x02, 0xed, 0xba, 0xcd, 0xfc, 0xdf, 0xbb, 0x6a, 0xff, 0x66, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0xf9, 0xb3, 0x96, 0xc0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompanyHandlerClient is the client API for CompanyHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyHandlerClient interface {
	CreateCompany(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Company, error)
	Find(ctx context.Context, in *CompanyID, opts ...grpc.CallOption) (*Company, error)
	Edit(ctx context.Context, in *CompanyWithUser, opts ...grpc.CallOption) (*Nothing, error)
}

type companyHandlerClient struct {
	cc *grpc.ClientConn
}

func NewCompanyHandlerClient(cc *grpc.ClientConn) CompanyHandlerClient {
	return &companyHandlerClient{cc}
}

func (c *companyHandlerClient) CreateCompany(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/company_grpc.CompanyHandler/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyHandlerClient) Find(ctx context.Context, in *CompanyID, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/company_grpc.CompanyHandler/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyHandlerClient) Edit(ctx context.Context, in *CompanyWithUser, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/company_grpc.CompanyHandler/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyHandlerServer is the server API for CompanyHandler service.
type CompanyHandlerServer interface {
	CreateCompany(context.Context, *UserID) (*Company, error)
	Find(context.Context, *CompanyID) (*Company, error)
	Edit(context.Context, *CompanyWithUser) (*Nothing, error)
}

// UnimplementedCompanyHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedCompanyHandlerServer struct {
}

func (*UnimplementedCompanyHandlerServer) CreateCompany(ctx context.Context, req *UserID) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (*UnimplementedCompanyHandlerServer) Find(ctx context.Context, req *CompanyID) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (*UnimplementedCompanyHandlerServer) Edit(ctx context.Context, req *CompanyWithUser) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}

func RegisterCompanyHandlerServer(s *grpc.Server, srv CompanyHandlerServer) {
	s.RegisterService(&_CompanyHandler_serviceDesc, srv)
}

func _CompanyHandler_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyHandlerServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company_grpc.CompanyHandler/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyHandlerServer).CreateCompany(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyHandler_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyHandlerServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company_grpc.CompanyHandler/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyHandlerServer).Find(ctx, req.(*CompanyID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyHandler_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyWithUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyHandlerServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company_grpc.CompanyHandler/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyHandlerServer).Edit(ctx, req.(*CompanyWithUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompanyHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "company_grpc.CompanyHandler",
	HandlerType: (*CompanyHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyHandler_CreateCompany_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _CompanyHandler_Find_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _CompanyHandler_Edit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/company/delivery/grpc/company_grpc/company.proto",
}
