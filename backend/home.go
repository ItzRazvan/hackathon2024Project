package backend

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	if isLoggedIn(c) {
		return c.Render(http.StatusOK, "home.html", nil)
	}
	return c.Redirect(http.StatusFound, "/login")
}
