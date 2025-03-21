package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maximilianoarevalo/zapping_test/backend/src/controllers"
)

// Cors
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Evitar conflicto cors
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.Use(corsMiddleware)

	// Rutas API
	router.HandleFunc("/create-user", controllers.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")

	router.HandleFunc("/livestream", controllers.Player).Methods("GET") // probado en postman

	return router
}
