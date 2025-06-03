package storage

import (
	"errors"
	"shortly/internal/models"
	"sync"
)

var (
	inMemoryURLs  = make(map[string]models.URL)
	inMemoryUsers = make(map[string]models.User) // clave: email

	memLock = sync.RWMutex{}
)

// Crear una nueva URL (en DB o en memoria)
func SaveURL(shortID, original string) error {
	if UseDB {
		url := models.URL{
			ShortID:     shortID,
			OriginalURL: original,
		}
		return DB.Create(&url).Error
	}

	// Fallback en memoria
	memLock.Lock()
	defer memLock.Unlock()

	inMemoryURLs[shortID] = models.URL{
		ShortID:     shortID,
		OriginalURL: original,
	}
	return nil
}

// Buscar una URL por su shortID
func GetOriginalURL(shortID string) (string, error) {
	if UseDB {
		var url models.URL
		result := DB.First(&url, "short_id = ?", shortID)
		if result.Error != nil {
			return "", errors.New("URL no encontrada")
		}
		return url.OriginalURL, nil
	}

	// Fallback en memoria
	memLock.RLock()
	defer memLock.RUnlock()

	url, exists := inMemoryURLs[shortID]
	if !exists {
		return "", errors.New("URL no encontrada")
	}
	return url.OriginalURL, nil
}

// Guardar usuario (para login o anónimos)
func SaveUser(user models.User) error {
	if UseDB {
		return DB.Create(&user).Error
	}

	memLock.Lock()
	defer memLock.Unlock()
	inMemoryUsers[user.Email] = user
	return nil
}

// Buscar usuario por email
func FindUserByEmail(email string) (models.User, error) {
	if UseDB {
		var user models.User
		result := DB.First(&user, "email = ?", email)
		if result.Error != nil {
			return user, errors.New("usuario no encontrado")
		}
		return user, nil
	}

	memLock.RLock()
	defer memLock.RUnlock()
	user, exists := inMemoryUsers[email]
	if !exists {
		return user, errors.New("usuario no encontrado")
	}
	return user, nil
}
