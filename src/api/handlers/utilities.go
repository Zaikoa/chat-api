package handlers

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func openConnection() (db *sql.DB) {
	connStr := "postgres://postgres:secret@localhost:5432/system?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	// Checks for errors
	if err != nil {
		log.Fatal(err)
	}

	// Checks if we have a connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
