package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/price/handler"
	pb "m3o.com/price/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("price"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPriceHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
