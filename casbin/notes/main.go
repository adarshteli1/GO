package main

import (
	"log"
	"notes/casbin"
	"notes/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	casbin.InitCasbin()
	routes.Routes(router)

	router.Run(":" + port)
}
