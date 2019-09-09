package main

import (
	"fmt"
	"github.com/gzlj/micro-demo/cmd/helloworld/handlers"
	"github.com/gzlj/micro-demo/cmd/helloworld/proto"
	"github.com/micro/go-micro"
	"time"
)

func main() {

	// no need to create registry object here
	// because we specify registry in command line: --registry consul --registry_address 192.168.35.105:8500
	//reg := consul.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{
	//		"192.168.1.71:8500",
	//	}
	//})
	service := micro.NewService(
		//micro.Registry(reg),
		micro.Name("helloworld"),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*15))
	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(handlers.Greeter))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
