package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/ping/handler"
	pb "m3o.com/ping/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("ping"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPingHandler(srv.Server(), new(handler.Ping))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
