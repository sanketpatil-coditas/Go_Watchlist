package service

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/repos"
	"fmt"
)

type WatchlistDeleteService struct {
	repo repo.WatchlistDeleteRepository
}

func WatchlistDeleteServiceInterface(r repo.WatchlistDeleteRepository) *WatchlistDeleteService {
	return &WatchlistDeleteService{repo: r}
}

func (s *WatchlistDeleteService) DeleteWatchlist(req model.DeleteWatchlistRequest) error {
	err := s.repo.DeleteWatchlist(req.UserID, req.WatchlistName)
	if err != nil {
		return fmt.Errorf("could not delete watchlist %q for user %d: %v", req.WatchlistName, req.UserID, err)
	}
	return nil
}
