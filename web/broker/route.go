package broker

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoadControllers(route *gin.Engine) {

	route.GET("/broker/list", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})
	route.POST("/broker/add", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})
	route.POST("/broker/edit", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})
	route.POST("/broker/cancel", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})

	route.GET("/broker/:idOrPhone/status", func(c *gin.Context) {
		idOrPhone := c.Param("idOrPhone")
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})
	route.GET("/broker/:idOrPhone/customers", func(c *gin.Context) {
		idOrPhone := c.Param("idOrPhone")
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})
	route.GET("/broker/:idOrPhone/subbrokers", func(c *gin.Context) {
		idOrPhone := c.Param("idOrPhone")
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})
}
