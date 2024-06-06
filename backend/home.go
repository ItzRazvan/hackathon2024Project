package backend

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	if isLoggedIn(c) {
		now := time.Now()

		lunaNeformatata := now.Month().String()
		month := strings.ToLower(lunaNeformatata)

		anNeformatat := now.Year()
		year := "y" + strconv.Itoa(anNeformatat)

		absenteLuna, absenteAn := getAbsenteNow(getId(c), month, year)

		return c.Render(http.StatusOK, "home.html", map[string]interface{}{
			"luna":        lunaNeformatata,
			"absenteLuna": absenteLuna,
			"an":          anNeformatat,
			"absenteAn":   absenteAn,
		},
		)
	}
	return c.Redirect(http.StatusFound, "/login")
}
