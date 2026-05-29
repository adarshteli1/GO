package controller

import (
	"context"
	"net/http"
	"restaurant_management/database"
	"restaurant_management/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodItemCollection *mongo.Collection = database.OpenCollection(database.Client, "fooditem")
var validate = validator.New()

func GetFoodItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		foodItemId := c.Param("fooditem_id")

		var fooditem models.FoodItem
		err := foodItemCollection.FindOne(ctx, bson.M{"fooditem_id": foodItemId}).Decode(&fooditem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FoodItem not found"})
		}
		c.JSON(http.StatusOK, fooditem)
	}
}

func CreateFoodItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
	}
}
