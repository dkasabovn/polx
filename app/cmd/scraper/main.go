package main

import (
	"context"
	"polx/app/services/scraper"
	"polx/app/system/log"
)

func main() {
	ctx := context.Background()
	log.InfoStruct("Running svc")
	if err := scraper.GetScraperSvc().RunTask(ctx); err != nil {
		log.Error(err)
	}
}
