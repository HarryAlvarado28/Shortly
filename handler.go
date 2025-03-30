package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"shortly/internal/models"
	"shortly/internal/storage"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	shortID := GenerateID(6)

	err := storage.SaveURL(shortID, req.URL)
	if err != nil {
		http.Error(w, "Error al guardar en DB", http.StatusInternalServerError)
		return
	}

	baseURL := GetEnv("BASE_URL", "http://localhost:8080")
	resp := map[string]string{"short_url": baseURL + "/" + shortID}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")

	var url models.URL
	result := storage.DB.First(&url, "short_id = ?", id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	// Incrementar contador de clics
	url.Clicks++
	storage.DB.Save(&url)

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/stats/")

	var url models.URL
	result := storage.DB.First(&url, "short_id = ?", id)
	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	resp := map[string]interface{}{
		"short_id":     url.ShortID,
		"original_url": url.OriginalURL,
		"clicks":       url.Clicks,
		"created_at":   url.CreatedAt,
		"expires_at":   url.ExpiresAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
