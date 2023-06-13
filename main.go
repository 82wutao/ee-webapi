package main

import (
	"log"
	"os"
	"strconv"

	"github.com/82wutao/ee-webapi/web"
	"github.com/gin-gonic/gin"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("webapi listen_addr listen_port\n")
		return
	}
	addr := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("get port argument error %+v\n", err)
		return
	}

	route := web.GinMakeRoute()
	route.GET("/order", func(ctx *gin.Context) {
		ctx.JSON(200, os.Args)
	})
	web.Launch(addr, int16(port), route)
	// webapi.Launch(addr, int16(port))
}
