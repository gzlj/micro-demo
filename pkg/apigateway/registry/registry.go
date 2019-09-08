package registry

import (
	"fmt"
	"github.com/gzlj/micro-demo/pkg/apigateway/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"
)

var ThisService micro.Service

func init() {

	fmt.Println("----------------------------package registry init()")
	reg := consul.NewRegistry(func(options *registry.Options) {
		url := config.RegistryUrl
		options.Addrs = []string{
			url,
		}
	})
	ThisService = micro.NewService(
		micro.Registry(reg),
		micro.Name("apigateway"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Flags(
			cli.StringFlag{
				Name:  "port",
				Value: "8080",
				Usage: "service port",
			},
		))

	ThisService.Init(
		micro.Action(func(c *cli.Context) {
			config.Port = c.String("port")
		}))

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
