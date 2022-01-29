package pg

import (
	"context"
	"database/sql"
	"polx/app/datastore"
	"polx/app/domain/bo"
	"polx/app/domain/definition"
	"polx/app/system/log"
	"sync"
)

var (
	scraperOnce sync.Once
	scraperInst *scraperRepo
)

type scraperRepo struct {
	db *sql.DB
}

func GetScraperRepo() definition.ScraperRepo {
	scraperOnce.Do(func() {
		scraperInst = &scraperRepo{
			db: datastore.RwInstance(),
		}
	})
	return scraperInst
}

func (s *scraperRepo) GetShills(ctx context.Context, query string) ([]bo.Shill, error) {
	statement := "SELECT shill_name FROM shills WHERE shill_name LIKE '$1%' LIMIT 5"
	res, err := s.db.QueryContext(ctx, statement, query)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	var shills []bo.Shill
	for res.Next() {
		var shill bo.Shill
		if err := res.Scan(
			&shill.Name,
		); err != nil {
			log.Error(err)
			return nil, err
		}

		shills = append(shills, shill)
	}
	return shills, nil
}

func (s *scraperRepo) GetTradesByShill(ctx context.Context, shillName string) ([]bo.TradeEntry, error) {
	statement := "SELECT publication_date, shill_name, ticker, transaction_date, transaction_type, shares, price_per_share FROM trades WHERE shill_name = $1"
	res, err := s.db.QueryContext(ctx, statement, shillName)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	var trades []bo.TradeEntry
	for res.Next() {
		var trade bo.TradeEntry
		if err := res.Scan(
			&trade.PublicationDate,
			&trade.Name,
			&trade.Ticker,
			&trade.TransactionDate,
			&trade.TransactionType,
			&trade.TransactionType,
			&trade.Shares,
			&trade.PricePerShare,
		); err != nil {
			log.Error(err)
			return nil, err
		}
		trades = append(trades, trade)
	}

	return trades, nil
}
