package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"shortly/internal/models"
	"shortly/internal/storage"
	"shortly/internal/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validar campos mínimos
	input.Username = strings.TrimSpace(input.Username)
	input.Email = strings.ToLower(strings.TrimSpace(input.Email))
	if input.Username == "" || input.Email == "" || len(input.Password) < 6 {
		http.Error(w, "Datos inválidos o incompletos", http.StatusBadRequest)
		return
	}

	// Hashear la contraseña
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Error al hashear contraseña", http.StatusInternalServerError)
		return
	}

	// Guardar usuario
	user := models.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	if err := storage.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error al registrar usuario", http.StatusInternalServerError)
		return
	}

	// Éxito
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Usuario registrado con éxito",
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Buscar usuario por email
	var user models.User
	result := storage.DB.First(&user, "email = ?", input.Email)
	if result.Error != nil {
		http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
		return
	}

	// Validar contraseña
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
		return
	}

	// Generar token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Error al generar token", http.StatusInternalServerError)
		return
	}

	// Éxito
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
