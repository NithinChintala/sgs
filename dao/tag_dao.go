package dao

import (
	"database/sql"
	"github.com/NithinChintala/sgs/model"
	"log"
	"net/http"
	"encoding/json"
)

func ReadTags(result *sql.Rows) []model.Tag {
	tags := make([]model.Tag, 0)
	for result.Next() {
		var tag model.Tag
		err := result.Scan(&tag.Id, &tag.Word, &tag.Searches, &tag.LastSearch)
		if err != nil {
			log.Fatal(err)
		}
		tags = append(tags, tag)
	}

	return tags
}

func GetTags(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	results, err := db.Query("SELECT * FROM tags")
	if err != nil {
		log.Fatal(err)
	}
	tags := ReadTags(results)
	json.NewEncoder(w).Encode(tags)
}
