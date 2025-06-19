package model

type DeleteWatchlistRequest struct {
	UserID        int64  `json:"userId" validate:"required" example:"101"`
	WatchlistName string `json:"watchlistName" validate:"required"  example:"Tech Stocks"`
}

type DeleteWatchlistResponse struct {
	Message string `json:"message" example:"Watchlist Deleted successfully"`
}
