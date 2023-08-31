package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Control struct {
	*bun.BaseModel `bun:"controls"`
	Id             int       `bun:"id,pk,autoincrement"`
	StudentId      int       `bun:"student_id"`
	Time           time.Time `bun:"time"`
	Status         int       `bun:"status"`
}
