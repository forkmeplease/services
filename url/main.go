package main

import (
	admin "m3o.com/pkg/service/proto"
	"m3o.com/url/handler"
	pb "m3o.com/url/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("url"),
		service.Version("latest"),
	)
	h := handler.NewUrl()
	// Register handler
	pb.RegisterUrlHandler(srv.Server(), h)
	admin.RegisterAdminHandler(srv.Server(), h)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
