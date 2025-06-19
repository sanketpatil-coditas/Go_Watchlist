// package repo

// import (
// 	db "Go_Watchlist/dbConn"
// 	model "Go_Watchlist/watchlist/models"
// 	"Go_Watchlist/dbConfig"
// )

// func GetWatchlistsByUser(userID int64) ([]model.WatchlistInfo, error) {
// 	var raw []dbConfig.CreateWatchlist
// 	if err := db.DB.Where("user_id = ?", userID).Find(&raw).Error; err != nil {
// 		return nil, err
// 	}

// 	var list []model.WatchlistInfo
// 	for _, wl := range raw {
// 		list = append(list, model.WatchlistInfo{
// 			ID:            wl.ID,
// 			WatchlistName: wl.WatchlistName,
// 			LastUpdatedAt: wl.LastUpdatedAt,
// 			ScripCount:    wl.ScripCount,
// 		})
// 	}
// 	return list, nil
// }
package repo

import (
	db "Go_Watchlist/dbConn"
	model "Go_Watchlist/watchlist/models"
	"Go_Watchlist/dbConfig"
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
