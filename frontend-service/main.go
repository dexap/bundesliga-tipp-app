package main

import (
	"github.com/dexap/bundesliga-tipp-app/frontend-service/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return templates.Home().Render(c.Response().Writer, c.Request())
	})

	e.Logger.Fatal(e.Start(":3000"))
}
