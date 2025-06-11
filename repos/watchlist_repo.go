package repo

import (
	"Go_Watchlist/dbConn"
	"Go_Watchlist/models"
	
)

func AddToWatchlist(item model.WatchlistItem) error {

	return dbConn.DB.Create(&item).Error
}

func RemoveFromWatchlist(item model.WatchlistItem) error {
	return dbConn.DB.Where("pan = ? AND stock_symbol = ?", item.PAN, item.StockSymbol).Delete(&model.WatchlistItem{}).Error
}

// func GetWatchlist(pan string) ([]string, error) {
// 	var results []model.WatchlistItem
// 	err := dbConn.DB.Where("pan = ?", pan).Find(&results).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	var symbols []string
// 	for _, item := range results {
// 		symbols = append(symbols, item.StockSymbol)
// 	}
// 	return symbols, nil
// }
