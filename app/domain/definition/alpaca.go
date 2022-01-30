package definition

import (
	"polx/app/domain/bo"
)

type AlpacaRepo interface {
	GetBars(ticker, startDate, endDate string ) (*bo.AlpacaResponse, error) 
}
