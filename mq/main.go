package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/mq/handler"
	pb "m3o.com/mq/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("mq"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMqHandler(srv.Server(), new(handler.Mq))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
