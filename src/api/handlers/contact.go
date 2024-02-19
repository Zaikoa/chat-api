package handlers

import (
	. "api/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InsertContact(db *sql.DB, contact Contact) int {
	query := `
	INSERT INTO contact
	(first_name, last_name, phone_number)
	VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, contact.First_name, contact.Last_name, contact.Phone_number).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func ChangeFirstName(db *sql.DB, old string, new string, contact_id int) bool {
	query := `
	UPDATE contact
	SET first_name=%s
	WHERE first_name=%s AND contact_id=%s
	`
	_, err := db.Exec(query, new, old, contact_id)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func ChangeLastName(db *sql.DB, old string, new string, contact_id int) bool {
	query := `
	UPDATE contact
	SET last_name=%s
	WHERE last_name=%s AND contact_id=%s
	`
	_, err := db.Exec(query, new, old, contact_id)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func ChangePhoneNumber(db *sql.DB, old int, new int, contact_id int) bool {
	query := `
	UPDATE contact
	SET phone_number=%s
	WHERE phone_number=%s AND contact_id=%s
	`
	_, err := db.Exec(query, new, old, contact_id)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
