package service

import (
	"Go_Watchlist/dbConfig"
	model "Go_Watchlist/watchlist/models"
	repo "Go_Watchlist/watchlist/repos"
	"fmt"
	"time"
)

type WatchlistCreateService struct {
	repo repo.WatchlistCreateRepository
}

func WatchlistCreateServiceInterface(r repo.WatchlistCreateRepository) *WatchlistCreateService {
	return &WatchlistCreateService{repo: r}
}

func (s *WatchlistCreateService) CreateWatchlist(req model.CreateWatchlistRequest) (int64, error) {
	existing, err := s.repo.GetWatchlistByName(req.UserID, req.WatchlistName)
	if err != nil {
		return 0, err
	}
	if existing != nil {
		return 0, fmt.Errorf("watchlist with name '%s' already exists", req.WatchlistName)
	}
	watchlist := dbConfig.CreateWatchlist{
		UserID:        req.UserID,
		WatchlistName: req.WatchlistName,
		LastUpdatedAt: time.Now(),
	}
	err = s.repo.CreateWatchlist(&watchlist)
	if err != nil {
		return 0, err
	}
	return watchlist.ID, nil
}
