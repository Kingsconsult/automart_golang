package routes

import (
	controller "automart/controllers"
	"github.com/gin-gonic/gin"
)

func CarRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/cars", controller.GetCars())
	incomingRoutes.GET("/cars/:car_id", controller.GetCar())
	incomingRoutes.POST("/cars", controller.CreateCar())
	incomingRoutes.PATCH("/cars/:car_id", controller.UpdateCar())
}
