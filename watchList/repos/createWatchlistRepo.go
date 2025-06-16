package repo

import (
	"Go_Watchlist/watchlist/dbConn"
	"Go_Watchlist/watchlist/models"
	"time"
	"gorm.io/gorm"
)

func CreateWatchlist(w model.CreateWatchlist) error {
	w.LastUpdatedAt = time.Now()
	return db.DB.Create(&w).Error
}
func GetWatchlistByName(userID int64, name string) (*model.CreateWatchlist, error) {
	var watchlist model.CreateWatchlist
	err := db.DB.Where("user_id = ? AND watchlist_name = ?", userID, name).First(&watchlist).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // not found = no duplicate
		}
		return nil, err // real error
	}
	return &watchlist, nil // found = duplicate
}


