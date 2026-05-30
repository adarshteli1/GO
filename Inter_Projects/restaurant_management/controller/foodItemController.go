package controller

import (
	"context"
	"fmt"
	"net/http"
	"restaurant_management/database"
	"restaurant_management/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodItemCollection *mongo.Collection = database.OpenCollection(database.Client, "fooditem")
var validate = validator.New()

func GetFoodItemId() gin.HandlerFunc {
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
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		var foodItem models.FoodItem

		if err := c.BindJSON(&foodItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding with JSON"})
		}

		validateErr := validate.Struct(foodItem)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		}

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": foodItem.Menu_Id}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FoodITem Not Found	"})
		}

		foodItem.Created_at, _ = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		foodItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		foodItem.ID = primitive.NewObjectID()
		foodItem.FoodItem_ID = foodItem.ID.Hex()
		var num = toFixed(*foodItem.FoodItem_Price, 2)
		foodItem.FoodItem_Price = &num

		result, inserErr := foodItemCollection.InsertOne(ctx, foodItem)
		if inserErr != nil {
			msg := fmt.Sprintf("FoodItem Not Inserted")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetFoodItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}

}

func UpdateFoodItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
