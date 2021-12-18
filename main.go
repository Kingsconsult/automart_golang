package main


import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"automart/database"
	"automart/middleware"
	"routes"
	"go.mongodb.org/mongo-driver/mongo"
)


var manufacturerCollection *mongo.Collection = database.OpenCollection(database.Client, "manufacturer")

func main(){
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.new()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.ManufacturerRoutes(router)

	router.Run(":" + port)
}