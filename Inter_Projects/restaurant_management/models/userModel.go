package models

import "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.EngineE) {

	incomingRoutes.GET("/user", controller.GetUser())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.POST("/login", controller.Login())
	incomingRoutes.GET("/signup", controller.Signup())

}
