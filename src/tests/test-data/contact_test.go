package testdata

import (
	"api/handlers"
	"database/sql"
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func init() {

}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
func TestRetrievetContact(t *testing.T) {
	t.Log("Seeing if user is correctly added to the ")
	db, mock := NewMock()

	query := "SELECT first_name, last_name, phone_number FROM contact WHERE contact_id=%s"

	rows := sqlmock.NewRows([]string{"first_name", "last_name", "phone_number"}).
		AddRow("John", "Smith", 5182223456)

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	result := handlers.RetrieveContactByID(db, 1)
	t.Log(result)
	assert.NotNil(t, result)

}
