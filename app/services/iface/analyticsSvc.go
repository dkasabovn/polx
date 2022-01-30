package iface

import (
	"context"
	"polx/app/domain/bo"
)

type AnalyticsService interface {
	GetShillTrades(ctx context.Context, name string) (map[string]bo.StockResult, error)
}