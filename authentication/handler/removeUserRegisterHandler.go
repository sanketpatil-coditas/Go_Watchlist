package handler

import (
	model "Go_Watchlist/userAuthentication/models"
	service "Go_Watchlist/userAuthentication/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validateRemoveUser = validator.New()

func RemoveUserHandler(c *gin.Context) {
	var req model.RemoveUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validateRemoveUser.Struct(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = customErrorMessageRemoveUser(fe)
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.RemoveRegisterUser(req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	resp := model.RemoveUserResponse{
		Name:   req.Name,
		PAN:    req.PAN,
		Mobile: req.Mobile,
		Age:    req.Age,
		DOB:    req.DOB,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User removed successfully",
		"User":    resp,
	})
}

func customErrorMessageRemoveUser(fe validator.FieldError) string {
	switch fe.Field() {
	case "Name":
		return "Name is required"
	case "PAN":
		return "PAN must be exactly 10 alphanumeric characters"
	case "Mobile":
		return "Mobile number must be exactly 10 digits"
	case "Age":
		return "Age must be greater than 0"
	case "DOB":
		return "DOB is required"
	default:
		return "Invalid field"
	}
}
