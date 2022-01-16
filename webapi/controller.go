package webapi

import (
	"github.com/gin-gonic/gin"
)

func orderQuery(c *gin.Context) {
	c.JSON(200, "orderQuery")
}
func orderSubmmit(c *gin.Context) {
	c.JSON(200, "orderSubmmit")
}
func orderCancel(c *gin.Context) {
	c.JSON(200, "orderCancel")
}
