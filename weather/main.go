package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/weather/handler"
	pb "m3o.com/weather/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("weather"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterWeatherHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
