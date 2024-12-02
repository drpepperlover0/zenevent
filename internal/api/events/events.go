package events

import (
	"html/template"
	"net/http"
	"os"

	"github.com/drpepperlover0/internal/api"
	"github.com/drpepperlover0/internal/structs"
	"github.com/drpepperlover0/storage"
	"github.com/labstack/echo/v4"
)

func (h *EventHandler) ShowEvents(c echo.Context) error {

	events, err := os.ReadFile("internal/frontend/events/all_events.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, string(events))
}

func (h *EventHandler) Join(c echo.Context) error {

	var themeRedirect = c.Request().URL.Query().Get("theme")
	var eventId = c.Request().URL.Query().Get("event_id")

	cookie, err := c.Cookie("token")
	if err != nil || cookie == nil {
		return c.String(http.StatusUnauthorized, "You're not authorized")
	}

	name, err := api.ParseNameJWT(cookie.Value)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Access denied.")
	}

	if api.ValidateOrg(name) {
		return c.Redirect(http.StatusSeeOther, "/auth/login/participant")
	}

	if err := storage.AddToEvent(name, eventId); err != nil {
		return c.Redirect(http.StatusSeeOther, "/events/"+themeRedirect)
	}

	return c.Redirect(http.StatusSeeOther, "/events/"+themeRedirect)
}

func (h *EventHandler) ThemeParty(c echo.Context) error {

	var tmp = template.Must(template.ParseFiles("internal/frontend/events/theme_party.html"))

	events, err := storage.FindEvents("Theme Party")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return tmp.Execute(c.Response().Writer, struct{ Events []structs.Event }{Events: events})
}
