package manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"net/http"
)


var (
	Service micro.Service
)




func ListRegisteredSvcs(ctx *gin.Context) {
	//say := ctx.Query("say")
	reg := Service.Client().Options().Registry
	services, err := reg.ListServices()
	names := []string{}

	maps := map[string][]string{}

	if err != nil {

	}
	for _,s := range services {
		names = append(names,s.Name)

		nodes := s.Nodes
		fmt.Println(nodes)
		fmt.Println(s.Endpoints)
		addrs := []string{}
		for _, n := range nodes {
			addrs = append(addrs, n.Address)
		}
		ss := *s
		fmt.Println(ss)
		bbb, _ :=reg.GetService(s.Name)
		for hehe := range bbb  {
			fmt.Println(hehe)
		}
		//fmt.Println(*bbb)
		maps[s.Name] = addrs
	}

	ctx.JSON(http.StatusOK, gin.H{
		"services": maps,
	})
}
