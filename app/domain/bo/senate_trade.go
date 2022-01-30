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
	Timestamp	 time.Time `json:"t"`
	OpenPrice	 float32   `json:"o"`
	HighPrice	 float32   `json:"h"`
	LowPrice	 float32   `json:"l"`
	ClosePrice	 float32   `json:"c"`
	Volume   	 int       `json:"v"`
	NumberTrades int 	   `json:"n"`
	VW       	 int       `json:"vw"`
}

type AlpacaResponse struct {
	Bars 	[]Bar `json:"bars"`
}

type StockResult struct {
	Ticker        string
	Position	 float32
	
	SenatorSales         float32
	SenatorTotalSpent    float32
	SenatorAvgSharePrice float32

	RetailSales         float32
	RetailTotalSpent    float32
	RetailAvgSharePrice float32

	CurrentPrice  float32
	StartDate     time.Time
	EndDate		  time.Time
}
