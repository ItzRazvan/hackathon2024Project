package backend

import (
	"github.com/labstack/echo/v4"
)

func logout(c echo.Context) error {
	sessionDestroy(c)
	return nil
}
