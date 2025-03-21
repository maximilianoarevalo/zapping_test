package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/maximilianoarevalo/zapping_test/backend/auth"
	"github.com/maximilianoarevalo/zapping_test/backend/models"
	"github.com/maximilianoarevalo/zapping_test/backend/repository"
)

// Create user controller
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear usuario DB
	err := repository.CreateUser(user)
	if err != nil {
		if err.Error() == fmt.Sprintf("el usuario con el correo %s ya está registrado", user.Email) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  "Correo ya registrado!",
				"status":   400,
				"userMail": user.Email,
			})
			return
		} else {
			http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Cuenta creada!",
		"status":   200,
		"userName": user.Name,
		"userMail": user.Email,
	}) //redirigir al login
}

// Login user controller
func Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener cuenta usuario por email
	user, err := repository.GetUserByEmail(loginData.Email)
	if err != nil {
		http.Error(w, "Usuario no encontrado ", http.StatusNotFound)
		return
	}

	// Comparar password vs encriptada
	if !repository.VerifyPassword(loginData.Password, user.Password) {
		//http.Error(w, "Contraseña incorrecta ", http.StatusUnauthorized)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  "Contraseña incorrecta!",
			"status":   400,
			"userMail": user.Email,
		})
		return
	}

	// Generar el token JWT
	claims := &jwt.RegisteredClaims{
		Issuer:    "zapping_test",
		Subject:   user.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(4 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token
	signedToken, err := token.SignedString(auth.JwtKey)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Login verificado!",
		"status":   200,
		"userName": user.Name,
		"userMail": user.Email,
		"token":    signedToken, // retornar token al front
	})
}
