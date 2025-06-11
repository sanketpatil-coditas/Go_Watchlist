package handler

import (
	"Go_Watchlist/models"
	"Go_Watchlist/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := service.RegisterUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
func RemoveRegisterUser(c *gin.Context) {
	var item model.User
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := service.RemoveRegisterUser(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove stock"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stock removed from watchlist"})
}
