package routers

import (
"github.com/gin-gonic/gin"
"Tugas2golang-gin/controllers"
)

func GetRoutes() *gin.Engine {
	routers := gin.Default()
	routers.POST("/order", controllers.CreateOrder)
	routers.GET("/order", controllers.GetOrderAll)
	routers.PUT("/order/{orderID}",controllers.UpdateOrder)
	routers.DELETE("/order/{orderID}",controllers.DeleteOrder)
	return routers


}