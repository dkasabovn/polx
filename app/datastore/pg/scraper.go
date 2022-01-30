package pg

import (
	"context"
	"database/sql"
	"fmt"
	"polx/app/datastore"
	"polx/app/domain/bo"
	"polx/app/domain/definition"
	"polx/app/system/log"
	"strings"
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
	statement := "SELECT shill_name FROM trades WHERE shill_name LIKE '$1%' LIMIT 5"
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

func (s *scraperRepo) BulkInsert(ctx context.Context, entries []bo.TradeEntry) error {

	args := make([]interface{}, 7*len(entries))
	valueStrings := make([]string, len(entries))

	for k, v := range entries {
		valueStrings[k] = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", k*7+1, k*7+2, k*7+3, k*7+4, k*7+5, k*7+6, k*7+7)
		args[k*7] = v.PublicationDate
		args[k*7+1] = v.Name
		args[k*7+2] = v.Ticker
		args[k*7+3] = v.TransactionDate
		args[k*7+4] = v.TransactionType
		args[k*7+5] = v.Shares
		args[k*7+6] = v.PricePerShare
	}

	statement := fmt.Sprintf("INSERT INTO trades (publication_date, shill_name, ticker, transaction_date, transaction_type, shares, price_per_share) VALUES %s", strings.Join(valueStrings, ", "))

	_, err := s.db.ExecContext(ctx, statement, args...)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
