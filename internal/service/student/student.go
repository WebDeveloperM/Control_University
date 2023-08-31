package student

import (
	"context"
	"controlUniversity/internal/entity"
)

type Service struct {
	repo Repository
}

func ServiceStudent(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) AddStudent(ctx context.Context, data Create) (entity.Student, error) {
	return s.repo.AddStudent(ctx, data)
}

func (s Service) GetDetailStudent(ctx context.Context, id int) (entity.Student, error) {
	return s.repo.GetDetailStudent(ctx, id)
}
