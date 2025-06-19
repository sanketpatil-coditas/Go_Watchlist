package handler

import (
	"Go_Watchlist/userAuthentication/constants"
	model "Go_Watchlist/watchlist/models"
	service "Go_Watchlist/watchlist/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteWatchlistHandler struct {
	service *service.WatchlistDeleteService
}

func DeleteWatchlistHandlerInterface(s *service.WatchlistDeleteService) *DeleteWatchlistHandler {
	return &DeleteWatchlistHandler{service: s}
}

// DeleteWatchlist godoc
// @Summary Delete a watchlist
// @Description Delete a watchlist for a specific user by name
// @Tags Watchlist
// @Accept  json
// @Produce  json
// @Param request body model.DeleteWatchlistRequest true "Delete Watchlist Request"
// @Success 200 {object} model.DeleteWatchlistResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /watchlist/remove [post]
func (h *DeleteWatchlistHandler) DeleteWatchlist(c *gin.Context) {
	var req model.DeleteWatchlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.DeleteWatchlist(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.DeleteWatchlistResponse{
		Message: constants.MsgUserRegisterremove,
	})
}
