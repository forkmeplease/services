package main

import (
	"m3o.com/spam/handler"
	pb "m3o.com/spam/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("spam"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterSpamHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
