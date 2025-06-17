
package handler

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserWatchlistsHandler(c *gin.Context) {
	var req model.GetUserWatchlistsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	resp, err := service.GetUserWatchlists(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
