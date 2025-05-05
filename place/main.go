package main

import (
	"m3o.com/place/handler"
	pb "m3o.com/place/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("place"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPlaceHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
