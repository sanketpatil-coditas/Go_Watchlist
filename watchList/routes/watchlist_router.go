package routes

import (
	"Go_Watchlist/watchlist/handler"

	"github.com/gin-gonic/gin"
)

func WatchlistRoutes(r *gin.Engine) {
	r.POST("/watchlist/add", handler.AddToWatchlistHandler)
	r.POST("/watchlist/remove", handler.RemoveFromWatchlistHandler)
	r.POST("/watchlist/create", handler.CreateWatchlistHandler)

}
