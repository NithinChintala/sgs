package model

import (
	"time"
	"net/http"
)

type Tag struct {
	Id         int       `json:"id"`
	Word       string    `json:"word"`
	Searches   int       `json:"searches"`
	LastSearch *time.Time `json:"lastSearch"`
}

func TagFromForm(r *http.Request) Tag {

	id := mustAtoi(r.FormValue("id"))
	word := r.FormValue("word")
	searches := mustAtoi(r.FormValue("searches"))

	// TODO probably don't do this
	var lastSearch *time.Time = nil
	//lastSearch := nilifyStr(r.FormValue("lastSearch"))

	return Tag{id, word, searches, lastSearch}
}
