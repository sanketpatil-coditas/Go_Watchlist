package handler

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/services"
	"net/http"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)
var validate = validator.New()
func CreateWatchlistHandler(c *gin.Context) {
	var req model.CreateWatchlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateWatchlist(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Watchlist created successfully"})
}
