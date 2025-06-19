package service

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/repos"
	"Go_Watchlist/dbConfig"
	"time"
	"fmt"
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
	if err := s.repo.CreateWatchlist(&watchlist); err != nil {
		return 0, err
	}
	return watchlist.ID, nil
}
