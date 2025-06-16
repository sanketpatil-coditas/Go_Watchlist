package service

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/repos"
	"time"
	"fmt"
)

func CreateWatchlist(req model.CreateWatchlistRequest) error {

	existing, err := repo.GetWatchlistByName(req.UserID, req.WatchlistName)
	if err != nil {
		return err
	}
	if existing != nil {
		return fmt.Errorf("watchlist with name '%s' already exists for this user", req.WatchlistName)
	}
	watchlist := model.CreateWatchlist{
		UserID:        req.UserID,
		WatchlistName: req.WatchlistName,
		LastUpdatedAt: time.Now(),
	}
	return repo.CreateWatchlist(watchlist)
}
