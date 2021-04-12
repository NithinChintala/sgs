package dao

import (
	"database/sql"
	"github.com/NithinChintala/sgs/model"
	"log"
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
