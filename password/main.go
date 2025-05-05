package main

import (
	"m3o.com/password/handler"
	pb "m3o.com/password/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("password"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPasswordHandler(srv.Server(), new(handler.Password))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
