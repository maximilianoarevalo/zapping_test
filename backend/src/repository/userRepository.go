package repository

import (
	"fmt"
	"log"

	"github.com/maximilianoarevalo/zapping_test/backend/src/db"
	"github.com/maximilianoarevalo/zapping_test/backend/src/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	// Verificar si el usuario ya existe
	var existingUserCount int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", user.Email).Scan(&existingUserCount)
	if err != nil {
		log.Println("Error al verificar usuario:", err)
		return err
	}

	// Si el usuario ya existe, devolver un error
	if existingUserCount > 0 {
		return fmt.Errorf("el usuario con el correo %s ya está registrado", user.Email)
	}

	encryptedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal("Error al encriptar contraseña")
	}
	_, err = db.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, encryptedPassword)
	if err != nil {
		log.Println("Error al insertar usuario:", err)
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println("Error al obtener usuario por email:", err)
		return nil, err
	}
	return &user, nil
}

// Fuente: https://medium.com/@rnp0728/secure-password-hashing-in-go-a-comprehensive-guide-5500e19e7c1f
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
