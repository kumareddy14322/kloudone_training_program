package router

import (
	"github.com/gin-gonic/gin"
	
	"github.com/kumareddy14322/kloudone_training_program/payment_api/src/controller"
)

func ChargeRouter() *gin.Engine {
	r := gin.Default()
	g1 := r.Group("/stripe")
	{
		g1.POST("payment", controller.Payment)
	}
	return r
}
