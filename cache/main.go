package main

import (
	"m3o.com/cache/handler"
	pb "m3o.com/cache/proto"
	adminpb "m3o.com/pkg/service/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("cache"),
		service.Version("latest"),
	)

	// Register handler
	c := new(handler.Cache)
	pb.RegisterCacheHandler(srv.Server(), c)
	adminpb.RegisterAdminHandler(srv.Server(), c)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
