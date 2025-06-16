package handler

import (
	model "Go_Watchlist/watchlist/models"
	service "Go_Watchlist/watchlist/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Remove stock from watchlist
// @Description  Remove Unwanted stocks from Watchlist
// @Tags Watchlist
// @Accept json
// @Produce json
// @Param item body model.WatchlistRequest true "Watchlist Input"
// @Success 200  {object} model.WatchlistResponse
// @Failure 400 {object} map[string]string
// @Router /watchlist/remove [post]
func RemoveFromWatchlistHandler(c *gin.Context) {
	var req model.AddWatchListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := service.RemoveStockFromWatchlist(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove stock"})
		return
	}
	resp := model.AddWatchListResponse{
		Stocks: []string{req.StockSymbol},
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stock removed from watchlist",
		"Watchlist": resp})
}

