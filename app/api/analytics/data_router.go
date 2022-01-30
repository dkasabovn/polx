package analytics

import (
	"github.com/labstack/echo/v4"
)

func addDataRoutes(e *echo.Echo) {
	g := e.Group("/shills")

	routerShillAutocomplete(g)
}

func routerShillAutocomplete(g *echo.Group) {
	g.GET("/autocomplete", func(c echo.Context) error { return nil })
}

// TODO(gabe)
func routerGabesShitTODO(g *echo.Group) {
	g.GET("/gabe", func(c echo.Context) error {
		return nil
	})
}
