package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	// DB config
	dbCredentials := "user=test dbname=zapping_test sslmode=disable password=test host=postgres port=5432" //-> Docker compose test
	DB, err = sql.Open("postgres", dbCredentials)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al hacer ping a la base de datos: ", err)
	}
	log.Println("Conectado a la base de datos PostgreSQL")
}
