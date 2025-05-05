package main

import (
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/youtube/handler"
	pb "m3o.com/youtube/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("youtube"),
		service.Version("latest"),
	)

	// Setup google maps
	c, err := config.Get("google.apikey")
	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}
	apiKey := c.String("")
	if len(apiKey) == 0 {
		logger.Fatalf("Missing required config: google.apikey")
	}

	// Register handler
	pb.RegisterYoutubeHandler(srv.Server(), handler.New(apiKey))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
