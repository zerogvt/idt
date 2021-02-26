package data

import (
	"encoding/json"
	"errors"
	"io"
)

// User models a user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ToJSON serializes a user
func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// UserStore is an in-memory "DB" of users
type UserStore struct {
	store map[int]*User
}

//NewUserStore creates a new user store
func NewUserStore() *UserStore {
	return &UserStore{
		store: make(map[int]*User),
	}

}

// Read retrieves a user from the store
func (us *UserStore) Read(id int) (*User, error) {
	var user *User
	var found bool
	if user, found = us.store[id]; !found {
		return nil, errors.New("Not found")
	}
	return user, nil
}

// Update puts a user into the the store
func (us *UserStore) Update(id int, name string) error {
	user := User{
		ID:   id,
		Name: name,
	}
	us.store[id] = &user
	usr, ok := us.store[id]
	if !ok || *usr != user {
		return errors.New("DB Error")
	}
	return nil
}
