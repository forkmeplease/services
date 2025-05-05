package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/dns/handler"
	pb "m3o.com/dns/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("dns"),
	)

	// Register handler
	pb.RegisterDnsHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
