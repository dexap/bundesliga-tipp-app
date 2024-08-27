package main

import (
	"context"
	"log"

	"github.com/dexap/bundesliga-tipp-app/frontend-service/templates"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return templates.Home().Render(context.Background(), c.Response().Writer)
	})

	e.GET("/login", func(c echo.Context) error {
		return templates.Login().Render(context.Background(), c.Response().Writer)
	})

	e.POST("/login", handleLogin)

	log.Fatal(e.Start(":3000"))
}

func handleLogin(c echo.Context) error {
	// Implementieren Sie hier Ihre Login-Logik
	// Für dieses Beispiel geben wir einfach die Home-Seite zurück
	return templates.Home().Render(context.Background(), c.Response().Writer)
}
