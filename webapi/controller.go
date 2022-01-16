package webapi

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func orderQuery(c *gin.Context) {
	c.JSON(200, "orderQuery")
}

type OrderSubmit struct {
	TradePair int
	Side      int
	Price     decimal.Decimal
	Vol       decimal.Decimal
}

func (os OrderSubmit) toString() string {
	return fmt.Sprintf("TradePair:%d,Side:%d,Price:%s,Vol:%s",
		os.TradePair, os.Side, os.Price.String(), os.Vol.String())
}

func orderSubmmit(c *gin.Context) {
	var submit OrderSubmit
	if err := c.Bind(&submit); err != nil {
		log.Printf("gin bind request error %+v\n", err)
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "param error",
			"result": err.Error(),
		})
		return
	}
	log.Printf("gin bind request is %+v\n", submit)

	ctx := c.Request.Context()
	key := submit.toString()
	duration := time.Minute * 5
	suc, err := redisTryAcquire(ctx, key, "existed", duration)
	if err != nil {
		log.Printf("redis try setnx error %+v\n", err)

		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "redis has errror",
		})
		return
	}
	if !suc {
		log.Printf("redis try setnx ,but existed \n")

		c.JSON(200, gin.H{
			"code": 409,
			"msg":  "redis has errror",
		})
		return
	}

	c.JSON(200, "orderSubmmit")
}
func orderCancel(c *gin.Context) {
	c.JSON(200, "orderCancel")
}
