package main

import "net/http"

func (app *Application) ContactHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "contact")
}

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "index")
}

func (app *Application) AboutHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "about")
}
