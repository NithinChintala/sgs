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
