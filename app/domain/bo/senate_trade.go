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
