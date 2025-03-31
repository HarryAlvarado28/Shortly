package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"shortly/internal/handlers"
	"shortly/internal/middleware"
	"shortly/internal/storage"
)

func main() {
	_ = godotenv.Load() // Cargar .env local
	storage.InitDB()

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/anon", handlers.AnonymousSessionHandler)

	http.HandleFunc("/stats/", middleware.ValidateJWT(handlers.HandleStats))
	http.HandleFunc("/my/urls", middleware.ValidateJWT(handlers.HandleMyUrls))
	http.HandleFunc("/shorten", middleware.ValidateJWT(handlers.HandleShorten))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "public/index.html")
			return
		}

		// Verifica si el archivo solicitado existe
		staticPath := filepath.Join("public", r.URL.Path)
		if _, err := os.Stat(staticPath); err == nil {
			http.ServeFile(w, r, staticPath)
			return
		}

		// Si no es un archivo estático, asumir que es un ID y redirigir
		handlers.HandleRedirect(w, r)
	})

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
