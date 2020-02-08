package db

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	UUID  uuid.UUID
	Login string
	Date  string
}

func (p *Postgres) Add(u *User) error {
	exist, err := p.isLoginExist(u.Login)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("login %q already exist", u.Login)
	}

	_, err = p.conn.Exec(
		"INSERT INTO users VALUES ($1, $2, $3)",
		u.UUID, u.Login, u.Date,
	)

	return err
}

func (p *Postgres) Get(login string) (*User, error) {
	exist, err := p.isLoginExist(login)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, fmt.Errorf("login %q not exist", login)
	}

	var u User
	err = p.conn.QueryRow("SELECT * FROM users WHERE login=$1", login).
		Scan(&u.UUID, &u.Login, &u.Date)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (p *Postgres) Set(u *User) error {
	exist, err := p.isLoginExist(u.Login)
	if err != nil {
		return err
	}

	if !exist {
		return fmt.Errorf("login %q not exist", u.Login)
	}

	_, err = p.conn.Exec(
		"UPDATE users SET uuid=$1, date=$3 WHERE login=$2",
		u.UUID, u.Login, u.Date,
	)

	return err
}

func (p *Postgres) isLoginExist(login string) (bool, error) {
	var exist bool
	err := p.conn.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM users WHERE login=$1);", login,
	).Scan(&exist)

	return exist, err
}
