package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var funcs = template.FuncMap{
	"GetYear": func() int {
		return time.Now().Year()
	},
}

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, pages ...string) (*View, error) {
	files := []string{"templates/base.layout.tmpl.html", "templates/navbar.layout.html"}
	for _, p := range pages {
		files = append(files, fmt.Sprintf("templates/%s.tmpl.html", p))
	}
	t, err := template.New("").Funcs(funcs).ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	return &View{
		Template: t,
		Layout:   layout,
	}, nil
}

func (v *View) Render(w http.ResponseWriter, data any) error {
	v.Template.ExecuteTemplate(w, v.Layout, TemplateData{Route: "index"})
}
