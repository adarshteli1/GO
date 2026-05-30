package controller

import (
	"context"
	"log"
	"net/http"
	"restaurant_management/database"
	"restaurant_management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenuId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro while listinhg the menu"})
			return
		}
		var allmenus []bson.M

		if err = result.All(ctx, &allmenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allmenus)
	}
}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		menu_Id := c.Param("menu_id")

		var menu models.Menu
		err := foodItemCollection.FindOne(ctx, bson.M{"menu_id": menu_Id}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FoodItem not found"})
		}
		c.JSON(http.StatusOK, menu)
	}
}
func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
