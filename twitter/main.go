package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/twitter/handler"
	pb "m3o.com/twitter/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("twitter"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTwitterHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
