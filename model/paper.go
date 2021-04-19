package model

import (
	"net/http"
	"strconv"
	"log"
)

type Paper struct {
	Id      int     `json:"id"`
	Year    int     `json:"year"`
	Title   string  `json:"title"`
	Journal *string  `json:"journal"`
	Volume  *int    `json:"volume"`
	Issue   *int    `json:"issue"`
	Pages   *string `json:"pages"`
	Doi     *string `json:"doi"`
}

func nilifyStr(str string) *string {
	if str == "<nil>" || str == "" {
		return nil
	}
	return &str
}

func mustAtoi(str string) int {
	conv, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return conv
}

func nilifyInt(str string) *int {
	if str == "<nil>" || str == "" {
		return nil
	}
	conv := mustAtoi(str)
	return &conv
}


func PaperFromForm(r *http.Request) Paper {

	id := mustAtoi(r.FormValue("id"))
	year := mustAtoi(r.FormValue("year"))
	title := r.FormValue("title")

	journal := nilifyStr(r.FormValue("journal"))
	volume := nilifyInt(r.FormValue("volume"))
	issue := nilifyInt(r.FormValue("issue"))
	pages := nilifyStr(r.FormValue("pages"))
	doi := nilifyStr(r.FormValue("doi"))

	return Paper{id, year, title, journal, volume, issue, pages, doi}
}
