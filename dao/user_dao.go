package dao

import (
	"database/sql"
	"github.com/NithinChintala/sgs/model"
	"log"
	"encoding/json"
	"net/http"
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

func GetUsersAPI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetUsers())
}
