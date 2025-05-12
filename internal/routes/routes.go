package routes

import (
	"github.com/drpepperlover0/internal/api"
	_ "github.com/drpepperlover0/internal/api/events"
	"github.com/drpepperlover0/internal/api/home"
	"github.com/labstack/echo/v4"
)

func InitRoutes() *echo.Echo {

	router := echo.New()

	authHandle := api.NewAuthHandler()
	homeHandle := home.NewHomeHandler()
	// eventHandle := events.NewEventHandler()

	// ~~~~~~~~~~AUTH HANDLERS~~~~~~~~~~
	auth := router.Group("/auth") // localhost:8080/auth
	{
		reg := auth.Group("/register") // localhost:8080/auth/register
		{

			reg.GET("/participant", authHandle.SignUp) // localhost:8080/auth/register/participant
			{
				reg.POST("/participant/check", authHandle.RegCheckUser) // localhost:8080/auth/register/participant/check
			}
			reg.GET("/organizer", authHandle.SignUpOrg) // localhost:8080/auth/register/organizer
			{
				reg.POST("/organizer/check-org", authHandle.RegCheckOrg) // localhost:8080/auth/register/organizer/check-org
			}

		}
		login := auth.Group("/login") // localhost:8080/auth/login
		{

			login.GET("/participant", authHandle.LogIn) // localhost:8080/auth/login/participant
			{
				login.POST("/participant/check", authHandle.LoginCheckUser) // localhost:8080/auth/login/participant/check
			}
			login.GET("/organizer", authHandle.LogInOrg) // localhost:8080/auth/login/organizer
			{
				login.POST("/organizer/check-org", authHandle.LoginCheckOrg) // localhost:8080/auth/login/organizer/check-org
			}

		}
		auth.GET("/logout", authHandle.LogOut)
	}
	// ~~~~~~~~~~AUTH HANDLERS~~~~~~~~~~

	// -------------------HOME HANDLERS-------------------
	info := router.Group("/info")
	{
		info.GET("/organizers", homeHandle.OrgInfo) // localhost:8080/info/organizers
		{
			info.GET("/organizers/from-home", homeHandle.OrgInfoSetter) // localhost:8080/info/organizers/from-home
		}
	}
	{
		router.GET("/", homeHandle.ShowHome)     // localhost:8080
		router.GET("/home", homeHandle.ShowHome) // localhost:8080/home

		router.GET("/profile", homeHandle.Profile) // localhost:8080/profile
		router.GET("/home/org", nil)
	}
	// -------------------HOME HANDLERS-------------------

	// ____________________EVENTS HANDLERS____________________
	/*
	events := router.Group("/events")
	{
		events.GET("/all", eventHandle.ShowEvents)
		events.GET("/theme_party", eventHandle.ThemeParty)
		events.GET("/quest", eventHandle.Quest)

		events.GET("/make", eventHandle.MakeEventForm)
		{
			events.POST("/make/process", eventHandle.MakeEvent)
		}
		events.GET("/join", eventHandle.Join)
	}
	*/
	// ____________________EVENTS HANDLERS____________________

	return router
}
