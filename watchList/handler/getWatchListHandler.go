
// package handler

// import (
// 	"Go_Watchlist/watchlist/models"
// 	"Go_Watchlist/watchlist/services"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetUserWatchlistsHandler(c *gin.Context) {
// 	var req model.GetUserWatchlistsRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
// 		return
// 	}

// 	resp, err := service.GetUserWatchlists(req.UserID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, resp)
// }
package handler

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetWatchlistHandler struct {
	service *service.WatchlistGetService
}

func GetWatchlistHandlerInterface(s *service.WatchlistGetService) *GetWatchlistHandler {
	return &GetWatchlistHandler{service: s}
}
// CreateWatchlist godoc
// @Summary Create a new watchlist
// @Description Creates a new stock watchlist for a user.
// @Tags Watchlist
// @Accept  json
// @Produce  json
// @Param request body model.GetUserWatchlistsRequest true "Get Watchlist Payload"
// @Success 200 {object} model.GetUserWatchlistsResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /watchlist/get [post]
func (h *GetWatchlistHandler) GetUserWatchlistsHandler(c *gin.Context) {
	var req model.GetUserWatchlistsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	resp, err := h.service.GetUserWatchlists(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
