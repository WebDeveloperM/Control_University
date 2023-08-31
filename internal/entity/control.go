package entity

import (
	"github.com/uptrace/bun"
)

type Control struct {
	*bun.BaseModel `bun:"controls"`
	Id             int    `bun:"id,pk,autoincrement"`
	StudentId      int    `bun:"student_id"`
	Time           string `bun:"time"`
	Status         int    `bun:"status"`
}
