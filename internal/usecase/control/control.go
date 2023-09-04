package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/control"
	"github.com/gin-gonic/gin"
)

type UseCase struct {
	control Control
	student Student
	user    User
}

func ControlUseCase(control Control) UseCase {
	return UseCase{control: control}
}

func (s UseCase) AddControl(ctx context.Context, data control.Create) (entity.Control, error) {
	return s.control.AddControl(ctx, data)
}

func (s UseCase) DeleteControl(ctx *gin.Context, id int) (bool, error) {
	return s.control.DeleteControl(ctx, id)
}

func (s UseCase) UpadateControl(ctx *gin.Context, data control.Update, id int) (entity.Control, error) {
	return s.control.UpadateControl(ctx, data, id)
}
