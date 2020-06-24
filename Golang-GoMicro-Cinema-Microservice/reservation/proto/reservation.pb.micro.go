// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: reservation/proto/reservation.proto

package reservation

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

// Client API for Reservation service

type ReservationService interface {
	InitReservation(ctx context.Context, in *InitReservationRequest, opts ...client.CallOption) (*InitReservationResponse, error)
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...client.CallOption) (*CreateReservationResponse, error)
	ReadReservation(ctx context.Context, in *ReadReservationRequest, opts ...client.CallOption) (*ReadReservationResponse, error)
	UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...client.CallOption) (*UpdateReservationResponse, error)
	DeleteReservation(ctx context.Context, in *DeleteReservationRequest, opts ...client.CallOption) (*DeleteReservationResponse, error)
	GetReservationsForUser(ctx context.Context, in *GetReservationForUserRequest, opts ...client.CallOption) (*GetReservationForUserResponse, error)
	DeleteReservationsForShowing(ctx context.Context, in *DeleteReservationsForShowingRequest, opts ...client.CallOption) (*DeleteReservationsForShowingResponse, error)
}

type reservationService struct {
	c    client.Client
	name string
}

func NewReservationService(name string, c client.Client) ReservationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "reservation"
	}
	return &reservationService{
		c:    c,
		name: name,
	}
}

func (c *reservationService) InitReservation(ctx context.Context, in *InitReservationRequest, opts ...client.CallOption) (*InitReservationResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.InitReservation", in)
	out := new(InitReservationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...client.CallOption) (*CreateReservationResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.CreateReservation", in)
	out := new(CreateReservationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) ReadReservation(ctx context.Context, in *ReadReservationRequest, opts ...client.CallOption) (*ReadReservationResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.ReadReservation", in)
	out := new(ReadReservationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...client.CallOption) (*UpdateReservationResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.UpdateReservation", in)
	out := new(UpdateReservationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) DeleteReservation(ctx context.Context, in *DeleteReservationRequest, opts ...client.CallOption) (*DeleteReservationResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.DeleteReservation", in)
	out := new(DeleteReservationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) GetReservationsForUser(ctx context.Context, in *GetReservationForUserRequest, opts ...client.CallOption) (*GetReservationForUserResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.GetReservationsForUser", in)
	out := new(GetReservationForUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationService) DeleteReservationsForShowing(ctx context.Context, in *DeleteReservationsForShowingRequest, opts ...client.CallOption) (*DeleteReservationsForShowingResponse, error) {
	req := c.c.NewRequest(c.name, "Reservation.DeleteReservationsForShowing", in)
	out := new(DeleteReservationsForShowingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Reservation service

type ReservationHandler interface {
	InitReservation(context.Context, *InitReservationRequest, *InitReservationResponse) error
	CreateReservation(context.Context, *CreateReservationRequest, *CreateReservationResponse) error
	ReadReservation(context.Context, *ReadReservationRequest, *ReadReservationResponse) error
	UpdateReservation(context.Context, *UpdateReservationRequest, *UpdateReservationResponse) error
	DeleteReservation(context.Context, *DeleteReservationRequest, *DeleteReservationResponse) error
	GetReservationsForUser(context.Context, *GetReservationForUserRequest, *GetReservationForUserResponse) error
	DeleteReservationsForShowing(context.Context, *DeleteReservationsForShowingRequest, *DeleteReservationsForShowingResponse) error
}

func RegisterReservationHandler(s server.Server, hdlr ReservationHandler, opts ...server.HandlerOption) error {
	type reservation interface {
		InitReservation(ctx context.Context, in *InitReservationRequest, out *InitReservationResponse) error
		CreateReservation(ctx context.Context, in *CreateReservationRequest, out *CreateReservationResponse) error
		ReadReservation(ctx context.Context, in *ReadReservationRequest, out *ReadReservationResponse) error
		UpdateReservation(ctx context.Context, in *UpdateReservationRequest, out *UpdateReservationResponse) error
		DeleteReservation(ctx context.Context, in *DeleteReservationRequest, out *DeleteReservationResponse) error
		GetReservationsForUser(ctx context.Context, in *GetReservationForUserRequest, out *GetReservationForUserResponse) error
		DeleteReservationsForShowing(ctx context.Context, in *DeleteReservationsForShowingRequest, out *DeleteReservationsForShowingResponse) error
	}
	type Reservation struct {
		reservation
	}
	h := &reservationHandler{hdlr}
	return s.Handle(s.NewHandler(&Reservation{h}, opts...))
}

type reservationHandler struct {
	ReservationHandler
}

func (h *reservationHandler) InitReservation(ctx context.Context, in *InitReservationRequest, out *InitReservationResponse) error {
	return h.ReservationHandler.InitReservation(ctx, in, out)
}

func (h *reservationHandler) CreateReservation(ctx context.Context, in *CreateReservationRequest, out *CreateReservationResponse) error {
	return h.ReservationHandler.CreateReservation(ctx, in, out)
}

func (h *reservationHandler) ReadReservation(ctx context.Context, in *ReadReservationRequest, out *ReadReservationResponse) error {
	return h.ReservationHandler.ReadReservation(ctx, in, out)
}

func (h *reservationHandler) UpdateReservation(ctx context.Context, in *UpdateReservationRequest, out *UpdateReservationResponse) error {
	return h.ReservationHandler.UpdateReservation(ctx, in, out)
}

func (h *reservationHandler) DeleteReservation(ctx context.Context, in *DeleteReservationRequest, out *DeleteReservationResponse) error {
	return h.ReservationHandler.DeleteReservation(ctx, in, out)
}

func (h *reservationHandler) GetReservationsForUser(ctx context.Context, in *GetReservationForUserRequest, out *GetReservationForUserResponse) error {
	return h.ReservationHandler.GetReservationsForUser(ctx, in, out)
}

func (h *reservationHandler) DeleteReservationsForShowing(ctx context.Context, in *DeleteReservationsForShowingRequest, out *DeleteReservationsForShowingResponse) error {
	return h.ReservationHandler.DeleteReservationsForShowing(ctx, in, out)
}
