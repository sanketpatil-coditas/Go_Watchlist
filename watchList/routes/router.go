package routes

import (
	"Go_Watchlist/watchlist/constants"
	"Go_Watchlist/watchlist/handler"
	repo "Go_Watchlist/watchlist/repos"
	service "Go_Watchlist/watchlist/services"

	"github.com/gin-gonic/gin"
)

func WatchlistRoutes(r *gin.Engine) {

	repoCreate := repo.WatchlistCreateRepoInterface()
	serviceCreate := service.WatchlistCreateServiceInterface(repoCreate)
	createHandler := handler.CreateWatchlistHandlerInterface(serviceCreate)

	repoDelete := repo.WatchlistDeleteRepoInterface()
	serviceDelete := service.WatchlistDeleteServiceInterface(repoDelete)
	deleteHandler := handler.DeleteWatchlistHandlerInterface(serviceDelete)

	repoGet := repo.WatchlistGetRepoInterface()
	serviceGet := service.WatchlistGetServiceInterface(repoGet)
	getHandler := handler.GetWatchlistHandlerInterface(serviceGet)

	r.POST(constants.RemoveWatchlistRoute, deleteHandler.DeleteWatchlist)
	r.POST(constants.CreateWatchlistRoute, createHandler.CreateWatchlist)
	r.POST(constants.GetUserWatchlistRoute, getHandler.GetUserWatchlistsHandler)

}
