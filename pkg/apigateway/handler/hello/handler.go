package hello

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gzlj/micro-demo/cmd/helloworld/proto"
	r "github.com/gzlj/micro-demo/pkg/apigateway/registry"
	"net/http"
)

var (
	client proto.GreeterService
)

func init() {
	client = proto.NewGreeterService("helloworld", r.ThisService.Client())
}

func Greet(ctx *gin.Context) {
	say := ctx.Query("say")
	resp, err := client.SayHello(context.TODO(), &proto.HelloRequest{
		Name: say,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "helloworld service is not running well.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "helloworld service return: " + resp.Massae,
	})
}
