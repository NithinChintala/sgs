package dao

import (
	"database/sql"
	"github.com/NithinChintala/sgs/model"
	"log"
)

func ReadUsers(result *sql.Rows) []model.User {
	users := make([]model.User, 0)
	for result.Next() {
		var user model.User
		err := result.Scan(&user.Id, &user.FirstName, &user.LastName,
			&user.Username, &user.Password, &user.Email, &user.DateOfBirth)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}

func GetUsers() []model.User {
	connect()
	defer db.Close()
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	return ReadUsers(results)
}

func GetUsersById(id int) []model.User {
	connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	return ReadUsers(results)
}

func CreateUser(user model.User) {
	connect()
	defer db.Close()

	insert :=
	`
	INSERT INTO users (first_name, last_name, username, password, email, date_of_birth)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(insert, user.FirstName, user.LastName, user.Username, user.Password, user.Email, user.DateOfBirth)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUser(id int, user model.User) {
	connect()
	defer db.Close()

	update := `
	UPDATE users
	SET first_name=?, last_name=?, username=?, password=?, email=?, date_of_birth=?
	WHERE id=?
	`

	_, err := db.Exec(update, user.FirstName, user.LastName, user.Username, user.Password, user.Email, user.DateOfBirth, user.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUser(id int) {
	connect()
	defer db.Close()

	delete := "DELETE FROM users WHERE id=?"
	_, err := db.Exec(delete, id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUsersByPaperId(paperId int) []model.User {
	connect()
	defer db.Close()

	query := `
	SELECT users.* FROM papers, authors, users
	WHERE papers.id = authors.paper_id 
	AND authors.user_id = users.id 
	AND papers.id = ?
	`
	results, err := db.Query(query, paperId)
	if err != nil {
		log.Fatal(err)
	}
	return ReadUsers(results)
}