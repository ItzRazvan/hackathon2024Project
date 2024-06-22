package backend

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Absenta struct {
	IdUser uint   `json:"id_user"`
	Minut  string `json:"minut"`
	Ora    int    `json:"ora"`
	Ziua   int    `json:"ziua"`
	Luna   string `json:"luna"`
	An     int    `json:"an"`
}

func addAbsenta(c echo.Context) error {
	//daca un request e facut, vom adauga o absenta la aceasta luna si an
	if c.Request().Method == "POST" {
		fmt.Println("Absenta adaugata")

		absenta := new(Absenta)

		err := c.Bind(absenta)

		fmt.Println(absenta)
		if err != nil {
			return err
		}

		//transform absenta.minut in int
		minut, err := strconv.Atoi(absenta.Minut)
		if err != nil {
			return err
		}
		if minut < 10 {
			absenta.Minut = "0" + absenta.Minut
		}

		err = addAbsentaToDB(absenta.IdUser, strings.ToLower(absenta.Luna), "y"+strconv.Itoa(absenta.An))

		if err != nil {
			return err
		}

		err = addAbsenteToAllAbsences(absenta.IdUser, absenta.Ziua, absenta.Ora, absenta.Minut, strings.ToLower(absenta.Luna), absenta.An)

		if err != nil {
			return err
		}

		return c.JSON(200, map[string]interface{}{
			"message": "Absenta adaugata cu succes",
		})

	}
	return c.JSON(400, map[string]interface{}{
		"message": "Eroare la adaugarea absentei",
	})
}
