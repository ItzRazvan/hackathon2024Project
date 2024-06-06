package backend

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func tabele(c echo.Context) error {
	if isLoggedIn(c) && isAdmin(c) {

		absenteLuna := getAllAbsenteMonth()
		absenteAn := getAllAbsenteYear()

		var AbsLuna []UserAbsenceMonth
		var AbsAn []UserAbsenceYear

		AbsLuna = append(AbsLuna, absenteLuna...)

		AbsAn = append(AbsAn, absenteAn...)

		return c.Render(http.StatusOK, "tabel.html", map[string]interface{}{
			"absenteLuna": AbsLuna,
			"absenteAn":   AbsAn,
		})
	}
	return c.Redirect(http.StatusFound, "/")
}
