package student

import (
	"controlUniversity/internal/entity"
)

type Filter struct {
	StudentId *int
	Search    *string
}

type Create struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserId    int    `json:"user_id"`
}

type Detail struct {
	Id        int              `json:"id"`
	FirstName string           `json:"firstname"`
	Lastname  string           `json:"lastname"`
	UserId    int              `json:"user_id"`
	Controls  []entity.Control `json:"controls"`
}
