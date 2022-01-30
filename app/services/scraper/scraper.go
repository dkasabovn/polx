package scraper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"polx/app/datastore/pg"
	"polx/app/domain/bo"
	"polx/app/domain/definition"
	"polx/app/services/iface"
	"polx/app/system/log"
	"sync"
	"time"
)

var (
	scraperSvcOnce sync.Once
	scraperSvcInst *scraperSvc
)

type scraperSvc struct {
	scraperRepo definition.ScraperRepo
}

func GetScraperSvc() iface.ScraperSvc {
	scraperSvcOnce.Do(func() {
		scraperSvcInst = &scraperSvc{
			scraperRepo: pg.GetScraperRepo(),
		}
	})
	return scraperSvcInst
}

func (s *scraperSvc) GetShills(ctx context.Context, query string) ([]bo.Shill, error) {
	return s.scraperRepo.GetShills(ctx, query)
}

func (s *scraperSvc) BulkInsert(ctx context.Context, entries []bo.TradeEntry) ([]int, error) {
	return s.scraperRepo.BulkInsert(ctx, entries)
}

func (s *scraperSvc) GetTradesByShill(ctx context.Context, shillName string) ([]bo.TradeEntry, error) {
	return s.scraperRepo.GetTradesByShill(ctx, shillName)
}

func (s *scraperSvc) RunTask(ctx context.Context, pageNum int) error {
	entries, err := s.fetchTradeEntries(ctx, pageNum)
	if err != nil {
		log.Error(err)
		return err
	}

	converted, err := s.hashTradeEntries(entries)

	if err != nil {
		log.Error(err)
		return err
	}

	if _, err := s.BulkInsert(ctx, converted); err != nil {
		log.Error(err)
		return err
	}

	// TODO(dk): Add notifications here

	return nil
}

func assemblePostPayload(pageNum int) ([]byte, error) {
	req := map[string]interface{}{
		"congressType":    "Both",
		"pageNumber":      pageNum,
		"pageSize":        100,
		"politicianParty": "Both",
		"shareTypes":      []string{"Stock"},
		"ticker":          false,
	}

	bytePayload, err := json.Marshal(req)
	return bytePayload, err
}

func (s *scraperSvc) fetchTradeEntries(ctx context.Context, pageNum int) ([]bo.TradeEntryRaw, error) {
	payload, err := assemblePostPayload(pageNum)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	resp, err := http.Post("https://api.capitoltrades.com/trades", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.InfoStruct(resp)
		log.Error(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Error(errors.New("Status code is not 200"))
		return nil, nil
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var newData []bo.TradeEntryRaw
	if err := json.Unmarshal(bytes, &newData); err != nil {
		log.Error(err)
		return nil, err
	}

	return newData, nil
}

func (s *scraperSvc) hashTradeEntries(entries []bo.TradeEntryRaw) ([]bo.TradeEntry, error) {
	out := make([]bo.TradeEntry, len(entries))
	for i, v := range entries {
		stringified, err := json.Marshal(v)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		hashed := sha256.Sum256(stringified)
		encoded := base64.StdEncoding.EncodeToString(hashed[:])

		// TODO(dk): Fix dates here
		entry := &bo.TradeEntry{
			Name:          v.Name,
			Ticker:        v.Ticker,
			Shares:        int(v.Shares),
			PricePerShare: float32(v.Price),
			Hash:          encoded,
		}

		parsedTrans, err := time.Parse("2006-01-02", v.TransactionDate)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		parsedPub, err := time.Parse("2006-01-02 15:04:05", v.PublicationDate[0:19])
		if err != nil {
			log.Error(err)
			return nil, err
		}

		entry.PublicationDate = parsedPub
		entry.TransactionDate = parsedTrans

		if v.TradeType == "Sale" || v.TradeType == "Sale (Partial)" {
			entry.TransactionType = bo.Sell
		} else {
			entry.TransactionType = bo.Buy
		}

		out[i] = *entry
	}

	return out, nil
}
