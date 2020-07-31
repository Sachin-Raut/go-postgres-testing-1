package users

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	columns := []string{"id", "email"}
	rows := sqlmock.NewRows(columns).AddRow(27, "a@abc.com")

	mock.ExpectQuery("select").
		WithArgs(27).
		WillReturnRows(rows)

	receiveduser, err := GetUser(db, 27)

	expecteduser := &User{
		ID:    27,
		Email: "a@abc.com",
	}

	assert.Nil(t, err)
	assert.EqualValues(t, expecteduser, receiveduser)

}

func TestGetUserNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	columns := []string{"id", "email"}
	rows := sqlmock.NewRows(columns).AddRow(27, "a@abc.com")

	mock.ExpectQuery("select").
		WithArgs(2).
		WillReturnRows(rows)

	receiveduser, err := GetUser(db, 27)

	assert.NotNil(t, err)
	assert.Nil(t, receiveduser)

}
