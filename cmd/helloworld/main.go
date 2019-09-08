package main

import (
	"fmt"
	"github.com/gzlj/micro-demo/cmd/helloworld/handlers"
	"github.com/gzlj/micro-demo/cmd/helloworld/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"
)

func main() {

	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.1.71:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("helloworld"),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*15))
	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(handlers.Greeter))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
