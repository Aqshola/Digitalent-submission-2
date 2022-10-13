package routers

import (
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/orders", controllers.GetOrder)
	router.GET("/orders/:id", controllers.GetOrderDetail)
	router.POST("/orders", controllers.CreateOrder)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)
	return router
}
