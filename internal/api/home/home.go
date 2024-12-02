package home

import (
	"html/template"
	"net/http"

	"github.com/drpepperlover0/internal/api"
	"github.com/labstack/echo/v4"
)

func (h *HomeHandler) ShowHome(c echo.Context) error {

	login := struct {
		IsLogin bool
	}{
		IsLogin: true,
	}
	tmp := template.Must(template.ParseFiles("internal/frontend/home/index.html"))

	cookie, err := c.Request().Cookie("token")
	if err != nil || cookie == nil {
		login.IsLogin = false
	}

	return tmp.Execute(c.Response().Writer, login)
}

func (h *HomeHandler) Profile(c echo.Context) error {

	tokenString, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	name, err := api.ParseNameJWT(tokenString.Value)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := struct {
		Name          string
		IsParticipant bool
	}{
		Name:          name,
		IsParticipant: false,
	}

	if !api.ValidateOrg(name) {
		data.Name = name
		data.IsParticipant = true
	}

	tmp := template.Must(template.ParseFiles("internal/frontend/home/profile.html"))

	return tmp.Execute(c.Response().Writer, data)
}
