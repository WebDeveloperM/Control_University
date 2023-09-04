package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/student"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	AddControl(ctx context.Context, data Create) (entity.Control, error)
	GetAllControls(ctx context.Context, filter student.Filter) ([]entity.Control, int, error)
	DeleteControl(ctx *gin.Context, id int) (bool, error)
	UpadateControl(ctx *gin.Context, data Update, id int) (entity.Control, error)
}
