package install

import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/micro-demo/cmd/pkg/apigateway/handler/hello"
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
	registryHello(s.Engine)
}

func registryHello(e *gin.Engine) {
	e.GET("/hello", hello.Greet)
}
