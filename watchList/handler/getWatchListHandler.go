package handler

import (
	"Go_Watchlist/watchlist/constants"
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/services"
	"Go_Watchlist/dbConfig"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

type GetWatchlistHandler struct {
	service *service.WatchlistGetService
}

func GetWatchlistHandlerInterface(s *service.WatchlistGetService) *GetWatchlistHandler {
	return &GetWatchlistHandler{service: s}
}

// GetUserWatchlistsHandler godoc
// @Summary Get all watchlists for a user
// @Description Returns all stock watchlists for a given user ID
// @Tags Watchlist
// @Accept  json
// @Produce  json
// @Param request body models.GetUserWatchlistsRequest true "User ID Payload"
// @Success 200 {object} models.GetUserWatchlistsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /watchlist/get [post]
func (h *GetWatchlistHandler) GetUserWatchlistsHandler(c *gin.Context) {
	var req model.GetUserWatchlistsRequest
	log.Println("Received get watchlist request")

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Validation Error", err.Error())

		c.JSON(http.StatusBadRequest, dbConfig.ErrorResponse{Error: constants.MsgInvalidInput + err.Error()})
		return
	}

	resp, err := h.service.GetUserWatchlists(req.UserID)
	if err != nil {
		log.Println("Validation Error", err.Error())

		c.JSON(http.StatusInternalServerError, dbConfig.ErrorResponse{Error: err.Error()})
		return
	}
	log.Println("Watchlist fetched successfully with ID")

	c.JSON(http.StatusOK, resp)
}
