package main

import (
	"time"

	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/service"
	"go-micro.dev/v5/store"
	admin "m3o.com/pkg/service/proto"

	"m3o.com/rss/handler"
	pb "m3o.com/rss/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("rss"),
	)

	st := store.DefaultStore
	crawl := handler.NewCrawl(st)
	rss := handler.NewRss(st, crawl)

	// crawl
	go func() {
		crawl.FetchAll()
		tick := time.NewTicker(1 * time.Minute)
		for _ = range tick.C {
			crawl.FetchAll()
		}
	}()

	// Register handler
	pb.RegisterRssHandler(srv.Server(), rss)
	admin.RegisterAdminHandler(srv.Server(), rss)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
