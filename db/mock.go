package db

import "fmt"

type MockDB map[string]*User

func (m MockDB) Add(u *User) error {
	_, exist := m[u.Login]
	if exist {
		return fmt.Errorf("login %q already exist", u.Login)
	}

	m[u.Login] = u

	return nil
}

func (m MockDB) Get(login string) (*User, error) {
	user, exist := m[login]
	if !exist {
		return nil, fmt.Errorf("login %q not exist", login)
	}

	return user, nil
}

func (m MockDB) Set(u *User) error {
	_, exist := m[u.Login]
	if !exist {
		return fmt.Errorf("login %q not exist", u.Login)
	}

	m[u.Login] = u

	return nil
}
