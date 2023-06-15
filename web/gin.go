package web

import (
	"github.com/gin-gonic/gin"
)

func GinMakeRoute(filePathPattern string) *gin.Engine {
	engin := gin.New()
	engin.LoadHTMLGlob(filePathPattern)
	return engin
}
