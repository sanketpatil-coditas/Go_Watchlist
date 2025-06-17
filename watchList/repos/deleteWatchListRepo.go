package repo

import (
	"Go_Watchlist/dbConn"
	"Go_Watchlist/dbConfig"
)

func DeleteWatchlist(userID int64, name string) error {
    return db.DB.
        Where("user_id = ? AND watchlist_name = ?", userID, name).
        Delete(&dbConfig.CreateWatchlist{}).
        Error
}
