// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/image.proto

package image

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

// Client API for Image service

type ImageService interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (*UploadResponse, error)
	Resize(ctx context.Context, in *ResizeRequest, opts ...client.CallOption) (*ResizeResponse, error)
	Convert(ctx context.Context, in *ConvertRequest, opts ...client.CallOption) (*ConvertResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type imageService struct {
	c    client.Client
	name string
}

func NewImageService(name string, c client.Client) ImageService {
	return &imageService{
		c:    c,
		name: name,
	}
}

func (c *imageService) Upload(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (*UploadResponse, error) {
	req := c.c.NewRequest(c.name, "Image.Upload", in)
	out := new(UploadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) Resize(ctx context.Context, in *ResizeRequest, opts ...client.CallOption) (*ResizeResponse, error) {
	req := c.c.NewRequest(c.name, "Image.Resize", in)
	out := new(ResizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) Convert(ctx context.Context, in *ConvertRequest, opts ...client.CallOption) (*ConvertResponse, error) {
	req := c.c.NewRequest(c.name, "Image.Convert", in)
	out := new(ConvertResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Image.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Image service

type ImageHandler interface {
	Upload(context.Context, *UploadRequest, *UploadResponse) error
	Resize(context.Context, *ResizeRequest, *ResizeResponse) error
	Convert(context.Context, *ConvertRequest, *ConvertResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterImageHandler(s server.Server, hdlr ImageHandler, opts ...server.HandlerOption) error {
	type image interface {
		Upload(ctx context.Context, in *UploadRequest, out *UploadResponse) error
		Resize(ctx context.Context, in *ResizeRequest, out *ResizeResponse) error
		Convert(ctx context.Context, in *ConvertRequest, out *ConvertResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type Image struct {
		image
	}
	h := &imageHandler{hdlr}
	return s.Handle(s.NewHandler(&Image{h}, opts...))
}

type imageHandler struct {
	ImageHandler
}

func (h *imageHandler) Upload(ctx context.Context, in *UploadRequest, out *UploadResponse) error {
	return h.ImageHandler.Upload(ctx, in, out)
}

func (h *imageHandler) Resize(ctx context.Context, in *ResizeRequest, out *ResizeResponse) error {
	return h.ImageHandler.Resize(ctx, in, out)
}

func (h *imageHandler) Convert(ctx context.Context, in *ConvertRequest, out *ConvertResponse) error {
	return h.ImageHandler.Convert(ctx, in, out)
}

func (h *imageHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ImageHandler.Delete(ctx, in, out)
}
