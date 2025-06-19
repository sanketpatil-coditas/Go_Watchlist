package repo

import (
	"Go_Watchlist/dbConn"
	"Go_Watchlist/dbConfig"

)

func AddToWatchlist(item dbConfig.AddWatchListItem) error {
	return db.DB.Create(&item).Error
}

