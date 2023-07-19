package main

import (
	"server/handler"
	pb "server/proto"

	"micro.dev/v4/service"
	"micro.dev/v4/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("server"),
	)

	// Register handler
	pb.RegisterServerHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
