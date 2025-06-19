package repo

import (
	"Go_Watchlist/dbConfig"
	"Go_Watchlist/dbConn"
	"errors"
	"gorm.io/gorm"
)

type watchlistCreateRepoImpl struct{}

func WatchlistCreateRepoInterface() WatchlistCreateRepository {
	return &watchlistCreateRepoImpl{}
}

func (r *watchlistCreateRepoImpl) CreateWatchlist(watchlist *dbConfig.CreateWatchlist) error {
	return db.DB.Create(watchlist).Error
}

func (r *watchlistCreateRepoImpl) GetWatchlistByName(userID int64, name string) (*dbConfig.CreateWatchlist, error) {
	var watchlist dbConfig.CreateWatchlist
	err := db.DB.Where("user_id = ? AND watchlist_name = ?", userID, name).First(&watchlist).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil 
	}

	if err != nil {
		return nil, err
	}
	return &watchlist, nil
}
