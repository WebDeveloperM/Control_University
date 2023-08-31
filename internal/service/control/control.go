package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/student"
)

type Service struct {
	repo Repository
}

func ServiceControl(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) AddControl(ctx context.Context, data Create) (entity.Control, error) {
	return s.repo.AddControl(ctx, data)
}

func (s Service) GetAll(ctx context.Context, filter student.Filter) ([]entity.Control, int, error) {
	return s.repo.GetAllControls(ctx, filter)
}
