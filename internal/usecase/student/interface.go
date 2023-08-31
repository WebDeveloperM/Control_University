package student

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/student"
)

type Student interface {
	AddStudent(ctx context.Context, data student.Create) (entity.Student, error)
	GetDetailStudent(ctx context.Context, id int) (entity.Student, error)
}

type Controls interface {
	GetAll(ctx context.Context, filter student.Filter) ([]entity.Control, int, error)
}
