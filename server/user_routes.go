package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/NithinChintala/sgs/dao"
	"github.com/NithinChintala/sgs/model"
)

func (s *server) createUser(w http.ResponseWriter, r *http.Request) {
	newUser := model.UserFromForm(r)
	dao.CreateUser(newUser)
	http.Redirect(w, r, "/users", 301)
}

func (s *server) getUsers(w http.ResponseWriter, r *http.Request) {
	s.tmpl.ExecuteTemplate(w, "Users", dao.GetUsers())
}

func (s *server) getUserById(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id, _ := strconv.Atoi(args["id"])
	user := dao.GetUsersById(id)

	s.tmpl.ExecuteTemplate(w, "EditUser", user)
}

func (s *server) updateUser(w http.ResponseWriter, r *http.Request) {
	newUser := model.UserFromForm(r)
	dao.UpdateUser(newUser.Id, newUser)
	http.Redirect(w, r, "/users", 301)
}

func (s *server) deleteUser(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id, _ := strconv.Atoi(args["id"])
	dao.DeleteUser(id)
	http.Redirect(w, r, "/users", 301)
}
