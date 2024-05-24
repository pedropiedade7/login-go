package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates
var TemplateFS embed.FS

func (a *Application) RenderTemplate(w http.ResponseWriter, page string) {
	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = template.ParseFS(
			TemplateFS,
			"templates/"+page+".tmpl.html",
			"templates/navbar.layout.html",
			"templates/base.layout.tmpl.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		a.Cache[page] = t
	} else {
		fmt.Println("Cache Hit")
		t = a.Cache[page]
	}

	contact := struct {
		Email    string
		Telefone string
	}{
		Email:    "pedro@gmail.com",
		Telefone: "33 9999-8888",
	}

	t.Execute(w, contact)
}
