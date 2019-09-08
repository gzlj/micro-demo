package main

import (
	r "github.com/gzlj/micro-demo/cmd/gateway/registry"
	"github.com/gzlj/micro-demo/cmd/pkg/apigateway/install"
)

func main() {
	server := install.Server()
	go func() {
		r.ThisService.Run()
	}()
	server.Engine.Run((":8080"))
}
