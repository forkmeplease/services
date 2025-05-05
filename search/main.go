package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	admin "m3o.com/pkg/service/proto"
	"m3o.com/search/handler"
	pb "m3o.com/search/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("search"),
		service.Version("latest"),
	)

	h := handler.New()
	// Register handler
	pb.RegisterSearchHandler(srv.Server(), h)
	admin.RegisterAdminHandler(srv.Server(), h)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
