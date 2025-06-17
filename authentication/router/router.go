package routes

import (
	"github.com/gin-gonic/gin"
	"Go_Watchlist/userAuthentication/handler"
	"Go_Watchlist/userAuthentication/constants"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST(constants.AddUserRoute, handler.RegisterUserHandler)
	r.POST(constants.RemoveUserRoute, handler.RemoveUserHandler)
}
