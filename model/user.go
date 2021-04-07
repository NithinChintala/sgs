package model

type User struct {
	Id          int     `json:"id"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Email       string  `json:"email"`
	DateOfBirth *string `json:"dateOfBirth"`
}
