package entity

import (
	"github.com/uptrace/bun"
)

type DeleteControl struct {
	*bun.BaseModel `bun:"delete_control"`
	Id             int    `bun:"id,pk,autoincrement"`
	ContorlId      int    `bun:"control_id"`
	UserId         int64  `bun:"user_id"`
	DeleteAt       string `bun:"delete_at"`
}
