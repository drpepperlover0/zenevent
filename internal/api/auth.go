package api

import (
	"log"
	"net/http"

	"github.com/drpepperlover0/internal/structs"
	"github.com/drpepperlover0/storage"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {

	user := structs.User{
		Username: c.Request().FormValue("username"),
		Password: c.Request().FormValue("password"),
		Role:     structs.Role1,
	}

	if err := storage.AddPart(user); err != nil {
		log.Println("add to table error")
		c.Redirect(http.StatusSeeOther, "http://127.0.0.1:3000/internal/frontend/participant/register_part.html")
	}

	return c.Redirect(http.StatusPermanentRedirect, "/auth/login/participant?to_login=true")
}

func (h *Handler) LogIn(c echo.Context) error {

	user := structs.User{
		Username: c.Request().FormValue("username"),
		Password: c.Request().FormValue("password"),
	}

	regRed := c.Request().URL.Query().Get("to_login")

	if err := storage.IsValidUser(user); err != nil {
		if regRed == "true" {
			return c.Redirect(http.StatusSeeOther, "http://127.0.0.1:3000/internal/frontend/participant/register_part.html")
		}
		return c.Redirect(http.StatusSeeOther, "http://127.0.0.1:3000/internal/frontend/participant/login_part.html")
	}

	userToken, err := GenerateUserJWT(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "error generate token",
		})
	}
	c.SetCookie(&http.Cookie{
		Name:  "token",
		Value: userToken,
		Path:  "/home",
	})

	if err := c.Redirect(http.StatusSeeOther, "/home"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "oops, you are in backrooms :(",
		})
	}

	return nil
}

func (h *Handler) SignUpOrg(c echo.Context) error {

	org := structs.Organizer{
		IndividEmail: c.FormValue("ind_email"),
		Name:         c.FormValue("org_name"),
		SID:          c.FormValue("org_id"),
	}

	if !(ValidateEmail(org.IndividEmail)) || !(ValidateOrg(org.Name)) || storage.AddOrg(org) != nil {
		return c.Redirect(http.StatusSeeOther, "http://127.0.0.1:3000/internal/frontend/organizer/register_org.html")
	}

	return c.Redirect(http.StatusPermanentRedirect, "/auth/login/organizer?to_login=true")
}

func (h *Handler) LogInOrg(c echo.Context) error {

	org := structs.Organizer{
		Name:         c.FormValue("org_name"),
		SID:          c.FormValue("org_id"),
	}
	isRedirect := c.Request().URL.Query().Get("to_login")

	if err := storage.IsValidOrg(org); err != nil {
		return SmartRedirect(c, isRedirect)
	}

	orgToken, err := GenerateOrgJWT(org)
	if err != nil {
		return SmartRedirect(c, isRedirect)
	}

	c.SetCookie(&http.Cookie{
		Name:  "token",
		Value: orgToken,
		Path:  "/home",
	})

	return c.Redirect(http.StatusSeeOther, "/home")
}

func SmartRedirect(c echo.Context, isRedirect string) error {
	if isRedirect == "true" {
		return c.Redirect(http.StatusSeeOther, "http://127.0.0.1:3000/internal/frontend/organizer/register_org.html")
	}
	return c.Redirect(http.StatusSeeOther, "http://127.0.0.1:3000/internal/frontend/organizer/login_org.html")
}
