package analytics

import (
	"github.com/labstack/echo/v4"
)

func addDataRoutes(e *echo.Echo) {
	g := e.Group("/shills")

	routerShillAutocomplete(g)
	routerGabesShitTODO(g)
}

func routerShillAutocomplete(g *echo.Group) {
	g.GET("/autocomplete", controllerShillAutocomplete)
}

// TODO(gabe)
func routerGabesShitTODO(g *echo.Group) {
	g.GET("/gabe", controllerShillStockResults)
}
