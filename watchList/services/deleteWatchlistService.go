package service

import (
    "fmt"
    "Go_Watchlist/watchlist/models"
    "Go_Watchlist/dbConfig"
	"Go_Watchlist/dbConn"
)

func DeleteWatchlist(req model.DeleteWatchlistRequest) error {
    err := db.DB.Where("user_id = ? AND watchlist_name = ?", req.UserID, req.WatchlistName).
        Delete(&dbConfig.CreateWatchlist{}).
        Error
    if err != nil {
        return fmt.Errorf("could not delete watchlist %q for user %d: %v", req.WatchlistName, req.UserID, err)
    }
    return nil
}
