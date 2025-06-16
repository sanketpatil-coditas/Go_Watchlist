package handler

import (
	model "Go_Watchlist/userAuthentication/models"
	service "Go_Watchlist/userAuthentication/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// RegisterUserHandler godoc
// @Summary      Register a new user
// @Description  Register a new user by providing name, PAN, mobile, age, and DOB
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserRequest  true  "User registration input"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  map[string]interface{}
// @Router       /register/add [post]
func RegisterUserHandler(c *gin.Context) {
	var req model.AddUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = customErrorMessage(fe)
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.RegisterUser(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := model.AddUserResponse{
		Name:   req.Name,
		PAN:    req.PAN,
		Mobile: req.Mobile,
		Age:    req.Age,
		DOB:    req.DOB,
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully",
		"User": resp})
}
func customErrorMessage(fe validator.FieldError) string {
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
		return "DOB is required and must be in YYYY-MM-DD format"
	default:
		return "Invalid value"
	}
}
