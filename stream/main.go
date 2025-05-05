package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/stream/handler"
	pb "m3o.com/stream/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("stream"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterStreamHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
