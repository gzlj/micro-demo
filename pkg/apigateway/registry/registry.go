package registry

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"
)

var ThisService micro.Service

func init() {

	fmt.Println("----------------------------package registry init()")
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.1.71:8500",
		}
	})
	ThisService = micro.NewService(
		micro.Registry(reg),
		micro.Name("apigateway"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10))
	ThisService.Init()

	fmt.Println("----------------------------init() ThisService is : ", ThisService)
}

func GatewayService() micro.Service {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.1.71:8500",
		}
	})
	s := micro.NewService(
		micro.Registry(reg),
		micro.Name("apigateway"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10))
	ThisService.Init()
	return s
}
