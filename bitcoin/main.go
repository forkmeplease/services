package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/bitcoin/handler"
	pb "m3o.com/bitcoin/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("bitcoin"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterBitcoinHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
