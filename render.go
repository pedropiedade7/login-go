package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (a *Application) RenderTemplate(w http.ResponseWriter, page string) {
	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = template.ParseFiles(
			"templates/"+page+".tmpl.html", "templates/base.layout.tmpl.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		a.Cache[page] = t
	} else {
		fmt.Println("Cache Hit")
		t = a.Cache[page]
	}

	t.Execute(w, nil)
}
