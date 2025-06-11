// package model

//	type WatchlistItem struct {
//		PAN         string `json:"pan"`
//		StockSymbol string `json:"stock_symbol"`
//	}
package model

type WatchlistItem struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	PAN         string `gorm:"not null" json:"pan"`
	StockSymbol string `gorm:"not null" json:"stock_symbol"`
}

