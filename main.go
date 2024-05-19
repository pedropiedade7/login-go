package main

import (
	"log"
	"net/http"
)

type Api struct{}

func (Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Now it's a func as a handler"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("<h1>Olá mundo GG</h1>"))
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.Handle("/api", Api{})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
