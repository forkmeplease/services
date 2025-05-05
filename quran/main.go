package main

import (
	"m3o.com/quran/handler"
	pb "m3o.com/quran/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("quran"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterQuranHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
