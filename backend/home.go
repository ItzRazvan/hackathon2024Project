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

		lunaRomana := transformaLunaInRo(lunaNeformatata)

		anNeformatat := now.Year()
		year := "y" + strconv.Itoa(anNeformatat)

		absenteLuna, absenteAn := getAbsenteNow(getId(c), month, year)

		return c.Render(http.StatusOK, "home.html", map[string]interface{}{
			"luna":        lunaRomana,
			"absenteLuna": absenteLuna,
			"an":          anNeformatat,
			"absenteAn":   absenteAn,
		},
		)
	}
	return c.Redirect(http.StatusFound, "/login")
}

func transformaLunaInRo(luna string) string {
	switch luna {
	case "January":
		return "Ianuarie"
	case "February":
		return "Februarie"
	case "March":
		return "Martie"
	case "April":
		return "Aprilie"
	case "May":
		return "Mai"
	case "June":
		return "Iunie"
	case "July":
		return "Iulie"
	case "August":
		return "August"
	case "September":
		return "Septembrie"
	case "October":
		return "Octombrie"
	case "November":
		return "Noiembrie"
	case "December":
		return "Decembrie"
	}
	return ""
}
