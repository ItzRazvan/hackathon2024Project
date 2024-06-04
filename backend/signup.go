package backend

import (
	"net/http"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/labstack/echo/v4"
	passwordValidator "github.com/wagslane/go-password-validator"
)

func signup(c echo.Context) error {
	if isLoggedIn(c) && isAdmin(c) {
		return c.Render(http.StatusOK, "signup.html", nil)
	}
	return c.Redirect(http.StatusFound, "/login")
}

var verifier = emailverifier.NewVerifier()

func signupPost(c echo.Context) error {
	if isLoggedIn(c) && isAdmin(c) {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		access := c.FormValue("access")

		ret, err := verifier.Verify(email)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		if !ret.Syntax.Valid {
			return c.String(http.StatusBadRequest, "Invalid Email")
		}

		const minEntropy float64 = 50
		err = passwordValidator.Validate(email, minEntropy)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid Email")
		}

		id, err := createUser(name, email, password, access)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		//we create a user_absence_month and user_absence_year for the new user

		err = createUserAbsenceMonth(id)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		err = createUserAbsenceYear(id)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
	}
	return nil

}
