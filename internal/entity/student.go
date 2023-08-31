package entity

import "github.com/uptrace/bun"

type Student struct {
	bun.BaseModel `bun:"students"`
	Id            int    `bun:"id,pk,autoincrement"`
	FirstName     string `bun:"firstname"`
	LastName      string `bun:"lastname"`
}
