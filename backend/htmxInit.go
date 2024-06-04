package backend

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, index string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, index, data)
}
