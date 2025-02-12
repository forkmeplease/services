// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/secret.proto

package secret

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/micro/v5/service/client"
	server "github.com/micro/micro/v5/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Secret service

type SecretService interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...client.CallOption) (*SetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type secretService struct {
	c    client.Client
	name string
}

func NewSecretService(name string, c client.Client) SecretService {
	return &secretService{
		c:    c,
		name: name,
	}
}

func (c *secretService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Secret.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretService) Set(ctx context.Context, in *SetRequest, opts ...client.CallOption) (*SetResponse, error) {
	req := c.c.NewRequest(c.name, "Secret.Set", in)
	out := new(SetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Secret.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Secret.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Secret service

type SecretHandler interface {
	Get(context.Context, *GetRequest, *GetResponse) error
	Set(context.Context, *SetRequest, *SetResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterSecretHandler(s server.Server, hdlr SecretHandler, opts ...server.HandlerOption) error {
	type secret interface {
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Set(ctx context.Context, in *SetRequest, out *SetResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Secret struct {
		secret
	}
	h := &secretHandler{hdlr}
	return s.Handle(s.NewHandler(&Secret{h}, opts...))
}

type secretHandler struct {
	SecretHandler
}

func (h *secretHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.SecretHandler.Get(ctx, in, out)
}

func (h *secretHandler) Set(ctx context.Context, in *SetRequest, out *SetResponse) error {
	return h.SecretHandler.Set(ctx, in, out)
}

func (h *secretHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.SecretHandler.Delete(ctx, in, out)
}

func (h *secretHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.SecretHandler.List(ctx, in, out)
}
