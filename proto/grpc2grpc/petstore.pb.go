// Code generated by protoc-gen-go. DO NOT EDIT.
// source: petstore.proto

package grpc2grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Pet struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pet) Reset()         { *m = Pet{} }
func (m *Pet) String() string { return proto.CompactTextString(m) }
func (*Pet) ProtoMessage()    {}
func (*Pet) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2ab5c8ed41385256, []int{0}
}
func (m *Pet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pet.Unmarshal(m, b)
}
func (m *Pet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pet.Marshal(b, m, deterministic)
}
func (dst *Pet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pet.Merge(dst, src)
}
func (m *Pet) XXX_Size() int {
	return xxx_messageInfo_Pet.Size(m)
}
func (m *Pet) XXX_DiscardUnknown() {
	xxx_messageInfo_Pet.DiscardUnknown(m)
}

var xxx_messageInfo_Pet proto.InternalMessageInfo

func (m *Pet) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2ab5c8ed41385256, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type PetByIdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PetByIdRequest) Reset()         { *m = PetByIdRequest{} }
func (m *PetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*PetByIdRequest) ProtoMessage()    {}
func (*PetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2ab5c8ed41385256, []int{2}
}
func (m *PetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PetByIdRequest.Unmarshal(m, b)
}
func (m *PetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PetByIdRequest.Marshal(b, m, deterministic)
}
func (dst *PetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PetByIdRequest.Merge(dst, src)
}
func (m *PetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_PetByIdRequest.Size(m)
}
func (m *PetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PetByIdRequest proto.InternalMessageInfo

func (m *PetByIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserByNameRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserByNameRequest) Reset()         { *m = UserByNameRequest{} }
func (m *UserByNameRequest) String() string { return proto.CompactTextString(m) }
func (*UserByNameRequest) ProtoMessage()    {}
func (*UserByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2ab5c8ed41385256, []int{3}
}
func (m *UserByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserByNameRequest.Unmarshal(m, b)
}
func (m *UserByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserByNameRequest.Marshal(b, m, deterministic)
}
func (dst *UserByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserByNameRequest.Merge(dst, src)
}
func (m *UserByNameRequest) XXX_Size() int {
	return xxx_messageInfo_UserByNameRequest.Size(m)
}
func (m *UserByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserByNameRequest proto.InternalMessageInfo

func (m *UserByNameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type PetResponse struct {
	Pet                  *Pet     `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PetResponse) Reset()         { *m = PetResponse{} }
func (m *PetResponse) String() string { return proto.CompactTextString(m) }
func (*PetResponse) ProtoMessage()    {}
func (*PetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2ab5c8ed41385256, []int{4}
}
func (m *PetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PetResponse.Unmarshal(m, b)
}
func (m *PetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PetResponse.Marshal(b, m, deterministic)
}
func (dst *PetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PetResponse.Merge(dst, src)
}
func (m *PetResponse) XXX_Size() int {
	return xxx_messageInfo_PetResponse.Size(m)
}
func (m *PetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PetResponse proto.InternalMessageInfo

func (m *PetResponse) GetPet() *Pet {
	if m != nil {
		return m.Pet
	}
	return nil
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2ab5c8ed41385256, []int{5}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type EmptyReq struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyReq) Reset()         { *m = EmptyReq{} }
func (m *EmptyReq) String() string { return proto.CompactTextString(m) }
func (*EmptyReq) ProtoMessage()    {}
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2ab5c8ed41385256, []int{6}
}
func (m *EmptyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyReq.Unmarshal(m, b)
}
func (m *EmptyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyReq.Marshal(b, m, deterministic)
}
func (dst *EmptyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyReq.Merge(dst, src)
}
func (m *EmptyReq) XXX_Size() int {
	return xxx_messageInfo_EmptyReq.Size(m)
}
func (m *EmptyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyReq.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyReq proto.InternalMessageInfo

func (m *EmptyReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EmptyRes struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRes) Reset()         { *m = EmptyRes{} }
func (m *EmptyRes) String() string { return proto.CompactTextString(m) }
func (*EmptyRes) ProtoMessage()    {}
func (*EmptyRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_2ab5c8ed41385256, []int{7}
}
func (m *EmptyRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRes.Unmarshal(m, b)
}
func (m *EmptyRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRes.Marshal(b, m, deterministic)
}
func (dst *EmptyRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRes.Merge(dst, src)
}
func (m *EmptyRes) XXX_Size() int {
	return xxx_messageInfo_EmptyRes.Size(m)
}
func (m *EmptyRes) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRes.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRes proto.InternalMessageInfo

func (m *EmptyRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Pet)(nil), "grpc2grpc.Pet")
	proto.RegisterType((*User)(nil), "grpc2grpc.User")
	proto.RegisterType((*PetByIdRequest)(nil), "grpc2grpc.PetByIdRequest")
	proto.RegisterType((*UserByNameRequest)(nil), "grpc2grpc.UserByNameRequest")
	proto.RegisterType((*PetResponse)(nil), "grpc2grpc.PetResponse")
	proto.RegisterType((*UserResponse)(nil), "grpc2grpc.UserResponse")
	proto.RegisterType((*EmptyReq)(nil), "grpc2grpc.EmptyReq")
	proto.RegisterType((*EmptyRes)(nil), "grpc2grpc.EmptyRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PetStoreServiceClient is the client API for PetStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PetStoreServiceClient interface {
	PetById(ctx context.Context, in *PetByIdRequest, opts ...grpc.CallOption) (*PetResponse, error)
	UserByName(ctx context.Context, in *UserByNameRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ListUsers(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (PetStoreService_ListUsersClient, error)
	StoreUsers(ctx context.Context, opts ...grpc.CallOption) (PetStoreService_StoreUsersClient, error)
	BulkUsers(ctx context.Context, opts ...grpc.CallOption) (PetStoreService_BulkUsersClient, error)
}

type petStoreServiceClient struct {
	cc *grpc.ClientConn
}

func NewPetStoreServiceClient(cc *grpc.ClientConn) PetStoreServiceClient {
	return &petStoreServiceClient{cc}
}

func (c *petStoreServiceClient) PetById(ctx context.Context, in *PetByIdRequest, opts ...grpc.CallOption) (*PetResponse, error) {
	out := new(PetResponse)
	err := c.cc.Invoke(ctx, "/grpc2grpc.PetStoreService/PetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) UserByName(ctx context.Context, in *UserByNameRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpc2grpc.PetStoreService/UserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) ListUsers(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (PetStoreService_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PetStoreService_serviceDesc.Streams[0], "/grpc2grpc.PetStoreService/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &petStoreServiceListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PetStoreService_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type petStoreServiceListUsersClient struct {
	grpc.ClientStream
}

func (x *petStoreServiceListUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *petStoreServiceClient) StoreUsers(ctx context.Context, opts ...grpc.CallOption) (PetStoreService_StoreUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PetStoreService_serviceDesc.Streams[1], "/grpc2grpc.PetStoreService/StoreUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &petStoreServiceStoreUsersClient{stream}
	return x, nil
}

type PetStoreService_StoreUsersClient interface {
	Send(*User) error
	CloseAndRecv() (*EmptyRes, error)
	grpc.ClientStream
}

type petStoreServiceStoreUsersClient struct {
	grpc.ClientStream
}

func (x *petStoreServiceStoreUsersClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *petStoreServiceStoreUsersClient) CloseAndRecv() (*EmptyRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *petStoreServiceClient) BulkUsers(ctx context.Context, opts ...grpc.CallOption) (PetStoreService_BulkUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PetStoreService_serviceDesc.Streams[2], "/grpc2grpc.PetStoreService/BulkUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &petStoreServiceBulkUsersClient{stream}
	return x, nil
}

type PetStoreService_BulkUsersClient interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ClientStream
}

type petStoreServiceBulkUsersClient struct {
	grpc.ClientStream
}

func (x *petStoreServiceBulkUsersClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *petStoreServiceBulkUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PetStoreServiceServer is the server API for PetStoreService service.
type PetStoreServiceServer interface {
	PetById(context.Context, *PetByIdRequest) (*PetResponse, error)
	UserByName(context.Context, *UserByNameRequest) (*UserResponse, error)
	ListUsers(*EmptyReq, PetStoreService_ListUsersServer) error
	StoreUsers(PetStoreService_StoreUsersServer) error
	BulkUsers(PetStoreService_BulkUsersServer) error
}

func RegisterPetStoreServiceServer(s *grpc.Server, srv PetStoreServiceServer) {
	s.RegisterService(&_PetStoreService_serviceDesc, srv)
}

func _PetStoreService_PetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).PetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc2grpc.PetStoreService/PetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).PetById(ctx, req.(*PetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStoreService_UserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).UserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc2grpc.PetStoreService/UserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).UserByName(ctx, req.(*UserByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStoreService_ListUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PetStoreServiceServer).ListUsers(m, &petStoreServiceListUsersServer{stream})
}

type PetStoreService_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type petStoreServiceListUsersServer struct {
	grpc.ServerStream
}

func (x *petStoreServiceListUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _PetStoreService_StoreUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PetStoreServiceServer).StoreUsers(&petStoreServiceStoreUsersServer{stream})
}

type PetStoreService_StoreUsersServer interface {
	SendAndClose(*EmptyRes) error
	Recv() (*User, error)
	grpc.ServerStream
}

type petStoreServiceStoreUsersServer struct {
	grpc.ServerStream
}

func (x *petStoreServiceStoreUsersServer) SendAndClose(m *EmptyRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *petStoreServiceStoreUsersServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PetStoreService_BulkUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PetStoreServiceServer).BulkUsers(&petStoreServiceBulkUsersServer{stream})
}

type PetStoreService_BulkUsersServer interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type petStoreServiceBulkUsersServer struct {
	grpc.ServerStream
}

func (x *petStoreServiceBulkUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *petStoreServiceBulkUsersServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PetStoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc2grpc.PetStoreService",
	HandlerType: (*PetStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PetById",
			Handler:    _PetStoreService_PetById_Handler,
		},
		{
			MethodName: "UserByName",
			Handler:    _PetStoreService_UserByName_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsers",
			Handler:       _PetStoreService_ListUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StoreUsers",
			Handler:       _PetStoreService_StoreUsers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BulkUsers",
			Handler:       _PetStoreService_BulkUsers_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "petstore.proto",
}

func init() { proto.RegisterFile("petstore.proto", fileDescriptor_petstore_2ab5c8ed41385256) }

var fileDescriptor_petstore_2ab5c8ed41385256 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x4f, 0xea, 0x40,
	0x10, 0xce, 0xb6, 0xe5, 0x3d, 0x3a, 0xbc, 0x94, 0xf7, 0xe6, 0x19, 0xad, 0x0d, 0x87, 0x66, 0xbd,
	0xd4, 0x0b, 0x20, 0x78, 0xf4, 0x84, 0xf1, 0x60, 0x62, 0x4c, 0x53, 0xe2, 0xd5, 0x04, 0x61, 0x82,
	0x8d, 0x94, 0x2e, 0xdd, 0xc5, 0x84, 0xbf, 0xe5, 0x2f, 0x34, 0xbb, 0x50, 0x2c, 0x45, 0x2f, 0x64,
	0x67, 0xbe, 0xef, 0x9b, 0xf9, 0xe6, 0xa3, 0xe0, 0x09, 0x52, 0x52, 0xe5, 0x05, 0x75, 0x45, 0x91,
	0xab, 0x1c, 0xdd, 0x79, 0x21, 0xa6, 0x03, 0xfd, 0xc3, 0x2f, 0xc1, 0x8e, 0x49, 0xa1, 0x07, 0x56,
	0x3a, 0xf3, 0x59, 0xc8, 0xa2, 0x46, 0x62, 0xa5, 0x33, 0x44, 0x70, 0x96, 0x93, 0x8c, 0x7c, 0x2b,
	0x64, 0x91, 0x9b, 0x98, 0x37, 0x7f, 0x06, 0xe7, 0x49, 0x52, 0x71, 0xc4, 0x0d, 0xa0, 0xb9, 0x96,
	0x54, 0x54, 0xf8, 0xfb, 0x1a, 0x4f, 0xa0, 0x41, 0xd9, 0x24, 0x5d, 0xf8, 0xb6, 0x01, 0xb6, 0x85,
	0xee, 0x8a, 0xd7, 0x7c, 0x49, 0xbe, 0xb3, 0xed, 0x9a, 0x82, 0x87, 0xe0, 0xc5, 0xa4, 0x46, 0x9b,
	0xfb, 0x59, 0x42, 0xab, 0x35, 0xc9, 0x23, 0x57, 0xbc, 0x07, 0xff, 0xb4, 0x83, 0xd1, 0xe6, 0x71,
	0x92, 0x51, 0x49, 0xaa, 0xae, 0x67, 0x87, 0xeb, 0x79, 0x0f, 0x5a, 0x31, 0xa9, 0x84, 0xa4, 0xc8,
	0x97, 0x92, 0x30, 0x04, 0x5b, 0x90, 0x32, 0xac, 0xd6, 0xc0, 0xeb, 0xee, 0x53, 0xe8, 0x6a, 0x92,
	0x86, 0xf8, 0x10, 0xfe, 0xe8, 0x0d, 0x7b, 0xc5, 0x05, 0x38, 0x7a, 0xd8, 0x4e, 0xd2, 0xae, 0x48,
	0x0c, 0xcd, 0x80, 0xbc, 0x03, 0xcd, 0xbb, 0x4c, 0xa8, 0x4d, 0x42, 0x2b, 0xfc, 0x0b, 0x76, 0x26,
	0xe7, 0x3b, 0x23, 0xfa, 0x59, 0x41, 0xe5, 0x31, 0x3a, 0xf8, 0xb0, 0xa0, 0x1d, 0x93, 0x1a, 0xeb,
	0x7f, 0x67, 0x4c, 0xc5, 0x7b, 0x3a, 0x25, 0xbc, 0x81, 0xdf, 0xbb, 0x20, 0xf0, 0xfc, 0xd0, 0x64,
	0x25, 0x9c, 0xe0, 0xb4, 0xe6, 0xbf, 0xb4, 0x7c, 0x0b, 0xf0, 0x15, 0x12, 0x76, 0x6a, 0x96, 0x0f,
	0xb2, 0x0b, 0xce, 0xea, 0x07, 0x95, 0x43, 0x86, 0xe0, 0x3e, 0xa4, 0x52, 0xe9, 0x9e, 0xc4, 0xff,
	0x15, 0x56, 0x79, 0x68, 0x50, 0xcf, 0xa2, 0xcf, 0xf0, 0x1a, 0xc0, 0xdc, 0xb1, 0x55, 0xd5, 0x09,
	0xc1, 0x37, 0x63, 0x64, 0xc4, 0xf0, 0x0a, 0xdc, 0xd1, 0x7a, 0xf1, 0xf6, 0x83, 0xa8, 0xde, 0x88,
	0x58, 0x9f, 0xbd, 0xfc, 0x32, 0x9f, 0xf1, 0xf0, 0x33, 0x00, 0x00, 0xff, 0xff, 0xab, 0x5a, 0xe5,
	0xe6, 0xd8, 0x02, 0x00, 0x00,
}
