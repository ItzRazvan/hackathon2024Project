package backend

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var (
	store = sessions.NewCookieStore([]byte("Hackathon2024-SecretKey"))
)

func sessionInit(c echo.Context, id uint) error {

	session, err := store.Get(c.Request(), "session")

	if err != nil {
		return err
	}

	session.Values["authenticated"] = true

	session.Values["id"] = id

	session.Options = &sessions.Options{
		MaxAge: 60 * 60 * 24,
		Path:   "/",
		Secure: false,
	}

	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return nil

}

func sessionDestroy(c echo.Context) error {
	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return nil
	}
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
	return nil
}

func isLoggedIn(c echo.Context) bool {
	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return false
	}
	return session.Values["authenticated"] == true
}

func getId(c echo.Context) uint {
	session, err := store.Get(c.Request(), "session")
	if err != nil {
		return 0
	}
	return session.Values["id"].(uint)
}
