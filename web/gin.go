package web

import "github.com/gin-gonic/gin"

func GinMakeRoute() *gin.Engine {
	return gin.New()
}
