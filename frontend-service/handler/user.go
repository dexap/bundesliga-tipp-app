package handler

import (
	"github.com/dexap/bundesliga-tipp-app/frontend-service/model"
	"github.com/dexap/bundesliga-tipp-app/frontend-service/templates/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "user@mail.com",
	}
	return render(c, user.Show(u))
}
