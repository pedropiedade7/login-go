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

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")
}

func RenderTemplate(w http.ResponseWriter, page string) {
	tmpl, err := template.ParseFiles("templates/" + page + ".tmpl.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {

	static := http.Dir("static")
	staticHandler := http.FileServer(static)

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)

	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	http.ListenAndServe(":3000", nil)
}
