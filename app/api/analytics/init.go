package analytics

import "github.com/labstack/echo/v4"

func Init() {
	e := echo.New()

	addDataRoutes(e)
}