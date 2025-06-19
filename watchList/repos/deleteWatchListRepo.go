package repo

import (
	"Go_Watchlist/dbConfig"
	"Go_Watchlist/dbConn"
)

type watchlistDeleteRepoImpl struct{}

func WatchlistDeleteRepoInterface() WatchlistDeleteRepository {
	return &watchlistDeleteRepoImpl{}
}

func (r *watchlistDeleteRepoImpl) DeleteWatchlist(userID int64, name string) error {
	return db.DB.
		Where("user_id = ? AND watchlist_name = ?", userID, name).
		Delete(&dbConfig.CreateWatchlist{}).
		Error
}
