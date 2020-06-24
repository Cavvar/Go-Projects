// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: movie/proto/movie.proto

package movie

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

// Client API for Movie service

type MovieService interface {
	CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...client.CallOption) (*CreateMovieResponse, error)
	ReadMovie(ctx context.Context, in *ReadMovieRequest, opts ...client.CallOption) (*ReadMovieResponse, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...client.CallOption) (*DeleteMovieResponse, error)
	GetAllMovies(ctx context.Context, in *GetAllMoviesRequest, opts ...client.CallOption) (*GetAllMoviesResponse, error)
}

type movieService struct {
	c    client.Client
	name string
}

func NewMovieService(name string, c client.Client) MovieService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "movie"
	}
	return &movieService{
		c:    c,
		name: name,
	}
}

func (c *movieService) CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...client.CallOption) (*CreateMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movie.CreateMovie", in)
	out := new(CreateMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) ReadMovie(ctx context.Context, in *ReadMovieRequest, opts ...client.CallOption) (*ReadMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movie.ReadMovie", in)
	out := new(ReadMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...client.CallOption) (*DeleteMovieResponse, error) {
	req := c.c.NewRequest(c.name, "Movie.DeleteMovie", in)
	out := new(DeleteMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieService) GetAllMovies(ctx context.Context, in *GetAllMoviesRequest, opts ...client.CallOption) (*GetAllMoviesResponse, error) {
	req := c.c.NewRequest(c.name, "Movie.GetAllMovies", in)
	out := new(GetAllMoviesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Movie service

type MovieHandler interface {
	CreateMovie(context.Context, *CreateMovieRequest, *CreateMovieResponse) error
	ReadMovie(context.Context, *ReadMovieRequest, *ReadMovieResponse) error
	DeleteMovie(context.Context, *DeleteMovieRequest, *DeleteMovieResponse) error
	GetAllMovies(context.Context, *GetAllMoviesRequest, *GetAllMoviesResponse) error
}

func RegisterMovieHandler(s server.Server, hdlr MovieHandler, opts ...server.HandlerOption) error {
	type movie interface {
		CreateMovie(ctx context.Context, in *CreateMovieRequest, out *CreateMovieResponse) error
		ReadMovie(ctx context.Context, in *ReadMovieRequest, out *ReadMovieResponse) error
		DeleteMovie(ctx context.Context, in *DeleteMovieRequest, out *DeleteMovieResponse) error
		GetAllMovies(ctx context.Context, in *GetAllMoviesRequest, out *GetAllMoviesResponse) error
	}
	type Movie struct {
		movie
	}
	h := &movieHandler{hdlr}
	return s.Handle(s.NewHandler(&Movie{h}, opts...))
}

type movieHandler struct {
	MovieHandler
}

func (h *movieHandler) CreateMovie(ctx context.Context, in *CreateMovieRequest, out *CreateMovieResponse) error {
	return h.MovieHandler.CreateMovie(ctx, in, out)
}

func (h *movieHandler) ReadMovie(ctx context.Context, in *ReadMovieRequest, out *ReadMovieResponse) error {
	return h.MovieHandler.ReadMovie(ctx, in, out)
}

func (h *movieHandler) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, out *DeleteMovieResponse) error {
	return h.MovieHandler.DeleteMovie(ctx, in, out)
}

func (h *movieHandler) GetAllMovies(ctx context.Context, in *GetAllMoviesRequest, out *GetAllMoviesResponse) error {
	return h.MovieHandler.GetAllMovies(ctx, in, out)
}