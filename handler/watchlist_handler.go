package handler

import (
	"Go_Watchlist/models"
	"Go_Watchlist/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToWatchlistHandler(c *gin.Context) {
	var item model.WatchlistItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if item.PAN == "" || len(item.PAN) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PAN is required or Invalid"})
		return
	}
	if err := service.AddStockToWatchlist(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stock added to watchlist"})
}


func RemoveFromWatchlistHandler(c *gin.Context) {
	var item model.WatchlistItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := service.RemoveStockFromWatchlist(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove stock"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stock removed from watchlist"})
}

// func GetWatchlistHandler(c *gin.Context) {
// 	pan := c.Query("pan")
// 	if pan == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "PAN is required"})
// 		return
// 	}
// 	watchlist, err := service.GetUserWatchlist(pan)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch watchlist"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"watchlist": watchlist})
// }
