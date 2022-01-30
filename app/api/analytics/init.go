package analytics

import "github.com/labstack/echo/v4"

func Init() {
	e := echo.New()

	addDataRoutes(e)

	e.Logger.Fatal(e.Start("0.0.0.0:6969"))
}
