package middleware

import (
	"context"
	"net/http"
	"strings"

	"shortly/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey = contextKey("userID")

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token no proporcionado", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.GetEnv("JWT_SECRET", "my_secret_key")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Token mal formado", http.StatusUnauthorized)
			return
		}

		userID := uint(claims["user_id"].(float64))
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
