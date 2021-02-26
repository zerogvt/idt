package data

import (
	"encoding/json"
	"errors"
	"io"
)

// User models a user
type User struct {
	ID int `json:"id"`
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

//Get retrieves a user from the store
func (us *UserStore) Get(id int) (*User, error) {
	var user *User
	var found bool
	if user, found = us.store[id]; !found {
		return nil, errors.New("Not found")
	}
	return user, nil
}

//Put puts a user into the the store
func (us *UserStore) Put(id int) error {
	user := User{
		ID: id,
	}
	_, exists := us.store[id]
	if exists {
		return errors.New("User id exists")
	}
	us.store[id] = &user
	return nil
}
