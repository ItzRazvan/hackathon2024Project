package backend

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func sendUserInfo(c echo.Context) error {
	user := getUserById(getId(c))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"name":   user.Name,
		"email":  user.Email,
		"access": user.Access,
	})
}

