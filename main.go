package main

import (
	"flag"
	"html/template"
	"log"
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

func main() {
	cache := make(map[string]*template.Template)

	config := Config{Version: "1.0.0"}

	flag.StringVar(&config.Port, "port", "8080", "porta do server")
	flag.StringVar(&config.Env, "env", "dev", "ambiente")

	flag.Parse()

	app := Application{
		Config: config,
		Cache:  cache,
	}

	log.Printf("Servidor de %s na versão %s escutando na porta :%s", config.Env, config.Version, config.Port)

	log.Fatal(app.Start())

}
