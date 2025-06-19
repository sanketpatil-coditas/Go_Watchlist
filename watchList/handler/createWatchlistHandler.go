package handler

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/services"
	"net/http"
	"fmt"
	"Go_Watchlist/watchlist/constants"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateWatchlistHandler struct {
	service *service.WatchlistCreateService
}

func CreateWatchlistHandlerInterface(s *service.WatchlistCreateService) *CreateWatchlistHandler {
	return &CreateWatchlistHandler{service: s}
}

var createValidator = validator.New()
// CreateWatchlist godoc
// @Summary Create a new watchlist
// @Description Create a new watchlist for a specific user
// @Tags Watchlist
// @Accept  json
// @Produce  json
// @Param request body model.CreateWatchlistRequest true "Watchlist Request"
// @Success 200 {object} model.CreateWatchlistSuccessResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /watchlist/create [post]
func (h *CreateWatchlistHandler) CreateWatchlist(c *gin.Context) {
var req model.CreateWatchlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := createValidator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateWatchlist(req)
	if err != nil {
		fmt.Println("Service Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.CreateWatchlistSuccessResponse{
		Message: constants.MsgWatchlistSuccess,
		ID:      id,
	})
}
