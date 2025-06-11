package handler

import (
	model "Go_Watchlist/models"
	service "Go_Watchlist/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUserHandler godoc
// @Summary      Register a new user
// @Description  Registers user with PAN, name, mobile, DOB, age
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body  model.User  true  "User Info"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /register/add [post]
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

// RemoveRegisterUser godoc
// @Summary      Remove an existing user
// @Description  Removes a user from the system by PAN
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body  model.User  true  "User Info"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /register/remove [post]
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
