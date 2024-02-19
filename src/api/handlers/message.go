package handlers

import (
	. "api/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

/*
Inserts a message into the message table using the struct
*/
func InsertMessage(db *sql.DB, message Message) int {
	query := `
	INSERT INTO message 
	(from_number, to_number, message_text, conversation_id 
	VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, message.From_number,
		message.To_number, message.Message_text, message.Conversation_id,
	).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

/*
Changes a message inside of the message table

Should return true or recieve a error
*/
func ChangeMessage(db *sql.DB, message string, message_id int) bool {
	if MessageExist(db, message_id) {
		query := `
		UPDATE message
		SET message_text=%s
		WHERE message_id=%s
		`
		_, err := db.Exec(query, message, message_id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return true
}

/*
Checks to see if message id exists

Returns true if it does and false if not
*/
func MessageExist(db *sql.DB, message_id int) bool {
	query := `
	SELECT message_text
	FROM message
	WHERE id=%s
	`
	var result bool

	err := db.QueryRow(query, message_id).Scan(&result)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		return false
	}

	return result
}

/*
Changes the from number
Usage = if contact number changes

returns true if succeds and error elsewise
*/
func ChangeFromNumber(db *sql.DB, old int, new int, message_id int) bool {
	query := `
	UPDATE message
	SET from_number=%s
	WHERE from_number=%s
	`
	_, err := db.Exec(query, new, old)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

/*
Changes the to number
Usage = if contact number changes

returns true if succeds and error elsewise
*/
func ChangeToNumber(db *sql.DB, old int, new int, message_id int) bool {
	query := `
	UPDATE message
	SET to_number=%s
	WHERE to_number=%s
	`
	_, err := db.Exec(query, new, old)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
