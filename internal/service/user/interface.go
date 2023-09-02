package user

import (
	"context"
	"controlUniversity/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, data Create) (string, error)
	GetOneUser(ctx context.Context, id int) (entity.User, error)
}
