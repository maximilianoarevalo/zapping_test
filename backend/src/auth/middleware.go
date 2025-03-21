package auth

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("testKey")

// Middleware para validar el token
func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		// Si no hay token
		if tokenString == "" {
			http.Error(w, "Token de autorización faltante", http.StatusUnauthorized)
			return
		}

		// Limpiar token
		tokenString = tokenString[len("Bearer "):]

		// Parsear el token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			return JwtKey, nil
		})

		// Verificar token valido
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido o expirado", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
