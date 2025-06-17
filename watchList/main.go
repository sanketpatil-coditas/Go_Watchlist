package main

import (
	"Go_Watchlist/dbConn"
	_ "Go_Watchlist/watchlist/docs"
	// "Go_Watchlist/utils/postgres"

	// "Go_Watchlist/watchlist/handler"
	"Go_Watchlist/watchlist/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Watchlist Service
// @version 1.0
// @description Manage user stock watchlist
// @host localhost:8081
// @BasePath /
func main() {
	db.Connect()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.WatchlistRoutes(r)

	r.Run(":8081")
}
