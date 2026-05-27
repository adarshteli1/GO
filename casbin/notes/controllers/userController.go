package controllers

import (
	"net/http"
	"notes/helpers"
	"notes/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// var userCollection *mongo.Collection = database.openCollection(database.Client, "user")
var validate = validator.New()
var Users []models.User

func VerifyPassword(userPass string, providedPass string) (bool, string) {

	if userPass != providedPass {
		return false, "Invalid Password"
	}
	return true, ""

}
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// defer cancel()
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := validate.Struct(user)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User"})
			return
		}
		// _, insertErr := userCollection.InsertOne(ctx, user)

		// if insertErr != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		// 	return
		// }
		Users = append(Users, user)

		c.JSON(http.StatusOK, gin.H{"message": "User Regstered"})

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "no user found"})
			return
		}

		for _, u := range Users {
			if u.Username == user.Username {
				foundUser = u
				break
			}
		}

		if foundUser.Username == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "User Not found"})
			return
		}

		passwordIsValid, _ := VerifyPassword(user.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Password"})
			return
		}
		// if foundUser.Email == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"messgae": "User not found"})
		// }
		token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.Username, foundUser.Email, foundUser.Role)
		// helpers.UpdateTokens(token, refreshToken, foundUser.Username)

		// err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// }

		c.JSON(http.StatusOK, gin.H{
			"token":        token,
			"refreshToken": refreshToken,
		})

	}

}

func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "signup route",
		})

	}
}

func Notes() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "notes route",
		})

	}
}
