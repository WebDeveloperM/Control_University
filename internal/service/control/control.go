package control

import (
	"context"
	"controlUniversity/internal/entity"
	"controlUniversity/internal/service/student"
	"github.com/gin-gonic/gin"
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

func (s Service) DeleteControl(ctx *gin.Context, id int) (bool, error) {
	return s.repo.DeleteControl(ctx, id)
}

func (s Service) UpadateControl(ctx *gin.Context, data Update, id int) (entity.Control, error) {
	return s.repo.UpadateControl(ctx, data, id)
}
