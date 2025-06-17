package model

// Request payload:
type DeleteWatchlistRequest struct {
    UserID        int64  `json:"userId" validate:"required"`
    WatchlistName string `json:"watchlistName" validate:"required"`
}

// Response payload:
type DeleteWatchlistResponse struct {
    Message         string `json:"message"`
    DeletedWatchlist string `json:"watchlistName"`
}
