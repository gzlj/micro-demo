package main

import (
	"github.com/gzlj/micro-demo/pkg/apigateway/config"
	"github.com/gzlj/micro-demo/pkg/apigateway/install"
	r "github.com/gzlj/micro-demo/pkg/apigateway/registry"
)

func main() {
	server := install.Server()
	go func() {
		r.ThisService.Run()
	}()

	server.Engine.Run((":" + config.Port))
}
