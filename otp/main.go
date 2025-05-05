package main

import (
	"m3o.com/otp/handler"
	pb "m3o.com/otp/proto"
	admin "m3o.com/pkg/service/proto"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("otp"),
	)

	h := new(handler.Otp)
	// Register handler
	pb.RegisterOtpHandler(srv.Server(), h)
	admin.RegisterAdminHandler(srv.Server(), h)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
