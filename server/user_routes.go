package server

import (
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/NithinChintala/sgs/dao"
	_ "github.com/NithinChintala/sgs/model"
)

func (s *server) createUser(w http.ResponseWriter, r *http.Request) {
}

func (s *server) getUsers(w http.ResponseWriter, r *http.Request) {
	s.tmpl.ExecuteTemplate(w, "Users", dao.GetUsers())
}

func (s *server) getUserById(w http.ResponseWriter, r *http.Request) {
}

func (s *server) updateUser(w http.ResponseWriter, r *http.Request) {
}

func (s *server) deleteUser(w http.ResponseWriter, r *http.Request) {
}