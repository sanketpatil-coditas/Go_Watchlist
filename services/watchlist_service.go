package service

import (
	"Go_Watchlist/models"
	"Go_Watchlist/repos"
	
)

func AddStockToWatchlist(item model.WatchlistItem) error {
	return repo.AddToWatchlist(item)
}

func RemoveStockFromWatchlist(item model.WatchlistItem) error {
	return repo.RemoveFromWatchlist(item)
}

// func GetUserWatchlist(pan string) ([]string, error) {
// 	return repo.GetWatchlist(pan)
// }
