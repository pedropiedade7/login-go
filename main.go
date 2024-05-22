package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// func StaticHandler(w http.ResponseWriter, r *http.Request) {
// 	f, err := os.Open("static" + r.URL.Path)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	if strings.HasSuffix(r.URL.Path, ".css") {
// 		w.Header().Add("Content-Type", "text/css")
// 	}

// 	io.Copy(w, f)

// }

var cache map[string]*template.Template

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about")
}

func RenderTemplate(w http.ResponseWriter, page string) {
	env := "dev"

	var t *template.Template
	var err error

	_, exists := cache[page]

	if !exists || env == "dev" {
		t, err = template.ParseFiles(
			"templates/"+page+".tmpl.html", "templates/base.layout.tmpl.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		cache[page] = t
	} else {
		fmt.Println("Cache Hit")
		t = cache[page]
	}

	t.Execute(w, nil)
}

func main() {

	cache = make(map[string]*template.Template)

	static := http.Dir("static")
	staticHandler := http.FileServer(static)

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)
	http.HandleFunc("/about", AboutHandler)

	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	http.ListenAndServe(":3000", nil)
}
