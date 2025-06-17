package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "Go_Watchlist/watchlist/models"
    "Go_Watchlist/watchlist/services"
)

func DeleteWatchlistHandler(c *gin.Context) {
    var req model.DeleteWatchlistRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := service.DeleteWatchlist(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    resp := model.DeleteWatchlistResponse{
        Message:         "Watchlist deleted successfully",
        DeletedWatchlist: req.WatchlistName,
    }
    c.JSON(http.StatusOK, resp)
}
