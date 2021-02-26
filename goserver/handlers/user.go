package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zerogvt/idt/goserver/data"
)

// UserHandler is a http.Handler
type UserHandler struct {
	log *log.Logger
	db  *data.UserStore
}

// NewUserHandler creates a users handler with the given logger and db backend
func NewUserHandler(db *data.UserStore, l *log.Logger) *UserHandler {
	return &UserHandler{
		log: l,
		db:  db,
	}
}

// GetUser fetches a user from data store
func (uh *UserHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id int
	var err error
	if id, err = strconv.Atoi(vars["id"]); err != nil {
		http.Error(rw, "Bad parameter", http.StatusBadRequest)
		return
	}
	var u *data.User
	if u, err = uh.db.Get(id); err != nil {
		http.Error(rw, "User not found", http.StatusNotFound)
		return
	}
	// serialize the list to JSON
	if err = u.ToJSON(rw); err != nil {
		http.Error(rw, "Error while serializing user", http.StatusInternalServerError)
		return
	}
}

// PutUser puts a user to data store if her id is not already there
// in case of conflict returns an error
func (uh *UserHandler) PutUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id int
	var err error
	if id, err = strconv.Atoi(vars["id"]); err != nil {
		http.Error(rw, "Bad parameter", http.StatusBadRequest)
		return
	}
	if err = uh.db.Put(id); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
	}
	rw.WriteHeader(http.StatusCreated)
}
