package entity

import (
	"github.com/uptrace/bun"
)

type UpdateControl struct {
	*bun.BaseModel `bun:"update_control"`
	Id             int    `bun:"id,pk,autoincrement"`
	ContorlId      int    `bun:"control_id"`
	UserId         int64  `bun:"user_id"`
	UpdateAt       string `bun:"update_at"`
	Changes        string `bun:"changes"`
}
