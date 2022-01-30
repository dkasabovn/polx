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
	Ticker      string  `json:"ticker"`
	Position    float32 `json:"position"`
	ShareDelta  float32 `json:"shareDelta"`
	PeakShares  float32 `json:"peakShares"`
	TotalOrders float32 `json:"totalOrders"`

	SenatorInitalPrice   float32 `json:"senatorInititalPrice"`
	SenatorSales         float32 `json:"senatorSales"`
	SenatorTotalSpent    float32 `json:"senatorTotalSpent"`
	SenatorAvgSharePrice float32 `json:"senatorAvgSharePrice"`
	SenatorValue         float32 `json:"senatorValue"`

	RetailInitalPrice   float32 `json:"retailInitialPrice"`
	RetailSales         float32 `json:"retailSales"`
	RetailTotalSpent    float32 `json:"retailTotalSpent"`
	RetailAvgSharePrice float32 `json:"retailAvgSharePrice"`
	RetailValue         float32 `json:"retailValue"`

	CurrentPrice float32   `json:"currentPrice"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
}
