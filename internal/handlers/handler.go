package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

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

	if userID == nil {
		exp := time.Now().AddDate(0, 0, 15)
		url.ExpiresAt = &exp
	}

	// ✅ Usa SaveURL que funciona con DB o memoria
	if err := storage.SaveURL(url.ShortID, url.OriginalURL); err != nil {
		http.Error(w, "Error al guardar el enlace", http.StatusInternalServerError)
		return
	}

	baseURL := utils.GetEnv("BASE_URL", "http://localhost:8080")
	resp := ShortenResponse{ShortURL: baseURL + "/" + shortID}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")

	// ✅ Usa GetOriginalURL compatible con DB o memoria
	original, err := storage.GetOriginalURL(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// ⚠ Aquí no se puede verificar expiración ni aumentar clics en memoria (solo en DB)
	if storage.UseDB {
		var url models.URL
		result := storage.DB.First(&url, "short_id = ?", id)
		if result.Error == nil {
			if url.ExpiresAt != nil && url.ExpiresAt.Before(time.Now()) {
				http.Error(w, "Este enlace ha expirado", http.StatusGone)
				return
			}
			url.Clicks++
			storage.DB.Save(&url)
		}
	}

	http.Redirect(w, r, original, http.StatusFound)
}

func HandleStats(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/stats/")

	if !storage.UseDB {
		http.Error(w, "Estadísticas no disponibles en modo memoria", http.StatusNotImplemented)
		return
	}

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

	if !storage.UseDB {
		http.Error(w, "Función no disponible en modo memoria", http.StatusNotImplemented)
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
