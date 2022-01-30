package pg_test

import (
	"context"
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

	When("Inserting bulk records", func() {
		It("Should not fail", func() {
			trades := []bo.TradeEntry{
				{
					PublicationDate: time.Now(),
					Name:            "BTS Jung Kook",
					Ticker:          "COK",
					TransactionDate: time.Now(),
					TransactionType: bo.Buy,
					Shares:          1,
					PricePerShare:   100.0,
				},
				{
					PublicationDate: time.Now(),
					Name:            "BTS Jung Kook",
					Ticker:          "COK",
					TransactionDate: time.Now(),
					TransactionType: bo.Buy,
					Shares:          20,
					PricePerShare:   2000.0,
				},
			}
			err := pg.GetScraperRepo().BulkInsert(context.Background(), trades)
			Ω(err).ShouldNot(HaveOccurred())
			entries, err := pg.GetScraperRepo().GetTradesByShill(context.Background(), "BTS Jung Kook")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(entries).Should(HaveLen(2))
		})
	})
})
