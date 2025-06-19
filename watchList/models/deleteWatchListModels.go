package model

type DeleteWatchlistRequest struct {
    UserID        int64  `json:"userId" validate:"required"`
    WatchlistName string `json:"watchlistName" validate:"required"`
}

type DeleteWatchlistResponse struct {
    Message         string `json:"message"`
}
