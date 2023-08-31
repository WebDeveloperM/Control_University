package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/student"
)

type Repository interface {
	AddControl(ctx context.Context, data Create) (entity.Control, error)
	GetAllControls(ctx context.Context, filter student.Filter) ([]entity.Control, int, error)
}
