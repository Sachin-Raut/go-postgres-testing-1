package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Sachin-Raut/go-postgres-testing-1/users"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

var db *sql.DB

const (
	pgHost     = "host"
	pgPort     = "port"
	pgUsername = "user"
	pgPassword = "password"
	pgDBname   = "dbname"
	pgSSLmode  = "sslmode"
)

func init() {
	host := os.Getenv(pgHost)
	port := os.Getenv(pgPort)
	user := os.Getenv(pgUsername)
	password := os.Getenv(pgPassword)
	dbname := os.Getenv(pgDBname)
	sslmode := os.Getenv(pgSSLmode)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error 1-", err)
		return
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Error 2-", err)
		return
	}
	fmt.Println("database connection successful")
}

func main() {
	user, err := users.GetUser(db, 27) // getUser(28)
	if err != nil {
		//panic(err)
		fmt.Println("User not found")
		return
	}
	fmt.Println(user.ID)
	fmt.Println(user.Email)
}
