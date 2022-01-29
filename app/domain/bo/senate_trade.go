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
	SharesLow       int
	SharesHigh      int
	PricePerShare   float32
}

type Shill struct {
	Id   int
	Name string
}
