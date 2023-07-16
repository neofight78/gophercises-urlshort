package main

import (
	"fmt"
	"net/http"

	"github.com/neofight78/gophercises-urlshort/handlers"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/handlers-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /handlers
  url: https://github.com/gophercises/urlshort
- path: /handlers-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := handlers.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
