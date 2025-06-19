package dbConfig

import "time"

type AddWatchListItem struct {
	ID          uint   `gorm:"primaryKey"`
	PAN         string `gorm:"not null"`
	StockSymbol string `gorm:"not null"`
}

type CreateWatchlist struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	UserID        int64     `gorm:"not null;index:idx_user_watchlist,unique"`
	WatchlistName string    `gorm:"not null;index:idx_user_watchlist,unique"`
	LastUpdatedAt time.Time `gorm:"not null"`
	ScripCount    *int64    `gorm:"default:0"`
}

type AddUserDB struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	PAN    string `gorm:"unique;not null"`
	Mobile string `gorm:"not null"`
	Age    int    `gorm:"not null"`
	DOB    string `gorm:"not null"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}