package handler

import (
	"Go_Watchlist/watchlist/models"
	service "Go_Watchlist/watchlist/services"
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Add stock to watchlist
// @Description  Add Favorite stocks in watchlist
// @Tags Watchlist
// @Accept json
// @Produce json
// @Param item body model.WatchlistRequest true "Watchlist Input"
// @Success 200 {object} model.WatchlistResponse
// @Failure 400 {object} map[string]string
// @Router /watchlist/add [post]
func AddToWatchlistHandler(c *gin.Context) {
	var req model.AddWatchListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = "Invalid or missing field"
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.AddStockToWatchlist(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := model.AddWatchListResponse{
		Stocks: []string{req.StockSymbol},
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stock added to watchlist",
		"Watchlist": resp})
}
