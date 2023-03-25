package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"house/handler"
	"house/model"
	"house/subscriber"

	house "house/proto/house"
)

func main() {
	// 初始化MySQL连接池
	model.InitDb()
	// 初始化Redis连接池
	model.InitRedis()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.house"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	house.RegisterHouseHandler(service.Server(), new(handler.House))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.house", service.Server(), new(subscriber.House))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.house", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
