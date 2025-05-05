package main

import (
	log "go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"m3o.com/notes/handler"
	pb "m3o.com/notes/proto"
	admin "m3o.com/pkg/service/proto"
)

func main() {
	// New Service
	srv := service.New(
		service.Name("notes"),
		service.Version("latest"),
	)

	// Initialise service
	srv.Init()

	h := handler.New(srv.Client())
	// Register Handler
	pb.RegisterNotesHandler(srv.Server(), h)
	admin.RegisterAdminHandler(srv.Server(), h)

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
