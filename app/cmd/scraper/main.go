package main

import (
	"context"
	"polx/app/services/scraper"
	"polx/app/system/log"
)

func main() {
	ctx := context.Background()
	i := 1
	for i < 10 {
		if err := scraper.GetScraperSvc().RunTask(ctx, i); err != nil {
			log.Error(err)
		}
		i++
	}
}
