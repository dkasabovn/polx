package definition

import (
	"context"
	"polx/app/domain/bo"
	"time"
)

type ScraperRepo interface {
	BulkInsert(ctx context.Context, entries []bo.TradeEntry) ([]int, error)
	GetTradesByShill(ctx context.Context, shillName string) ([]bo.TradeEntry, error)
	GetShillsTickers(ctx context.Context, shillName string) ([]string, error)
	GetShillsDates(ctx context.Context, shillName string) (time.Time, error)
	GetShills(ctx context.Context, query string) ([]bo.Shill, error)
}
