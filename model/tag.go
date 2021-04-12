package model

import (
	"time"
)

type Tag struct {
	Id         int       `json:"id"`
	Word       string    `json:"firstName"`
	Searches   int       `json:"searches"`
	LastSearch *time.Time `json:"lastSearch"`
}
