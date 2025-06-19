package model

type CreateWatchlistRequest struct {
	UserID        int64  `json:"userId" validate:"required" example:"101"`
	WatchlistName string `json:"watchlistName" validate:"required" example:"Tech Stocks"`
}

type CreateWatchlistSuccessResponse struct {
	Message string `json:"message" example:"Watchlist created successfully"`
	ID      int64  `json:"id" example:"1"`
}
