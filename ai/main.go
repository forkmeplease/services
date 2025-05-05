package main

import (
	"m3o.com/ai/handler"
	pb "m3o.com/ai/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("ai"),
	)

	// Register handler
	pb.RegisterAiHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
