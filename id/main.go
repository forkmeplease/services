package main

import (
	"math/rand"
	"time"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/id/handler"
	pb "m3o.com/id/proto"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create service
	srv := service.New(
		service.Name("id"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterIdHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
