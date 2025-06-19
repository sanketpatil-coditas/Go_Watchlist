package service

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/repos"
	"fmt"
	"log"
)

type WatchlistDeleteService struct {
	repo repo.WatchlistDeleteRepository
}

func WatchlistDeleteServiceInterface(r repo.WatchlistDeleteRepository) *WatchlistDeleteService {
	return &WatchlistDeleteService{repo: r}
}

func (s *WatchlistDeleteService) DeleteWatchlist(req model.DeleteWatchlistRequest) error {
	log.Println("Service: Deleting watchlist for user")

	err := s.repo.DeleteWatchlist(req.UserID, req.WatchlistName)
	if err != nil {
		return fmt.Errorf("could not delete watchlist '%s' for user %d: %v", req.WatchlistName, req.UserID, err)
	}
	return nil
}
