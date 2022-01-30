package bo

type TradeEntryRaw struct {
	PublicationDate string  `json:"publicationDate"`
	Name            string  `json:"politicianName"`
	Ticker          string  `json:"ticker"`
	TransactionDate string  `json:"transactionDate"`
	Shares          float64 `json:"shares"`
	Price           float64 `json:"sharePrice"`
	TradeType       string  `json:"tradeType"`
}
