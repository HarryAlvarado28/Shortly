package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "public/index.html")
			return
		}

		http.ServeFile(w, r, "public"+r.URL.Path)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
