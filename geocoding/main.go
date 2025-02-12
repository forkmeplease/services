package main

import (
	"github.com/micro/services/geocoding/handler"
	pb "github.com/micro/services/geocoding/proto"

	"github.com/micro/micro/v5/service"
	"github.com/micro/micro/v5/service/config"
	"github.com/micro/micro/v5/service/logger"
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
