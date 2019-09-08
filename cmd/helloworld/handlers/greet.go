package handlers

import (
	"context"
	"github.com/gzlj/micro-demo/cmd/helloworld/proto"
)

type Greeter struct {
}

//SayHello(context.Context, *HelloRequest, *HelloResponse) error

func (g Greeter) SayHello(ctx context.Context, req *proto.HelloRequest, resp *proto.HelloResponse) error {
	resp.Massae = "Hello " + req.Name
	return nil

}
