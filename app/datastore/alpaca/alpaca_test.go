package alpaca_test

import (
	"polx/app/datastore/alpaca"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Ω

var _ = Describe("Alpaca", func() {

	When("Getting Trades", func() {
		It("Should not fail", func() {
			bars, err := alpaca.GetAlpacaRepo().GetBars("AAPL", "2020-01-01", "2022-01-01")
			Ω(len(bars.Bars)).ShouldNot(Equal(0), "LEN NOT ZERO")
			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
