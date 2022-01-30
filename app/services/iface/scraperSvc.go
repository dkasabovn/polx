package iface

import "context"

type ScraperSvc interface {
	RunTask(ctx context.Context) error
}
