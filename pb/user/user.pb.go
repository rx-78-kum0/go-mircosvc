// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usersvc.proto

package pb

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

type GetUserRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetUserReply struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReply) Reset()         { *m = GetUserReply{} }
func (m *GetUserReply) String() string { return proto.CompactTextString(m) }
func (*GetUserReply) ProtoMessage()    {}
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReply.Unmarshal(m, b)
}
func (m *GetUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReply.Marshal(b, m, deterministic)
}
func (m *GetUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReply.Merge(m, src)
}
func (m *GetUserReply) XXX_Size() int {
	return xxx_messageInfo_GetUserReply.Size(m)
}
func (m *GetUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReply proto.InternalMessageInfo

func (m *GetUserReply) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Rs                   string   `protobuf:"bytes,1,opt,name=rs,proto3" json:"rs,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetRs() string {
	if m != nil {
		return m.Rs
	}
	return ""
}

func (m *LoginReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "pb.GetUserRequest")
	proto.RegisterType((*GetUserReply)(nil), "pb.GetUserReply")
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "pb.LoginReply")
}

func init() { proto.RegisterFile("usersvc.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0xca, 0x82, 0x40,
	0x10, 0xc7, 0x3f, 0xf7, 0xa3, 0xac, 0x41, 0x44, 0xf6, 0x24, 0x9e, 0x64, 0x4f, 0x41, 0xb0, 0x50,
	0xbd, 0x43, 0x5d, 0x3a, 0x19, 0x3d, 0x80, 0xe6, 0x12, 0x82, 0xb9, 0xdb, 0xac, 0x16, 0xbe, 0x7d,
	0xec, 0xb6, 0x8a, 0x42, 0xb7, 0x99, 0xf9, 0xff, 0x67, 0x7e, 0x33, 0x03, 0xd0, 0x69, 0x81, 0x5c,
	0xa1, 0x6c, 0x25, 0x25, 0xaa, 0x60, 0x0c, 0xc2, 0x93, 0x68, 0xaf, 0x5a, 0x60, 0x26, 0x9e, 0x9d,
	0xd0, 0x2d, 0x8d, 0xe0, 0xbf, 0xab, 0xca, 0xd8, 0x4b, 0xbd, 0xcd, 0x3a, 0x33, 0x21, 0x4b, 0x21,
	0x18, 0x3d, 0xaa, 0xee, 0x7f, 0x38, 0x8e, 0x10, 0x9c, 0xe5, 0xbd, 0x6a, 0x86, 0x19, 0x09, 0xac,
	0x0c, 0xa7, 0xc9, 0x1f, 0xc2, 0xd9, 0xc6, 0xdc, 0x68, 0x2a, 0xd7, 0xfa, 0x2d, 0xb1, 0x8c, 0xc9,
	0x57, 0x1b, 0x72, 0xc6, 0x01, 0xdc, 0x1c, 0xc3, 0x09, 0x81, 0xa0, 0x76, 0xfd, 0x04, 0xb5, 0xe1,
	0x0a, 0x44, 0xd7, 0x64, 0xc2, 0x7d, 0x05, 0xbe, 0x59, 0xeb, 0xf2, 0xba, 0xd1, 0x1d, 0xf8, 0x6e,
	0x49, 0x4a, 0xb9, 0x2a, 0xf8, 0xfc, 0xaa, 0x24, 0x9a, 0xd5, 0x54, 0xdd, 0xb3, 0x3f, 0xba, 0x85,
	0x85, 0xa5, 0x51, 0x2b, 0x4e, 0x0f, 0x48, 0xc2, 0x49, 0xc5, 0x9a, 0x8b, 0xa5, 0xfd, 0xd9, 0xe1,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xb9, 0x15, 0xad, 0x41, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserSvcClient is the client API for UserSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSvcClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type userSvcClient struct {
	cc *grpc.ClientConn
}

func NewUserSvcClient(cc *grpc.ClientConn) UserSvcClient {
	return &userSvcClient{cc}
}

func (c *userSvcClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/pb.UserSvc/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/pb.UserSvc/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSvcServer is the server API for UserSvc service.
type UserSvcServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

// UnimplementedUserSvcServer can be embedded to have forward compatible implementations.
type UnimplementedUserSvcServer struct {
}

func (*UnimplementedUserSvcServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserSvcServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterUserSvcServer(s *grpc.Server, srv UserSvcServer) {
	s.RegisterService(&_UserSvc_serviceDesc, srv)
}

func _UserSvc_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserSvc/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserSvc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserSvc",
	HandlerType: (*UserSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserSvc_GetUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserSvc_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usersvc.proto",
}