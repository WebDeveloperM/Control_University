package entity

import (
	"github.com/uptrace/bun"
)

type Control struct {
	*bun.BaseModel `bun:"controls"`
	Id             int    `bun:"id,pk,autoincrement"`
	StudentId      int    `bun:"student_id"`
	Time           string `bun:"time"`
	Status         string `bun:"status"`
	UserId         int    `bun:"user_id"`
	IsDelete       bool   `bun:"is_delete"`
}
