package toolexchange

import (
	"time"
)

type Tool struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	ImageURL   string    `json:"imageURL"`
	IsBorrowed bool      `json:"isBorrowed"`
	BorrowedOn time.Time `json:"borrowedOn"`
}

type Tools []Tool
