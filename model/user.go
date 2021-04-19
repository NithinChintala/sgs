package model

import (
	"net/http"
)

type User struct {
	Id          int     `json:"id"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Email       string  `json:"email"`
	DateOfBirth *string `json:"dateOfBirth"`
}

func UserFromForm(r *http.Request) User {

	id := mustAtoi(r.FormValue("id"))
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	dob := nilifyStr(r.FormValue("dateOfBirth"))

	return User{id, firstName, lastName, username, password, email, dob}
}