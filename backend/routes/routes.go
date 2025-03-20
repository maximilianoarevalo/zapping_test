package routes

import (
	"github.com/gorilla/mux"
	"github.com/maximilianoarevalo/zapping_test/backend/controllers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Rutas API
	router.HandleFunc("/create_user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	return router
}
