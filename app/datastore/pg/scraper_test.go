package pg_test

import (
	"context"
	"fmt"
	"polx/app/datastore/pg"
	"polx/app/domain/bo"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Ω

var _ = AfterSuite(func() {
	Ω(pg.ClearAll()).ShouldNot(HaveOccurred())
})

var _ = Describe("Scraper", func() {
	BeforeEach(func() {
		Ω(pg.ClearAll()).ShouldNot(HaveOccurred())
	})

	date := time.Now()
	dateYear, dateMonth, dateDay := date.Date()

	trades := []bo.TradeEntry{
				{
					PublicationDate: date,
					Name:            "BTS Jung Kook",
					Ticker:          "MEN",
					TransactionDate: date.Add(-1*time.Hour*24),
					TransactionType: bo.Buy,
					Shares:          1,
					PricePerShare:   100.0,
					Hash:            "asd",
				},
				{
					PublicationDate: date,
					Name:            "BTS Jung Kook",
					Ticker:          "COK",
					TransactionDate: time.Now(),
					TransactionType: bo.Buy,
					Shares:          20,
					PricePerShare:   2000.0,
					Hash:            "asg",
				},
			}

	When("Inserting bulk records", func() {
		It("Should not fail", func() {
			pg.GetScraperRepo().BulkInsert(context.Background(), trades)
			// Ω(err).ShouldNot(HaveOccurred())
			// Ω(ids).Should(HaveLen(3))
			// entries, err := pg.GetScraperRepo().GetTradesByShill(context.Background(), "BTS Jung Kook")
			// Ω(err).ShouldNot(HaveOccurred())
			// Ω(entries).Should(HaveLen(1))
		})
	})

	When("Getting Dates", func() {
		
		It("Should the min and max", func() {
			pg.GetScraperRepo().BulkInsert(context.Background(), trades)
			start, end, err := pg.GetScraperRepo().GetShillsDates(context.Background(), "BTS Jung Kook")
			Ω(err).ShouldNot(HaveOccurred())

			startYear, startMonth, startDay := start.Date()
			Ω(startYear).Should(Equal(dateYear))
			Ω(startMonth).Should(Equal(dateMonth))
			Ω(startDay).Should(Equal(dateDay-1))

			endYear, endMonth, endDay := end.Date()
			Ω(endYear).Should(Equal(dateYear))
			Ω(endMonth).Should(Equal(dateMonth))
			Ω(endDay).Should(Equal(dateDay))
		})
	})

	When("Getting Distinct Tickers", func() {
		It("Should return distinct tickers", func() {
			pg.GetScraperRepo().BulkInsert(context.Background(), trades)
			tickers, err := pg.GetScraperRepo().GetShillsTickers(context.Background(), "BTS Jung Kook")
			Ω(err).ShouldNot(HaveOccurred())
			
			tickMap := make(map[string]bool)
			fmt.Print(tickers)
			for _, tick := range tickers {
				tickMap[tick] = true
			}

			Ω(len(tickMap)).Should(Equal(2))

			_, ok := tickMap["MEN"]
			Ω(ok).Should(Equal(true))

			_, ok = tickMap["COK"]
			Ω(ok).Should(Equal(true))
		})
	})
})
