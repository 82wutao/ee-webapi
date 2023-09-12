package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoadControllers(route *gin.Engine) {
	route.GET("/login", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.HTML(200, "login.html", nil)
	})
	route.POST("/login", func(c *gin.Context) {
		fmt.Printf(c.Request.URL.RawPath)
		c.JSON(200, gin.H{"name": "ggggggggggggggggggggggggggggggggggggggggg"})
	})
}
