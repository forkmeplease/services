package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/forex/handler"
	pb "m3o.com/forex/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("forex"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterForexHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
