// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/geocoding.proto

package geocoding

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

// Client API for Geocoding service

type GeocodingService interface {
	// Lookup an address, the result will be the normalized address which contains coordinates
	Lookup(ctx context.Context, in *LookupRequest, opts ...client.CallOption) (*LookupResponse, error)
	// Reverse geocode coordinates to an address
	Reverse(ctx context.Context, in *ReverseRequest, opts ...client.CallOption) (*ReverseResponse, error)
}

type geocodingService struct {
	c    client.Client
	name string
}

func NewGeocodingService(name string, c client.Client) GeocodingService {
	return &geocodingService{
		c:    c,
		name: name,
	}
}

func (c *geocodingService) Lookup(ctx context.Context, in *LookupRequest, opts ...client.CallOption) (*LookupResponse, error) {
	req := c.c.NewRequest(c.name, "Geocoding.Lookup", in)
	out := new(LookupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geocodingService) Reverse(ctx context.Context, in *ReverseRequest, opts ...client.CallOption) (*ReverseResponse, error) {
	req := c.c.NewRequest(c.name, "Geocoding.Reverse", in)
	out := new(ReverseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Geocoding service

type GeocodingHandler interface {
	// Lookup an address, the result will be the normalized address which contains coordinates
	Lookup(context.Context, *LookupRequest, *LookupResponse) error
	// Reverse geocode coordinates to an address
	Reverse(context.Context, *ReverseRequest, *ReverseResponse) error
}

func RegisterGeocodingHandler(s server.Server, hdlr GeocodingHandler, opts ...server.HandlerOption) error {
	type geocoding interface {
		Lookup(ctx context.Context, in *LookupRequest, out *LookupResponse) error
		Reverse(ctx context.Context, in *ReverseRequest, out *ReverseResponse) error
	}
	type Geocoding struct {
		geocoding
	}
	h := &geocodingHandler{hdlr}
	return s.Handle(s.NewHandler(&Geocoding{h}, opts...))
}

type geocodingHandler struct {
	GeocodingHandler
}

func (h *geocodingHandler) Lookup(ctx context.Context, in *LookupRequest, out *LookupResponse) error {
	return h.GeocodingHandler.Lookup(ctx, in, out)
}

func (h *geocodingHandler) Reverse(ctx context.Context, in *ReverseRequest, out *ReverseResponse) error {
	return h.GeocodingHandler.Reverse(ctx, in, out)
}
