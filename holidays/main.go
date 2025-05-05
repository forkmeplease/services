package main

import (
	"m3o.com/holidays/handler"
	pb "m3o.com/holidays/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("holidays"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterHolidaysHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
