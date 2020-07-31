package users

import (
	"database/sql"
)

//User is
type User struct {
	ID    int
	Email string
}

//GetUser is
func GetUser(db *sql.DB, id int) (*User, error) {
	var myuser User
	row := db.QueryRow("select id, email from users where id=$1", id)

	err := row.Scan(&myuser.ID, &myuser.Email)

	if err != nil {
		return nil, err
	}
	return &myuser, nil
}
