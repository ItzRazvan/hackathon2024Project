package backend

import (
	"net/http"
	"strconv"

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

type absenta struct {
	Id     uint   `json:"id"`
	Year   int    `json:"year"`
	Month  string `json:"month"`
	Day    int    `json:"day"`
	Hour   int    `json:"hour"`
	Minute int    `json:"minute"`
}

func showUserAbsences(c echo.Context) error {
	if isLoggedIn(c) && isAdmin(c) {
		id := c.QueryParam("id")
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid id",
			})
		}

		absente := getAllAbsencesForUser(uint(idUint))

		return c.Render(http.StatusOK, "user.html", map[string]interface{}{
			"absente": absente,
		})
	}
	return c.Redirect(http.StatusFound, "/login")
}
