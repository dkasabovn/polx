package analytics

import (
	"net/http"
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
