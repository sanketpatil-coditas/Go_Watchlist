package main

import (
	"Go_Watchlist/dbConn"
	_ "Go_Watchlist/watchlist/docs"
	"Go_Watchlist/watchlist/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Watchlist API
// @version 1.0
// @description This is an API for managing stock watchlists
// @termsOfService http://example.com/terms/

// @contact.name Sanket Patil
// @contact.email yourname@example.com

// @host localhost:8081
// @BasePath /

func main() {
	db.Connect()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.WatchlistRoutes(r)

	r.Run(":8081")
}
