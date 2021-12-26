package controller

import (
	"automart/database"
	"automart/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var carCollection *mongo.Collection = database.OpenCollection(database.Client, "car")
var validate = validator.New()

func GetCars() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := carCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching cars"})
		}
		var allCars []bson.M
		if err = result.All(ctx, &allCars); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allCars)
	}
}

func GetCar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		carId := c.Param("car_id")
		var car models.Car

		err := carCollection.FindOne(ctx, bson.M{"car_id": carId}).Decode(&car)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the car item"})
		}
		c.JSON(http.StatusOK, car)
	}
}

func CreateCar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var manufacturer models.Manufacturer
		var car models.Car

		if err := c.BindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(car)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		err := manufacturerCollection.FindOne(ctx, bson.M{"manufacturer_id": car.Manufacturer_id}).Decode(&car)
		defer cancel()

		if err != nil {
			msg := fmt.Sprintf("manufacturer was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		car.Created_at, _ = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		car.Updated_at, _ = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		car.ID = primitive.NewObjectID()
		var num = toFixed(*car.Price, 2)
		car.Price = &num
		// car.Name =

		result, insertErr := carCollection.InsertOne(ctx, car)
		if insertErr != nil {
			msg := fmt.Sprintf("car was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateCar() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
