package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // pg driver
)

type Postgres struct {
	conn *sql.DB

	dbName   string
	hostName string
	port     int
	user     string
	pass     string
}

func New(dbName, hostName, user, pass string, port int) *Postgres {
	return &Postgres{
		dbName:   dbName,
		hostName: hostName,
		port:     port,
		user:     user,
		pass:     pass,
	}
}

func (p *Postgres) Open() error {
	conn, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		p.user, p.dbName, p.pass, p.hostName, p.port,
	))
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		conn.Close()
		return err
	}

	p.conn = conn

	return nil
}

func (p *Postgres) Close() {
	p.conn.Close()
}

func (p *Postgres) Init() error {
	createUser := `
		CREATE TABLE IF NOT EXISTS users
		(
			uuid UUID NOT NULL PRIMARY KEY,
			login TEXT NOT NULL,
			date TEXT NOT NULL
		);
`

	rows, err := p.conn.Query(createUser)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Err() != nil {
		return rows.Err()
	}

	return nil
}
