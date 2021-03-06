// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: showing/proto/showing.proto

package showing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for Showing service

type ShowingService interface {
	CreateShowing(ctx context.Context, in *CreateShowingRequest, opts ...client.CallOption) (*CreateShowingResponse, error)
	ReadShowing(ctx context.Context, in *ReadShowingRequest, opts ...client.CallOption) (*ReadShowingResponse, error)
	UpdateShowing(ctx context.Context, in *UpdateShowingRequest, opts ...client.CallOption) (*UpdateShowingResponse, error)
	DeleteShowing(ctx context.Context, in *DeleteShowingRequest, opts ...client.CallOption) (*DeleteShowingResponse, error)
	GetAllShowings(ctx context.Context, in *GetAllShowingsRequest, opts ...client.CallOption) (*GetAllShowingsResponse, error)
	DeleteShowingsForMovie(ctx context.Context, in *DeleteShowingsForMovieRequest, opts ...client.CallOption) (*DeleteShowingsForMovieResponse, error)
	DeleteShowingsForRoom(ctx context.Context, in *DeleteShowingsForRoomRequest, opts ...client.CallOption) (*DeleteShowingsForRoomResponse, error)
}

type showingService struct {
	c    client.Client
	name string
}

func NewShowingService(name string, c client.Client) ShowingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "showing"
	}
	return &showingService{
		c:    c,
		name: name,
	}
}

func (c *showingService) CreateShowing(ctx context.Context, in *CreateShowingRequest, opts ...client.CallOption) (*CreateShowingResponse, error) {
	req := c.c.NewRequest(c.name, "Showing.CreateShowing", in)
	out := new(CreateShowingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showingService) ReadShowing(ctx context.Context, in *ReadShowingRequest, opts ...client.CallOption) (*ReadShowingResponse, error) {
	req := c.c.NewRequest(c.name, "Showing.ReadShowing", in)
	out := new(ReadShowingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showingService) UpdateShowing(ctx context.Context, in *UpdateShowingRequest, opts ...client.CallOption) (*UpdateShowingResponse, error) {
	req := c.c.NewRequest(c.name, "Showing.UpdateShowing", in)
	out := new(UpdateShowingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showingService) DeleteShowing(ctx context.Context, in *DeleteShowingRequest, opts ...client.CallOption) (*DeleteShowingResponse, error) {
	req := c.c.NewRequest(c.name, "Showing.DeleteShowing", in)
	out := new(DeleteShowingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showingService) GetAllShowings(ctx context.Context, in *GetAllShowingsRequest, opts ...client.CallOption) (*GetAllShowingsResponse, error) {
	req := c.c.NewRequest(c.name, "Showing.GetAllShowings", in)
	out := new(GetAllShowingsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showingService) DeleteShowingsForMovie(ctx context.Context, in *DeleteShowingsForMovieRequest, opts ...client.CallOption) (*DeleteShowingsForMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Showing.DeleteShowingsForMovie", in)
	out := new(DeleteShowingsForMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showingService) DeleteShowingsForRoom(ctx context.Context, in *DeleteShowingsForRoomRequest, opts ...client.CallOption) (*DeleteShowingsForRoomResponse, error) {
	req := c.c.NewRequest(c.name, "Showing.DeleteShowingsForRoom", in)
	out := new(DeleteShowingsForRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Showing service

type ShowingHandler interface {
	CreateShowing(context.Context, *CreateShowingRequest, *CreateShowingResponse) error
	ReadShowing(context.Context, *ReadShowingRequest, *ReadShowingResponse) error
	UpdateShowing(context.Context, *UpdateShowingRequest, *UpdateShowingResponse) error
	DeleteShowing(context.Context, *DeleteShowingRequest, *DeleteShowingResponse) error
	GetAllShowings(context.Context, *GetAllShowingsRequest, *GetAllShowingsResponse) error
	DeleteShowingsForMovie(context.Context, *DeleteShowingsForMovieRequest, *DeleteShowingsForMovieResponse) error
	DeleteShowingsForRoom(context.Context, *DeleteShowingsForRoomRequest, *DeleteShowingsForRoomResponse) error
}

func RegisterShowingHandler(s server.Server, hdlr ShowingHandler, opts ...server.HandlerOption) error {
	type showing interface {
		CreateShowing(ctx context.Context, in *CreateShowingRequest, out *CreateShowingResponse) error
		ReadShowing(ctx context.Context, in *ReadShowingRequest, out *ReadShowingResponse) error
		UpdateShowing(ctx context.Context, in *UpdateShowingRequest, out *UpdateShowingResponse) error
		DeleteShowing(ctx context.Context, in *DeleteShowingRequest, out *DeleteShowingResponse) error
		GetAllShowings(ctx context.Context, in *GetAllShowingsRequest, out *GetAllShowingsResponse) error
		DeleteShowingsForMovie(ctx context.Context, in *DeleteShowingsForMovieRequest, out *DeleteShowingsForMovieResponse) error
		DeleteShowingsForRoom(ctx context.Context, in *DeleteShowingsForRoomRequest, out *DeleteShowingsForRoomResponse) error
	}
	type Showing struct {
		showing
	}
	h := &showingHandler{hdlr}
	return s.Handle(s.NewHandler(&Showing{h}, opts...))
}

type showingHandler struct {
	ShowingHandler
}

func (h *showingHandler) CreateShowing(ctx context.Context, in *CreateShowingRequest, out *CreateShowingResponse) error {
	return h.ShowingHandler.CreateShowing(ctx, in, out)
}

func (h *showingHandler) ReadShowing(ctx context.Context, in *ReadShowingRequest, out *ReadShowingResponse) error {
	return h.ShowingHandler.ReadShowing(ctx, in, out)
}

func (h *showingHandler) UpdateShowing(ctx context.Context, in *UpdateShowingRequest, out *UpdateShowingResponse) error {
	return h.ShowingHandler.UpdateShowing(ctx, in, out)
}

func (h *showingHandler) DeleteShowing(ctx context.Context, in *DeleteShowingRequest, out *DeleteShowingResponse) error {
	return h.ShowingHandler.DeleteShowing(ctx, in, out)
}

func (h *showingHandler) GetAllShowings(ctx context.Context, in *GetAllShowingsRequest, out *GetAllShowingsResponse) error {
	return h.ShowingHandler.GetAllShowings(ctx, in, out)
}

func (h *showingHandler) DeleteShowingsForMovie(ctx context.Context, in *DeleteShowingsForMovieRequest, out *DeleteShowingsForMovieResponse) error {
	return h.ShowingHandler.DeleteShowingsForMovie(ctx, in, out)
}

func (h *showingHandler) DeleteShowingsForRoom(ctx context.Context, in *DeleteShowingsForRoomRequest, out *DeleteShowingsForRoomResponse) error {
	return h.ShowingHandler.DeleteShowingsForRoom(ctx, in, out)
}
