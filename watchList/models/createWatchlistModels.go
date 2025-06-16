package model

import "time"

type CreateWatchlist struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	UserID        int64     `gorm:"not null;index:idx_user_watchlist,unique"`
	WatchlistName string    `gorm:"not null;index:idx_user_watchlist,unique"`
	LastUpdatedAt time.Time `gorm:"not null"`
	ScripCount    *int64    `gorm:"default:0"`
}

type CreateWatchlistRequest struct {
	UserID        int64  `json:"user_id" binding:"required"`
	WatchlistName string `json:"watchlist_name" binding:"required"`
}


type CreateWatchlistResponse struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	WatchlistName string    `json:"watchlist_name"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
	ScripCount    *int64    `json:"scrip_count,omitempty"` // omit if nil
}