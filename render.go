package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates
var TemplateFS embed.FS

type TemplateData struct {
	Email    string
	Telefone string
	Route    string
}

func (a *Application) RenderTemplate(w http.ResponseWriter, page string, data any) error {
	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = parseTemplate(page, a.Config.Env)
		if err != nil {
			fmt.Println(err)
			return err
		}
		a.Cache[page] = t
	} else {
		fmt.Println("Cache Hit")
		t = a.Cache[page]
	}

	return t.ExecuteTemplate(w, "base", data)

}

func parseTemplate(page, env string) (*template.Template, error) {
	t := template.New("").Funcs(funcs)

	if env != "dev" {
		return t.ParseFS(
			TemplateFS,
			"templates/base.layout.tmpl.html",
			"templates/"+page+".tmpl.html",
			"templates/navbar.layout.html",
		)
	}
	return t.ParseFiles(
		"templates/base.layout.tmpl.html",
		"templates/"+page+".tmpl.html",
		"templates/navbar.layout.html",
	)

}
