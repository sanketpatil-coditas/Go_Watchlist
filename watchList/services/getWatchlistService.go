// package service

// import (
// 	"Go_Watchlist/watchlist/models"
	
// 	"Go_Watchlist/watchlist/repos"
// )


// func GetUserWatchlists(userID int64) (model.GetUserWatchlistsResponse, error) {
// 	watchlists, err := repo.GetWatchlistsByUser(userID)
// 	if err != nil {
// 		return model.GetUserWatchlistsResponse{}, err
// 	}
// 	return model.GetUserWatchlistsResponse{
// 		UserID:     userID,
// 		Watchlists: watchlists,
// 	}, nil
// }
package service

import (
	"Go_Watchlist/watchlist/models"
	"Go_Watchlist/watchlist/repos"
)

type WatchlistGetService struct {
	repo repo.WatchlistGetRepository
}

func WatchlistGetServiceInterface(r repo.WatchlistGetRepository) *WatchlistGetService {
	return &WatchlistGetService{repo: r}
}

func (s *WatchlistGetService) GetUserWatchlists(userID int64) (model.GetUserWatchlistsResponse, error) {
	watchlists, err := s.repo.GetWatchlistsByUser(userID)
	if err != nil {
		return model.GetUserWatchlistsResponse{}, err
	}
	return model.GetUserWatchlistsResponse{
		UserID:     userID,
		Watchlists: watchlists,
	}, nil
}
