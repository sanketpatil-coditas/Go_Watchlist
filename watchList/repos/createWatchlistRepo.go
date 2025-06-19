package repo

import (
	"Go_Watchlist/dbConfig"
	"Go_Watchlist/dbConn"
	"log"

	"gorm.io/gorm"
)

type watchlistCreateRepoImpl struct{}

func WatchlistCreateRepoInterface() WatchlistCreateRepository {
	return &watchlistCreateRepoImpl{}
}

func (r *watchlistCreateRepoImpl) CreateWatchlist(watchlist *dbConfig.CreateWatchlist) error {
	log.Println("Inserting watchlist into DB")
	return db.DB.Create(watchlist).Error
}

func (r *watchlistCreateRepoImpl) GetWatchlistByName(userID int64, name string) (*dbConfig.CreateWatchlist, error) {
	log.Printf("watchlist with name '%s' for user %d", name, userID)

	var watchlist dbConfig.CreateWatchlist
	err := db.DB.Where("user_id = ? AND watchlist_name = ?", userID, name).First(&watchlist).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Watchlist not found")
			return nil, nil
		}
		log.Println("DB Error", err.Error())
		return nil, err
	}
	return &watchlist, nil
}
