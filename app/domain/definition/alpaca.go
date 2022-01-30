package definition

import (
	"polx/app/domain/bo"
)

type AlpacaRepo interface {
	GetBars(ticker, startDate string) (*bo.AlpacaResponse, error)
}
