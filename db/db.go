package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // pg driver
)

type DB interface {
	Add(u *User) error
	Get(login string) (*User, error)
	Set(u *User) error
}

type Postgres struct {
	conn *sql.DB
}

func New(dbName, dbHost, dbUser, dbPass string, dbPort int) DB {
	conn, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		dbUser, dbName, dbPass, dbHost, dbPort,
	))
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	createUser := `
		CREATE TABLE IF NOT EXISTS users
		(
			uuid UUID NOT NULL PRIMARY KEY,
			login TEXT NOT NULL,
			date TEXT NOT NULL
		);
`

	rows, err := conn.Query(createUser)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Err() != nil {
		panic(rows.Err())
	}

	return &Postgres{
		conn: conn,
	}
}
