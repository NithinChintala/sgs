package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/NithinChintala/sgs/dao"
	"github.com/NithinChintala/sgs/model"
)

func (s *server) createPaper(w http.ResponseWriter, r *http.Request) {
	newPaper := model.PaperFromForm(r)
	dao.CreatePaper(newPaper)
	http.Redirect(w, r, "/papers", 301)
}

func (s *server) getPapers(w http.ResponseWriter, r *http.Request) {
	s.tmpl.ExecuteTemplate(w, "Papers", dao.GetPapers())
}

func (s *server) getPaperById(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id, _ := strconv.Atoi(args["id"])
	paper := dao.GetPapersById(id)

	s.tmpl.ExecuteTemplate(w, "EditPaper", paper)
}

func (s *server) updatePaper(w http.ResponseWriter, r *http.Request) {
	newPaper := model.PaperFromForm(r)
	dao.UpdatePaper(newPaper.Id, newPaper)
	http.Redirect(w, r, "/papers", 301)
}

func (s *server) deletePaper(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id, _ := strconv.Atoi(args["id"])
	dao.DeletePaper(id)
	http.Redirect(w, r, "/papers", 301)
}