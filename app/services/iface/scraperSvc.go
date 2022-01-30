package iface

import (
	"context"
	"polx/app/domain/bo"
)

type ScraperSvc interface {
	RunTask(ctx context.Context) error
	GetShills(ctx context.Context, query string) ([]bo.Shill, error)
}
