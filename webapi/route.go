package webapi

import "github.com/gin-gonic/gin"

func getRoute() *gin.Engine {
	route := gin.Default()

	route.GET("/order/query", orderQuery)
	route.POST("/order/submit", orderSubmmit)
	route.DELETE("/order/cancel", orderCancel)
	return route
}
