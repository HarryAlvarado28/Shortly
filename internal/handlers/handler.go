package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"shortly/internal/middleware"
	"shortly/internal/models"
	"shortly/internal/storage"
	"shortly/internal/utils"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	shortID := utils.GenerateID(6)

	// Verifica si hay un usuario autenticado en el contexto
	var userID *uint = nil
	if uidVal := r.Context().Value(middleware.UserIDKey); uidVal != nil {
		uid := uidVal.(uint)
		userID = &uid
	}

	url := models.URL{
		ShortID:     shortID,
		OriginalURL: req.URL,
		UserID:      userID,
	}

	if err := storage.DB.Create(&url).Error; err != nil {
		http.Error(w, "Error al guardar en DB", http.StatusInternalServerError)
		return
	}

	baseURL := utils.GetEnv("BASE_URL", "http://localhost:8080")
	resp := ShortenResponse{ShortURL: baseURL + "/" + shortID}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
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

func HandleStats(w http.ResponseWriter, r *http.Request) {
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

func HandleMyUrls(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(uint)

	var urls []models.URL
	err := storage.DB.Where("user_id = ?", userID).Find(&urls).Error
	if err != nil {
		http.Error(w, "Error al consultar URLs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}
