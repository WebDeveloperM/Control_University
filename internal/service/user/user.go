package user

import (
	"context"
	"controlUniversity/internal/entity"
)

type Service struct {
	repo Repository
}

func ServiceUser(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CreateUser(ctx context.Context, data Create) (string, error) {
	return s.repo.CreateUser(ctx, data)
}

func (s Service) GetOneUser(ctx context.Context, id int) (entity.User, error) {
	return s.repo.GetOneUser(ctx, id)
}
