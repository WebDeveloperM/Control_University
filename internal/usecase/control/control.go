package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/control"
)

type UseCase struct {
	control Control
}

func ControlUseCase(control Control) UseCase {
	return UseCase{control: control}
}

func (s UseCase) AddControl(ctx context.Context, data control.Create) (entity.Control, error) {
	return s.control.AddControl(ctx, data)
}
