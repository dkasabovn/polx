package analytics

import (
	"context"
	"fmt"
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
	analyticsOnce       sync.Once
	userServiceInstance *analyticsService
)

type analyticsService struct {
	tradeRepo  definition.ScraperRepo
	alpacaRepo definition.AlpacaRepo
}

func GetAnalyticsService() iface.AnalyticsService {
	analyticsOnce.Do(func() {
		userServiceInstance = &analyticsService{
			tradeRepo:  pg.GetScraperRepo(),
			alpacaRepo: alpaca.GetAlpacaRepo(),
		}
	})
	return userServiceInstance
}

func (a *analyticsService) GetShillTrades(ctx context.Context, shillName string) (map[string]bo.StockResult, int, error) {

	tickers, err := a.tradeRepo.GetShillsTickers(ctx, shillName)
	if err != nil {
		return nil, 0, err
	}

	trades, err := a.tradeRepo.GetTradesByShill(ctx, shillName)
	if err != nil {
		return nil, 0, err
	}

	startDate, err := a.tradeRepo.GetShillsDates(ctx, shillName)
	if err != nil {
		return nil, 0, err
	}

	alpacaData := make(map[string][]bo.Bar)
	for _, tick := range tickers {
		data, err := a.alpacaRepo.GetBars(tick, startDate.Format(time.RFC3339))
		if err != nil {
			return nil, 0, err
		}
		alpacaData[tick] = data.Bars
	}

	var ordersAmount float32

	stockResults := make(map[string]bo.StockResult)
	for _, senatorTrade := range trades {
		currentStatus, ok := stockResults[senatorTrade.Ticker]
		if len(alpacaData[senatorTrade.Ticker]) == 0 {
			fmt.Printf("len 0: %s\n", senatorTrade.Ticker)
			continue
		}
		index := math.Min(float64(senatorTrade.PublicationDate.Sub(startDate).Hours()/24), float64(len(alpacaData[senatorTrade.Ticker])-1))
		retailPrice := alpacaData[senatorTrade.Ticker][int(math.Max(0.0, index))].ClosePrice

		if !ok {
			currentStatus = bo.StockResult{}
			currentStatus.Ticker = senatorTrade.Ticker
			currentStatus.StartDate = senatorTrade.TransactionDate

			currentStatus.SenatorInitalPrice = senatorTrade.PricePerShare
			currentStatus.RetailInitalPrice = retailPrice
		}

		if senatorTrade.TransactionType == bo.Buy {

			if currentStatus.Position == 0 {
				currentStatus.RetailAvgSharePrice = retailPrice
				currentStatus.SenatorAvgSharePrice = currentStatus.RetailInitalPrice
			} else {
				currentStatus.RetailAvgSharePrice = (currentStatus.Position*currentStatus.RetailAvgSharePrice + retailPrice*float32(senatorTrade.Shares)) / (float32(senatorTrade.Shares) + currentStatus.Position)
				currentStatus.SenatorAvgSharePrice = (currentStatus.Position*currentStatus.SenatorAvgSharePrice + senatorTrade.PricePerShare*float32(senatorTrade.Shares)) / (float32(senatorTrade.Shares) + currentStatus.Position)
			}
			currentStatus.Position += float32(senatorTrade.Shares)
			currentStatus.ShareDelta += float32(senatorTrade.Shares)
			currentStatus.PeakShares += float32(math.Max(float64(currentStatus.Position), float64(currentStatus.PeakShares)))

			currentStatus.SenatorTotalSpent += float32(senatorTrade.Shares) * senatorTrade.PricePerShare
			currentStatus.RetailTotalSpent += float32(senatorTrade.Shares) * retailPrice

		} else {
			currentStatus.SenatorSales += (senatorTrade.PricePerShare - currentStatus.SenatorAvgSharePrice) * float32(senatorTrade.Shares)
			currentStatus.RetailSales += (retailPrice - float32(currentStatus.RetailAvgSharePrice)) * float32(senatorTrade.Shares)

			if currentStatus.Position-float32(senatorTrade.Shares) < 0 {
				currentStatus.PeakShares = float32(math.Max(float64(currentStatus.Position)+float64(senatorTrade.Shares), float64(currentStatus.PeakShares)))
			}
			currentStatus.ShareDelta -= float32(senatorTrade.Shares)
			currentStatus.Position -= float32(math.Max(0, float64(currentStatus.Position)-float64(senatorTrade.Shares)))
		}
		currentStatus.TotalOrders += float32(senatorTrade.Shares)
		ordersAmount += float32(senatorTrade.Shares)
		currentStatus.EndDate = senatorTrade.TransactionDate
		stockResults[senatorTrade.Ticker] = currentStatus
	}

	for tick, result := range stockResults {
		arr := alpacaData[tick]
		result.CurrentPrice = arr[len(arr)-1].ClosePrice

		result.SenatorValue = (result.SenatorSales - result.SenatorTotalSpent) + (result.CurrentPrice-result.SenatorAvgSharePrice)*result.Position
		result.RetailValue = (result.RetailSales - result.RetailTotalSpent) + (result.CurrentPrice-result.RetailAvgSharePrice)*result.Position

		result.SenatorValue = result.SenatorSales + (result.CurrentPrice-result.SenatorAvgSharePrice)*result.Position
		result.RetailValue = result.RetailSales + (result.CurrentPrice-result.RetailAvgSharePrice)*result.Position

		stockResults[tick] = result
	}

	return stockResults, int(ordersAmount), nil
}

// }

// 	if !ok {
// 		currentStatus = bo.StockResult{}
// 		currentStatus.Ticker = senatorTrade.Ticker
// 		currentStatus.StartDate = senatorTrade.TransactionDate

// 		currentStatus.SenatorInitalPrice = senatorTrade.PricePerShare
// 		currentStatus.SenatorAvgSharePrice = senatorTrade.PricePerShare

// 		index := math.Min(float64(senatorTrade.PublicationDate.Sub(startDate).Hours()/24), float64(len(alpacaData[senatorTrade.Ticker])-1))
// 		currentStatus.RetailInitalPrice = alpacaData[senatorTrade.Ticker][int(math.Max(0.0, index))].ClosePrice
// 		currentStatus.RetailAvgSharePrice = currentStatus.RetailInitalPrice
// 	}

// 	// publicationIndex := int(math.Min(senatorTrade.PublicationDate.Sub(startDate).Hours()/24, float64(alpacaDays-1)))
// 	data := alpacaData[senatorTrade.Ticker]
// 	lastDate := data[len(data)-1].Timestamp
// 	publicationIndex := int(lastDate.Sub(senatorTrade.PublicationDate).Hours()/24) + 1
// 	retailTrade := data[len(data)-publicationIndex]

// 	if senatorTrade.TransactionType == bo.Buy {
// 		if (float32(senatorTrade.Shares) + currentStatus.Position) != 0 {
// 			currentStatus.SenatorAvgSharePrice = (currentStatus.Position*currentStatus.SenatorAvgSharePrice + senatorTrade.PricePerShare*float32(senatorTrade.Shares)) / (float32(senatorTrade.Shares) + currentStatus.Position)
// 		}

// 		if (float32(senatorTrade.Shares) + currentStatus.Position) != 0 {
// 			currentStatus.RetailAvgSharePrice = (currentStatus.Position*currentStatus.RetailAvgSharePrice + retailTrade.ClosePrice*float32(senatorTrade.Shares)) / (float32(senatorTrade.Shares) + currentStatus.Position)
// 		}

// 		currentStatus.Position += float32(senatorTrade.Shares)
// 		currentStatus.ShareDelta += float32(senatorTrade.Shares)

// 		currentStatus.SenatorTotalSpent += senatorTrade.PricePerShare * float32(senatorTrade.Shares)
// 		currentStatus.RetailTotalSpent += retailTrade.ClosePrice * float32(senatorTrade.Shares)

// 	} else {
// 		currentStatus.ShareDelta -= float32(senatorTrade.Shares)
// 		currentStatus.SenatorSales += (senatorTrade.PricePerShare - float32(currentStatus.SenatorAvgSharePrice)) * float32(senatorTrade.Shares)
// 		currentStatus.RetailSales += (retailTrade.ClosePrice - float32(currentStatus.RetailAvgSharePrice)) * float32(senatorTrade.Shares)

// 		currentStatus.Position -= float32(math.Min(float64(senatorTrade.Shares), float64(currentStatus.Position)))
// 	}

// 	currentStatus.EndDate = senatorTrade.TransactionDate
// 	stockResults[senatorTrade.Ticker] = currentStatus
// }

// for tick, result := range stockResults {
// 	arr := alpacaData[tick]
// 	result.CurrentPrice = arr[len(arr)-1].ClosePrice

// 	result.SenatorValue = (result.SenatorSales - result.SenatorTotalSpent) + (result.CurrentPrice-result.SenatorAvgSharePrice)*result.Position
// 	result.RetailValue = (result.RetailSales - result.RetailTotalSpent) + (result.CurrentPrice-result.RetailAvgSharePrice)*result.Position

// 	// result.SenatorValue = result.SenatorSales  + (result.CurrentPrice - result.SenatorAvgSharePrice) * result.Position
// 	// result.RetailValue  = result.RetailSales   + (result.CurrentPrice - result.RetailAvgSharePrice)  * result.Position

// 	stockResults[tick] = result

// }

// return stockResults, nil
// }
