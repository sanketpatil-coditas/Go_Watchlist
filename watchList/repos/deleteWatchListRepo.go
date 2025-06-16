package repo

import (
	"Go_Watchlist/watchlist/dbConn"
	"Go_Watchlist/watchlist/models"
)

// func AddToWatchlist(item model.AddWatchListItem) error {
// 	return db.DB.Create(&item).Error
// }

func RemoveFromWatchlist(item model.AddWatchListItem) error {
	return db.DB.Where("pan = ? AND stock_symbol = ?", item.PAN, item.StockSymbol).Delete(&model.AddWatchListItem{}).Error
}

func GetWatchlist(pan string) ([]string, error) {
	var results []model.AddWatchListItem
	err := db.DB.Where("pan = ?", pan).Find(&results).Error
	if err != nil {
		return nil, err
	}

	var symbols []string
	for _, item := range results {
		symbols = append(symbols, item.StockSymbol)
	}
	return symbols, nil
}
