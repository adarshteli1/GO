package routes

import (
	"notes/controllers"
	"notes/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(incomingrRoutes *gin.Engine) {

	incomingrRoutes.POST("/register", controllers.Register())
	incomingrRoutes.GET("/login", controllers.Login())

	protected := incomingrRoutes.Group("/")

	protected.Use(middleware.AuthMiddleware())
	protected.Use(middleware.CasbinMiddleware())
	protected.POST("/signup", controllers.Signup())
	protected.POST("/notes", controllers.Notes())
}
