package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/dexap/bundesliga-tipp-app/frontend-service/handler"
	"github.com/dexap/bundesliga-tipp-app/frontend-service/model"
	"github.com/dexap/bundesliga-tipp-app/frontend-service/service"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := handler.UserHandler{}
	baseHandler := handler.BaseHandler{}

	e.GET("/", baseHandler.HandleLandingShow)
	e.GET("/user", userHandler.HandleUserShow)

	service.GenerateSchedule(model.GetTestTeams())

	log.Fatal(e.Start(":4000"))
}
