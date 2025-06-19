package repo

import (
	db "Go_Watchlist/dbConn"
	"Go_Watchlist/dbConfig"
	"Go_Watchlist/watchlist/models"
)

type watchlistGetRepoImpl struct{}

func WatchlistGetRepoInterface() WatchlistGetRepository {
	return &watchlistGetRepoImpl{}
}

func (r *watchlistGetRepoImpl) GetWatchlistsByUser(userID int64) ([]model.WatchlistInfo, error) {
	var raw []dbConfig.CreateWatchlist
	if err := db.DB.Where("user_id = ?", userID).Find(&raw).Error; err != nil {
		return nil, err
	}

	var list []model.WatchlistInfo
	for _, wl := range raw {
		list = append(list, model.WatchlistInfo{
			ID:            wl.ID,
			WatchlistName: wl.WatchlistName,
		})
	}
	return list, nil
}
