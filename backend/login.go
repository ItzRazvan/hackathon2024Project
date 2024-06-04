package backend

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func login(c echo.Context) error {
	if isLoggedIn(c) {
		return c.Redirect(http.StatusFound, "/")
	}
	return c.Render(http.StatusOK, "login.html", nil)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func loginPost(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	allGood, id := comparePasswordAndEmail(email, password)

	if allGood {
		err := sessionInit(c, id)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return c.Redirect(http.StatusFound, "/")
	}

	return c.String(http.StatusUnauthorized, "Unauthorized")

}
