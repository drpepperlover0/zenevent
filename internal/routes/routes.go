package routes

import (
	"net/http"

	"github.com/drpepperlover0/internal/api"
	"github.com/labstack/echo/v4"
)

func InitRoutes() *echo.Echo {

	router := echo.New()
	h := api.NewHandler()

	auth := router.Group("/auth") // localhost:8080/auth
	{
		reg := auth.Group("/register") // localhost:8080/auth/register
		{
			reg.POST("/participant", h.SignUp)  // localhost:8080/auth/register/participant
			reg.POST("/organizer", h.SignUpOrg) // localhost:8080/auth/register/organizer
		}
		login := auth.Group("/login")
		{
			login.POST("/participant", h.LogIn)
			login.POST("/organizer", h.LogInOrg)
		}
	}
	router.GET("/home", func(c echo.Context) error {
		cookie, err := c.Request().Cookie("token")
		if err != nil {
			return echo.NewHTTPError(400, map[string]interface{}{
				"message": "sss",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": cookie.Value,
		})
	})

	return router
}
