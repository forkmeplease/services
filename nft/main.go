package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/nft/handler"
	pb "m3o.com/nft/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("nft"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterNftHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
