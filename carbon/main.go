package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/carbon/handler"
	pb "m3o.com/carbon/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("carbon"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterCarbonHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
