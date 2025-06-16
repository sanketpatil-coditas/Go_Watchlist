package handler

import (
	model "Go_Watchlist/userAuthentication/models"
	service "Go_Watchlist/userAuthentication/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveUserHandler(c *gin.Context) {
	var req model.RemoveUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := service.RemoveRegisterUser(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove user"})
		return
	}
	resp := model.RemoveUserResponse{
		Name:   req.Name,
		PAN:    req.PAN,
		Mobile: req.Mobile,
		Age:    req.Age,
		DOB:    req.DOB,
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed successfully",
		"User": resp})
}
