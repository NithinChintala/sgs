package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/NithinChintala/sgs/dao"
	"github.com/NithinChintala/sgs/model"
)

func (s *server) createTag(w http.ResponseWriter, r *http.Request) {
	newTag := model.TagFromForm(r)
	dao.CreateTag(newTag)
	http.Redirect(w, r, "/tags", 301)
}

func (s *server) getTags(w http.ResponseWriter, r *http.Request) {
	s.tmpl.ExecuteTemplate(w, "Tags", dao.GetTags())
}

func (s *server) getTagById(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id, _ := strconv.Atoi(args["id"])
	tag := dao.GetTagsById(id)

	s.tmpl.ExecuteTemplate(w, "EditTag", tag)
}

func (s *server) updateTag(w http.ResponseWriter, r *http.Request) {
	newTag := model.TagFromForm(r)
	dao.UpdateTag(newTag.Id, newTag)
	http.Redirect(w, r, "/tags", 301)
}

func (s *server) deleteTag(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id, _ := strconv.Atoi(args["id"])
	dao.DeleteTag(id)
	http.Redirect(w, r, "/tags", 301)
}

func (s *server) getTagsByPaperId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("paper"))
	s.tmpl.ExecuteTemplate(w, "Tags", dao.GetTagsByPaperId(id))
}