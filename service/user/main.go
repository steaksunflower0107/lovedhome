package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"ihome/service/user/handler"
	"ihome/service/user/model"
	user "ihome/service/user/proto/user"
)

func main() {

	// 初始化MySQL连接池
	model.InitDb()
	// 初始化Redis连接池
	model.InitRedis()

	// 初始化consul
	consulReg := consul.NewRegistry()
	// 指定consul服务发现
	model.InitRedis()
	// New Service
	service := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
		micro.Address("10.135.103.247:12342"),
		micro.Registry(consulReg),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
