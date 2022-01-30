package analytics

import (
	"context"
	"math"
	"polx/app/datastore/alpaca"
	"polx/app/datastore/pg"
	"time"

	"polx/app/domain/bo"
	"polx/app/domain/definition"
	"polx/app/services/iface"
	"sync"
)

var (
	analyticsOnce sync.Once
	userServiceInstance *analyticsService
)

type analyticsService struct {
	tradeRepo definition.ScraperRepo
	alpacaRepo definition.AlpacaRepo
}

func GetAnalyticsService() iface.AnalyticsService {
	analyticsOnce.Do(func() {
		userServiceInstance = &analyticsService{
			tradeRepo: pg.GetScraperRepo(),
			alpacaRepo: alpaca.GetAlpacaRepo(),
		}
	})
	return userServiceInstance
}

func (a *analyticsService) GetShillTrades(ctx context.Context, shillName string) (map[string]bo.StockResult, error) {

	tickers, err := a.tradeRepo.GetShillsTickers(ctx, shillName)
	if err != nil {
		return nil, err
	}

	trades, err := a.tradeRepo.GetTradesByShill(ctx, shillName)
	if err != nil {
		return nil, err
	}

	startDate, endDate, err := a.tradeRepo.GetShillsDates(ctx, shillName)
	if err != nil {
		return nil, err
	}

	alpacaData := make(map[string][]bo.Bar)
	for _, tick := range tickers {
		data, err := a.alpacaRepo.GetBars(tick, startDate.Format(time.RFC3339), endDate.Format(time.RFC3339))
		if err != nil{ 
			return nil, err
		}

		alpacaData[tick] = data.Bars
	}


	stockResults := make(map[string]bo.StockResult)
	for _, senatorTrade := range trades {
		currentStatus, ok := stockResults[senatorTrade.Ticker]

		if !ok {
			currentStatus = bo.StockResult{}
			currentStatus.Ticker = senatorTrade.Ticker
			currentStatus.StartDate = senatorTrade.TransactionDate
		}

		publicationIndex := int(senatorTrade.PublicationDate.Sub(startDate).Hours()/24)
		retailTrade := alpacaData[senatorTrade.Ticker][publicationIndex]

		if senatorTrade.TransactionType == bo.Buy {
			currentStatus.SenatorAvgSharePrice = (currentStatus.Position * currentStatus.SenatorAvgSharePrice + senatorTrade.PricePerShare * float32(senatorTrade.Shares)) / (float32(senatorTrade.Shares) + currentStatus.Position)
			currentStatus.Position += float32(senatorTrade.Shares)
			currentStatus.SenatorTotalSpent += senatorTrade.PricePerShare * float32(senatorTrade.Shares)

			currentStatus.RetailAvgSharePrice = (currentStatus.Position * currentStatus.RetailAvgSharePrice + retailTrade.ClosePrice * float32(senatorTrade.Shares)) / (float32(senatorTrade.Shares) + currentStatus.Position)
			currentStatus.Position += float32(senatorTrade.Shares)
			currentStatus.RetailTotalSpent += retailTrade.ClosePrice * float32(senatorTrade.Shares)

		} else {
			currentStatus.SenatorSales += (senatorTrade.PricePerShare - float32(currentStatus.SenatorAvgSharePrice)) * float32(senatorTrade.Shares) 
			currentStatus.RetailSales += (retailTrade.ClosePrice - float32(currentStatus.RetailAvgSharePrice)) * float32(senatorTrade.Shares) 
			
			currentStatus.Position -= float32(math.Min(float64(senatorTrade.Shares), float64(currentStatus.Position)))
		}

		currentStatus.EndDate = senatorTrade.TransactionDate
		stockResults[senatorTrade.Ticker] = currentStatus
	}

	return stockResults, nil
}

