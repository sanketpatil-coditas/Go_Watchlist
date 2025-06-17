package routes

import (
	"Go_Watchlist/watchlist/constants"
	"Go_Watchlist/watchlist/handler"

	"github.com/gin-gonic/gin"
)

func WatchlistRoutes(r *gin.Engine) {
	// r.POST("/watchlist/add", handler.AddToWatchlistHandler)
	// r.POST(constants.RemoveWatchlistRoute, handler.RemoveFromWatchlistHandler)
	r.POST(constants.RemoveWatchlistRoute, handler.DeleteWatchlistHandler)

	r.POST(constants.CreateWatchlistRoute, handler.CreateWatchlistHandler)
	r.POST(constants.GetUserWatchlistRoute, handler.GetUserWatchlistsHandler)

}
