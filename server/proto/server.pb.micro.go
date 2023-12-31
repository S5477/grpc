// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/server.proto

package server

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "micro.dev/v4/service/client"
	server "micro.dev/v4/service/server"
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

// Client API for Server service

type ServerService interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type serverService struct {
	c    client.Client
	name string
}

func NewServerService(name string, c client.Client) ServerService {
	return &serverService{
		c:    c,
		name: name,
	}
}

func (c *serverService) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "Server.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterServerHandler(s server.Server, hdlr ServerHandler, opts ...server.HandlerOption) error {
	type server interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
	}
	type Server struct {
		server
	}
	h := &serverHandler{hdlr}
	return s.Handle(s.NewHandler(&Server{h}, opts...))
}

type serverHandler struct {
	ServerHandler
}

func (h *serverHandler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.ServerHandler.Hello(ctx, in, out)
}
