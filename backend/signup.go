package backend

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"mime/multipart"
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
		image, err := c.FormFile("image")

		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		ret, err := verifier.Verify(email)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		if !ret.Syntax.Valid {
			return c.String(http.StatusBadRequest, "Invalid Email")
		}

		const minEntropy float64 = 60
		err = passwordValidator.Validate(password, minEntropy)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid Email")
		}

		id, err := createUser(name, email, password, access)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		err = createUserAbsenceMonth(id)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		err = createUserAbsenceYear(id)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		_ = sendUserData(id, name, email, image)
	}
	return nil

}

func sendUserData(id uint, name string, email string, image *multipart.FileHeader) error {
	file, err := image.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	imageBytes, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	//transform image bytes to base64

	imageBase64 := base64.StdEncoding.EncodeToString([]byte(imageBytes))

	data := map[string]interface{}{
		"id":    id,
		"name":  name,
		"email": email,
		"image": imageBase64,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://192.168.178.1:6969/newUser", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
