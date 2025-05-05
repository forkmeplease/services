package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/crypto/handler"
	pb "m3o.com/crypto/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("crypto"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterCryptoHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
