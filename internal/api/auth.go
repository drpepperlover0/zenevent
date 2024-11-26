package api

import (
	"log"
	"net/http"
	"os"

	"github.com/drpepperlover0/internal/structs"
	"github.com/drpepperlover0/storage"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {

	reg_html, err := os.ReadFile("internal/frontend/participant/register_part.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "parse register_part.html error")
	}

	return c.HTML(http.StatusOK, string(reg_html))
}

func (h *Handler) RegCheckUser(c echo.Context) error {

	user := structs.User{
		Username: c.Request().FormValue("username"),
		Password: c.Request().FormValue("password"),
		Role:     structs.Role1,
	}

	if err := storage.AddPart(user); err != nil {
		log.Println("add to table error")
		c.Redirect(http.StatusSeeOther, "/auth/register/participant")
	}

	return c.Redirect(http.StatusPermanentRedirect, "/auth/login/participant/check?to_login=true")
}

func (h *Handler) LogIn(c echo.Context) error {

	login_html, err := os.ReadFile("internal/frontend/participant/login_part.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "parse login_part.html error")
	}

	return c.HTML(http.StatusOK, string(login_html))
}

func (h *Handler) LoginCheckUser(c echo.Context) error {

	user := structs.User{
		Username: c.Request().FormValue("username"),
		Password: c.Request().FormValue("password"),
	}

	regRed := c.Request().URL.Query().Get("to_login")

	if err := storage.IsValidUser(user); err != nil {
		return SmartRedirect(c, regRed, structs.Role1)
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
		Path: "/",
	})

	if err := c.Redirect(http.StatusSeeOther, "/"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "oops, you are in backrooms :(",
		})
	}

	return nil
}

func (h *Handler) SignUpOrg(c echo.Context) error {

	regOrg_html, err := os.ReadFile("internal/frontend/organizer/register_org.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "parse register_org.html error")
	}

	return c.HTML(http.StatusOK, string(regOrg_html))
}

func (h *Handler) RegCheckOrg(c echo.Context) error {

	org := structs.Organizer{
		IndividEmail: c.FormValue("ind_email"),
		Name:         c.FormValue("org_name"),
		SID:          c.FormValue("org_id"),
	}

	if !(ValidateEmail(org.IndividEmail)) || !(ValidateOrg(org.Name)) || storage.AddOrg(org) != nil {
		return c.Redirect(http.StatusSeeOther, "/auth/register/organizer")
	}

	return c.Redirect(http.StatusPermanentRedirect, "/auth/login/organizer/check-org?to_login=true")
}

func (h *Handler) LogInOrg(c echo.Context) error {

	loginOrg_html, err := os.ReadFile("internal/frontend/organizer/login_org.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "parse login_org.html error")
	}

	return c.HTML(http.StatusOK, string(loginOrg_html))
}

func (h *Handler) LoginCheckOrg(c echo.Context) error {

	org := structs.Organizer{
		Name: c.FormValue("org_name"),
		SID:  c.FormValue("org_id"),
	}
	isRedirect := c.Request().URL.Query().Get("to_login")

	if err := storage.IsValidOrg(org); err != nil {
		return SmartRedirect(c, isRedirect, structs.Role2)
	}

	orgToken, err := GenerateOrgJWT(org)
	if err != nil {
		return SmartRedirect(c, isRedirect, structs.Role2)
	}

	c.SetCookie(&http.Cookie{
		Name:  "token",
		Value: orgToken,
		Path: "/",
	})

	return c.Redirect(http.StatusSeeOther, "/")
}

func SmartRedirect(c echo.Context, isRedirect string, role string) error {
	if isRedirect == "true" {
		return c.Redirect(http.StatusSeeOther, "/auth/register/"+role)
	}
	return c.Redirect(http.StatusSeeOther, "/auth/login/"+role)
}

func (h *Handler) LogOut(c echo.Context) error {

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, "/")
}
