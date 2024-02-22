package models

import (
	"time"
)

type Message struct {
	From_number     int
	To_number       int
	Message_text    string
	Conversation_id int
}

type Contact struct {
	First_name   string
	Last_name    string
	Phone_number int
}

type GroupMember struct {
	Contact_id      int
	Conversation_id int
	Joined_date     time.Time
}
