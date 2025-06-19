package service

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/dbConfig"
	"Go_Watchlist/watchlist/repos"
)

func AddStockToWatchlist(req model.AddWatchListRequest) error {
	item := dbConfig.AddWatchListItem{
		PAN:         req.PAN,
		StockSymbol: req.StockSymbol,
	}
	return repo.AddToWatchlist(item)
}

