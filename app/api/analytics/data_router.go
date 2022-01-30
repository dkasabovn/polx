package analytics

import (
	"github.com/labstack/echo/v4"
)

func addDataRoutes(e *echo.Echo) {
	g := e.Group("/shills")

	routerShillAutocomplete(g)
	routerGabes(g)
	routerAllShill(g)
}

func routerShillAutocomplete(g *echo.Group) {
	g.GET("/autocomplete", controllerShillAutocomplete)
}

// TODO(gabe)
func routerGabes(g *echo.Group) {
	g.POST("/gabe", controllerShillStockResults)
}

func routerAllShill(g *echo.Group) {
	g.GET("/all", controllerAllShills)
}
