package server

import (
	"net/http"
	"strconv"

	"github.com/NithinChintala/sgs/dao"
	"github.com/NithinChintala/sgs/model"
	"github.com/gorilla/mux"
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

func (s *server) getPapersByTagWord(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("tag")
	s.tmpl.ExecuteTemplate(w, "Papers", dao.GetPapersByTagWord(word))
}

func (s *server) getPapersByCiterId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("citer"))
	s.tmpl.ExecuteTemplate(w, "Papers", dao.GetPapersByCiterId(id))
}

func (s *server) getPapersByCiteeId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("citee"))
	s.tmpl.ExecuteTemplate(w, "Papers", dao.GetPapersByCiteeId(id))
}

func (s *server) getPapersByUserId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("author"))
	s.tmpl.ExecuteTemplate(w, "Papers", dao.GetPapersByUserId(id))
}
