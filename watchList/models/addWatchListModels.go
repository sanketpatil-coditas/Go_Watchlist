package model

type AddWatchListResponse struct {
	Stocks []string `json:"stocks"`
}

type AddWatchListRequest struct {
	PAN         string `json:"pan" binding:"required,len=10"`
	StockSymbol string `json:"stock_symbol" binding:"required"`
}