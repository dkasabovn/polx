package iface

import (
	"context"
	"polx/app/domain/bo"
)

type ScraperSvc interface {
	RunTask(ctx context.Context, pageNum int) error
	GetShills(ctx context.Context, query string) ([]bo.Shill, error)
	GetShillsAll(ctx context.Context) ([]string, error)
}
