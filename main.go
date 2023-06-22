package main

import (
	"fmt"
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

	route := web.GinMakeRoute("html/*")
	route.GET("/order", func(ctx *gin.Context) {
		ctx.JSON(200, os.Args)
	})
	route.GET("/records", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "index.html", nil)
	})
	route.GET("/login", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})
	route.GET("/", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "index.html", nil)
	})
	web.Launch(addr, int16(port), route)
	// webapi.Launch(addr, int16(port))
}
