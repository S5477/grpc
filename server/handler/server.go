package handler

import (
	"context"

	log "micro.dev/v4/service/logger"

	pb "server/proto"
)

type Server struct{}

// Return a new handler
func New() *Server {
	return &Server{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Server) Hello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloResponse) error {
	log.Info("Received Server.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
