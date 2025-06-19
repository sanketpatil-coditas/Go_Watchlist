package handler

import (
	"Go_Watchlist/watchlist/constants"
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/services"
	"net/http"
	"log"
	"Go_Watchlist/dbConfig"
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
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /watchlist/remove [post]
func (h *DeleteWatchlistHandler) DeleteWatchlist(c *gin.Context) {
	var req model.DeleteWatchlistRequest
	log.Println("Received request to delete watchlist")

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Validation error", err.Error())

		c.JSON(http.StatusBadRequest, dbConfig.ErrorResponse{Error: "Invalid input: " + err.Error()})
		return
	}

	if err := h.service.DeleteWatchlist(req); err != nil {
		log.Println("Service error", err.Error())

		c.JSON(http.StatusInternalServerError, dbConfig.ErrorResponse{Error: err.Error()})
		return
	}
	log.Println("Successfully deleted watchlist for user")

	c.JSON(http.StatusOK, model.DeleteWatchlistResponse{
		Message: constants.MsgWatchlistremove,
	})
}
