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

func GetTags() []model.Tag {
	connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM tags")
	if err != nil {
		log.Fatal(err)
	}
	return ReadTags(results)
}

func GetTagsById(id int) []model.Tag {
	connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM tags WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	return ReadTags(results)
}

func CreateTag(tag model.Tag) {
	connect()
	defer db.Close()

	insert := "INSERT INTO tags (word, searches, last_search) VALUES (?, ?, ?)"
	_, err := db.Exec(insert, tag.Word, tag.Searches, tag.LastSearch)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateTag(id int, tag model.Tag) {
	connect()
	defer db.Close()

	update := `
	UPDATE tags
	SET word=?, searches=?, last_search=?
	WHERE id=?
	`

	_, err := db.Exec(update, tag.Word, tag.Searches, tag.LastSearch, tag.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTag(id int) {
	connect()
	defer db.Close()

	delete := "DELETE FROM tags WHERE id=?"
	_, err := db.Exec(delete, id)
	if err != nil {
		log.Fatal(err)
	}
}
