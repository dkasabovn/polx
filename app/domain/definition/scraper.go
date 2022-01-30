package definition

import (
	"context"
	"polx/app/domain/bo"
)

type ScraperRepo interface {
	BulkInsert(ctx context.Context, entries []bo.TradeEntry) error
	GetTradesByShill(ctx context.Context, shillName string) ([]bo.TradeEntry, error)
}
