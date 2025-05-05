package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/ethereum/handler"
	pb "m3o.com/ethereum/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("ethereum"),
	)

	// Register handler
	pb.RegisterEthereumHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
