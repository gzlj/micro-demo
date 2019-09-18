package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/micro-demo/pkg/apigateway/handler/hello"
	"github.com/gzlj/micro-demo/pkg/apigateway/handler/manage"

	//r "github.com/gzlj/micro-demo/pkg/apigateway/registry"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"time"
)

func main() {
	//server := install.Server()
	//go func() {
	//	r.ThisService.Run()
	//}()


	//server := install.Server()
	//go func() {
	//	server.Engine.Run()
	//}()

	go RunHttp()

	//r.ThisService.Run()

	service := micro.NewService(
		//micro.Registry(reg),
		micro.Name("gateway"),
		micro.RegisterTTL(time.Second*60),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init()
	hello.Init(service)
	manage.Service = service
	if err := service.Run(); err != nil {
	}

}

func Router() *gin.Engine{
	r := gin.Default()
	r.GET("/hello", hello.Greet)
	r.GET("/services", manage.ListRegisteredSvcs)
	return r
}

func RunHttp() {
	r := Router()
	r.Run(":8080")
}
