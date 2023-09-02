package user

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/user"
)

type User interface {
	CreateUser(ctx context.Context, data user.Create) (string, error)
}

type Students interface {
	GetDetailStudent(ctx context.Context, id int) (entity.Student, error)
}

type Controls interface {
}
