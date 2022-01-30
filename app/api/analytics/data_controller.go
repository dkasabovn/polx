package analytics

import (
	"fmt"
	"net/http"
	"polx/app/domain/bo"
	"polx/app/services/analytics"
	"polx/app/services/scraper"
	"polx/app/system/log"

	"github.com/labstack/echo/v4"
)

func controllerShillAutocomplete(c echo.Context) error {
	name := c.QueryParam("value")
	shills, err := scraper.GetScraperSvc().GetShills(c.Request().Context(), name)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data": []string{},
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": shills,
	})
}

func controllerShillStockResults(c echo.Context) error {
	shill := new(bo.Shill)
	if err := c.Bind(shill); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data":  []string{},
			"error": "Binding Issue",
		})
	}

	stockResults, err := analytics.GetAnalyticsService().GetShillTrades(c.Request().Context(), shill.Name)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data":  []string{},
			"error": "None",
		})
	}

	fmt.Println("stockResults")

	// fmt.Println(stockResults)
	// var result []bo.StockResult
	// for _, val := range stockResults {
	// 	result = append(result, val)
	// 	fmt.Println(result)
	// }
	fmt.Println(stockResults)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"data": stockResults,
	})

}
