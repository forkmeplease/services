package main

import (
	"m3o.com/translate/handler"
	pb "m3o.com/translate/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("translate"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTranslateHandler(srv.Server(), handler.NewTranslation())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
