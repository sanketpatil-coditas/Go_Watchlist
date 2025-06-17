package model

import "time"

type GetUserWatchlistsRequest struct {
	UserID int64 `json:"userId" validate:"required"`
}

type WatchlistInfo struct {
	ID            int64     `json:"id"`
	WatchlistName string    `json:"watchlistName"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	ScripCount    *int64    `json:"scripCount,omitempty"`
}

type GetUserWatchlistsResponse struct {
	UserID     int64           `json:"userId"`
	Watchlists []WatchlistInfo `json:"watchlists"`
}
