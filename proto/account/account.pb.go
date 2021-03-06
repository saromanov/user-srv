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
	Id                  string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username            string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email               string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Created             int64    `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Updated             int64    `protobuf:"varint,5,opt,name=updated" json:"updated,omitempty"`
	AppId               string   `protobuf:"bytes,6,opt,name=appId" json:"appId,omitempty"`
	AppSecret           string   `protobuf:"bytes,7,opt,name=appSecret" json:"appSecret,omitempty"`
	Platform            string   `protobuf:"bytes,8,opt,name=platform" json:"platform,omitempty"`
	Password            string   `protobuf:"bytes,9,opt,name=password" json:"password,omitempty"`
	MobileNumber        string   `protobuf:"bytes,10,opt,name=mobileNumber" json:"mobileNumber,omitempty"`
	PhoneNumber         string   `protobuf:"bytes,11,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	ReceiveNotification bool     `protobuf:"varint,12,opt,name=receiveNotification" json:"receiveNotification,omitempty"`
	Address             string   `protobuf:"bytes,13,opt,name=address" json:"address,omitempty"`
	City                string   `protobuf:"bytes,14,opt,name=city" json:"city,omitempty"`
	State               string   `protobuf:"bytes,15,opt,name=state" json:"state,omitempty"`
	Country             string   `protobuf:"bytes,16,opt,name=country" json:"country,omitempty"`
	Zip                 string   `protobuf:"bytes,17,opt,name=zip" json:"zip,omitempty"`
	Salt                string   `protobuf:"bytes,18,opt,name=salt" json:"salt,omitempty"`
	ProfilePhoto        string   `protobuf:"bytes,19,opt,name=profilePhoto" json:"profilePhoto,omitempty"`
	PromoCodeID         string   `protobuf:"bytes,20,opt,name=promoCodeID" json:"promoCodeID,omitempty"`
	JoinMailList        bool     `protobuf:"varint,21,opt,name=joinMailList" json:"joinMailList,omitempty"`
	Organization        string   `protobuf:"bytes,22,opt,name=organization" json:"organization,omitempty"`
	WhitelistHosts      []string `protobuf:"bytes,23,rep,name=whitelistHosts" json:"whitelistHosts,omitempty"`
	WhitelistIPs        []string `protobuf:"bytes,24,rep,name=whitelistIPs" json:"whitelistIPs,omitempty"`
	SubscriptionStart   string   `protobuf:"bytes,25,opt,name=subscriptionStart" json:"subscriptionStart,omitempty"`
	SubscriptionEnd     float64  `protobuf:"fixed64,26,opt,name=subscriptionEnd" json:"subscriptionEnd,omitempty"`
	SubscriptionPlan    string   `protobuf:"bytes,27,opt,name=subscriptionPlan" json:"subscriptionPlan,omitempty"`
	IsSubscriptionValid bool     `protobuf:"varint,28,opt,name=isSubscriptionValid" json:"isSubscriptionValid,omitempty"`
	Role                string   `protobuf:"bytes,29,opt,name=role" json:"role,omitempty"`
	IsEnabled           bool     `protobuf:"varint,30,opt,name=isEnabled" json:"isEnabled,omitempty"`
	IsVerified          bool     `protobuf:"varint,31,opt,name=isVerified" json:"isVerified,omitempty"`
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
	Id                  string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username            string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email               string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Created             int64    `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Updated             int64    `protobuf:"varint,5,opt,name=updated" json:"updated,omitempty"`
	AppId               string   `protobuf:"bytes,6,opt,name=appId" json:"appId,omitempty"`
	AppSecret           string   `protobuf:"bytes,7,opt,name=appSecret" json:"appSecret,omitempty"`
	Platform            string   `protobuf:"bytes,8,opt,name=platform" json:"platform,omitempty"`
	Password            string   `protobuf:"bytes,9,opt,name=password" json:"password,omitempty"`
	MobileNumber        string   `protobuf:"bytes,10,opt,name=mobileNumber" json:"mobileNumber,omitempty"`
	PhoneNumber         string   `protobuf:"bytes,11,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	ReceiveNotification bool     `protobuf:"varint,12,opt,name=receiveNotification" json:"receiveNotification,omitempty"`
	Address             string   `protobuf:"bytes,13,opt,name=address" json:"address,omitempty"`
	City                string   `protobuf:"bytes,14,opt,name=city" json:"city,omitempty"`
	State               string   `protobuf:"bytes,15,opt,name=state" json:"state,omitempty"`
	Country             string   `protobuf:"bytes,16,opt,name=country" json:"country,omitempty"`
	Zip                 string   `protobuf:"bytes,17,opt,name=zip" json:"zip,omitempty"`
	Salt                string   `protobuf:"bytes,18,opt,name=salt" json:"salt,omitempty"`
	ProfilePhoto        string   `protobuf:"bytes,19,opt,name=profilePhoto" json:"profilePhoto,omitempty"`
	PromoCodeID         string   `protobuf:"bytes,20,opt,name=promoCodeID" json:"promoCodeID,omitempty"`
	JoinMailList        bool     `protobuf:"varint,21,opt,name=joinMailList" json:"joinMailList,omitempty"`
	Organization        string   `protobuf:"bytes,22,opt,name=organization" json:"organization,omitempty"`
	WhitelistHosts      []string `protobuf:"bytes,23,rep,name=whitelistHosts" json:"whitelistHosts,omitempty"`
	WhitelistIPs        []string `protobuf:"bytes,24,rep,name=whitelistIPs" json:"whitelistIPs,omitempty"`
	SubscriptionStart   string   `protobuf:"bytes,25,opt,name=subscriptionStart" json:"subscriptionStart,omitempty"`
	SubscriptionEnd     float64  `protobuf:"fixed64,26,opt,name=subscriptionEnd" json:"subscriptionEnd,omitempty"`
	SubscriptionPlan    string   `protobuf:"bytes,27,opt,name=subscriptionPlan" json:"subscriptionPlan,omitempty"`
	IsSubscriptionValid bool     `protobuf:"varint,28,opt,name=isSubscriptionValid" json:"isSubscriptionValid,omitempty"`
	Role                string   `protobuf:"bytes,29,opt,name=role" json:"role,omitempty"`
	IsEnabled           bool     `protobuf:"varint,30,opt,name=isEnabled" json:"isEnabled,omitempty"`
	IsVerified          bool     `protobuf:"varint,31,opt,name=isVerified" json:"isVerified,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	UpdateName(ctx context.Context, in *UpdateNameRequest, opts ...client.CallOption) (*UpdateNameResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error)
	LoginNative(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
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

func (c *accountClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*CreateResponse, error) {
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

func (c *accountClient) LoginNative(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Account.LoginNative", in)
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
	Create(context.Context, *User, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	UpdateName(context.Context, *UpdateNameRequest, *UpdateNameResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	UpdatePassword(context.Context, *UpdatePasswordRequest, *UpdatePasswordResponse) error
	LoginNative(context.Context, *LoginRequest, *LoginResponse) error
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
	ReadSession(context.Context, *ReadSessionRequest, *ReadSessionResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler) {
	s.Handle(s.NewHandler(&Account{hdlr}))
}

type Account struct {
	AccountHandler
}

func (h *Account) Create(ctx context.Context, in *User, out *CreateResponse) error {
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

func (h *Account) LoginNative(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.AccountHandler.LoginNative(ctx, in, out)
}

func (h *Account) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.AccountHandler.Logout(ctx, in, out)
}

func (h *Account) ReadSession(ctx context.Context, in *ReadSessionRequest, out *ReadSessionResponse) error {
	return h.AccountHandler.ReadSession(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 1034 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x57, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0xb6, 0x2c, 0x59, 0xb2, 0x47, 0x3f, 0xb6, 0x29, 0xc7, 0x61, 0x94, 0x3f, 0x83, 0x40, 0x8b,
	0x34, 0x6d, 0x16, 0x85, 0x73, 0x28, 0xda, 0x5b, 0xe1, 0x04, 0xa8, 0x81, 0xd4, 0x30, 0x24, 0x24,
	0xa7, 0x5e, 0x56, 0x5a, 0xca, 0x66, 0xb1, 0x5a, 0x6e, 0x97, 0xab, 0xb8, 0xc9, 0xbb, 0xf4, 0x75,
	0x7a, 0xed, 0xbb, 0xf4, 0x09, 0x3a, 0xfc, 0x93, 0xb8, 0x92, 0x5c, 0x34, 0xbd, 0x05, 0xf0, 0x8d,
	0xf3, 0xcd, 0xc7, 0xe1, 0x90, 0x33, 0xdf, 0xac, 0x04, 0xdd, 0x78, 0x32, 0x91, 0xf3, 0xac, 0x8c,
	0xf2, 0x42, 0x96, 0x92, 0xfd, 0xd9, 0x82, 0xc6, 0x5b, 0xc5, 0x0b, 0xd2, 0x83, 0x6d, 0x91, 0xd0,
	0xda, 0x49, 0xed, 0xd9, 0xde, 0x10, 0x57, 0x64, 0x00, 0xbb, 0x73, 0xc4, 0xb3, 0x78, 0xc6, 0xe9,
	0xb6, 0x41, 0x17, 0x36, 0x39, 0x82, 0x1d, 0x3e, 0x8b, 0x45, 0x4a, 0xeb, 0xc6, 0x61, 0x0d, 0x42,
	0xa1, 0x35, 0x29, 0x78, 0x5c, 0xf2, 0x84, 0x36, 0x10, 0xaf, 0x0f, 0xbd, 0xa9, 0x3d, 0xf3, 0x3c,
	0x31, 0x9e, 0x1d, 0xeb, 0x71, 0xa6, 0x8e, 0x14, 0xe7, 0xf9, 0x79, 0x42, 0x9b, 0x36, 0x92, 0x31,
	0xc8, 0x23, 0xd8, 0xc3, 0xc5, 0x88, 0xe3, 0xfe, 0x92, 0xb6, 0x8c, 0x67, 0x09, 0xe8, 0xcc, 0xf2,
	0x34, 0x2e, 0xa7, 0xb2, 0x98, 0xd1, 0x5d, 0x9b, 0x99, 0xb7, 0x8d, 0x2f, 0x56, 0xea, 0x46, 0x16,
	0x09, 0xdd, 0x73, 0x3e, 0x67, 0x13, 0x06, 0x9d, 0x99, 0x1c, 0x8b, 0x94, 0x5f, 0xcc, 0x67, 0x63,
	0x5e, 0x50, 0x30, 0xfe, 0x0a, 0x46, 0x4e, 0xa0, 0x9d, 0x5f, 0xcb, 0xcc, 0x53, 0xda, 0x86, 0x12,
	0x42, 0xe4, 0x5b, 0xe8, 0x17, 0x7c, 0xc2, 0xc5, 0x7b, 0x7e, 0x21, 0x4b, 0x31, 0x15, 0x93, 0xb8,
	0x14, 0x32, 0xa3, 0x1d, 0x64, 0xee, 0x0e, 0x37, 0xb9, 0xf4, 0xed, 0xe3, 0x24, 0x29, 0xb8, 0x52,
	0xb4, 0x6b, 0xe2, 0x79, 0x93, 0x10, 0x68, 0x4c, 0x44, 0xf9, 0x81, 0xf6, 0x0c, 0x6c, 0xd6, 0xfa,
	0x45, 0x54, 0x89, 0x6f, 0x43, 0xf7, 0xed, 0x8b, 0x18, 0xc3, 0xbc, 0xad, 0xae, 0x5a, 0xf1, 0x81,
	0x1e, 0xd8, 0x18, 0xce, 0x24, 0x07, 0x50, 0xff, 0x28, 0x72, 0x7a, 0x68, 0x50, 0xbd, 0xd4, 0x51,
	0x55, 0x9c, 0x96, 0x94, 0xd8, 0xa8, 0x7a, 0xad, 0xef, 0x8e, 0xf5, 0x9e, 0xe2, 0x45, 0x2f, 0xaf,
	0xb1, 0xec, 0xb4, 0x6f, 0xef, 0x1e, 0x62, 0xe6, 0xee, 0x85, 0x9c, 0xc9, 0x33, 0x99, 0xf0, 0xf3,
	0x57, 0xf4, 0xc8, 0xdd, 0x7d, 0x09, 0xe9, 0x28, 0xbf, 0x4a, 0x91, 0xfd, 0x8c, 0xd5, 0x7e, 0x23,
	0x54, 0x49, 0xef, 0x99, 0x4b, 0x57, 0x30, 0xcd, 0x91, 0xc5, 0x55, 0x9c, 0x89, 0x8f, 0xf6, 0x61,
	0x8e, 0xed, 0x49, 0x21, 0x46, 0xbe, 0x84, 0xde, 0xcd, 0xb5, 0x28, 0x79, 0x8a, 0x1b, 0x7e, 0x92,
	0xaa, 0x54, 0xf4, 0xfe, 0x49, 0x1d, 0x59, 0x2b, 0xa8, 0x8e, 0xb5, 0x40, 0xce, 0x2f, 0x15, 0xa5,
	0x86, 0x55, 0xc1, 0xc8, 0x37, 0x70, 0xa8, 0xe6, 0x63, 0x35, 0x29, 0x44, 0xae, 0x63, 0x8f, 0xca,
	0xb8, 0x28, 0xe9, 0x03, 0x73, 0xe8, 0xba, 0x83, 0x3c, 0x83, 0xfd, 0x10, 0x7c, 0x9d, 0x25, 0x74,
	0x80, 0xdc, 0xda, 0x70, 0x15, 0x26, 0xcf, 0xe1, 0x20, 0x84, 0x2e, 0xd3, 0x38, 0xa3, 0x0f, 0x4d,
	0xd8, 0x35, 0x5c, 0xf7, 0x84, 0x50, 0xa3, 0x00, 0x7d, 0x17, 0xa7, 0x28, 0xa6, 0x47, 0xb6, 0x27,
	0x36, 0xb8, 0x74, 0x8d, 0x0a, 0x99, 0x72, 0xfa, 0xd8, 0xd6, 0x48, 0xaf, 0x75, 0xd7, 0x0b, 0xf5,
	0x3a, 0x8b, 0xc7, 0x29, 0xea, 0xe4, 0x89, 0xd9, 0xbb, 0x04, 0xc8, 0x13, 0x00, 0xa1, 0xde, 0xf1,
	0x02, 0xfb, 0x0a, 0xdd, 0x4f, 0x8d, 0x3b, 0x40, 0x98, 0x80, 0xd6, 0x08, 0x7b, 0x4a, 0x3f, 0xef,
	0xa7, 0x48, 0x39, 0x10, 0x6d, 0x7d, 0x4d, 0xb4, 0xfc, 0xf7, 0x5c, 0x60, 0xa3, 0x7a, 0x39, 0x3b,
	0x93, 0xfd, 0xd5, 0x82, 0xee, 0x99, 0x61, 0x0d, 0xf9, 0x6f, 0x73, 0x8e, 0x45, 0xbf, 0x1b, 0x1e,
	0x77, 0xc3, 0xe3, 0x6e, 0x78, 0x7c, 0xbe, 0xc3, 0xe3, 0x04, 0x7a, 0x5e, 0xd0, 0x2a, 0x97, 0x99,
	0xe2, 0xab, 0x8a, 0x66, 0x4f, 0xa1, 0xfb, 0x8a, 0xa7, 0xfc, 0x56, 0xc9, 0xb3, 0x03, 0xe8, 0x79,
	0x82, 0x0d, 0xc1, 0x1e, 0x43, 0x7b, 0xc8, 0xe3, 0xe4, 0xb6, 0x0d, 0x5f, 0x41, 0xc7, 0xba, 0xdd,
	0x89, 0x0f, 0xa0, 0xa1, 0x67, 0x84, 0x61, 0xb4, 0x4f, 0x77, 0x22, 0xfd, 0xab, 0x64, 0x68, 0x20,
	0xf6, 0x1c, 0xba, 0x6f, 0x8d, 0xe6, 0x7d, 0xac, 0x7f, 0xe1, 0x62, 0x1e, 0x9e, 0xeb, 0xf2, 0xb8,
	0x82, 0x43, 0x8b, 0x5c, 0xe0, 0xf8, 0xf1, 0x11, 0x16, 0x53, 0xa8, 0x16, 0x4e, 0x21, 0xec, 0x62,
	0x99, 0x26, 0x2b, 0xa3, 0x2b, 0x84, 0x2a, 0x93, 0xad, 0x5e, 0x9d, 0x6c, 0xec, 0x08, 0x48, 0x78,
	0x90, 0x3b, 0xfe, 0x8f, 0x1a, 0xdc, 0xb3, 0xf0, 0xa5, 0x1b, 0x26, 0x3e, 0x87, 0x63, 0x68, 0xea,
	0xbd, 0xe7, 0xfe, 0x55, 0x9c, 0xe5, 0xb2, 0xf0, 0xec, 0x20, 0x0b, 0x0f, 0x69, 0x46, 0xc6, 0x6f,
	0x16, 0x0c, 0x9b, 0x48, 0x08, 0xe9, 0x5e, 0x9d, 0xc8, 0x6c, 0x2a, 0x8a, 0xd9, 0x82, 0xd5, 0x30,
	0xac, 0x55, 0x98, 0x51, 0x38, 0x5e, 0x4d, 0xcf, 0x65, 0x2e, 0xa1, 0x3b, 0xe2, 0x71, 0x31, 0xb9,
	0xf6, 0x09, 0x87, 0x97, 0xaf, 0xdd, 0x36, 0xd6, 0xb7, 0xc3, 0x07, 0x45, 0x34, 0x15, 0x33, 0x51,
	0xba, 0x8f, 0x8b, 0x35, 0xf4, 0xc5, 0xe5, 0x74, 0xaa, 0x70, 0x3e, 0xdb, 0x59, 0xef, 0x2c, 0xf6,
	0x02, 0x7a, 0xfe, 0x40, 0xd7, 0x14, 0x0f, 0x61, 0x47, 0x9f, 0xa0, 0xf0, 0xb8, 0xfa, 0xb2, 0xd2,
	0x16, 0x63, 0xa7, 0x40, 0x74, 0x07, 0xb9, 0xcf, 0x9e, 0x4f, 0x12, 0x95, 0xa0, 0x2c, 0xb2, 0x78,
	0xd8, 0x25, 0xc0, 0xbe, 0x87, 0x7e, 0x65, 0x8f, 0x3b, 0x87, 0x41, 0xcb, 0x71, 0x5c, 0x4f, 0xed,
	0x46, 0x9e, 0xe2, 0x1d, 0xec, 0x17, 0xe8, 0xbc, 0x91, 0x57, 0x22, 0xfb, 0xff, 0xaf, 0x11, 0x7e,
	0x60, 0xea, 0xd5, 0x0f, 0x0c, 0x7b, 0x09, 0x5d, 0x17, 0xfd, 0x13, 0x52, 0x7a, 0x61, 0x36, 0xc9,
	0x79, 0xf9, 0xdf, 0x2e, 0x8f, 0xda, 0xf0, 0x74, 0x7b, 0xc8, 0xe9, 0xdf, 0x75, 0x68, 0xfd, 0x68,
	0xff, 0x10, 0xe0, 0x81, 0x4d, 0x3b, 0x04, 0x88, 0x7d, 0xe6, 0xc1, 0x7e, 0x54, 0x1d, 0x0a, 0x6c,
	0x8b, 0x7c, 0x01, 0x0d, 0xfd, 0x7c, 0xa4, 0x13, 0x05, 0xd2, 0x1e, 0x74, 0xa3, 0x50, 0xc9, 0x48,
	0xfb, 0x0e, 0x60, 0xa9, 0x04, 0x42, 0xa2, 0x35, 0xfd, 0x0d, 0xfa, 0xd1, 0x06, 0xa9, 0x6c, 0x91,
	0xaf, 0xa1, 0x69, 0x71, 0xd2, 0x8b, 0x2a, 0x92, 0xc7, 0x64, 0x56, 0x64, 0x6d, 0xc8, 0x76, 0xe4,
	0x20, 0xb9, 0x32, 0x9c, 0x90, 0xbc, 0x32, 0x8b, 0x0c, 0xd9, 0xf6, 0x16, 0x92, 0x2b, 0x5d, 0x8d,
	0xe4, 0x6a, 0xd3, 0x21, 0xf9, 0xcc, 0x0f, 0x91, 0x85, 0x9e, 0x8e, 0xa3, 0x8d, 0x1a, 0x1e, 0xdc,
	0x8f, 0x6e, 0x11, 0xcf, 0x16, 0x89, 0xa0, 0x6d, 0x2a, 0x7a, 0x81, 0xdf, 0xad, 0xf7, 0x9c, 0x74,
	0xa3, 0xb0, 0x7b, 0x06, 0xbd, 0xa8, 0x52, 0x6e, 0x9b, 0xa1, 0xad, 0x0e, 0x31, 0xbe, 0x65, 0x55,
	0x31, 0xc3, 0x6a, 0xd9, 0x90, 0xfc, 0x83, 0x1d, 0xae, 0xfe, 0x27, 0x5f, 0x3f, 0x5a, 0x57, 0xc2,
	0xe0, 0x28, 0xda, 0xd0, 0xea, 0x6c, 0x6b, 0xdc, 0x34, 0x7f, 0xfd, 0x5e, 0xfe, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x21, 0x26, 0x92, 0x71, 0x0b, 0x0e, 0x00, 0x00,
}
