package model

type GetUserWatchlistsRequest struct {
	UserID int64 `json:"userId" validate:"required" example:"101"`
}

type WatchlistInfo struct {
	ID            int64  `json:"id"`
	WatchlistName string `json:"watchlistName"`
}

type GetUserWatchlistsResponse struct {
	UserID     int64           `json:"userId"`
	Watchlists []WatchlistInfo `json:"watchlists"`
}
