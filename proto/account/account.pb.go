// Code generated by protoc-gen-go.
// source: account.proto
// DO NOT EDIT!

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	account.proto

It has these top-level messages:
	User
	Session
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	UpdateNameRequest
	UpdateNameResponse
	UpdatePasswordRequest
	UpdatePasswordResponse
	SearchRequest
	SearchResponse
	ReadSessionRequest
	ReadSessionResponse
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
*/
package account

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type User struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Created   int64  `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Updated   int64  `protobuf:"varint,5,opt,name=updated" json:"updated,omitempty"`
	AppId     string `protobuf:"bytes,6,opt,name=appId" json:"appId,omitempty"`
	AppSecret string `protobuf:"bytes,7,opt,name=appSecret" json:"appSecret,omitempty"`
	Platform  string `protobuf:"bytes,8,opt,name=platform" json:"platform,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Session struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Created  int64  `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	Expires  int64  `protobuf:"varint,4,opt,name=expires" json:"expires,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateRequest struct {
	User     *User  `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ReadRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ReadResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type UpdateNameRequest struct {
	Email       string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Oldusername string `protobuf:"bytes,2,opt,name=oldusername" json:"oldusername,omitempty"`
	Username    string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
}

func (m *UpdateNameRequest) Reset()                    { *m = UpdateNameRequest{} }
func (m *UpdateNameRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNameRequest) ProtoMessage()               {}
func (*UpdateNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type UpdateNameResponse struct {
}

func (m *UpdateNameResponse) Reset()                    { *m = UpdateNameResponse{} }
func (m *UpdateNameResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNameResponse) ProtoMessage()               {}
func (*UpdateNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type UpdatePasswordRequest struct {
	UserId          string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	OldPassword     string `protobuf:"bytes,2,opt,name=oldPassword" json:"oldPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=newPassword" json:"newPassword,omitempty"`
	ConfirmPassword string `protobuf:"bytes,4,opt,name=confirmPassword" json:"confirmPassword,omitempty"`
}

func (m *UpdatePasswordRequest) Reset()                    { *m = UpdatePasswordRequest{} }
func (m *UpdatePasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordRequest) ProtoMessage()               {}
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type UpdatePasswordResponse struct {
}

func (m *UpdatePasswordResponse) Reset()                    { *m = UpdatePasswordResponse{} }
func (m *UpdatePasswordResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordResponse) ProtoMessage()               {}
func (*UpdatePasswordResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type SearchRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Limit    int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset   int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type SearchResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *SearchResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ReadSessionRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *ReadSessionRequest) Reset()                    { *m = ReadSessionRequest{} }
func (m *ReadSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadSessionRequest) ProtoMessage()               {}
func (*ReadSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type ReadSessionResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *ReadSessionResponse) Reset()                    { *m = ReadSessionResponse{} }
func (m *ReadSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadSessionResponse) ProtoMessage()               {}
func (*ReadSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ReadSessionResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type LoginResponse struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *LoginResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type LogoutRequest struct {
	SessionId string `protobuf:"bytes,1,opt,name=sessionId" json:"sessionId,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

type LogoutResponse struct {
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*Session)(nil), "Session")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*UpdateNameRequest)(nil), "UpdateNameRequest")
	proto.RegisterType((*UpdateNameResponse)(nil), "UpdateNameResponse")
	proto.RegisterType((*UpdatePasswordRequest)(nil), "UpdatePasswordRequest")
	proto.RegisterType((*UpdatePasswordResponse)(nil), "UpdatePasswordResponse")
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "SearchResponse")
	proto.RegisterType((*ReadSessionRequest)(nil), "ReadSessionRequest")
	proto.RegisterType((*ReadSessionResponse)(nil), "ReadSessionResponse")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "LogoutResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...client.CallOption) (*UpdateNameResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
	ReadSession(ctx context.Context, in *ReadSessionRequest, opts ...client.CallOption) (*ReadSessionResponse, error)
}

type accountClient struct {
	c           client.Client
	serviceName string
}

func NewAccountClient(serviceName string, c client.Client) AccountClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "account"
	}
	return &accountClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...client.CallOption) (*UpdateNameResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.UpdateName", in)
	out := new(UpdateNameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.UpdatePassword", in)
	out := new(UpdatePasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) ReadSession(ctx context.Context, in *ReadSessionRequest, opts ...client.CallOption) (*ReadSessionResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.ReadSession", in)
	out := new(ReadSessionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	UpdateName(context.Context, *UpdateNameRequest, *UpdateNameResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	UpdatePassword(context.Context, *UpdatePasswordRequest, *UpdatePasswordResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
	ReadSession(context.Context, *ReadSessionRequest, *ReadSessionResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler) {
	s.Handle(s.NewHandler(&Account{hdlr}))
}

type Account struct {
	AccountHandler
}

func (h *Account) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AccountHandler.Create(ctx, in, out)
}

func (h *Account) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.AccountHandler.Read(ctx, in, out)
}

func (h *Account) UpdateName(ctx context.Context, in *UpdateNameRequest, out *UpdateNameResponse) error {
	return h.AccountHandler.UpdateName(ctx, in, out)
}

func (h *Account) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.AccountHandler.Update(ctx, in, out)
}

func (h *Account) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.AccountHandler.Delete(ctx, in, out)
}

func (h *Account) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.AccountHandler.Search(ctx, in, out)
}

func (h *Account) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *UpdatePasswordResponse) error {
	return h.AccountHandler.UpdatePassword(ctx, in, out)
}

func (h *Account) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *Account) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.AccountHandler.Logout(ctx, in, out)
}

func (h *Account) ReadSession(ctx context.Context, in *ReadSessionRequest, out *ReadSessionResponse) error {
	return h.AccountHandler.ReadSession(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0x9d, 0x34, 0xf3, 0xd3, 0xde, 0x36, 0xe9, 0xf7, 0xb9, 0xa5, 0x84, 0x00, 0xa2, 0xb2, 0x84,
	0x54, 0x40, 0xf5, 0xa2, 0x5d, 0x20, 0xd8, 0xa1, 0x22, 0xa4, 0x4a, 0x08, 0x55, 0x53, 0x75, 0xc7,
	0x26, 0x4c, 0x3c, 0x25, 0xd2, 0x4c, 0x1c, 0xe2, 0x8c, 0xca, 0xcb, 0xf0, 0x46, 0xbc, 0x0e, 0x7b,
	0xfc, 0x9b, 0xd8, 0x69, 0x47, 0x50, 0x76, 0xb9, 0xc7, 0xc7, 0xbe, 0xc7, 0xbe, 0xf7, 0xdc, 0x40,
	0x94, 0xcd, 0x66, 0x6c, 0x55, 0x36, 0xa4, 0xaa, 0x59, 0xc3, 0xf0, 0xcf, 0x00, 0x86, 0x57, 0x9c,
	0xd6, 0x28, 0x86, 0x8d, 0x22, 0x4f, 0x82, 0xc3, 0xe0, 0x68, 0x6b, 0x2a, 0xbe, 0x50, 0x0a, 0x9b,
	0x2b, 0x81, 0x97, 0xd9, 0x92, 0x26, 0x1b, 0x0a, 0x6d, 0x63, 0xb4, 0x0f, 0x23, 0xba, 0xcc, 0x8a,
	0x45, 0x12, 0xaa, 0x05, 0x1d, 0xa0, 0x04, 0x26, 0xb3, 0x9a, 0x66, 0x0d, 0xcd, 0x93, 0xa1, 0xc0,
	0xc3, 0xa9, 0x0d, 0xe5, 0xca, 0xaa, 0xca, 0xd5, 0xca, 0x48, 0xaf, 0x98, 0x50, 0x9e, 0x94, 0x55,
	0xd5, 0x79, 0x9e, 0x8c, 0xf5, 0x49, 0x2a, 0x40, 0x4f, 0x60, 0x4b, 0x7c, 0x5c, 0x52, 0xb1, 0xbf,
	0x49, 0x26, 0x6a, 0xa5, 0x03, 0xa4, 0xb2, 0x6a, 0x91, 0x35, 0x73, 0x56, 0x2f, 0x93, 0x4d, 0xad,
	0xcc, 0xc6, 0xb8, 0x80, 0xc9, 0x25, 0xe5, 0xbc, 0x60, 0xe5, 0xbd, 0x2e, 0xe4, 0x48, 0x0f, 0x6f,
	0x49, 0xa7, 0xdf, 0xab, 0xa2, 0xa6, 0xdc, 0x5e, 0xca, 0x84, 0xf8, 0x03, 0x44, 0x67, 0x8a, 0x34,
	0xa5, 0xdf, 0x56, 0x94, 0x37, 0xe8, 0x11, 0x0c, 0xe5, 0x81, 0x2a, 0xe5, 0xf6, 0xc9, 0x88, 0xc8,
	0x67, 0x9d, 0x2a, 0x48, 0x49, 0xce, 0x38, 0xbf, 0x61, 0x75, 0x6e, 0x73, 0xdb, 0x18, 0x1f, 0x42,
	0x6c, 0xcf, 0xe1, 0x15, 0x2b, 0x39, 0xed, 0x2b, 0xc7, 0xcf, 0x20, 0x7a, 0x4f, 0x17, 0xb4, 0xcb,
	0xd4, 0x27, 0xfc, 0x07, 0xb1, 0x25, 0xe8, 0x23, 0xf0, 0x53, 0xd8, 0x9e, 0xd2, 0x2c, 0x5f, 0xb7,
	0xe1, 0x05, 0xec, 0xe8, 0x65, 0x93, 0x71, 0xbd, 0x74, 0xfc, 0x12, 0xa2, 0x2b, 0x55, 0xac, 0x3f,
	0x5f, 0x53, 0xea, 0xb0, 0x5c, 0xa3, 0xe3, 0x1a, 0xfe, 0xd7, 0xc8, 0x27, 0xf1, 0xcc, 0xf6, 0x84,
	0xb6, 0x7d, 0x02, 0xb7, 0x7d, 0x0e, 0x61, 0x9b, 0x2d, 0xf2, 0x5e, 0x89, 0x5c, 0xc8, 0xab, 0x60,
	0xe8, 0x57, 0x10, 0xef, 0x03, 0x72, 0x13, 0x99, 0xf4, 0x3f, 0x02, 0x78, 0xa0, 0xe1, 0x0b, 0xf3,
	0xdc, 0x56, 0xc3, 0x01, 0x8c, 0xe5, 0xde, 0x73, 0xfb, 0x2a, 0x26, 0x32, 0x2a, 0x2e, 0xfc, 0x62,
	0xb9, 0x90, 0x64, 0x94, 0xf4, 0xa6, 0x65, 0x68, 0x21, 0x2e, 0x84, 0x8e, 0x60, 0x77, 0xc6, 0xca,
	0x79, 0x51, 0x2f, 0x5b, 0xd6, 0x50, 0xb1, 0xfa, 0x30, 0x4e, 0xe0, 0xa0, 0x2f, 0xcf, 0x28, 0x67,
	0x10, 0x5d, 0xd2, 0xac, 0x9e, 0x7d, 0xb5, 0x82, 0xdd, 0xcb, 0x07, 0xeb, 0xfc, 0xb8, 0xe1, 0x3e,
	0xa8, 0x40, 0x17, 0xc5, 0xb2, 0x68, 0x4c, 0x4b, 0xeb, 0x40, 0x5e, 0x9c, 0xcd, 0xe7, 0x5c, 0x18,
	0x4b, 0xf7, 0xb3, 0x89, 0xf0, 0x31, 0xc4, 0x36, 0xa1, 0x69, 0x8a, 0xc7, 0x30, 0x92, 0x19, 0xb8,
	0x48, 0x17, 0x76, 0x95, 0xd6, 0x18, 0x3e, 0x01, 0x24, 0x3b, 0xc8, 0x98, 0xcd, 0x8a, 0x14, 0xc6,
	0xe5, 0x1a, 0x69, 0x1f, 0xb6, 0x03, 0xf0, 0x1b, 0xd8, 0xf3, 0xf6, 0x98, 0x3c, 0x18, 0x26, 0x86,
	0x63, 0x7a, 0x6a, 0x93, 0x58, 0x8a, 0x5d, 0xc0, 0x9f, 0x61, 0xe7, 0x23, 0xbb, 0x2e, 0xca, 0x7f,
	0x7f, 0x0d, 0xd7, 0x82, 0x61, 0xcf, 0x82, 0xa7, 0x10, 0x99, 0xd3, 0xef, 0x21, 0xe9, 0x58, 0x6d,
	0x62, 0xab, 0xe6, 0xef, 0x2e, 0x2f, 0xbc, 0x61, 0xe9, 0x3a, 0xc9, 0xc9, 0xaf, 0x10, 0x26, 0xef,
	0xf4, 0x30, 0x46, 0xaf, 0x60, 0xac, 0x87, 0x00, 0x8a, 0x89, 0x37, 0x55, 0xd2, 0x5d, 0xe2, 0x4f,
	0x07, 0x3c, 0x40, 0xcf, 0x61, 0x28, 0xdf, 0x11, 0xed, 0x10, 0xc7, 0xe3, 0x69, 0x44, 0x5c, 0x4b,
	0x0b, 0xda, 0x6b, 0x80, 0xce, 0x12, 0x08, 0x91, 0x5b, 0x46, 0x4c, 0xf7, 0xc8, 0x1d, 0x9e, 0x19,
	0x48, 0x31, 0x1a, 0x17, 0x62, 0x3c, 0xef, 0x0b, 0x31, 0x3d, 0x7f, 0x2b, 0xb2, 0x9e, 0x3d, 0x82,
	0xec, 0x4d, 0x29, 0x41, 0xee, 0x0d, 0x25, 0x45, 0xd6, 0x4d, 0x26, 0xc8, 0x5e, 0x7b, 0x0b, 0xb2,
	0xdf, 0x7d, 0x82, 0x7c, 0x66, 0xa7, 0x49, 0x6b, 0xac, 0x03, 0x72, 0xa7, 0x99, 0xd3, 0x87, 0x64,
	0x8d, 0x8b, 0x06, 0xc2, 0x8b, 0x23, 0x55, 0x5a, 0x14, 0x11, 0xb7, 0x81, 0xd2, 0x98, 0x78, 0x15,
	0xd7, 0xda, 0x74, 0x81, 0x90, 0x5a, 0xeb, 0x0a, 0x2b, 0xb4, 0xf9, 0x95, 0x13, 0xe4, 0xb7, 0x7a,
	0xbe, 0xda, 0x7f, 0xcd, 0x1e, 0xb9, 0x6d, 0x86, 0x74, 0x9f, 0xdc, 0xd1, 0xed, 0x78, 0xf0, 0x65,
	0xac, 0xfe, 0xbc, 0xa7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xe1, 0xba, 0x65, 0x8a, 0x07,
	0x00, 0x00,
}
