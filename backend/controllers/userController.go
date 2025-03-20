package controllers

import (
	"encoding/json"
	"net/http"

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
		http.Error(w, "Error al crear usuario ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cuenta creada!"})
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
		http.Error(w, "Usuario no encontrado ", http.StatusUnauthorized)
		return
	}

	// Comparar password vs encriptada
	if !repository.VerifyPassword(loginData.Password, user.Password) {
		http.Error(w, "Contraseña incorrecta ", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login verificado!", "userName": user.Email, "userMail": user.Email}) //TODO: retornar token al front
}
