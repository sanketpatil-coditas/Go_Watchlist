package service

import (
	"Go_Watchlist/dbConfig"
	model "Go_Watchlist/watchlist/models"
	repo "Go_Watchlist/watchlist/repos"
	"fmt"
	"log"
	"time"
)

type WatchlistCreateService struct {
	repo repo.WatchlistCreateRepository
}

func WatchlistCreateServiceInterface(r repo.WatchlistCreateRepository) *WatchlistCreateService {
	return &WatchlistCreateService{repo: r}
}

func (s *WatchlistCreateService) CreateWatchlist(req model.CreateWatchlistRequest) (int64, error) {
	log.Printf("Checking if watchlist '%s' already exists for user %d", req.WatchlistName, req.UserID)

	existing, err := s.repo.GetWatchlistByName(req.UserID, req.WatchlistName)
	if err != nil {
		log.Printf("checking existing watchlist: %v", err)
		return 0, err
	}
	if existing != nil {
		msg := fmt.Sprintf("watchlist with name '%s' already exists", req.WatchlistName)
		log.Println(msg)
		return 0, fmt.Errorf(msg)
	}

	watchlist := dbConfig.CreateWatchlist{
		UserID:        req.UserID,
		WatchlistName: req.WatchlistName,
		LastUpdatedAt: time.Now(),
	}

	log.Println("Creating new watchlist record")
	if err := s.repo.CreateWatchlist(&watchlist); err != nil {
		log.Printf("Error creating watchlist: %v", err)
		return 0, err
	}
	return watchlist.ID, nil
}
