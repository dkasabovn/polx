package analytics

import (
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
			"data": []string{},
		})
	}

	stockResults, err := analytics.GetAnalyticsService().GetShillTrades(c.Request().Context(), shill.Name )
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data": []string{},
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": stockResults,
	})


}
