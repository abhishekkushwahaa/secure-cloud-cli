package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	var err error
	connStr := "host=localhost port=5432 dbname=secure-cloud user=postgres sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	log.Println("Successfully connected to the database!")
}
