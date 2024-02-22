package handlers

import (
	. "api/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InsertGroupMember(db *sql.DB, member GroupMember) int {
	query := `
	INSERT INTO group_member
	(contact_id, conversation_id, joined_datetime)
	VALUES ($1, $2, $3) RETURNING id
	`

	var pk int
	err := db.QueryRow(query, member.Contact_id, member.Conversation_id, member.Joined_date).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func SetLeftDateTime(db *sql.DB, contact_id int) {
	query := `
	UPDATE group_member
	SET left_datetime=CURRENT_DATE
	
	WHERE contact_id=%s
	`
	_, err := db.Exec(query, contact_id)

	if err != nil {
		log.Fatal(err)
	}
}
