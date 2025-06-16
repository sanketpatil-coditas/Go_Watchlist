package main

import (
	"Go_Watchlist/userAuthentication/db"
	"Go_Watchlist/userAuthentication/router"
	_ "Go_Watchlist/userAuthentication/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Watchlist API
// @version 1.0
// @description This is a sample server for user registration.
// @host localhost:8080
// @BasePath /
func main() {
	db.Connect() // make sure your db.Connect() initializes GORM
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterRoutes(router)
	router.Run(":8080")
}
