package backend

import "github.com/labstack/echo/v4"

type hours struct {
	OraMaxima string `json:"oraMaxima"`
}

var hoursData = hours{
	OraMaxima: "8",
}

func sendHours(c echo.Context) error {
	return c.JSON(200, hoursData)
}
