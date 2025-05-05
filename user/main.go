package main

import (
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"go-micro.dev/v5/store"

	otp "m3o.com/otp/proto"
	adminpb "m3o.com/pkg/service/proto"
	"m3o.com/user/handler"
	proto "m3o.com/user/proto"
)

func main() {
	srv := service.New(
		service.Name("user"),
	)
	srv.Init()

	hd := handler.NewUser(
		store.DefaultStore,
		otp.NewOtpService("otp", srv.Client()),
	)

	proto.RegisterUserHandler(srv.Server(), hd)
	adminpb.RegisterAdminHandler(srv.Server(), hd)

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
