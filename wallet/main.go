package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/wallet/handler"
	pb "m3o.com/wallet/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("wallet"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterWalletHandler(srv.Server(), handler.NewHandler())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
