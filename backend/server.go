package backend

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ServerStart() {
	app := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("./Views/html/*.html")),
	}

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowCredentials: true,
	}))

	app.Renderer = t

	app.Static("/", "/Views")

	app.GET("/user", sendUserInfo)

	app.GET("/login", login)
	app.POST("/login", loginPost)
	app.DELETE("/login", logout)

	app.GET("/", home)

	app.GET("/signup", signup)
	app.POST("/signup", signupPost)

	app.GET("/tabele", tabele)

	app.POST("/absenta", addAbsenta)

	app.GET("/getData", getData)

	app.Logger.Fatal(app.Start(":8080"))
}
