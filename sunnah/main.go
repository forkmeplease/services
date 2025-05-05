package main

import (
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/sunnah/handler"
	pb "m3o.com/sunnah/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("sunnah"),
		service.Version("latest"),
	)

	v, err := config.Get("sunnah.api_key")
	if err != nil {
		logger.Fatalf("sunnha.api_key config not found: %v", err)
	}
	key := v.String("")
	if len(key) == 0 {
		logger.Fatal("sunnah.api_key config not found")
	}

	// Register handler
	pb.RegisterSunnahHandler(srv.Server(), handler.New(key))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
