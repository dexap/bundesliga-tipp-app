package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	r := echo.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.GET("/", func(c echo.Context) error {
		return c.String(200, "bundesliga-service says: Hello, World!")
	})

	r.Logger.Fatal(r.Start(":3000"))

}
