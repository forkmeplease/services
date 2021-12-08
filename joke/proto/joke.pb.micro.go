// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/joke.proto

package joke

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Joke service

func NewJokeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Joke service

type JokeService interface {
	// get n random jokes
	Random(ctx context.Context, in *RandomRequest, opts ...client.CallOption) (*RandomResponse, error)
}

type jokeService struct {
	c    client.Client
	name string
}

func NewJokeService(name string, c client.Client) JokeService {
	return &jokeService{
		c:    c,
		name: name,
	}
}

func (c *jokeService) Random(ctx context.Context, in *RandomRequest, opts ...client.CallOption) (*RandomResponse, error) {
	req := c.c.NewRequest(c.name, "Joke.Random", in)
	out := new(RandomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Joke service

type JokeHandler interface {
	// get n random jokes
	Random(context.Context, *RandomRequest, *RandomResponse) error
}

func RegisterJokeHandler(s server.Server, hdlr JokeHandler, opts ...server.HandlerOption) error {
	type joke interface {
		Random(ctx context.Context, in *RandomRequest, out *RandomResponse) error
	}
	type Joke struct {
		joke
	}
	h := &jokeHandler{hdlr}
	return s.Handle(s.NewHandler(&Joke{h}, opts...))
}

type jokeHandler struct {
	JokeHandler
}

func (h *jokeHandler) Random(ctx context.Context, in *RandomRequest, out *RandomResponse) error {
	return h.JokeHandler.Random(ctx, in, out)
}
