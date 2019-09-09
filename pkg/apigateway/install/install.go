package install

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/micro-demo/pkg/apigateway/handler/hello"
	"github.com/gzlj/micro-demo/pkg/apigateway/handler/middleware"
)

func Server() *ApiServer {
	server := &ApiServer{}
	server.Engine = gin.Default()
	server.registryApis()
	return server
}

type ApiServer struct {
	Engine *gin.Engine
}

func (s *ApiServer) registryApis() {
	//1)add public api

	//2)add auth middle software
	s.Engine.Use(middleware.Auth())

	//3)add auth api
	registryHello(s.Engine)
}

func registryHello(e *gin.Engine) {
	e.GET("/hello", hello.Greet)
}




