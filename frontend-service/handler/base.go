package handler

import (
	landing "github.com/dexap/bundesliga-tipp-app/frontend-service/templates/landing"
	"github.com/labstack/echo/v4"
)

type BaseHandler struct{}

func (h BaseHandler) HandleLandingShow(c echo.Context) error {
	return render(c, landing.Landing())
}
