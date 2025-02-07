// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/rss.proto

package rss

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

// Client API for Rss service

type RssService interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error)
	Feed(ctx context.Context, in *FeedRequest, opts ...client.CallOption) (*FeedResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type rssService struct {
	c    client.Client
	name string
}

func NewRssService(name string, c client.Client) RssService {
	return &rssService{
		c:    c,
		name: name,
	}
}

func (c *rssService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "Rss.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rssService) Remove(ctx context.Context, in *RemoveRequest, opts ...client.CallOption) (*RemoveResponse, error) {
	req := c.c.NewRequest(c.name, "Rss.Remove", in)
	out := new(RemoveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rssService) Feed(ctx context.Context, in *FeedRequest, opts ...client.CallOption) (*FeedResponse, error) {
	req := c.c.NewRequest(c.name, "Rss.Feed", in)
	out := new(FeedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rssService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Rss.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rss service

type RssHandler interface {
	Add(context.Context, *AddRequest, *AddResponse) error
	Remove(context.Context, *RemoveRequest, *RemoveResponse) error
	Feed(context.Context, *FeedRequest, *FeedResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterRssHandler(s server.Server, hdlr RssHandler, opts ...server.HandlerOption) error {
	type rss interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error
		Feed(ctx context.Context, in *FeedRequest, out *FeedResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Rss struct {
		rss
	}
	h := &rssHandler{hdlr}
	return s.Handle(s.NewHandler(&Rss{h}, opts...))
}

type rssHandler struct {
	RssHandler
}

func (h *rssHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.RssHandler.Add(ctx, in, out)
}

func (h *rssHandler) Remove(ctx context.Context, in *RemoveRequest, out *RemoveResponse) error {
	return h.RssHandler.Remove(ctx, in, out)
}

func (h *rssHandler) Feed(ctx context.Context, in *FeedRequest, out *FeedResponse) error {
	return h.RssHandler.Feed(ctx, in, out)
}

func (h *rssHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.RssHandler.List(ctx, in, out)
}
