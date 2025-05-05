package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/time/handler"
	pb "m3o.com/time/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("time"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTimeHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
