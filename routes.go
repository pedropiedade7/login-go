package main

import "net/http"

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.HomeHandler)
	mux.HandleFunc("/contact", app.ContactHandler)
	mux.HandleFunc("/about", app.AboutHandler)

	static := http.Dir("static")
	staticHandler := http.FileServer(static)

	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	return mux
}
