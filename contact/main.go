package main

import (
	"go-micro.dev/v5/store"
	admin "m3o.com/pkg/service/proto"

	"m3o.com/contact/domain"
	"m3o.com/contact/handler"
	pb "m3o.com/contact/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("contact"),
		service.Version("latest"),
	)

	contactDomain := domain.NewContactDomain(store.DefaultStore)

	h := handler.NewContact(contactDomain)
	// Register handler
	pb.RegisterContactHandler(srv.Server(), h)
	admin.RegisterAdminHandler(srv.Server(), h)
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
