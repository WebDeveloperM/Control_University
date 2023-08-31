package student

import (
	"context"
	"controlUniversity/internal/entity"
)

type Repository interface {
	AddStudent(ctx context.Context, data Create) (entity.Student, error)
	GetDetailStudent(ctx context.Context, id int) (entity.Student, error)
}
