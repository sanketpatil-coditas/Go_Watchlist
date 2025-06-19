package handler

import (
	"Go_Watchlist/dbConfig"
	"Go_Watchlist/watchlist/constants"
	model "Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/services"
	"log"
	"net/http"

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
// @Failure 400 {object} dbConfig.ErrorResponse
// @Failure 500 {object} dbConfig.ErrorResponse
// @Router /watchlist/create [post]
func (h *CreateWatchlistHandler) CreateWatchlist(c *gin.Context) {
	var req model.CreateWatchlistRequest
	log.Println("Received CreateWatchlist request")

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Validation Error: %v", err)
		c.JSON(http.StatusBadRequest, dbConfig.ErrorResponse{Error: err.Error()})
		return
	}

	if err := createValidator.Struct(req); err != nil {
		log.Println("Validation Failed", err.Error())
		c.JSON(http.StatusBadRequest, dbConfig.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := h.service.CreateWatchlist(req)
	if err != nil {
		log.Println("Service Error", err.Error())
		c.JSON(http.StatusInternalServerError, dbConfig.ErrorResponse{Error: err.Error()})
		return
	}

	log.Printf("Watchlist created successfully with ID %d", id)
	c.JSON(http.StatusOK, model.CreateWatchlistSuccessResponse{
		Message: constants.MsgWatchlistSuccess,
		ID:      id,
	})
}
