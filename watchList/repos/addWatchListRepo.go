package repo

import (
	"Go_Watchlist/watchlist/dbConn"
	"Go_Watchlist/watchlist/models"
)

func AddToWatchlist(item model.AddWatchListItem) error {
	return db.DB.Create(&item).Error
}

