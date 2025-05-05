package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/email/handler"
	pb "m3o.com/email/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("email"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterEmailHandler(srv.Server(), handler.NewEmailHandler())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
