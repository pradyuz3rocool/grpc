// Code generated by protoc-gen-go. DO NOT EDIT.
// source: petstore.proto

package grpc2rest

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
	return fileDescriptor_petstore_4bf5cd587f28fc9d, []int{0}
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
	return fileDescriptor_petstore_4bf5cd587f28fc9d, []int{1}
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
	return fileDescriptor_petstore_4bf5cd587f28fc9d, []int{2}
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
	return fileDescriptor_petstore_4bf5cd587f28fc9d, []int{3}
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
	return fileDescriptor_petstore_4bf5cd587f28fc9d, []int{4}
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
	return fileDescriptor_petstore_4bf5cd587f28fc9d, []int{5}
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

type PetRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PetRequest) Reset()         { *m = PetRequest{} }
func (m *PetRequest) String() string { return proto.CompactTextString(m) }
func (*PetRequest) ProtoMessage()    {}
func (*PetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4bf5cd587f28fc9d, []int{6}
}
func (m *PetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PetRequest.Unmarshal(m, b)
}
func (m *PetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PetRequest.Marshal(b, m, deterministic)
}
func (dst *PetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PetRequest.Merge(dst, src)
}
func (m *PetRequest) XXX_Size() int {
	return xxx_messageInfo_PetRequest.Size(m)
}
func (m *PetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PetRequest proto.InternalMessageInfo

func (m *PetRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_petstore_4bf5cd587f28fc9d, []int{7}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (dst *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(dst, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*Pet)(nil), "grpc2rest.Pet")
	proto.RegisterType((*User)(nil), "grpc2rest.User")
	proto.RegisterType((*PetByIdRequest)(nil), "grpc2rest.PetByIdRequest")
	proto.RegisterType((*UserByNameRequest)(nil), "grpc2rest.UserByNameRequest")
	proto.RegisterType((*PetResponse)(nil), "grpc2rest.PetResponse")
	proto.RegisterType((*UserResponse)(nil), "grpc2rest.UserResponse")
	proto.RegisterType((*PetRequest)(nil), "grpc2rest.PetRequest")
	proto.RegisterType((*UserRequest)(nil), "grpc2rest.UserRequest")
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
	PetPUT(ctx context.Context, in *PetRequest, opts ...grpc.CallOption) (*PetResponse, error)
	UserPUT(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type petStoreServiceClient struct {
	cc *grpc.ClientConn
}

func NewPetStoreServiceClient(cc *grpc.ClientConn) PetStoreServiceClient {
	return &petStoreServiceClient{cc}
}

func (c *petStoreServiceClient) PetById(ctx context.Context, in *PetByIdRequest, opts ...grpc.CallOption) (*PetResponse, error) {
	out := new(PetResponse)
	err := c.cc.Invoke(ctx, "/grpc2rest.PetStoreService/PetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) UserByName(ctx context.Context, in *UserByNameRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpc2rest.PetStoreService/UserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) PetPUT(ctx context.Context, in *PetRequest, opts ...grpc.CallOption) (*PetResponse, error) {
	out := new(PetResponse)
	err := c.cc.Invoke(ctx, "/grpc2rest.PetStoreService/PetPUT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petStoreServiceClient) UserPUT(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/grpc2rest.PetStoreService/UserPUT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetStoreServiceServer is the server API for PetStoreService service.
type PetStoreServiceServer interface {
	PetById(context.Context, *PetByIdRequest) (*PetResponse, error)
	UserByName(context.Context, *UserByNameRequest) (*UserResponse, error)
	PetPUT(context.Context, *PetRequest) (*PetResponse, error)
	UserPUT(context.Context, *UserRequest) (*UserResponse, error)
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
		FullMethod: "/grpc2rest.PetStoreService/PetById",
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
		FullMethod: "/grpc2rest.PetStoreService/UserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).UserByName(ctx, req.(*UserByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStoreService_PetPUT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).PetPUT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc2rest.PetStoreService/PetPUT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).PetPUT(ctx, req.(*PetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetStoreService_UserPUT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServiceServer).UserPUT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc2rest.PetStoreService/UserPUT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServiceServer).UserPUT(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PetStoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc2rest.PetStoreService",
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
		{
			MethodName: "PetPUT",
			Handler:    _PetStoreService_PetPUT_Handler,
		},
		{
			MethodName: "UserPUT",
			Handler:    _PetStoreService_UserPUT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petstore.proto",
}

func init() { proto.RegisterFile("petstore.proto", fileDescriptor_petstore_4bf5cd587f28fc9d) }

var fileDescriptor_petstore_4bf5cd587f28fc9d = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xd3, 0x52, 0xe0, 0x65, 0x78, 0x53, 0xe2, 0x46, 0xb1, 0x12, 0x0f, 0xcd, 0x7a, 0xc1,
	0x0b, 0x18, 0x38, 0x98, 0x18, 0x4f, 0x78, 0xf2, 0x62, 0x36, 0x45, 0xae, 0x26, 0x08, 0x13, 0x6d,
	0x62, 0xdb, 0x75, 0x77, 0x31, 0xe1, 0x23, 0xf8, 0xad, 0xcd, 0x6e, 0x69, 0xed, 0xbf, 0x70, 0xf2,
	0xd6, 0xd9, 0x79, 0x9e, 0x79, 0x26, 0xbf, 0x29, 0xb8, 0x1c, 0x95, 0x54, 0x89, 0xc0, 0x09, 0x17,
	0x89, 0x4a, 0x48, 0xef, 0x4d, 0xf0, 0xcd, 0x4c, 0xa0, 0x54, 0xf4, 0x1a, 0x5a, 0x0c, 0x15, 0x71,
	0xc1, 0x0e, 0xb7, 0x9e, 0xe5, 0x5b, 0xe3, 0x76, 0x60, 0x87, 0x5b, 0x42, 0xc0, 0x89, 0xd7, 0x11,
	0x7a, 0xb6, 0x6f, 0x8d, 0x7b, 0x81, 0xf9, 0xa6, 0x2f, 0xe0, 0xac, 0x24, 0x8a, 0x9a, 0x76, 0x04,
	0xff, 0x76, 0x12, 0x45, 0x41, 0x9f, 0xd7, 0xe4, 0x14, 0xda, 0x18, 0xad, 0xc3, 0x0f, 0xaf, 0x65,
	0x1a, 0x69, 0xa1, 0x5f, 0xf9, 0x7b, 0x12, 0xa3, 0xe7, 0xa4, 0xaf, 0xa6, 0xa0, 0x3e, 0xb8, 0x0c,
	0xd5, 0x62, 0xff, 0xb8, 0x0d, 0xf0, 0x73, 0x87, 0xb2, 0xb6, 0x15, 0x9d, 0xc2, 0x89, 0xde, 0x60,
	0xb1, 0x7f, 0x5a, 0x47, 0x98, 0x89, 0x8a, 0xf1, 0x56, 0x39, 0x9e, 0x4e, 0xa1, 0xcf, 0x50, 0x05,
	0x28, 0x79, 0x12, 0x4b, 0x24, 0x3e, 0xb4, 0x38, 0x2a, 0xa3, 0xea, 0xcf, 0xdc, 0x49, 0x4e, 0x61,
	0xa2, 0x45, 0xba, 0x45, 0xe7, 0xf0, 0x5f, 0x27, 0xe4, 0x8e, 0x2b, 0x70, 0xf4, 0xb0, 0x83, 0x65,
	0x50, 0xb0, 0x18, 0x99, 0x69, 0xd2, 0x1b, 0x00, 0x93, 0xd2, 0xb8, 0x74, 0x23, 0x4a, 0x84, 0x7e,
	0x1a, 0xd3, 0x6c, 0xf9, 0x23, 0xa2, 0xb3, 0x6f, 0x1b, 0x06, 0x0c, 0xd5, 0x52, 0x9f, 0x7e, 0x89,
	0xe2, 0x2b, 0xdc, 0x20, 0xb9, 0x87, 0xee, 0x81, 0x32, 0xb9, 0x28, 0x13, 0x28, 0x90, 0x1f, 0x0d,
	0x2b, 0x70, 0x32, 0x1e, 0x0f, 0x00, 0xbf, 0x17, 0x20, 0x97, 0x15, 0x1e, 0xa5, 0xc3, 0x8c, 0xce,
	0xab, 0xb4, 0xb2, 0x21, 0xb7, 0xd0, 0x61, 0xa8, 0xd8, 0xea, 0x99, 0x9c, 0x55, 0x63, 0x8e, 0xa7,
	0xdf, 0x41, 0x57, 0x0f, 0xd2, 0xce, 0x61, 0x6d, 0xf8, 0xf1, 0xd0, 0xd7, 0x8e, 0xf9, 0xf5, 0xe7,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x55, 0x2b, 0xe5, 0x0c, 0x03, 0x00, 0x00,
}
