package routes

import (
	controller "automart/controllers"
	"github.com/gin-gonic/gin"
)

func ManufacturerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/manufacturers", controller.GetManufacturers())
	incomingRoutes.GET("/manufacturers/:manufacturer_id", controller.GetManufacturer())
	incomingRoutes.POST("/manufacturers", controller.CreateManufacturer())
	incomingRoutes.PATCH("/manufacturers/:manufacturer_id", controller.UpdateManufacturer())
}
