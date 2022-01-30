package bo

import "time"

type TradeType int

const (
	Buy TradeType = iota
	Sell
)

type TradeEntry struct {
	Id              int
	PublicationDate time.Time
	Name            string
	Ticker          string
	TransactionDate time.Time
	TransactionType TradeType
	Shares          int
	PricePerShare   float32
	Hash            string
}

type Shill struct {
	Id   int    `json:"id",omitempty`
	Name string `json:"name",omitempty`
}

type Bar struct {
	Timestamp    time.Time `json:"t"`
	OpenPrice    float32   `json:"o"`
	HighPrice    float32   `json:"h"`
	LowPrice     float32   `json:"l"`
	ClosePrice   float32   `json:"c"`
	Volume       int       `json:"v"`
	NumberTrades int       `json:"n"`
	VW           int       `json:"vw"`
}

type AlpacaResponse struct {
	Bars []Bar `json:"bars"`
}

type StockResult struct {
	Ticker   string  `json:"ticker"`
	Position float32 `json:"position"`

	SenatorSales         float32 `json:"senatorSales"`
	SenatorTotalSpent    float32 `json:"senatorTotalSpent"`
	SenatorAvgSharePrice float32 `json:"senatorAvgSharePrice"`

	RetailSales         float32 `json:"retailSales"`
	RetailTotalSpent    float32 `json:"retailTotalSpent"`
	RetailAvgSharePrice float32 `json:"retailAvgSharePrice"`

	CurrentPrice float32   `json:"currentPrice"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
}
