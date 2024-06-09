package backend

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type chartDataStruct struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

func getData(c echo.Context) error {

	if isLoggedIn(c) {
		var id uint

		//we check if the id is in the url

		if c.QueryParam("id") != "" {
			idUint, err := strconv.ParseUint(c.QueryParam("id"), 10, 32)
			if err != nil {
				return c.JSON(400, "Invalid id")
			}
			id = uint(idUint)
		} else {
			id = getId(c)
		}

		absenteLuna := getAbsenteMonthById(id)

		dataAbsente := []int{
			absenteLuna.January,
			absenteLuna.February,
			absenteLuna.March,
			absenteLuna.April,
			absenteLuna.May,
			absenteLuna.June,
			absenteLuna.July,
			absenteLuna.August,
			absenteLuna.September,
			absenteLuna.October,
			absenteLuna.November,
			absenteLuna.December,
		}

		chartData := chartDataStruct{
			Labels: []string{"Ianuarie", "Februarie", "Martie", "Aprilie", "Mai", "Iunie", "Iulie", "August", "Septembrie", "Octombrie", "Noiembrie", "Decembrie"},
			Data:   dataAbsente,
		}

		return c.JSON(200, chartData)
	}
	return c.JSON(401, "Unauthorized")
}
