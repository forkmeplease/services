package main

import (
	"m3o.com/cron/handler"
	pb "m3o.com/cron/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("cron"),
	)

	// Register handler
	pb.RegisterCronHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
