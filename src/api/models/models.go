package models

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
