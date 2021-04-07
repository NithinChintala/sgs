package model

type Reference struct {
	Id int `json:"id"`
	CiterId int `json:"citerId"` 
	CiteeId int `json:"citeeId"` 
}