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

// Get fetches a user from data store and returns a JSON representation of her
func (uh *UserHandler) Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	var id int
	var err error
	if id, err = strconv.Atoi(vars["id"]); err != nil {
		http.Error(rw, "Bad parameter", http.StatusBadRequest)
		return
	}
	var u *data.User
	if u, err = uh.db.Read(id); err != nil {
		http.Error(rw, "User not found", http.StatusNotFound)
		return
	}
	if err = u.ToJSON(rw); err != nil {
		http.Error(rw, "Error while serializing user", http.StatusInternalServerError)
	}
}

// Put updates or creates a user in data store.
func (uh *UserHandler) Put(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id int
	var name string
	var err error
	if id, err = strconv.Atoi(vars["id"]); err != nil {
		http.Error(rw, "Bad id parameter", http.StatusBadRequest)
		return
	}
	var ok bool
	if name, ok = vars["name"]; !ok {
		http.Error(rw, "Bad name parameter", http.StatusBadRequest)
		return
	}
	if err = uh.db.Update(id, name); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
