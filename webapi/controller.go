package webapi

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/82wutao/ee-rpcdeclare/order"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func orderQuery(c *gin.Context) {

	userID := c.Query("userid")
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("gin request query error %+v\n", err)
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "param error",
			"result": err.Error(),
		})
		return
	}

	req := &order.OrderQueryReq{UserID: id}
	resp, err := order.OrderQuery(c.Request.Context(), req)
	if err != nil {
		log.Printf("order rpc error %+v\n", err)
		c.JSON(200, gin.H{
			"code":   500,
			"msg":    "inner error",
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, resp)
}

type OrderSubmit struct {
	UserID    int
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

	req := &order.OrderSubmitReq{UserID: submit.UserID, OrderParam: nil}
	resp, err := order.OrderSubmit(c.Request.Context(), req)
	if err != nil {
		log.Printf("order rpc error %+v\n", err)
		c.JSON(200, gin.H{
			"code":   500,
			"msg":    "inner error",
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, resp)
}
func orderCancel(c *gin.Context) {
	userID := c.Query("userid")
	id, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("gin request query error %+v\n", err)
		c.JSON(200, gin.H{
			"code":   400,
			"msg":    "param error",
			"result": err.Error(),
		})
		return
	}

	req := &order.OrderCancelReq{UserID: id}
	resp, err := order.OrderCancel(c.Request.Context(), req)
	if err != nil {
		log.Printf("order rpc error %+v\n", err)
		c.JSON(200, gin.H{
			"code":   500,
			"msg":    "inner error",
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, resp)
}
