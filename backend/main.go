package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maximilianoarevalo/zapping_test/backend/db"
	"github.com/maximilianoarevalo/zapping_test/backend/routes"
)

func main() {
	// DB Connection
	db.Connect()

	// Routes
	rt := routes.SetupRoutes()

	fmt.Println("Running on :3000")
	log.Fatal(http.ListenAndServe(":3000", rt))
}
