package storage

import (
	"errors"
	"shortly/internal/models"
)

// Crear una nueva URL en la base de datos
func SaveURL(shortID, original string) error {
	url := models.URL{
		ShortID:     shortID,
		OriginalURL: original,
	}
	return DB.Create(&url).Error
}

// Buscar una URL por su shortID
func GetOriginalURL(shortID string) (string, error) {
	var url models.URL
	result := DB.First(&url, "short_id = ?", shortID)
	if result.Error != nil {
		return "", errors.New("URL no encontrada")
	}
	return url.OriginalURL, nil
}
