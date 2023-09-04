package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/control"
	"github.com/gin-gonic/gin"
)

type Control interface {
	AddControl(ctx context.Context, data control.Create) (entity.Control, error)
	DeleteControl(ctx *gin.Context, id int) (bool, error)
	UpadateControl(ctx *gin.Context, data control.Update, id int) (entity.Control, error)
}

type Student struct {
}
type User struct {
}
