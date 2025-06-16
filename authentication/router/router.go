package routes

import (
	"github.com/gin-gonic/gin"
	"Go_Watchlist/userAuthentication/handler"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register/add", handler.RegisterUserHandler)
	r.POST("/register/remove", handler.RemoveUserHandler)
}
