package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/postcode/handler"
	pb "m3o.com/postcode/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("postcode"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPostcodeHandler(srv.Server(), new(handler.Postcode))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
