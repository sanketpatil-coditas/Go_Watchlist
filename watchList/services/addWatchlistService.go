package service

import (
	"Go_Watchlist/watchlist/models"
	repo "Go_Watchlist/watchlist/repos"
)

func AddStockToWatchlist(req model.AddWatchListRequest) error {
	item := model.AddWatchListItem{
		PAN:         req.PAN,
		StockSymbol: req.StockSymbol,
	}
	return repo.AddToWatchlist(item)
}

