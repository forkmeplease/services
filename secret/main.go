package main

import (
	"m3o.com/secret/handler"
	pb "m3o.com/secret/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("secret"),
	)

	// Register handler
	pb.RegisterSecretHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
