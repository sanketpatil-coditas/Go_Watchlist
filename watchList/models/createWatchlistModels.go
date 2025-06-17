package model

import "time"


type CreateWatchlistRequest struct {
	UserID        int64  `json:"userId" validate:"required"`
	WatchlistName string `json:"watchlistName" validate:"required"`
}


type CreateWatchlistResponse struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"userId"`
	WatchlistName string    `json:"watchlistName"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	ScripCount    *int64    `json:"scripCount,omitempty"`
}