package main

import (
	"m3o.com/geocoding/handler"
	pb "m3o.com/geocoding/proto"

	"go-micro.dev/v5/config"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"googlemaps.github.io/maps"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("geocoding"),
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
	m, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		logger.Fatalf("Error configuring google maps client: %v", err)
	}

	// Register handler
	pb.RegisterGeocodingHandler(srv.Server(), &handler.Geocoding{Maps: m})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
