package main

import (
	"log"
	"net/http"
)

func (app *Application) ContactHandler(w http.ResponseWriter, r *http.Request) {
	err := app.RenderTemplate(w, "contact", TemplateData{
		Email:    "pedro@pedro.pedro",
		Telefone: "99 9999-9999",
		Route:    "contact",
	})
	if err != nil {
		log.Print(err)
	}
}

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// err := app.RenderTemplate(w, "index", TemplateData{Route: "index"})
	// if err != nil {
	// 	log.Println(err)
	// }
	view, err := NewView("index")
	if err != nil {
		log.Fatal(err)
	}

	err = view.Render(w, TemplateData{Route: "index"})
	if err != nil {
		log.Fatal(err)
	}
}

func (app *Application) AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := app.RenderTemplate(w, "about", TemplateData{Route: "about"})
	if err != nil {
		log.Println(err)
	}
}
func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := app.RenderTemplate(w, "login", TemplateData{Route: "login"})
	if err != nil {
		log.Println(err)
	}
}
