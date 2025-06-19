package repo

import (
	"Go_Watchlist/dbConfig"
	"Go_Watchlist/watchlist/models"
)

type WatchlistCreateRepository  interface {
	CreateWatchlist(watchlist *dbConfig.CreateWatchlist) error
	GetWatchlistByName(userID int64, name string) (*dbConfig.CreateWatchlist, error)
}

type WatchlistDeleteRepository interface {
	DeleteWatchlist(userID int64, name string) error
}

type WatchlistGetRepository interface {
	GetWatchlistsByUser(userID int64) ([]model.WatchlistInfo, error)
}