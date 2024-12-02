package events

import (
	"net/http"
	"os"

	"github.com/drpepperlover0/internal/api"
	"github.com/drpepperlover0/internal/structs"
	"github.com/drpepperlover0/storage"
	"github.com/labstack/echo/v4"
)

func (h *EventHandler) MakeEventForm(c echo.Context) error {

	makeForm, err := os.ReadFile("internal/frontend/events/make_events.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.HTML(http.StatusOK, string(makeForm))
}

func (h *EventHandler) MakeEvent(c echo.Context) error {

	cookie, err := c.Cookie("token")
	if err != nil || cookie == nil {
		return c.String(http.StatusUnauthorized, "Access denied")
	}

	name, err := api.ParseNameJWT(cookie.Value)
	if err != nil || name == "" {
		return c.String(http.StatusBadRequest, "Parse error")
	}

	if !api.ValidateOrg(name) {
		return c.String(http.StatusUnauthorized, "Access denied")
	}

	event := structs.Event{
		OrgName:     name,
		EventName:   c.FormValue("event_name"),
		Description: c.FormValue("event_desc"),
		EventTheme:  c.FormValue("picked_theme"),
		EventDate:   c.FormValue("event-date"),
	}

	if err := storage.AddEvent(event); err != nil {
		return c.Redirect(http.StatusSeeOther, "/events/make?error=true")
	}

	return c.Redirect(http.StatusSeeOther, "/home")
}
