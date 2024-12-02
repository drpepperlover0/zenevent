package home

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (h *HomeHandler) OrgInfoSetter(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/info/organizers?isHome=true")
}

func (h *HomeHandler) OrgInfo(c echo.Context) error {

	// cookie, err := c.Cookie("token")
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }
	home := c.Request().URL.Query().Get("isHome")

	orgInfo, err := os.ReadFile("internal/frontend/home/org_info.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if home != "true" {
		c.SetCookie(&http.Cookie{
			Name:   "token",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
	}

	return c.HTML(http.StatusOK, string(orgInfo))
}
