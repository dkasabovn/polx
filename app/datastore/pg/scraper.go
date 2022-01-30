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
	"time"
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
	statement := "SELECT DISTINCT shill_name FROM trades WHERE shill_name LIKE $1 LIMIT 10"
	res, err := s.db.QueryContext(ctx, statement, fmt.Sprintf("%%%s%%", query))

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

func (s *scraperRepo) BulkInsert(ctx context.Context, entries []bo.TradeEntry) ([]int, error) {
	args := make([]interface{}, 8*len(entries))
	valueStrings := make([]string, len(entries))

	for k, v := range entries {
		valueStrings[k] = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", k*8+1, k*8+2, k*8+3, k*8+4, k*8+5, k*8+6, k*8+7, k*8+8)
		args[k*8] = v.PublicationDate
		args[k*8+1] = v.Name
		args[k*8+2] = v.Ticker
		args[k*8+3] = v.TransactionDate
		args[k*8+4] = v.TransactionType
		args[k*8+5] = v.Shares
		args[k*8+6] = v.PricePerShare
		args[k*8+7] = v.Hash
	}

	statement := fmt.Sprintf("INSERT INTO trades (publication_date, shill_name, ticker, transaction_date, transaction_type, shares, price_per_share, trade_hash) VALUES %s ON CONFLICT DO NOTHING RETURNING id", strings.Join(valueStrings, ", "))

	rows, err := s.db.QueryContext(ctx, statement, args...)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	var ids []int

	for rows.Next() {
		var id int
		if err := rows.Scan(
			&id,
		); err != nil {
			log.Error(err)
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (s *scraperRepo) GetShillsTickers(ctx context.Context, shillName string) ([]string, error){
	statement := "SELECT DISTINCT ticker FROM trades WHERE shill_name = $1"
	res, err := s.db.QueryContext(ctx, statement, shillName)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	var tickers []string
	for res.Next() {
		var tick string
		if err := res.Scan(&tick); err != nil {
			log.Error(err)
			return nil, err
		}
		tickers = append(tickers, tick)
	}

	return tickers, nil	

}

func (s *scraperRepo) GetShillsDates(ctx context.Context, shillName string) (time.Time, time.Time, error) {
	statement := "SELECT min(transaction_date), max(transaction_date) FROM trades WHERE shill_name = $1"
	res, err := s.db.QueryContext(ctx, statement, shillName)

	if err != nil {
		log.Error(err)
		return time.Time{}, time.Time{}, err
	}

	var min time.Time
	var max time.Time
	// var max time.Time
	for res.Next() {
		if err := res.Scan(&min, &max); err != nil {
			log.Error(err)
			return time.Time{}, time.Time{}, err
		}
	}
	// fmt.Print(minTick)

	return min, max,  nil	
}
