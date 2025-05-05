package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/prayer/handler"
	pb "m3o.com/prayer/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("prayer"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPrayerHandler(srv.Server(), handler.New(srv.Client()))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
