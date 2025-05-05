package main

import (
	"m3o.com/news/handler"
	pb "m3o.com/news/proto"

	"go-micro.dev/v5/config"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("news"),
		service.Version("latest"),
	)

	// Setup google maps
	c, err := config.Get("news.apikey")
	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
	}

	apiKey := c.String("")
	if len(apiKey) == 0 {
		logger.Fatalf("Missing required config: news.apikey")
	}

	// Register handler
	pb.RegisterNewsHandler(srv.Server(), handler.New(apiKey))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
