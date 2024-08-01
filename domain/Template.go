package domain

import (
	"html/template"
	"io"
)

type Templates struct {
	Templates *template.Template
}

func NewTemplates() *Templates {
	return &Templates{
		Templates: template.Must(template.ParseGlob("views/*.gohtml")),
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
