package backend

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	//emailverifier "github.com/AfterShip/email-verifier"
	"github.com/labstack/echo/v4"
	passwordValidator "github.com/wagslane/go-password-validator"
)

func signup(c echo.Context) error {
	if isLoggedIn(c) && isAdmin(c) {
		return c.Render(http.StatusOK, "signup.html", nil)
	}
	return c.Redirect(http.StatusFound, "/login")
}

//var verifier = emailverifier.NewVerifier()

func signupPost(c echo.Context) error {
	imSendingTheReq := c.RealIP() == "192.168.137.50"
	if (isLoggedIn(c) && isAdmin(c)) || imSendingTheReq {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		access := c.FormValue("access")
		image, err := c.FormFile("image")

		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		/*
				ret, err := verifier.Verify(email)
				if err != nil {
					return c.String(http.StatusInternalServerError, "Internal Server Error")
				}
			if !ret.Syntax.Valid {
				return c.String(http.StatusBadRequest, "Invalid Email")
			}
		*/

		const minEntropy float64 = 40
		err = passwordValidator.Validate(password, minEntropy)
		if err != nil {
			return c.String(http.StatusForbidden, "Invalid Email")
		}

		id, err := createUser(name, email, password, access)
		if err != nil {
			return c.String(http.StatusForbidden, "Internal Server Error")
		}

		err = sendUserData(id, name, email, image)

		if err != nil {
			fmt.Println("Requestul nu a putut fi trimis")

			deleteUser(id)

			return c.String(http.StatusBadRequest, "Requestul nu a putut fi trimis")
		}

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

	req, err := http.NewRequest("POST", "http://192.168.137.1:6969/newUser", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	_, err = client.Do(req)

	if err != nil {
		return err
	}

	return nil
}

type data struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	IsOk  bool   `json:"isOk"`
}

func reqIsOk(c echo.Context) error {
	var d data
	err := c.Bind(&d)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if !d.IsOk {
		deleteUser(d.Id)

		fmt.Println("Userul nu a fost adaugat")

		req, _ := http.NewRequest("PUT", "http://localhost:8080/singup", bytes.NewBuffer([]byte("0")))

		client := &http.Client{}

		_, _ = client.Do(req)

		return nil
	}

	err = createUserAbsenceMonth(d.Id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	err = createUserAbsenceYear(d.Id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	fmt.Println("Userul a fost adaugat cu succes")

	req, _ := http.NewRequest("PUT", "http://localhost:8080/singup", bytes.NewBuffer([]byte("1")))

	client := &http.Client{}

	_, _ = client.Do(req)

	return nil
}
