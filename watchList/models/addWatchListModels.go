package model

// type AddWatchListItem struct {
// 	ID          uint   `gorm:"primaryKey"`
// 	PAN         string `gorm:"not null"`
// 	StockSymbol string `gorm:"not null"`
// }

type AddWatchListResponse struct {
	Stocks []string `json:"stocks"`
}

type AddWatchListRequest struct {
	PAN         string `json:"pan" binding:"required,len=10"`
	StockSymbol string `json:"stock_symbol" binding:"required"`
}